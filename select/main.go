package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendMessage(ch1, "Sending message to channel 1", time.Second*3)
	go sendMessage(ch2, "Sending message to channel 2", time.Second*4)

	for range 10 {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(time.Second * 2):
			fmt.Println("No message received")
		}
	}
}

func sendMessage(ch chan<- string, mensagem string, interval time.Duration) {
	for {
		time.Sleep(interval)
		ch <- mensagem
	}
}
