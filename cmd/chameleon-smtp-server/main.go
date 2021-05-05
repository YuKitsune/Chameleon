package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yukitsune/chameleon/cmd"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/handlers"
	"github.com/yukitsune/chameleon/pkg/smtp"
	"io/ioutil"
	"net/url"
	"strings"
)

const (
	configKey = "config"

	apiBaseUrlKey = "chameleon-api-base-url"
	allowedHostsKey = "allowed-hosts"

	hostnameKey = "smtp-hostname"
	listenInterfaceKey = "smtp-listen-interface"
	maxMailSizeKey = "smtp-max-mail-size"
	timeoutKey  = "smtp-timeout"
	maxClientsKey = "smtp-max-clients"
	xClientOnKey = "smtp-xclient-on"

	startTLSOnKey = "smtp-tls-start-tls-on"
	tlsAlwaysOnKey = "smtp-tls-always-on"
	privateKeyFileKey = "smtp-tls-private-key-file"
	publicKeyFileKey = "smtp-tls-public-key-file-key"
)

func main() {
	serveCmd := &cobra.Command{
		Use: "serve",
		Short: "Starts the SMTP server",
		RunE: serve,
	}

	// Todo: Make a nicer function for this
	// 	Maybe move to a separate abstraction
	serveCmd.Flags().String(configKey, "", "the path to the config file. If provided, then all other flags will be ignored")
	_ = viper.BindPFlag(configKey, serveCmd.Flags().Lookup(configKey))

	serveCmd.Flags().String(apiBaseUrlKey, "", "the base URL for the chameleon API server (E.g. https://api.chameleon.io)")
	_ = viper.BindPFlag(apiBaseUrlKey, serveCmd.Flags().Lookup(apiBaseUrlKey))

	// Todo: Need a better explanation
	serveCmd.Flags().String(allowedHostsKey, "relay.chameleon.io", "the allowed host")
	_ = viper.BindPFlag(allowedHostsKey, serveCmd.Flags().Lookup(allowedHostsKey))

	serveCmd.Flags().String(hostnameKey, "", "the DNS name of the SMTP server")
	_ = viper.BindPFlag(hostnameKey, serveCmd.Flags().Lookup(hostnameKey))

	serveCmd.Flags().String(listenInterfaceKey, "127.0.0.1:2525", "Listen interface specified in <ip>:<port>")
	_ = viper.BindPFlag(listenInterfaceKey, serveCmd.Flags().Lookup(listenInterfaceKey))

	serveCmd.Flags().Int(maxMailSizeKey, 50000000, "the maximum size (in bytes) that an email can be")
	_ = viper.BindPFlag(maxMailSizeKey, serveCmd.Flags().Lookup(maxMailSizeKey))

	// Todo: Need a better explanation
	serveCmd.Flags().Int64(timeoutKey, 30, "the amount of seconds before timing out the client connection")
	_ = viper.BindPFlag(timeoutKey, serveCmd.Flags().Lookup(timeoutKey))

	serveCmd.Flags().Int(maxClientsKey, 1000, "the maximum amount of clients that can be connected at once")
	_ = viper.BindPFlag(maxClientsKey, serveCmd.Flags().Lookup(maxClientsKey))

	serveCmd.Flags().Bool(xClientOnKey, false, "when using a proxy such as Nginx, XCLIENT command is used to pass the original client's IP address & client's HELO")
	_ = viper.BindPFlag(xClientOnKey, serveCmd.Flags().Lookup(xClientOnKey))

	rootCmd := &cobra.Command{
		Use: "chameleon-smtp-server <command> [flags]",
		Short: "The Chameleon SMTP server is the entry point for all mail.",
	}

	rootCmd.PersistentFlags().String(log.LogLevelKey, "info", fmt.Sprintf("the logging level, options are %s", strings.Join(cmd.GetValidLogLevels(), ", ")))
	_ = viper.BindPFlag(log.LogLevelKey, rootCmd.Flags().Lookup(log.LogLevelKey))

	rootCmd.PersistentFlags().String(log.LogDirKey, "./logs", "the log directory")
	_ = viper.BindPFlag(log.LogDirKey, rootCmd.Flags().Lookup(log.LogDirKey))

	rootCmd.AddCommand(serveCmd)
	if err := rootCmd.Execute(); err != nil {
		cmd.ExitFromError(err)
	}
}

func serve(command *cobra.Command, args []string) error {

	logger := cmd.MakeLogger(viper.GetString(log.LogLevelKey), viper.GetString(log.LogDirKey))

	chameleonConfig, err := getConfig()
	if err != nil {
		return err
	}

	apiUrl, err := url.Parse(chameleonConfig.ApiUrl)
	if err != nil {
		return err
	}

	handler := handlers.NewDefaultHandler(apiUrl, logger)

	server, err := smtp.NewServer(chameleonConfig.SmtpConfig, handler, logger)
	if err != nil {
		return err
	}

	server.SetAllowedHosts(chameleonConfig.AllowedHosts)

	go func(logger log.ChameleonLogger) {
		if err := server.Start(); err != nil {
			logger.Fatal(err)
		}
	}(logger)

	// server.Start doesn't block, wait for exit signal
	cmd.WaitForShutdownSignal(logger)
	server.Shutdown()

	return nil
}

var configFile string

func initConfig() {

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("chameleond")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/chameleond")
		viper.AddConfigPath(".")

		viper.SetEnvPrefix("chameleon")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		viper.AutomaticEnv()
	}

	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			// Config file not found, try to build the config based on flags and environment variables
		} else {
			// Config file was found, but some error occurred
		}
	}
}

func readConfigFromFlags() {

}

func getConfig() (*ChameleonSmtpServerConfig, error) {

	// If the --config flag was provided, then pull the config from there
	configFile := viper.GetString(configKey)
	if configFile != "" {

		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, err
		}

		chameleonConfig := &ChameleonSmtpServerConfig{}
		err = json.Unmarshal(data, chameleonConfig)
		if err != nil {
			return nil, err
		}

		return chameleonConfig, nil
	}

	// --config flag not provided,
	// gonna have to fill it out manually

	startTlSOn := viper.GetBool(startTLSOnKey)
	tlsAlwaysOn := viper.GetBool(tlsAlwaysOnKey)
	privateKeyFile := viper.GetString(privateKeyFileKey)
	publicKeyFile := viper.GetString(publicKeyFileKey)

	// Todo: Add these as flags
	rootCAs := ""
	protocols := []string{"tls1.0", "tls1.2"}
	ciphers := []string{"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305", "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305", "TLS_RSA_WITH_RC4_128_SHA", "TLS_RSA_WITH_AES_128_GCM_SHA256", "TLS_RSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA", "TLS_ECDHE_RSA_WITH_RC4_128_SHA", "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"}
	curves := []string{"P256", "P384", "P521", "X25519"}
	clientAuthType := ""

	tlsConfig := smtp.ServerTLSConfig{
		StartTLSOn:               startTlSOn,
		AlwaysOn:                 tlsAlwaysOn,
		PrivateKeyFile:           privateKeyFile,
		PublicKeyFile:            publicKeyFile,
		Protocols:                protocols,
		Ciphers:                  ciphers,
		Curves:                   curves,
		RootCAs:                  rootCAs,
		ClientAuthType:           clientAuthType,
		PreferServerCipherSuites: false,
	}

	hostName := viper.GetString(hostnameKey)
	listenInterface := viper.GetString(listenInterfaceKey)
	maxSize := viper.GetInt64(maxMailSizeKey)
	timeout := viper.GetInt(timeoutKey) // Todo: Ensure this is read as seconds
	maxClients := viper.GetInt(maxClientsKey)
	xClientOn := viper.GetBool(xClientOnKey)

	smtpConfig := smtp.ServerConfig{
		TLS:             tlsConfig,
		Hostname:        hostName,
		ListenInterface: listenInterface,
		MaxSize:         maxSize,
		Timeout:         timeout,
		MaxClients:      maxClients,
		XClientOn:       xClientOn,
	}
	
	chameleonConfig := &ChameleonSmtpServerConfig{
		ApiUrl:     viper.GetString(apiBaseUrlKey),
		AllowedHosts: viper.GetStringSlice(allowedHostsKey),
		SmtpConfig: smtpConfig,
	}

	return chameleonConfig, nil
}