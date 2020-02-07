package backend

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

const (
	DefaultCleanupInterval = 5 * time.Minute
)

// NewTTLCache create cache with expiration and cleanup interval,
// if cleanupInterval is 0, will use DefaultCleanupInterval
func newTTLCache(expiration time.Duration, cleanupInterval time.Duration) *gocache.Cache {
	if cleanupInterval == 0 {
		cleanupInterval = DefaultCleanupInterval
	}

	return gocache.New(expiration, cleanupInterval)
}

type MemoryBackend struct {
	name  string
	cache *gocache.Cache

	defaultExpiration time.Duration
}

func (c *MemoryBackend) Set(key string, value interface{}, duration time.Duration) {
	if duration == time.Duration(0) {
		duration = c.defaultExpiration
	}

	c.cache.Set(key, value, duration)
}

func (c *MemoryBackend) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

func NewMemoryBackend(name string, expiration time.Duration, cleanupInterval time.Duration) *MemoryBackend {
	return &MemoryBackend{
		name:              name,
		cache:             newTTLCache(expiration, cleanupInterval),
		defaultExpiration: expiration,
	}
}
