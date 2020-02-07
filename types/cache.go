package types

import (
	"fmt"
	"sync"
	"time"

	"cache/backend"
	"cache/key"
)

const EmptyCacheExpiration = 5 * time.Second

type RetrieveFunc func(k key.Key) (interface{}, error)

type Cache interface {
	Get(k key.Key) (interface{}, error)

	GetString(k key.Key) (string, error)
	GetBool(k key.Key) (bool, error)
	GetTime(k key.Key) (time.Time, error)

	Disabled() bool
}

type BaseCache struct {
	backend backend.Backend

	disabled     bool
	retrieveFunc RetrieveFunc
	retrieveMu   sync.RWMutex
}

// TODO: 内存上可以优化, error相同的话使用同一个对象
type EmptyCache struct {
	err error
}

func (c *BaseCache) Get(k key.Key) (interface{}, error) {
	// 1. if cache is disabled, fetch and return
	if c.disabled {
		value, err := c.retrieveFunc(k)
		if err != nil {
			return nil, err
		}
		return value, nil
	}

	key := k.Key()

	// 2. get from cache
	value, ok := c.backend.Get(key)
	if ok {
		// if retrieve fail from retrieveFunc
		if emptyCache, isEmptyCache := value.(EmptyCache); isEmptyCache {
			return nil, emptyCache.err
		}
		return value, nil
	}

	// 3. if not exists in cache, retrieve it
	return c.doRetrieve(k)
}

func (c *BaseCache) doRetrieve(k key.Key) (interface{}, error) {
	// 3 lock and unlock
	c.retrieveMu.Lock()
	defer c.retrieveMu.Unlock()

	key := k.Key()

	// 3.1 check again
	value, ok := c.backend.Get(key)
	if ok {
		// if retrieve fail from retrieveFunc
		if emptyCache, isEmptyCache := value.(EmptyCache); isEmptyCache {
			return nil, emptyCache.err
		}
		return value, nil
	}
	// 3.2 fetch
	value, err := c.retrieveFunc(k)
	if err != nil {
		// ! if error, cache it too, make it short enough(5s)
		c.backend.Set(key, EmptyCache{err: err}, EmptyCacheExpiration)
		return nil, err
	}

	// 4. set value to cache, use default expiration
	c.backend.Set(key, value, 0)

	return value, nil
}

// TODO: 这里需要实现所有类型的 GetXXXX

// ! if retrieve fail, will return ("", err) for expire time
func (c *BaseCache) GetString(k key.Key) (string, error) {
	value, err := c.Get(k)
	if err != nil {
		return "", err
	}

	v, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("not a string value. key=%s, value=%v(%T)", k.Key(), value, value)
	}
	return v, nil
}

func (c *BaseCache) GetBool(k key.Key) (bool, error) {
	value, err := c.Get(k)
	if err != nil {
		return false, err
	}

	v, ok := value.(bool)
	if !ok {
		return false, fmt.Errorf("not a string value. key=%s, value=%v(%T)", k.Key(), value, value)
	}
	return v, nil
}

var defaultZeroTime = time.Time{}

func (c *BaseCache) GetTime(k key.Key) (time.Time, error) {
	value, err := c.Get(k)
	if err != nil {
		return defaultZeroTime, err
	}

	v, ok := value.(time.Time)
	if !ok {
		return defaultZeroTime, fmt.Errorf("not a string value. key=%s, value=%v(%T)", k.Key(), value, value)
	}
	return v, nil
}

func (c *BaseCache) Disabled() bool {
	return c.disabled
}

func NewBaseCache(disabled bool, retrieveFunc RetrieveFunc, backend backend.Backend) Cache {
	return &BaseCache{
		backend:      backend,
		disabled:     disabled,
		retrieveFunc: retrieveFunc,
	}
}
