package main

import (
	"github.com/spf13/cobra"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api"
	"github.com/yukitsune/chameleon/internal/config"
	"github.com/yukitsune/chameleon/internal/grace"
	"github.com/yukitsune/chameleon/internal/log"
)

type ChameleonApiConfig struct {
	Api     *api.Config `mapstructure:"api"`
	Logging *log.Config `mapstructure:"log"`
}

func (c *ChameleonApiConfig) SetDefaults() error {
	if c.Api == nil {
		c.Api = &api.Config{}
	}
	err := c.Api.SetDefaults()
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

	return container.Resolve(func(svr *api.ChameleonApiServer, logger log.ChameleonLogger) {

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

func setupContainer(cfg *ChameleonApiConfig) (camogo.Container, error) {
	c := camogo.New()
	err := c.Register(func (r *camogo.Registrar) error {

		// Todo: Move to module

		// Configuration
		err := r.RegisterInstance(cfg.Api)
		if err != nil {
			return err
		}

		err = r.RegisterInstance(cfg.Logging)
		if err != nil {
			return err
		}

		// Services
		err = r.RegisterFactory(log.New, camogo.SingletonLifetime)
		if err != nil {
			return err
		}

		err = r.RegisterFactory(api.NewChameleonApiServer, camogo.SingletonLifetime)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}
