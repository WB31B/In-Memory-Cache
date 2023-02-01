package main

import (
	"fmt"
	"golang-cache/cache/cache"
	"time"
)

func main() {
	cache := cache.New()

	_, err := cache.Get("test")
	if err != nil {
		fmt.Println(err)
	}
	// Result: unknown key

	fmt.Println(cache.Delete("user"))
	// Result: unknown key

	cache.Set("username", "Golang", time.Second * 5)
	value, _ := cache.Get("username")
	fmt.Println(value)
}
