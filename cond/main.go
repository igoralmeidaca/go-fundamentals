package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	buffer := &Buffer{
		data: make([]int, 0),
	}
	buffer.cond = sync.NewCond(&buffer.mu)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range 5 {
			buffer.mu.Lock()
			fmt.Printf("Producing item %d\n", i)
			buffer.data = append(buffer.data, i)
			buffer.cond.Signal() // Notify a waiting consumer
			buffer.mu.Unlock()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		defer wg.Done()
		for range 5 {
			buffer.mu.Lock()
			for len(buffer.data) == 0 {
				buffer.cond.Wait() // Wait for a signal from the producer
			}
			item := buffer.data[0]
			buffer.data = buffer.data[1:] // create a new slice without the consumed item
			fmt.Printf("Consuming item %d\n", item)
			buffer.mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("End")
}

type Buffer struct {
	data []int
	cond *sync.Cond
	mu   sync.Mutex
}
