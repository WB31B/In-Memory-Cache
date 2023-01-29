package main

import (
	"fmt"
	"golang-cache/cache/cache"
	"time"
)

func main() {
	cache := cache.New()

	cache.Set("Danil", 18, time.Second * 5)
	cache.Set("Sofii", 19, time.Second * 5)
	cache.Set("Artem", 1, time.Second * 5)
	cache.Set("logged", true, time.Second * 5)
	cache.Set("title", "Title", time.Second * 5)

	fmt.Println("[Start programm] MAP:", cache)

	userId := cache.Get("Danil")
	fmt.Println("[find] userId:", userId)

	time.Sleep(time.Second * 6)

	userId = cache.Get("Danil")

	fmt.Println("[delete] userId:", userId)
	fmt.Println("[End programm] MAP:", cache)
}
