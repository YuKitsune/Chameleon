package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitsune/chameleon/cmd"
	"strings"
)

var (
	logLevel string
	logDir string

	port int
)

func main() {

	serveCmd := &cobra.Command{
		Use: "serve",
		Short: "Starts the REST API.",
		RunE: serve,
	}

	serveCmd.Flags().IntVar(&port, "port", 80, "the port number to listen for requests on")

	rootCmd := &cobra.Command{
		Use: "chameleon-api-server <command> [flags]",
		Short: "The Chameleon API is the REST API that powers Chameleon.",
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

	// Todo: Start HTTP server
	// 	Wait for either an error from ListenAndServe or OS signal

	cmd.WaitForShutdownSignal(logger)
	return nil
}
