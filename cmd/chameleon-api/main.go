package main

import (
	"github.com/spf13/cobra"
	"github.com/yukitsune/chameleon/cmd"
	"github.com/yukitsune/chameleon/internal/api"
	"github.com/yukitsune/chameleon/internal/log"
	"gopkg.in/yaml.v2"
	"os"
)

var configFile string

type ChameleonApiConfig struct {
	Api *api.ApiConfig `yaml:"api"`
	Logging *log.LogConfig     `yaml:"log"`
}

func (c *ChameleonApiConfig) SetDefaults() error {
	err := c.Api.SetDefaults()
	if err != nil {
		return err
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
		Short: "Starts the REST API",
		RunE:  serve,
	}

	serveCmd.Flags().StringVar(&configFile, "config", "", "the path to the config file")

	rootCmd := &cobra.Command{
		Use:   "chameleon-api <command> [flags]",
		Short: "The Chameleon REST API",
	}

	rootCmd.AddCommand(serveCmd)
	if err := rootCmd.Execute(); err != nil {
		cmd.ExitFromError(err)
	}
}

func serve(command *cobra.Command, args []string) error {

	config := ChameleonApiConfig{}
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
		config.Api = &api.ApiConfig{}
		config.Logging = &log.LogConfig{}
	}

	err := config.SetDefaults()
	if err != nil {
		return err
	}

	logger := cmd.MakeLogger(config.Logging.Level, config.Logging.Directory)

	// Run our server in a goroutine so that it doesn't block.
	apiHandler := api.NewChameleonApiServer(config.Api)
	errorChan := make(chan error, 1)
	go func() {

		// Todo: TLS
		if err = apiHandler.Start(); err != nil {
			errorChan <- err
		}
	}()

	cmd.WaitForShutdownSignalOrError(errorChan, logger)
	_ = apiHandler.Shutdown()

	return nil
}
