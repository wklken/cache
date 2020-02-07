# cache

go cache with fetch hook

- retrieveFunc will be called if the key not in cache
- TTL required
- support backend: memory (via [go-redis/cache](https://github.com/go-redis/cache))
- support backend: redis(via [patrickmn/go-cache](https://github.com/patrickmn/go-cache))


## usage

#### use string key

```go

```



build your own Key


```go
# 1. impl the keyer
func NewExampleKey(field1 string, field2 int64) ExampleKey {
	return ExampleKey{
		Field1: field1,
		Field2: field2,
	}
}

func (k ExampleKey) Key() string {
	return fmt.Sprintf("%s:%d", k.Field1, k.Field2)
}



# 2. impl the reterive func
func RetrieveExample(key ExampleKey) (interface{}, error) {
    f1 := key.Field1
    f2 := key.Field2

	username, err := GetFromDatabase(f1, f2)
	if err != nil {
		return nil, err
	}
	return username, nil
}



# 3. new a cache
var exampleCache cache.CacheFactory
exampleCache = cache.NewCache(
	"example",
	false,
    RetrieveExample,
	5*time.Minute,
	6*time.Minute)

# 4. use it
k := ExampleKey{
    f1: "aaaa",
    f2: 1,
}
exmaplCache.Get(k)
exmaplCache.GetString(k)
exmapleCache.Set(k, xxx)
```
