package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)

	go stage1(ch1, numbers)
	go stage2(ch1, ch2)
	stage3(ch2)
}

func stage1(ch1 chan<- int, numbers []int) {
	for _, number := range numbers {
		ch1 <- number
	}
	close(ch1)
}

func stage2(ch1 <-chan int, ch2 chan<- int) {
	for number := range ch1 {
		ch2 <- number * number
	}
	close(ch2)
}

func stage3(ch2 <-chan int) {
	for number := range ch2 {
		fmt.Println("Value received at stage 3: ", number)
	}
}
