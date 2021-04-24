package cmd

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SigHandler(log *logrus.Logger) {
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