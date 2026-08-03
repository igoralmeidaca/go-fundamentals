package main

import (
	"fmt"
	"time"
)

func main() {
	stopChannel := make(chan struct{})
	doneChannel := make(chan struct{})

	go schedule(1*time.Second, func() {
		// Your scheduled function logic here
		println("Scheduled function executed")
	}, stopChannel, doneChannel)

	// Simulate some work in the main function
	time.Sleep(5 * time.Second)

	close(stopChannel)

	<-doneChannel
}

func schedule(interval time.Duration, f func(), stopChannel <-chan struct{}, doneChannel chan<- struct{}) {
	ticker := time.NewTicker(interval)
	defer close(doneChannel)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			f()
		case <-stopChannel:
			fmt.Println("Stopping the scheduled function")
			return
		}
	}
}
