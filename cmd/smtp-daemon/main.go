package main

import (
	"fmt"
	"github.com/yukitsune/chameleon/cmd"
	"github.com/yukitsune/chameleon/internal/log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		cmd.ExitFromError(err)
	}

	logFactory := log.NewLogFactory(
		fmt.Sprintf("%s/logs", wd),
		log.DebugLevel,
		log.DefaultFileNameProvider,
		log.LogrusLoggerProvider,
	)

	logger, err := logFactory.Make()
	if err != nil {
		cmd.ExitFromError(err)
	}

	cmd.SigHandler(logger)
}
