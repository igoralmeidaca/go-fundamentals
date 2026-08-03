package main

import (
	"fmt"
	"time"
)

func main() {
	channels := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}

	for i, ch := range channels {
		go consumer(i+1, ch)
	}

	broadcast(channels, 100)
	time.Sleep(1 * time.Second) // Allow some time for consumers to process the value
}

func consumer(id int, ch <-chan int) {
	for value := range ch {
		fmt.Printf("Consumer %d received value: %d\n", id, value)
	}
}

func broadcast(channels []chan int, value int) {
	for _, ch := range channels {
		ch <- value
	}
}
