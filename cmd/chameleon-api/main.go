package main

import (
	"github.com/spf13/cobra"
	"github.com/yukitsune/chameleon/internal/api"
	"github.com/yukitsune/chameleon/internal/config"
	"github.com/yukitsune/chameleon/internal/grace"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/ioc"
)

func main() {

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts the REST API",
		RunE:  serve,
	}

	err := config.SetupFlagsForConfig(serveCmd, &ChameleonApiConfig{})
	if err != nil {
		grace.ExitFromError(err)
	}

	// Automatically setup the command-line flags based on our config struct
	rootCmd := &cobra.Command{
		Use:   "chameleon-api <command> [flags]",
		Short: "The Chameleon REST API",
	}

	rootCmd.AddCommand(serveCmd)

	err = rootCmd.Execute()
	if err != nil {
		grace.ExitFromError(err)
	}
}

func serve(command *cobra.Command, args []string) error {

	// Load the config
	apiConfig := &ChameleonApiConfig{}
	err := config.LoadConfig("api", apiConfig)
	if err != nil {
		return err
	}

	// Setup the IoC container
	container, err := setupContainer(apiConfig)
	if err != nil {
		return err
	}

	return container.ResolveInScope(func(svr *api.ChameleonApiServer, logger log.ChameleonLogger) {

		// Run our server in a goroutine so that it doesn't block.
		errorChan := make(chan error, 1)
		go func() {

			// Todo: TLS
			if err = svr.Start(); err != nil {
				errorChan <- err
			}
		}()

		grace.WaitForShutdownSignalOrError(errorChan, logger, func() { _ = svr.Shutdown() })
	})
}

func setupContainer(cfg *ChameleonApiConfig) (ioc.Container, error) {
	c := ioc.NewGolobbyContainer()
	var err error

	// Configuration
	err = c.RegisterSingletonInstance(cfg.Logging)
	if err != nil {
		return nil, err
	}

	err = c.RegisterSingletonInstance(cfg.Api)
	if err != nil {
		return nil, err
	}

	// Services
	err = c.RegisterTransientFactory(log.New)
	if err != nil {
		return nil, err
	}

	err = c.RegisterSingletonFactory(api.NewChameleonApiServer)
	if err != nil {
		return nil, err
	}

	return c, nil
}
