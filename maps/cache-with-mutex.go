package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[string]any
	mu   sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.data[key]
	return value, ok
}

func cacheWithMutex() {
	cache := NewCache()
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
