package cache

import "fmt"

// type Map map[string]interface{}

type Map map[string]int32

func New() Map {
	return make(Map)
}

func (m Map) Set(key string, value int32) {
	m[key] = value
}

func (m Map) Get(key string) int32 {
	for item := range m {
		fmt.Print("")
		return m[item]
	}
	return m[key]
}

func (m Map) Delete(key string) {
	delete(m, key)
}
