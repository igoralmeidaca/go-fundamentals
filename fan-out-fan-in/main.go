package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // workload simulation

	in := make(chan int)

	outs := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}

	go produce(numbers, in)

	for i, out := range outs {
		go fanOut(i+1, in, out)
	}

	outsReadOnly := make([]<-chan int, len(outs))
	for i, out := range outs {
		outsReadOnly[i] = out
	}

	outFinal := fanIn(outsReadOnly)
	for result := range outFinal {
		fmt.Printf("Result received: %d\n", result)
	}
}

func produce(numbers []int, in chan<- int) {
	for _, number := range numbers {
		in <- number
	}
	close(in)
}

func fanOut(id int, in <-chan int, out chan<- int) {
	for number := range in {
		fmt.Printf("Channel %d processing number %d\n", id, number)
		out <- number * number
	}
	close(out)
}

func fanIn(outsReadOnly []<-chan int) <-chan int {
	outFinal := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(outsReadOnly))

	for _, out := range outsReadOnly {
		go func() {
			defer wg.Done()

			for number := range out {
				outFinal <- number
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outFinal)
	}()

	return outFinal
}
