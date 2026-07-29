package main

import (
	"fmt"
	"time"
)

func main() {
	// channelWithoutBuffer()
	channelWithBuffer()
}

func channelWithBuffer() {
	// communication betwenn gorountines
	ch := make(chan int, 3) // buffer

	for range 3 {
		go send(ch)
	}

	time.Sleep(time.Second)

	for range 3 {
		go receive(ch)
	}

	time.Sleep(time.Second * 5)
}

func channelWithoutBuffer() {
	// communication betwenn gorountines
	ch := make(chan int)

	for range 3 {
		go send(ch)
	}

	time.Sleep(time.Second)

	for range 3 {
		go receive(ch)
	}

	time.Sleep(time.Second * 5)
}

func send(ch chan<- int) {
	fmt.Println("Sending value to channel")
	ch <- 100
	fmt.Println("Value successfully sent")
}

func receive(ch <-chan int) {
	time.Sleep(time.Second)
	value := <-ch
	fmt.Println("Value received:", value)
}
