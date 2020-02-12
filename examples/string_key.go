package main

import (
	"fmt"
	"time"

	"github.com/wklken/cache"
)

// 1. impl the reterive func
func RetrieveOK(k cache.Key) (interface{}, error) {
	arg := k.(cache.StringKey)
	fmt.Println("arg: ", arg)
	// you can use the arg to fetch data from database or http request
	// username, err := GetFromDatabase(arg)
	// if err != nil {
	//     return nil, err
	// }
	return "ok", nil
}

func main() {
	// 2. new a cache
	c := cache.NewCache(
		"example",
		false,
		RetrieveOK,
		5*time.Minute)

	// 4. use it
	k := cache.NewStringKey("hello")

	data, err := c.Get(k)
	fmt.Println("err == nil: ", err == nil)
	fmt.Println("data from cache: ", data)
}
