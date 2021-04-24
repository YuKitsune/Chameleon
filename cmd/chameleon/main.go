package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	fmt.Println("Hello, World!")

	WaitForCtrlC()
}


func WaitForCtrlC() {
	var endWaiter sync.WaitGroup
	endWaiter.Add(1)

	var signalChannel chan os.Signal
	signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		endWaiter.Done()
	}()
	endWaiter.Wait()
}