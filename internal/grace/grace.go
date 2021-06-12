package grace

import (
	"fmt"
	"github.com/yukitsune/chameleon/internal/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)


type shutdownHook func()
type Grace struct {
	shutdownHooks []shutdownHook
}

func WaitForShutdownSignal(logger log.ChameleonLogger, shutdownHooks ...shutdownHook) {
	waitForSignal(getShutdownSignalChan(), make(chan error, 1), logger)
	for _, hook := range shutdownHooks {
		hook()
	}
}

func WaitForShutdownSignalOrError(errorChan chan error, logger log.ChameleonLogger, shutdownHooks ...shutdownHook) {
	waitForSignal(getShutdownSignalChan(), errorChan, logger)
	for _, hook := range shutdownHooks {
		hook()
	}
}

func waitForSignal(shutdownSignalChan chan os.Signal, errorChan chan error, logger log.ChameleonLogger) {
	select {
	case sig := <-shutdownSignalChan:
		handleShutdownSignal(sig, logger)
		break

	case err := <-errorChan:
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
		logger.Infof("Shutdown completed, exiting")
		return
	} else {
		logger.Infof("Shutdown, unknown signal caught")
		return
	}
}

func handleError(err error, logger log.ChameleonLogger) {
	logger.Fatalf("A fatal error has occurred: %v", err)
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
	if _, printErr := fmt.Fprintf(os.Stderr, format, err); printErr != nil {
		// couldn't print to stderr, just print normally i guess
		fmt.Printf(format, err)
	}

	os.Exit(1)
}