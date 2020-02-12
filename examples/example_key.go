package main

import (
	"fmt"
	"time"

	"github.com/wklken/cache"
)

// 1. impl the key interface, Key() string
type ExampleKey struct {
	Field1 string
	Field2 int64
}

func (k ExampleKey) Key() string {
	return fmt.Sprintf("%s:%d", k.Field1, k.Field2)
}

// 2. impl the reterive func
func RetrieveExample(inKey cache.Key) (interface{}, error) {
	k := inKey.(ExampleKey)
	fmt.Println("ExampleKey Field1 and Field2 value:", k.Field1, k.Field2)
	// data, err := GetFromDatabase(k.Field1, k.Field2)
	// if err != nil {
	//     return nil, err
	// }
	return "world", nil
}

func main() {
	// 3. new a cache
	c := cache.NewCache(
		"example",
		false,
		RetrieveExample,
		5*time.Minute)

	// 4. use it
	k := ExampleKey{
		Field1: "hello",
		Field2: 42,
	}

	data, err := c.Get(k)
	fmt.Println("err == nil: ", err == nil)
	fmt.Println("data from cache: ", data)

	dataStr, err := c.GetString(k)
	fmt.Println("err == nil: ", err == nil)
	fmt.Printf("data type is %T, value is %s\n", dataStr, dataStr)
}
