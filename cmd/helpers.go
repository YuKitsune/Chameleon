package cmd

import (
	"fmt"
	"github.com/yukitsune/chameleon/internal/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WaitForShutdownSignal(log log.ChameleonLogger) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGKILL,
		os.Kill,
	)

	for sig := range signalChannel {
		if sig == syscall.SIGTERM || sig == syscall.SIGQUIT || sig == syscall.SIGINT || sig == os.Kill {
			log.Infof("Shutdown signal caught")
			go func() {
				select {
				// exit if graceful shutdown not finished in 60 sec.
				case <-time.After(time.Second * 60):
					log.Error("graceful shutdown timed out")
					os.Exit(1)
				}
			}()
			log.Infof("Shutdown completed, exiting.")
			return
		} else {
			log.Infof("Shutdown, unknown signal caught")
			return
		}
	}

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
