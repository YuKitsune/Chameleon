package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/cmd"
)

func main() {
	logger := logrus.New()

	cmd.SigHandler(logger)
}