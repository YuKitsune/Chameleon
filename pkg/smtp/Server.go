package smtp

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/yukitsune/chameleon"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/internal/rfc5321"
	"io"
	"io/ioutil"
	"net"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	CommandVerbMaxLength = 16
	CommandLineMaxLength = 1024
	// Number of allowed unrecognized commands before we terminate the connection
	MaxUnrecognizedCommands = 5
)

const (
	// server has just been created
	ServerStateNew = iota
	// Server has just been stopped
	ServerStateStopped
	// Server has been started and is running
	ServerStateRunning
	// Server could not start due to an error
	ServerStateStartError
)

// Server listens for SMTP clients on the port specified in its config
type Server struct {
	config		    *ServerConfig // stores guerrilla.ServerConfig
	tlsConfig		*tls.Config
	timeout         time.Duration // stores time.Duration
	listenInterface string
	clientPool      *Pool
	wg              sync.WaitGroup // for waiting to shutdown
	listener        net.Listener
	closedListener  chan bool
	hosts           allowedHosts // stores map[string]bool for faster lookup
	state           int
	// If log changed after a config reload, newLogStore stores the value here until it's safe to change it
	log				log.ChameleonLogger
	handler		 	Handler
	envelopePool 	*EnvelopePool
}

type allowedHosts struct {
	table      map[string]bool // host lookup table
	wildcards  []string        // host wildcard list (* is used as a wildcard)
	sync.Mutex                 // guard access to the map
}

type command []byte

var (
	cmdHELO     command = []byte("HELO")
	cmdEHLO     command = []byte("EHLO")
	cmdHELP     command = []byte("HELP")
	cmdXCLIENT  command = []byte("XCLIENT")
	cmdMAIL     command = []byte("MAIL FROM:")
	cmdRCPT     command = []byte("RCPT TO:")
	cmdRSET     command = []byte("RSET")
	cmdVRFY     command = []byte("VRFY")
	cmdNOOP     command = []byte("NOOP")
	cmdQUIT     command = []byte("QUIT")
	cmdDATA     command = []byte("DATA")
	cmdSTARTTLS command = []byte("STARTTLS")
)

func (c command) match(in []byte) bool {
	return bytes.Index(in, []byte(c)) == 0
}

// Creates and returns a new ready-to-run Server from a ServerConfig configuration
func NewServer(sc *ServerConfig, handler Handler, log log.ChameleonLogger) (*Server, error) {
	server := &Server{
		config:  sc,
		clientPool:      NewPool(sc.MaxClients),
		closedListener:  make(chan bool, 1),
		listenInterface: sc.ListenInterface,
		state:           ServerStateNew,
		handler:         handler,
		log:             log,
		envelopePool:    NewEnvelopePool(sc.MaxClients),
	}

	server.setTimeout(sc.Timeout)
	if err := server.configureTLS(); err != nil {
		return server, err
	}
	return server, nil
}

func (s *Server) configureTLS() error {
	sConfig := s.config
	if sConfig.TLS.AlwaysOn || sConfig.TLS.StartTLSOn {
		cert, err := tls.LoadX509KeyPair(sConfig.TLS.PublicKeyFile, sConfig.TLS.PrivateKeyFile)
		if err != nil {
			return fmt.Errorf("error while loading the certificate: %s", err)
		}
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			ClientAuth:   tls.VerifyClientCertIfGiven,
			ServerName:   sConfig.Hostname,
		}
		if len(sConfig.TLS.Protocols) > 0 {
			if min, ok := TLSProtocols[sConfig.TLS.Protocols[0]]; ok {
				tlsConfig.MinVersion = min
			}
		}
		if len(sConfig.TLS.Protocols) > 1 {
			if max, ok := TLSProtocols[sConfig.TLS.Protocols[1]]; ok {
				tlsConfig.MaxVersion = max
			}
		}
		if len(sConfig.TLS.Ciphers) > 0 {
			for _, val := range sConfig.TLS.Ciphers {
				if c, ok := TLSCiphers[val]; ok {
					tlsConfig.CipherSuites = append(tlsConfig.CipherSuites, c)
				}
			}
		}
		if len(sConfig.TLS.Curves) > 0 {
			for _, val := range sConfig.TLS.Curves {
				if c, ok := TLSCurves[val]; ok {
					tlsConfig.CurvePreferences = append(tlsConfig.CurvePreferences, c)
				}
			}
		}
		if len(sConfig.TLS.RootCAs) > 0 {
			caCert, err := ioutil.ReadFile(sConfig.TLS.RootCAs)
			if err != nil {
				s.log.WithError(err).Errorf("failed opening TLSRootCAs file [%s]", sConfig.TLS.RootCAs)
			} else {
				caCertPool := x509.NewCertPool()
				caCertPool.AppendCertsFromPEM(caCert)
				tlsConfig.RootCAs = caCertPool
			}

		}
		if len(sConfig.TLS.ClientAuthType) > 0 {
			if ca, ok := TLSClientAuthTypes[sConfig.TLS.ClientAuthType]; ok {
				tlsConfig.ClientAuth = ca
			}
		}
		tlsConfig.PreferServerCipherSuites = sConfig.TLS.PreferServerCipherSuites
		tlsConfig.Rand = rand.Reader
		s.tlsConfig = tlsConfig
	}
	return nil
}

// Set the timeout for the server and all clients
func (s *Server) setTimeout(seconds int) {
	duration := time.Duration(int64(seconds))
	s.clientPool.SetTimeout(duration)
	s.timeout = duration
}

// Set the allowed hosts for the server
func (s *Server) setAllowedHosts(allowedHosts []string) {
	s.hosts.Lock()
	defer s.hosts.Unlock()
	s.hosts.table = make(map[string]bool, len(allowedHosts))
	s.hosts.wildcards = nil
	for _, h := range allowedHosts {
		if strings.Contains(h, "*") {
			s.hosts.wildcards = append(s.hosts.wildcards, strings.ToLower(h))
		} else if len(h) > 5 && h[0] == '[' && h[len(h)-1] == ']' {
			if ip := net.ParseIP(h[1 : len(h)-1]); ip != nil {
				// this will save the normalized ip, as ip.String always returns ipv6 in short form
				s.hosts.table["["+ip.String()+"]"] = true
			}
		} else {
			s.hosts.table[strings.ToLower(h)] = true
		}
	}
}

// Begin accepting SMTP clients. Will block unless there is an error or server.Shutdown() is called
func (s *Server) Start(startWG *sync.WaitGroup) error {
	var clientID uint64
	clientID = 0

	listener, err := net.Listen("tcp", s.listenInterface)
	s.listener = listener
	if err != nil {
		startWG.Done() // don't wait for me
		s.state = ServerStateStartError
		return fmt.Errorf("[%s] Cannot listen on port: %s ", s.listenInterface, err.Error())
	}

	s.log.Infof("Listening on TCP %s", s.listenInterface)
	s.state = ServerStateRunning
	startWG.Done() // start successful, don't wait for me

	for {
		s.log.Debugf("[%s] Waiting for a new client. Next Client ID: %d", s.listenInterface, clientID+1)
		conn, err := listener.Accept()
		clientID++
		if err != nil {
			if e, ok := err.(net.Error); ok && !e.Temporary() {
				s.log.Infof("Server [%s] has stopped accepting new clients", s.listenInterface)
				// the listener has been closed, wait for clients to exit
				s.log.Infof("shutting down pool [%s]", s.listenInterface)
				s.clientPool.ShutdownState()
				s.clientPool.ShutdownWait()
				s.state = ServerStateStopped
				s.closedListener <- true
				return nil
			}
			s.log.WithError(err).Info("Temporary error accepting client")
			continue
		}
		go func(p Poolable, borrowErr error) {
			c := p.(*client)
			if borrowErr == nil {
				s.handleClient(c)
				s.envelopePool.Return(c.Envelope)
				s.clientPool.Return(c)
			} else {
				s.log.WithError(borrowErr).Info("couldn't borrow a new client")
				// we could not get a client, so close the connection.
				_ = conn.Close()
			}
			// intentionally placed Borrow in args so that it's called in the
			// same main goroutine.
		}(s.clientPool.Borrow(conn, clientID, s.log, s.envelopePool))

	}
}

func (s *Server) Shutdown() {
	if s.listener != nil {
		// This will cause Start function to return, by causing an error on listener.Accept
		_ = s.listener.Close()
		// wait for the listener to listener.Accept
		<-s.closedListener
		// At this point Start will exit and close down the pool
	} else {
		s.clientPool.ShutdownState()
		// listener already closed, wait for clients to exit
		s.clientPool.ShutdownWait()
		s.state = ServerStateStopped
	}
}

func (s *Server) GetActiveClientsCount() int {
	return s.clientPool.GetActiveClientsCount()
}

// Verifies that the host is a valid recipient.
// host checking turned off if there is a single entry and it's a dot.
func (s *Server) allowsHost(host string) bool {
	s.hosts.Lock()
	defer s.hosts.Unlock()
	// if hosts contains a single dot, further processing is skipped
	if len(s.hosts.table) == 1 {
		if _, ok := s.hosts.table["."]; ok {
			return true
		}
	}
	if _, ok := s.hosts.table[strings.ToLower(host)]; ok {
		return true
	}
	// check the wildcards
	for _, w := range s.hosts.wildcards {
		if matched, err := filepath.Match(w, strings.ToLower(host)); matched && err == nil {
			return true
		}
	}
	return false
}

func (s *Server) allowsIp(ip net.IP) bool {
	ipStr := ip.String()
	return s.allowsHost("[" + ipStr + "]")
}

const commandSuffix = "\r\n"

// Reads from the client until a \n terminator is encountered,
// or until a timeout occurs.
func (s *Server) readCommand(client *client) ([]byte, error) {
	//var input string
	var err error
	var bs []byte
	// In command state, stop reading at line breaks
	bs, err = client.bufin.ReadSlice('\n')
	if err != nil {
		return bs, err
	} else if bytes.HasSuffix(bs, []byte(commandSuffix)) {
		return bs[:len(bs)-2], err
	}
	return bs[:len(bs)-1], err
}

// flushResponse a response to the client. Flushes the client.bufout buffer to the connection
func (s *Server) flushResponse(client *client) error {
	if err := client.setTimeout(s.timeout); err != nil {
		return err
	}
	return client.bufout.Flush()
}

func (s *Server) isShuttingDown() bool {
	return s.clientPool.IsShuttingDown()
}

// Handles an entire client SMTP exchange
func (s *Server) handleClient(client *client) {
	defer client.closeConn()
	sc := s.config
	s.log.Infof("Handle client [%s], id: %d", client.RemoteIP, client.ID)

	// Initial greeting
	greeting := fmt.Sprintf("220 %s SMTP Chameleon(%s) #%d (%d) %s",
		sc.Hostname, chameleon.Version, client.ID,
		s.clientPool.GetActiveClientsCount(), time.Now().Format(time.RFC3339))

	helo := fmt.Sprintf("250 %s Hello", sc.Hostname)
	// ehlo is a multi-line reply and need additional \r\n at the end
	ehlo := fmt.Sprintf("250-%s Hello\r\n", sc.Hostname)

	// Extended feature advertisements
	messageSize := fmt.Sprintf("250-SIZE %d\r\n", sc.MaxSize)
	pipelining := "250-PIPELINING\r\n"
	advertiseTLS := "250-STARTTLS\r\n"
	advertiseEnhancedStatusCodes := "250-ENHANCEDSTATUSCODES\r\n"
	// The last line doesn't need \r\n since string will be printed as a new line.
	// Also, Last line has no dash -
	help := "250 HELP"

	if sc.TLS.AlwaysOn {
		tlsConfig := s.tlsConfig
		if err := client.upgradeToTLS(tlsConfig); err == nil {
			advertiseTLS = ""
		} else {
			s.log.WithError(err).Warnf("[%s] Failed TLS handshake", client.RemoteIP)
			// server requires TLS, but can't handshake
			client.kill()
		}
	}
	if !sc.TLS.StartTLSOn {
		// STARTTLS turned off, don't advertise it
		advertiseTLS = ""
	}
	r := Canned
	for client.isAlive() {
		switch client.state {
		case ClientGreeting:
			client.sendResponse(greeting)
			client.state = ClientCmd
		case ClientCmd:
			client.bufin.setLimit(CommandLineMaxLength)
			input, err := s.readCommand(client)
			s.log.Debugf("Client sent: %s", input)
			if err == io.EOF {
				s.log.WithError(err).Warnf("Client closed the connection: %s", client.RemoteIP)
				return
			} else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				s.log.WithError(err).Warnf("Timeout: %s", client.RemoteIP)
				return
			} else if err == LineLimitExceeded {
				client.sendResponse(r.FailLineTooLong)
				client.kill()
				break
			} else if err != nil {
				s.log.WithError(err).Warnf("Read error: %s", client.RemoteIP)
				client.kill()
				break
			}
			if s.isShuttingDown() {
				client.state = ClientShutdown
				continue
			}

			cmdLen := len(input)
			if cmdLen > CommandVerbMaxLength {
				cmdLen = CommandVerbMaxLength
			}
			cmd := bytes.ToUpper(input[:cmdLen])
			switch {
			case cmdHELO.match(cmd):
				if h, err := client.parser.Helo(input[4:]); err == nil {
					client.Helo = h
				} else {
					s.log.WithFields(log.Fields{"helo": h, "client": client.ID}).Warn("invalid helo")
					client.sendResponse(r.FailSyntaxError)
					break
				}
				client.resetTransaction()
				client.sendResponse(helo)

			case cmdEHLO.match(cmd):
				if h, _, err := client.parser.Ehlo(input[4:]); err == nil {
					client.Helo = h
				} else {
					client.sendResponse(r.FailSyntaxError)
					s.log.WithFields(log.Fields{"ehlo": h, "client": client.ID}).Warn("invalid ehlo")
					client.sendResponse(r.FailSyntaxError)
					break
				}
				client.ESMTP = true
				client.resetTransaction()
				client.sendResponse(ehlo,
					messageSize,
					pipelining,
					advertiseTLS,
					advertiseEnhancedStatusCodes,
					help)

			case cmdHELP.match(cmd):
				client.sendResponse("214-OK\r\n", "https://www.youtube.com/watch?v=dQw4w9WgXcQ")

			case sc.XClientOn && cmdXCLIENT.match(cmd):
				if toks := bytes.Split(input[8:], []byte{' '}); len(toks) > 0 {
					for i := range toks {
						if vals := bytes.Split(toks[i], []byte{'='}); len(vals) == 2 {
							if bytes.Equal(vals[1], []byte("[UNAVAILABLE]")) {
								// skip
								continue
							}
							if bytes.Equal(vals[0], []byte("ADDR")) {
								client.RemoteIP = string(vals[1])
							}
							if bytes.Equal(vals[0], []byte("HELO")) {
								client.Helo = string(vals[1])
							}
						}
					}
				}
				client.sendResponse(r.SuccessMailCmd)
			case cmdMAIL.match(cmd):
				if client.isInTransaction() {
					client.sendResponse(r.FailNestedMailCmd)
					break
				}
				client.MailFrom, err = client.parsePath(input[10:], client.parser.MailFrom)
				if err != nil {
					s.log.WithError(err).Error("MAIL parse error", "["+string(input[10:])+"]")
					client.sendResponse(err)
					break
				} else if client.parser.NullPath {
					// bounce has empty from address
					client.MailFrom = Address{}
				}
				client.sendResponse(r.SuccessMailCmd)

			case cmdRCPT.match(cmd):
				if len(client.RcptTo) > rfc5321.LimitRecipients {
					client.sendResponse(r.ErrorTooManyRecipients)
					break
				}
				to, err := client.parsePath(input[8:], client.parser.RcptTo)
				if err != nil {
					s.log.WithError(err).Error("RCPT parse error", "["+string(input[8:])+"]")
					client.sendResponse(err.Error())
					break
				}
				s.defaultHost(&to)
				if (to.IP != nil && !s.allowsIp(to.IP)) || (to.IP == nil && !s.allowsHost(to.Host)) {
					client.sendResponse(r.ErrorRelayDenied, " ", to.Host)
				} else {
					client.PushRcpt(to)
					rcptError := s.handler.ValidateRcpt(client.Envelope, s.log)
					if rcptError != nil {
						client.PopRcpt()
						client.sendResponse(r.FailRcptCmd, " ", rcptError)
					} else {
						client.sendResponse(r.SuccessRcptCmd)
					}
				}

			case cmdRSET.match(cmd):
				client.resetTransaction()
				client.sendResponse(r.SuccessResetCmd)

			case cmdVRFY.match(cmd):
				client.sendResponse(r.SuccessVerifyCmd)

			case cmdNOOP.match(cmd):
				client.sendResponse(r.SuccessNoopCmd)

			case cmdQUIT.match(cmd):
				client.sendResponse(r.SuccessQuitCmd)
				client.kill()

			case cmdDATA.match(cmd):
				if len(client.RcptTo) == 0 {
					client.sendResponse(r.FailNoRecipientsDataCmd)
					break
				}
				client.sendResponse(r.SuccessDataCmd)
				client.state = ClientData

			case sc.TLS.StartTLSOn && cmdSTARTTLS.match(cmd):

				client.sendResponse(r.SuccessStartTLSCmd)
				client.state = ClientStartTLS
			default:
				client.errors++
				if client.errors >= MaxUnrecognizedCommands {
					client.sendResponse(r.FailMaxUnrecognizedCmd)
					client.kill()
				} else {
					client.sendResponse(r.FailUnrecognizedCmd)
				}
			}

		case ClientData:

			// intentionally placed the limit 1MB above so that reading does not return with an error
			// if the client goes a little over. Anything above will err
			client.bufin.setLimit(sc.MaxSize + 1024000) // This a hard limit.

			n, err := client.Data.ReadFrom(client.smtpReader.DotReader())
			if n > sc.MaxSize {
				err = fmt.Errorf("maximum DATA size exceeded (%d)", sc.MaxSize)
			}
			if err != nil {
				if err == LineLimitExceeded {
					client.sendResponse(r.FailReadLimitExceededDataCmd, " ", LineLimitExceeded.Error())
					client.kill()
				} else if err == MessageSizeExceeded {
					client.sendResponse(r.FailMessageSizeExceeded, " ", MessageSizeExceeded.Error())
					client.kill()
				} else {
					client.sendResponse(r.FailReadErrorDataCmd, " ", err.Error())
					client.kill()
				}
				s.log.WithError(err).Warn("Error reading data")
				client.resetTransaction()
				break
			}

			res := s.handler.Handle(client.Envelope, s.log)
			if res.Code() < 300 {
				client.messagesSent++
			}
			client.sendResponse(res)
			client.state = ClientCmd
			if s.isShuttingDown() {
				client.state = ClientShutdown
			}
			client.resetTransaction()

		case ClientStartTLS:
			if !client.TLS && sc.TLS.StartTLSOn {
				tlsConfig := s.tlsConfig
				if err := client.upgradeToTLS(tlsConfig); err == nil {
					advertiseTLS = ""
					client.resetTransaction()
				} else {
					s.log.WithError(err).Warnf("[%s] Failed TLS handshake", client.RemoteIP)
					// Don't disconnect, let the client decide if it wants to continue
				}
			}
			// change to command state
			client.state = ClientCmd
		case ClientShutdown:
			// shutdown state
			client.sendResponse(r.ErrorShutdown)
			client.kill()
		}

		if client.bufErr != nil {
			s.log.WithError(client.bufErr).Debug("client could not buffer a response")
			return
		}
		// flush the response buffer
		if client.bufout.Buffered() > 0 {
			if s.log.IsDebug() {
				s.log.Debugf("Writing response to client: \n%s", client.response.String())
			}
			err := s.flushResponse(client)
			if err != nil {
				s.log.WithError(err).Debug("error writing response")
				return
			}
		}

	}
}

// defaultHost ensures that the host attribute is set, if addressed to Postmaster
func (s *Server) defaultHost(a *Address) {
	if a.Host == "" && a.IsPostmaster() {
		sc := s.config
		a.Host = sc.Hostname
		if !s.allowsHost(a.Host) {
			s.log.WithFields(
				log.Fields{"hostname": sc.Hostname}).
				Warn("the hostname is not present in AllowedHosts config setting")
		}
	}
}
