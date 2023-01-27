package main

import (
	"fmt"
	"golang-cache/cache/cache"
)

func main() {
	fmt.Println("[---Start---]")

	cache := cache.New()

	cache.Set("Danil", 18)
  cache.Set("userId", "id")
	cache.Set("Sofii", 19)
	cache.Set("", 4)

	fmt.Println(cache)

	userId := cache.Get("userId")
	fmt.Println(userId)

	cache.Delete("Danil")
	fmt.Println(cache)

}