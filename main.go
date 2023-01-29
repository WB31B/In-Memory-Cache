package main

import (
	"fmt"
	"golang-cache/cache/cache"
	"log"
	"time"
)

func main() {
	cache := cache.New()

	cache.Set("Danil", 18, time.Second * 5)
	fmt.Println("MAP:", cache)

	userId, err := cache.Get("Danil")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[find] userId:", userId)

	time.Sleep(time.Second * 6)

	userId, err = cache.Get("Danil")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[delete] userId:", userId)
}
