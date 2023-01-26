package main

import (
	"fmt"
	"golang-cache/cache/cache"
)

func main() {
	fmt.Println("[---Start---]")

	cache := cache.New()

	cache.Set("Danil", 18)
  cache.Set("userId", 919)
	cache.Set("Sofii", 19)

	fmt.Println(cache)

	userId := cache.Get("userId")
	fmt.Println(userId)

	cache.Delete("userId")

}