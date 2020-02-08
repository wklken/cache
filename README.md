# cache

go cache with multiple backends support, auto to fetch the data if missing.

- retrieveFunc will be called if the key not in cache
- TTL required
- support backend: memory (via [go-redis/cache](https://github.com/go-redis/cache))
- support backend: redis(via [patrickmn/go-cache](https://github.com/patrickmn/go-cache))


## usage

#### use string key

```go
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
		5*time.Minute,
		6*time.Minute)

	// 4. use it
	k := cache.NewStringKey("hello")

	data, err := c.Get(k)
	fmt.Println("err == nil: ", err == nil)
	fmt.Println("data from cache: ", data)
}
```

#### use your own key


```go
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
		5*time.Minute,
		6*time.Minute)

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
```

#### use redis backend

```go
func retrieveOK(k Key) (interface{}, error) {
	return "ok", nil
}

func main() {
    // 1. mock redis cli via miniredis
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	cli := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

    // 2. create redis backend
	be := NewRedisBackend("test", cli, 5*time.Second)

    // 3. new the cache
	c := NewRedisCache("test", false, retrieveOK, cli, 5 * time.Minute)
}
```
