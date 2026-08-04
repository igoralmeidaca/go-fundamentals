package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	logf("Starting cache with RWMutex example")

	cache := NewCache()
	var wg sync.WaitGroup

	wg.Add(4)
	cache.data["k"] = "v"

	go func() {
		defer wg.Done()
		logf("Goroutine 1: call Get()")
		_, _ = cache.Get("k")
		logf("Goroutine 1: returned from Get()")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 10)
		logf("Goroutine 2: call Get()")
		_, _ = cache.Get("k")
		logf("Goroutine 2: returned from Get()")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 20)
		logf("Goroutine 3: call Set()")
		cache.Set("k", "v2")
		logf("Goroutine 3: returned from Set()")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 300)
		logf("Goroutine 4: call Get()")
		_, _ = cache.Get("k")
		logf("Goroutine 4: returned from Get()")
	}()

	wg.Wait()
	logf("All goroutines completed")
}

type Cache struct {
	data map[string]any
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	logf(fmt.Sprintf("Set: Trying to set key: %s, value: %v", key, value))
	c.mu.Lock()
	logf("Set: Lock acquired for writing")
	time.Sleep(time.Millisecond * 300)
	c.data[key] = value
	c.mu.Unlock()
	logf("Set: Lock released for writing")
}

func (c *Cache) Get(key string) (any, bool) {
	logf(fmt.Sprintf("Get: Trying to read key: %s", key))
	c.mu.RLock() // block only readers, allowing multiple readers to access the data concurrently
	logf("Get: Lock acquired for reading")
	time.Sleep(time.Millisecond * 100)
	value, ok := c.data[key]
	c.mu.RUnlock()
	logf("Get: Lock released for reading")
	return value, ok
}

var start = time.Now()

func logf(msg string) {
	fmt.Printf("[%4dms] %s\n", time.Since(start).Milliseconds(), msg)
}
