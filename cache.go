package cache

import (
	"fmt"
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

const EmptyCacheExpiration = 5 * time.Second

// TODO: retrieveFunc 的参数, 不能只是 string!
type RetrieveFunc func(keyer Keyer) (interface{}, error)

type CacheFactory interface {
	Get(keyer Keyer) (interface{}, error)

	GetString(keyer Keyer) (string, error)
	GetBool(keyer Keyer) (bool, error)
	GetTime(keyer Keyer) (time.Time, error)

	Disabled() bool
}

type Cache struct {
	cache        *gocache.Cache
	disabled     bool
	retrieveFunc RetrieveFunc
	retrieveMu   sync.RWMutex
}

func NewCache(name string, disabled bool, retrieveFunc RetrieveFunc,
	expiration time.Duration, cleanupInterval time.Duration) *Cache {
	return &Cache{
		cache: NewTTLCache(
			name,
			expiration,
			cleanupInterval,
		),
		disabled:     disabled,
		retrieveFunc: retrieveFunc,
	}
}

// TODO: 内存上可以优化, error相同的话使用同一个对象
type EmptyCache struct {
	err error
}

func (c *Cache) Get(keyer Keyer) (interface{}, error) {
	// 1. if cache is disabled, fetch and return
	if c.disabled {
		value, err := c.retrieveFunc(keyer)
		if err != nil {
			return nil, err
		}
		return value, nil
	}

	key := keyer.Key()

	// 2. get from cache
	value, ok := c.cache.Get(key)
	if ok {
		// if retrieve fail from retrieveFunc
		if emptyCache, isEmptyCache := value.(EmptyCache); isEmptyCache {
			return nil, emptyCache.err
		}
		return value, nil
	}

	// 3. if not exists in cache, retrieve it
	return c.doRetrieve(keyer)
}

func (c *Cache) doRetrieve(keyer Keyer) (interface{}, error) {
	// 3 lock and unlock
	c.retrieveMu.Lock()
	defer c.retrieveMu.Unlock()

	key := keyer.Key()

	// 3.1 check again
	value, ok := c.cache.Get(key)
	if ok {
		// if retrieve fail from retrieveFunc
		if emptyCache, isEmptyCache := value.(EmptyCache); isEmptyCache {
			return nil, emptyCache.err
		}
		return value, nil
	}
	// 3.2 fetch
	value, err := c.retrieveFunc(keyer)
	if err != nil {
		// ! if error, cache it too, make it short enough(5s)
		c.cache.Set(key, EmptyCache{err: err}, EmptyCacheExpiration)
		return nil, err
	}

	// 4. set value to cache, use default expiration
	c.cache.Set(key, value, 0)

	return value, nil
}

// TODO: 这里需要实现所有类型的 GetXXXX

// ! if retrieve fail, will return ("", err) for expire time
func (c *Cache) GetString(keyer Keyer) (string, error) {
	value, err := c.Get(keyer)
	if err != nil {
		return "", err
	}

	v, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("not a string value. key=%s, value=%v(%T)", keyer.Key(), value, value)
	}
	return v, nil
}

func (c *Cache) GetBool(keyer Keyer) (bool, error) {
	value, err := c.Get(keyer)
	if err != nil {
		return false, err
	}

	v, ok := value.(bool)
	if !ok {
		return false, fmt.Errorf("not a string value. key=%s, value=%v(%T)", keyer.Key(), value, value)
	}
	return v, nil
}

var defaultZeroTime = time.Time{}

func (c *Cache) GetTime(keyer Keyer) (time.Time, error) {
	value, err := c.Get(keyer)
	if err != nil {
		return defaultZeroTime, err
	}

	v, ok := value.(time.Time)
	if !ok {
		return defaultZeroTime, fmt.Errorf("not a string value. key=%s, value=%v(%T)", keyer.Key(), value, value)
	}
	return v, nil
}

func (c *Cache) Disabled() bool {
	return c.disabled
}
