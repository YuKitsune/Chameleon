package smtp

import "crypto/tls"

// https://golang.org/pkg/crypto/tls/#pkg-constants
// Ciphers introduced before Go 1.7 are listed here,
// ciphers since Go 1.8, see tls_go1.8.go
// ....... since Go 1.13, see tls_go1.13.go
var TLSCiphers = map[string]uint16{

	// Note: Generally avoid using CBC unless for compatibility
	// The following ciphersuites are not configurable for TLS 1.3
	// see tls_go1.13.go for a list of ciphersuites always used in TLS 1.3

	"TLS_RSA_WITH_3DES_EDE_CBC_SHA":        tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
	"TLS_RSA_WITH_AES_128_CBC_SHA":         tls.TLS_RSA_WITH_AES_128_CBC_SHA,
	"TLS_RSA_WITH_AES_256_CBC_SHA":         tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA": tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA": tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA":  tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":   tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":   tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,

	"TLS_RSA_WITH_RC4_128_SHA":        tls.TLS_RSA_WITH_RC4_128_SHA,
	"TLS_RSA_WITH_AES_128_GCM_SHA256": tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
	"TLS_RSA_WITH_AES_256_GCM_SHA384": tls.TLS_RSA_WITH_AES_256_GCM_SHA384,

	"TLS_ECDHE_ECDSA_WITH_RC4_128_SHA":        tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
	"TLS_ECDHE_RSA_WITH_RC4_128_SHA":          tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
	"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256": tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384":   tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384": tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,

	// see tls_go1.13 for new TLS 1.3 ciphersuites
	// Note that TLS 1.3 ciphersuites are not configurable
}

// https://golang.org/pkg/crypto/tls/#pkg-constants
var TLSProtocols = map[string]uint16{
	"tls1.0": tls.VersionTLS10,
	"tls1.1": tls.VersionTLS11,
	"tls1.2": tls.VersionTLS12,
}

// https://golang.org/pkg/crypto/tls/#CurveID
var TLSCurves = map[string]tls.CurveID{
	"P256": tls.CurveP256,
	"P384": tls.CurveP384,
	"P521": tls.CurveP521,
}

// https://golang.org/pkg/crypto/tls/#ClientAuthType
var TLSClientAuthTypes = map[string]tls.ClientAuthType{
	"NoClientCert":               tls.NoClientCert,
	"RequestClientCert":          tls.RequestClientCert,
	"RequireAnyClientCert":       tls.RequireAnyClientCert,
	"VerifyClientCertIfGiven":    tls.VerifyClientCertIfGiven,
	"RequireAndVerifyClientCert": tls.RequireAndVerifyClientCert,
}

type ServerTLSConfig struct {
	// TLS Protocols to use. [0] = min, [1]max
	// Use Go's default if empty
	Protocols []string `json:"protocols,omitempty"`
	// TLS Ciphers to use.
	// Use Go's default if empty
	Ciphers []string `json:"ciphers,omitempty"`
	// TLS Curves to use.
	// Use Go's default if empty
	Curves []string `json:"curves,omitempty"`
	// PrivateKeyFile path to cert private key in PEM format.
	PrivateKeyFile string `json:"private_key_file"`
	// PublicKeyFile path to cert (public key) chain in PEM format.
	PublicKeyFile string `json:"public_key_file"`
	// TLS Root cert authorities to use. "A PEM encoded CA's certificate file.
	// Defaults to system's root CA file if empty
	RootCAs string `json:"root_cas_file,omitempty"`
	// declares the policy the server will follow for TLS Client Authentication.
	// Use Go's default if empty
	ClientAuthType string `json:"client_auth_type,omitempty"`
	// The following used to watch certificate changes so that the TLS can be reloaded
	_privateKeyFileMtime int64
	_publicKeyFileMtime  int64
	// controls whether the server selects the
	// client's most preferred cipher suite
	PreferServerCipherSuites bool `json:"prefer_server_cipher_suites,omitempty"`
	// StartTLSOn should we offer STARTTLS command. Cert must be valid.
	// False by default
	StartTLSOn bool `json:"start_tls_on,omitempty"`
	// AlwaysOn run this server as a pure TLS server, i.e. SMTPS
	AlwaysOn bool `json:"tls_always_on,omitempty"`
}

// Gets the timestamp of the TLS certificates. Returns a unix time of when they were last modified
// when the config was read. We use this info to determine if TLS needs to be re-loaded.
func (stc *ServerTLSConfig) getTlsKeyTimestamps() (int64, int64) {
	return stc._privateKeyFileMtime, stc._publicKeyFileMtime
}
