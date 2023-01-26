package cache

type Cache struct {
	cache map[string]interface{}
}

func New() Cache {
	return Cache{
		cache: make(map[string]interface{}),
	}
}

func (c Cache) Get(key string) interface{} {
	return c.cache[key]
}

func (c Cache) Set(key string, value interface{}) {
	c.cache[key] = value
}

func (c Cache) Delete(key string) {
	delete(c.cache, key)
}