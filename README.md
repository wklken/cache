# cache
go cache with fetch hook


## usage

use string key

```go
# 1. impl the reterive func
func RetrieveExample(key StringKey) (interface{}, error) {
    key := key.Key()

	username, err := GetFromDatabase(key)
	if err != nil {
		return nil, err
	}
	return username, nil
}

# 2. new a cache
exampleCache := cache.NewCache(
	"example",
	false,
    RetrieveExample,
	5*time.Minute,
	6*time.Minute)

# 4. use it
k := NewStringKey("aaa")

exmaplCache.Get(k)
exmapleCache.Set(k, xxx)
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

func (k *ExampleKey) Key() string {
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
