package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitsune/chameleon/cmd"
	"github.com/yukitsune/chameleon/internal/log"
	"strings"
)

var (
	logLevel string
	logDir string
)

func main() {

	var port int
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

	rootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "info", fmt.Sprintf("the logging level, options are %s", strings.Join(getValidLogLevels(), ", ")))
	rootCmd.PersistentFlags().StringVar(&logDir, "logDir", "./logs", "the log directory")

	rootCmd.AddCommand(serveCmd)
	if err := rootCmd.Execute(); err != nil {
		cmd.ExitFromError(err)
	}
}

func serve(command *cobra.Command, args []string) error {

	logger := makeLogger()

	cmd.SigHandler(logger)
	return nil
}

func makeLogger() log.ChameleonLogger {

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		cmd.ExitFromError(err)
	}

	logFactory := log.NewLogFactory(
		logDir,
		level,
		log.DefaultFileNameProvider,
		log.LogrusLoggerProvider,
	)

	logger, err := logFactory.Make()
	if err != nil {
		cmd.ExitFromError(err)
	}

	return logger
}

func getValidLogLevels() []string {
	logLevels := log.AllLevels
	var logLevelStrings []string
	for _, level := range logLevels {
		logLevelStrings = append(logLevelStrings, level.String())
	}

	return logLevelStrings
}