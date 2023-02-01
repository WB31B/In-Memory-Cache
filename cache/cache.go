package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]interface{}
	mu    sync.Mutex
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, err := c.cache[key]
	if !err {
		return nil, errors.New("Unknown key")
	}

	return value, nil
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.cache[key] = value
	c.mu.Unlock()

	go func() {
		time.Sleep(ttl)
		_ = c.Delete(key)
	}()
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, err := c.cache[key]; !err {
		return errors.New("unknown key")
	}

	delete(c.cache, key)
	return nil
}
