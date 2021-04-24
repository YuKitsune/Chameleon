package smtp

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"
)

// ServerConfig specifies config options for a single server
type ServerConfig struct {
	// TLS Configuration
	TLS ServerTLSConfig `json:"tls,omitempty"`
	// LogFile is where the logs go. Use path to file, or "stderr", "stdout" or "off".
	// defaults to AppConfig.Log file setting
	LogFile string `json:"log_file,omitempty"`
	// Hostname will be used in the server's reply to HELO/EHLO. If TLS enabled
	// make sure that the Hostname matches the cert. Defaults to os.Hostname()
	// Hostname will also be used to fill the 'Host' property when the "RCPT TO" address is
	// addressed to just <postmaster>
	Hostname string `json:"host_name"`
	// Listen interface specified in <ip>:<port> - defaults to 127.0.0.1:2525
	ListenInterface string `json:"listen_interface"`
	// MaxSize is the maximum size of an email that will be accepted for delivery.
	// Defaults to 10 Mebibytes
	MaxSize int64 `json:"max_size"`
	// Timeout specifies the connection timeout in seconds. Defaults to 30
	Timeout int `json:"timeout"`
	// MaxClients controls how many maximum clients we can handle at once.
	// Defaults to defaultMaxClients
	MaxClients int `json:"max_clients"`
	// IsEnabled set to true to start the server, false will ignore it
	IsEnabled bool `json:"is_enabled"`
	// XClientOn when using a proxy such as Nginx, XCLIENT command is used to pass the
	// original client's IP address & client's HELO
	XClientOn bool `json:"xclient_on,omitempty"`
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

// Validate validates the server's configuration.
func (sc *ServerConfig) Validate() error {
	var errs Errors

	if sc.TLS.StartTLSOn || sc.TLS.AlwaysOn {
		if sc.TLS.PublicKeyFile == "" {
			errs = append(errs, errors.New("PublicKeyFile is empty"))
		}
		if sc.TLS.PrivateKeyFile == "" {
			errs = append(errs, errors.New("PrivateKeyFile is empty"))
		}
		if _, err := tls.LoadX509KeyPair(sc.TLS.PublicKeyFile, sc.TLS.PrivateKeyFile); err != nil {
			errs = append(errs, fmt.Errorf("cannot use TLS config for [%s], %v", sc.ListenInterface, err))
		}
	}
	if len(errs) > 0 {
		return errs
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
