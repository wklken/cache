package cache

// in-memory cache
import (
	"time"

	gocache "github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

const (
	DefaultCleanupInterval = 5 * time.Minute
	NoCleanupInterval      = 0
)

// ? 定期清理内存锁了, 此时会不会导致请求等待? 或者波动?
// NOTE: 另外, 目前没有限制cache量, 没有使用LRU模型; 可能的问题是内存占用过多

func createTTLCache(expiration time.Duration, cleanupInterval time.Duration) *gocache.Cache {
	if cleanupInterval == 0 {
		cleanupInterval = DefaultCleanupInterval
	}

	return gocache.New(expiration, cleanupInterval)
}

// NewTTLCache create cache with expiration and cleanup interval,
// if cleanupInterval is 0, will use DefaultCleanupInterval
func NewTTLCache(name string, expiration time.Duration, cleanupInterval time.Duration) *gocache.Cache {
	log.Infof("Create TTLCache: %s, expiration: %s, cleanupInterval: %s", name, expiration, cleanupInterval)
	return createTTLCache(expiration, cleanupInterval)
}

// NewTTLCacheWithEvictedFunc create cache with expiration and cleanup interval
// if cleanupInterval is 0, will use DefaultCleanupInterval
// Bound a OnEvicted func
func NewTTLCacheWithEvictedFunc(name string, expiration time.Duration, cleanupInterval time.Duration,
	f func(string, interface{})) *gocache.Cache {
	c := createTTLCache(expiration, cleanupInterval)
	c.OnEvicted(f)

	log.Infof("Create TTLCacheWithEvictedFunc: %s, expiration: %s, cleanupInterval: %s",
		name, expiration, cleanupInterval)

	return c
}

// NewCache create in-memory cache without expiration and cleanup interval
// func NewCache(name string) *gocache.Cache {
// 	c := gocache.New(gocache.DefaultExpiration, NoCleanupInterval)
// 	log.Infof("Create Cache: %s", name)
// 	return c
// }
