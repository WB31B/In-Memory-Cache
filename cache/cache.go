package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	cache map[string]interface{}
}

func New() Cache {
	return Cache{
		cache: make(map[string]interface{}),
	}
}

func (c Cache) Get(key string) (interface{}, error) {
	for item := range c.cache {
		if key == item {
			return c.cache[key], nil
		}
	}

	return "This key is not found", fmt.Errorf("ERROR")
}

func (c Cache) Set(key string, value interface{}, ttl time.Duration) {
	// ttl time.Duration - clear cash
	fmt.Printf("time.Direction: %d : %s\n", ttl, key) // 5s
	c.cache[key] = value // ... = ...

	go c.Delete(key, ttl)
}

func (c Cache) Delete(key string, ttl time.Duration) {
	time.Sleep(ttl)
	delete(c.cache, key)
}