package main

import (
	"fmt"
	"github.com/yukitsune/chameleon/cmd"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/handlers"
	"github.com/yukitsune/chameleon/pkg/smtp"
	"net/url"
	"os"
	"time"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		cmd.ExitFromError(err)
	}

	logFactory := log.NewLogFactory(
		fmt.Sprintf("%s/logs", wd),
		log.DebugLevel,
		log.DefaultFileNameProvider,
		log.LogrusLoggerProvider,
	)

	logger, err := logFactory.Make()
	if err != nil {
		cmd.ExitFromError(err)
	}

	// Todo: Assign these via environment variables or command line args
	// Todo: Revisit TLS configuration
	startTlSOn := false
	tlsAlwaysOn := false
	protocols := []string  {"tls1.0", "tls1.2"}
	ciphers := []string {"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305", "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305", "TLS_RSA_WITH_RC4_128_SHA", "TLS_RSA_WITH_AES_128_GCM_SHA256", "TLS_RSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA", "TLS_ECDHE_RSA_WITH_RC4_128_SHA", "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"}
	curves := []string {"P256", "P384", "P521", "X25519"}
	privateKeyFile := ""
	publicKeyFile := ""
	rootCAs := ""
	clientAuthType := ""

	tlsConfig := smtp.ServerTLSConfig{
		StartTLSOn:               startTlSOn,
		AlwaysOn:                 tlsAlwaysOn,
		Protocols:                protocols,
		Ciphers:                  ciphers,
		Curves:                   curves,
		PrivateKeyFile:           privateKeyFile,
		PublicKeyFile:            publicKeyFile,
		RootCAs:                  rootCAs,
		ClientAuthType:           clientAuthType,
		PreferServerCipherSuites: false,
	}

	// Todo: Assign these via environment variables or command line args
	hostName := "relay.chameleon.io"
	var maxSize int64 = 50000000 // Bytes
	timeout := 30 * time.Second
	maxClients := 1000
	xClientOn := false

	config := &smtp.ServerConfig{
		TLS:             tlsConfig,
		Hostname:        hostName,
		ListenInterface: "",
		MaxSize:         maxSize,
		Timeout:         int(timeout.Seconds()),
		MaxClients:      maxClients,
		XClientOn:       xClientOn,
	}

	apiUrl, err := url.Parse("https://api.chameleon.io/")
	if err != nil {
		logger.Fatal(err)
	}

	handler := handlers.NewDefaultHandler(apiUrl, logger)

	server, err := smtp.NewServer(config, handler, logger)
	if err != nil {
		logger.Fatal(err)
	}

	go func(logger log.ChameleonLogger) {
		if err := server.Start(); err != nil {
			logger.Fatal(err)
		}
	}(logger)

	// server.Start doesn't block, wait for exit signal
	cmd.SigHandler(logger)
	server.Shutdown()
}
