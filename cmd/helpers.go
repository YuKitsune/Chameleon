package cmd

import (
	"fmt"
	"github.com/yukitsune/chameleon/internal/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WaitForShutdownSignal(logger log.ChameleonLogger) {
	WaitForSignal(getShutdownSignalChan(), make(chan error, 1), logger)
}

func WaitForShutdownSignalOrError(errorChan chan error, logger log.ChameleonLogger) {
	WaitForSignal(getShutdownSignalChan(), errorChan, logger)
}

func WaitForSignal(shutdownSignalChan chan os.Signal, errorChan chan error, logger log.ChameleonLogger) {
	select {
	case sig := <- shutdownSignalChan:
		handleShutdownSignal(sig, logger)
		break

	case err := <- errorChan:
		handleError(err, logger)
		break
	}
}

func handleShutdownSignal(sig os.Signal, logger log.ChameleonLogger) {
	if sig == syscall.SIGTERM || sig == syscall.SIGQUIT || sig == syscall.SIGINT || sig == os.Kill {
		logger.Infof("Shutdown signal caught")
		go func() {
			select {
			// exit if graceful shutdown not finished in 60 sec.
			case <-time.After(time.Second * 60):
				logger.Error("graceful shutdown timed out")
				os.Exit(1)
			}
		}()
		logger.Infof("Shutdown completed, exiting.")
		return
	} else {
		logger.Infof("Shutdown, unknown signal caught")
		return
	}
}

func handleError(err error, logger log.ChameleonLogger) {
	logger.Errorf("A fatal error has occurred: %v", err)
}

func getShutdownSignalChan() chan os.Signal {
	shutdownSignalChan := make(chan os.Signal, 1)
	signal.Notify(shutdownSignalChan,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGKILL,
		os.Kill,
	)

	return shutdownSignalChan
}

func ExitFromError(err error) {
	format := "error: %v\n"
	if _, err := fmt.Fprintf(os.Stderr, format, err); err != nil {
		fmt.Printf(format, err)
	}

	os.Exit(1)
}

func MakeLogger(logLevel string, logDir string) log.ChameleonLogger {

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		ExitFromError(err)
	}

	logFactory := log.NewLogFactory(
		logDir,
		level,
		log.DefaultFileNameProvider,
		log.LogrusLoggerProvider,
	)

	logger, err := logFactory.Make()
	if err != nil {
		ExitFromError(err)
	}

	return logger
}

func GetValidLogLevels() []string {
	logLevels := log.AllLevels
	var logLevelStrings []string
	for _, level := range logLevels {
		logLevelStrings = append(logLevelStrings, level.String())
	}

	return logLevelStrings
}
