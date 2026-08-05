package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := range 5 {
		go func() {
			defer wg.Done()
			fmt.Printf("Goroutine %d trying to access database...\n", i)
			_ = GetDatabase()
			fmt.Printf("Goroutine %d connected to database...\n", i)
		}()
	}

	wg.Wait()
	fmt.Println("End")
}

var (
	database *Database
	once     sync.Once
)

type Database struct {
	connected bool
}

func GetDatabase() *Database {
	once.Do(func() {
		fmt.Println("Connecting to the database...")
		time.Sleep(time.Second * 2) // Simulate a delay in connecting to the database
		database = &Database{connected: true}
		fmt.Println("Connected!")
	})
	return database
}
