package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/cmd"
)

func main() {
	logger := logrus.New()
	fmt.Println("Hello, World!")

	cmd.SigHandler(logger)
}