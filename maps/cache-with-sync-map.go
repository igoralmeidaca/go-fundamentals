package main

import (
	"fmt"
	"sync"
)

type Cache2 struct {
	data sync.Map
}

func NewCache2() *Cache2 {
	return &Cache2{}
}

func (c *Cache2) Set(key string, value any) {
	c.data.Store(key, value)
}

func (c *Cache2) Get(key string) (any, bool) {
	return c.data.Load(key)
}

func cacheWithSyncMap() {
	cache := NewCache2()
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			cache.Set(fmt.Sprintf("key_%d", i), i)
		})
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value, ok := cache.Get(fmt.Sprintf("key_%d", i))
			if ok {
				fmt.Printf("Value found for key_%d: %v\n", i, value)
			} else {
				fmt.Printf("Value not found for key_%d\n", i)
			}
		}()
	}

	wg.Wait()
}
