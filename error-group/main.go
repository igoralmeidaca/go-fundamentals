package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("-----WaitGroup-----")

	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("Done wg")

	fmt.Println("-----ErrorGroup-----")

	var eg errgroup.Group

	eg.SetLimit(2)

	for i := range 3 {
		eg.Go(func() error {
			return workerErrorGroup(i)
		})
	}

	err := eg.Wait()
	if err != nil {
		fmt.Println("Error eg:", err.Error()) // only the first error is returned
	} else {
		fmt.Println("Done eg")
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d initialized\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func workerErrorGroup(id int) error {
	fmt.Printf("Worker %d initialized\n", id)

	if id == 2 {
		return errors.New("unexpected error")
	}

	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)

	return nil
}
