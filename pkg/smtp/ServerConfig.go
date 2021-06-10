package smtp

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"
)

const defaultMaxClients = 100
const defaultTimeout = 30
const defaultInterface = "127.0.0.1:2525"
const defaultMaxSize = int64(10 << 20) // 10 Mebibytes

// ServerConfig specifies config options for a single server
type ServerConfig struct {
	// TLS Configuration
	TLS *ServerTLSConfig `yaml:"tls,omitempty"`

	// Hostname will be used in the server's reply to HELO/EHLO
	// If TLS enabled, make sure that the Hostname matches the cert
	// Defaults to os.Hostname()
	// Hostname will also be used to fill the 'Host' property when the "RCPT TO" address is addressed to just <postmaster>
	Hostname string `yaml:"hostname"`

	// Listen interface specified in <ip>:<port> - defaults to 127.0.0.1:2525
	ListenInterface string `yaml:"listen-interface"`

	// MaxSize is the maximum size of an email that will be accepted for delivery
	// Defaults to 10 Mebibytes
	MaxSize int64 `yaml:"max-mail-size"`

	// Timeout specifies the connection timeout in seconds. Defaults to 30
	Timeout int `yaml:"timeout"`

	// MaxClients controls how many maximum clients we can handle at once
	// Defaults to defaultMaxClients
	MaxClients int `yaml:"max-clients"`

	// Todo: Revisit
	// AllowedHosts lists which hosts to accept email for. Defaults to os.Hostname
	AllowedHosts []string `yaml:"allowed-hosts"`

	// XClientOn when using a proxy such as Nginx, XCLIENT command is used to pass the
	// original client's IP address & client's HELO
	XClientOn bool `yaml:"xclient-on,omitempty"`
}

// Loads in timestamps for the TLS keys
func (sc *ServerConfig) loadTlsKeyTimestamps() error {
	var statErr = func(iface string, err error) error {
		return fmt.Errorf(
			"could not stat key for server [%s], %s",
			iface,
			err.Error())
	}
	if sc.TLS.PrivateKeyFile == "" {
		sc.TLS._privateKeyFileMtime = time.Now().Unix()
		return nil
	}
	if sc.TLS.PublicKeyFile == "" {
		sc.TLS._publicKeyFileMtime = time.Now().Unix()
		return nil
	}
	if info, err := os.Stat(sc.TLS.PrivateKeyFile); err == nil {
		sc.TLS._privateKeyFileMtime = info.ModTime().Unix()
	} else {
		return statErr(sc.ListenInterface, err)
	}
	if info, err := os.Stat(sc.TLS.PublicKeyFile); err == nil {
		sc.TLS._publicKeyFileMtime = info.ModTime().Unix()
	} else {
		return statErr(sc.ListenInterface, err)
	}
	return nil
}

// Returns value changes between struct a & struct b.
// Results are returned in a map, where each key is the name of the field that was different.
// a and b are struct values, must not be pointer
// and of the same struct type
func getChanges(a interface{}, b interface{}) map[string]interface{} {
	ret := make(map[string]interface{}, 5)
	compareWith := structtomap(b)
	for key, val := range structtomap(a) {
		if sliceOfStr, ok := val.([]string); ok {
			val, _ = json.Marshal(sliceOfStr)
			val = string(val.([]uint8))
		}
		if sliceOfStr, ok := compareWith[key].([]string); ok {
			compareWith[key], _ = json.Marshal(sliceOfStr)
			compareWith[key] = string(compareWith[key].([]uint8))
		}
		if val != compareWith[key] {
			ret[key] = compareWith[key]
		}
	}
	// detect changes to TLS keys (have the key files been modified?)
	if oldTLS, ok := a.(ServerTLSConfig); ok {
		t1, t2 := oldTLS.getTlsKeyTimestamps()
		if newTLS, ok := b.(ServerTLSConfig); ok {
			t3, t4 := newTLS.getTlsKeyTimestamps()
			if t1 != t3 {
				ret["PrivateKeyFile"] = newTLS.PrivateKeyFile
			}
			if t2 != t4 {
				ret["PublicKeyFile"] = newTLS.PublicKeyFile
			}
		}
	}
	return ret
}

// Convert fields of a struct to a map
// only able to convert int, bool, slice-of-strings and string; not recursive
// slices are marshal'd to json for convenient comparison later
func structtomap(obj interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	t := v.Type()
	for index := 0; index < v.NumField(); index++ {
		vField := v.Field(index)
		fName := t.Field(index).Name
		k := vField.Kind()
		switch k {
		case reflect.Int:
			fallthrough
		case reflect.Int64:
			value := vField.Int()
			ret[fName] = value
		case reflect.String:
			value := vField.String()
			ret[fName] = value
		case reflect.Bool:
			value := vField.Bool()
			ret[fName] = value
		case reflect.Slice:
			ret[fName] = vField.Interface().([]string)
		}
	}
	return ret
}

// SetDefaults fills in default server settings for values that were not configured
// The defaults are:
// * Server listening to 127.0.0.1:2525
// * use your hostname to determine your which hosts to accept email for
// * 100 maximum clients
// * 10MB max message size
// * log to Stderr,
// * log level set to "`debug`"
// * timeout to 30 sec
// * Backend configured with the following processors: `HeadersParser|Header|Debugger`
// where it will log the received emails.
func (c *ServerConfig) SetDefaults() error {

	err := c.TLS.SetDefaults()
	if err != nil {
		return err
	}

	if len(c.AllowedHosts) == 0 {
		if h, err := os.Hostname(); err != nil {
			return err
		} else {
			c.AllowedHosts = append(c.AllowedHosts, h)
		}
	}

	h, err := os.Hostname()
	if err != nil {
		return err
	}

	if c.Hostname == "" {
		c.Hostname = h
	}
	if c.MaxClients == 0 {
		c.MaxClients = defaultMaxClients
	}
	if c.Timeout == 0 {
		c.Timeout = defaultTimeout
	}
	if c.MaxSize == 0 {
		c.MaxSize = defaultMaxSize // 10 Mebibytes
	}
	if c.ListenInterface == "" {
		c.ListenInterface = defaultInterface
	}

	return nil
}
