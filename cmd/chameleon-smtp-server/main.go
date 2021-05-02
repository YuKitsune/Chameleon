package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yukitsune/chameleon/cmd"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/handlers"
	"github.com/yukitsune/chameleon/pkg/smtp"
	"net/url"
	"strings"
)

var (
	logLevel string
	logDir string
)

const (
	configKey = "config"

	hostnameKey = "host"
	listenInterfaceKey = "listen-interface"
	maxMailSizeKey = "max-mail-size"
	timeoutKey  = "timeout"
	maxClientsKey = "max-clients"
	xClientOnKey = "xclient-on"

	startTLSOnKey = "start-tls-on"
	tlsAlwaysOnKey = "tls-always-on"
	privateKeyFileKey = "private-key-file"
	publicKeyFileKey = "public-key-file-key"
)

func main() {
	serveCmd := &cobra.Command{
		Use: "serve",
		Short: "Starts the SMTP server",
		RunE: serve,
	}

	// Todo: Make a nicer function for this
	// 	Maybe move to a separate abstraction

	serveCmd.Flags().String(configKey, "", "the config file for the SMTP server. If provided, then all other flags will be ignored")
	_ = viper.BindPFlag(configKey, serveCmd.Flags().Lookup(configKey))
	viper.SetDefault(configKey, nil)

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

	rootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "info", fmt.Sprintf("the logging level, options are %s", strings.Join(cmd.GetValidLogLevels(), ", ")))
	rootCmd.PersistentFlags().StringVar(&logDir, "logDir", "./logs", "the log directory")

	rootCmd.AddCommand(serveCmd)
	if err := rootCmd.Execute(); err != nil {
		cmd.ExitFromError(err)
	}
}

func serve(command *cobra.Command, args []string) error {

	logger := cmd.MakeLogger(logLevel, logDir)

	config, err := getConfig()
	if err != nil {
		logger.Fatal(err)
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
	cmd.WaitForShutdownSignal(logger)
	server.Shutdown()

	return nil
}

func getConfig() (*smtp.ServerConfig, error) {

	// If the --config flag was provided, then pull the config from there
	configValue := viper.GetString(configKey)
	if configValue != "" {

		viper.SetConfigFile(configValue)

		config := &smtp.ServerConfig{}
		err := viper.Unmarshal(config)
		if err != nil {
			return nil, err
		}

		return config, nil
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

	config := &smtp.ServerConfig{
		TLS:             tlsConfig,
		Hostname:        hostName,
		ListenInterface: listenInterface,
		MaxSize:         maxSize,
		Timeout:         timeout,
		MaxClients:      maxClients,
		XClientOn:       xClientOn,
	}

	return config, nil
}