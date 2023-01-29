package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]interface{}
	mu *sync.Mutex
}

func New() Cache {
	return Cache{
		cache: make(map[string]interface{}),
		mu: new(sync.Mutex),
	}
}

func (c Cache) Get(key string) interface{} {
	c.mu.Lock()
	for item := range c.cache {
		if key == item {
			defer c.mu.Unlock()
			return c.cache[key]
		}
	}

	defer c.mu.Unlock()
	return "This key is not found"
}

func (c Cache) Set(key string, value interface{}, ttl time.Duration) {
	// ttl time.Duration - clear cash
	fmt.Printf("time.Direction: %d : %s\n", ttl, key) // 5s
	c.cache[key] = value // ... = ...

	go func ()  {
		time.Sleep(ttl)
		c.mu.Lock()
		delete(c.cache, key)
		c.mu.Unlock()
	}()
}