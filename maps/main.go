package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// basic()
	// raceCondition()
	// mutex()
	// syncMap()
	// cacheWithMutex()
	cacheWithSyncMap()
}

func syncMap() {
	var wg sync.WaitGroup
	var m sync.Map

	wg.Add(100)

	for i := range 100 {
		go func() {
			defer wg.Done()
			m.Store(i, i)
		}()
	}

	wg.Wait()

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

func mutex() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	m := make(map[int]int)

	wg.Add(100)

	for i := range 100 {
		go func() {
			defer wg.Done()

			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}

	wg.Wait()
}

func raceCondition() {
	m := make(map[int]int)

	for i := range 1000 {
		m[i] = i
	}

	m2 := make(map[int]int)

	go func() {
		for i := range 1000 {
			m2[i] = i
		}
	}()

	// doesn't work, race condition
	// go func() {
	// 	for i := 1000; i < 2000; i++ {
	// 		m2[i] = i
	// 	}
	// }()

	time.Sleep(time.Second * 5)
}

func basic() {
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	fmt.Println("----------------")

	myMap2 := map[string]any{}
	myMap2["key"] = "value"

	myValue, ok := myMap2["key"].(string)
	if ok {
		fmt.Println(myValue)
	}
}
