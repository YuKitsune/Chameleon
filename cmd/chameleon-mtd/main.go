package main

import (
	"github.com/spf13/cobra"
	"github.com/yukitsune/chameleon/internal/config"
	"github.com/yukitsune/chameleon/internal/grace"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/handlers"
	"github.com/yukitsune/chameleon/pkg/smtp"
	"net/url"
)

type ChameleonMtdConfig struct {
	ApiUrl  string       `mapstructure:"api-url"`
	Smtp    *smtp.Config `mapstructure:"smtp"`
	Logging *log.Config  `mapstructure:"log"`
}

func (c *ChameleonMtdConfig) SetDefaults() error {
	if c.Smtp == nil {
		c.Smtp = &smtp.Config{}
	}
	err := c.Smtp.SetDefaults()
	if err != nil {
		return err
	}

	if c.Logging == nil {
		c.Logging = &log.Config{}
	}
	err = c.Logging.SetDefaults()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts the mail transfer daemon",
		RunE:  serve,
	}

	// Automatically setup the command-line flags based on our config struct
	err := config.SetupFlagsForConfig(serveCmd, &ChameleonMtdConfig{})
	if err != nil {
		grace.ExitFromError(err)
	}

	rootCmd := &cobra.Command{
		Use:   "chameleon-mtd <command> [flags]",
		Short: "The Chameleon Mail Transfer Daemon",
	}

	rootCmd.AddCommand(serveCmd)

	err = rootCmd.Execute()
	if err != nil {
		grace.ExitFromError(err)
	}
}

func serve(command *cobra.Command, args []string) error {

	// Load the config
	mtdConfig := &ChameleonMtdConfig{}
	err := config.LoadConfig("mtd", mtdConfig)
	if err != nil {
		return err
	}

	// Setup the logger
	logger, err := log.NewLogrusLogger(mtdConfig.Logging)
	if err != nil {
		return err
	}

	// Setup the handler
	apiUrl, err := url.Parse(mtdConfig.ApiUrl)
	if err != nil {
		return err
	}

	handler := handlers.NewDefaultHandler(apiUrl, logger)

	// Setup the server
	server, err := smtp.NewServer(mtdConfig.Smtp, handler, logger)
	if err != nil {
		return err
	}

	// Start the server and wait for any errors
	errorChan := make(chan error, 1)
	go func() {
		if err = server.Start(); err != nil {
			errorChan <- err
		}
	}()

	// server.Start doesn't block, wait for exit signal or error
	grace.WaitForShutdownSignalOrError(errorChan, server.Shutdown)

	return nil
}
