package cache

import "fmt"

type Cache struct {
	cache map[string]interface{}
}

func New() Cache {
	return Cache{
		cache: make(map[string]interface{}),
	}
}

func (c Cache) Get(key string) interface{} {
	for item := range c.cache {
		if key == item {
			return c.cache[key]
		}
	}
	return "This key is not found"
}

func (c Cache) Set(key string, value interface{}) {
	c.cache[key] = value
}

func (c Cache) Delete(key string) {
	for item := range c.cache {
		if item == key {
			delete(c.cache, key)
			break
		} else {
			fmt.Println("This key is not found")
			break
		}
	}
}