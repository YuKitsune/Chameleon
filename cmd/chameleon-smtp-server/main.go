package main

import (
	"github.com/spf13/cobra"
	"github.com/yukitsune/chameleon/cmd"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/handlers"
	"github.com/yukitsune/chameleon/pkg/smtp"
	"gopkg.in/yaml.v2"
	"net/url"
	"os"
)

var configFile string

func main() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts the SMTP server",
		RunE:  serve,
	}

	serveCmd.Flags().StringVar(&configFile, "config", "", "the path to the config file")

	rootCmd := &cobra.Command{
		Use:   "chameleon-smtp-server <command> [flags]",
		Short: "The Chameleon SMTP server is the entry point for all mail.",
	}

	rootCmd.AddCommand(serveCmd)

	err := rootCmd.Execute()
	if err != nil {
		cmd.ExitFromError(err)
	}
}

func serve(command *cobra.Command, args []string) error {

	config := ChameleonSmtpServerConfig{}
	if configFile != "" {
		data, err := os.ReadFile(configFile)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(data, config)
		if err != nil {
			return err
		}
	} else {
		config.Smtp = &smtp.ServerConfig{
			TLS: &smtp.ServerTLSConfig{},
		}
		config.Logging = &log.LogConfig{}
	}

	err := config.SetDefaults()
	if err != nil {
		return err
	}

	logger := cmd.MakeLogger(config.Logging.Level, config.Logging.Directory)

	apiUrl, err := url.Parse(config.ApiUrl)
	if err != nil {
		return err
	}

	handler := handlers.NewDefaultHandler(apiUrl, logger)

	server, err := smtp.NewServer(config.Smtp, handler, logger)
	if err != nil {
		return err
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
