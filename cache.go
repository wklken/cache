package cache

import (
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/wklken/cache/backend"
)

func NewCache(name string, disabled bool, retrieveFunc RetrieveFunc,
	expiration time.Duration, cleanupInterval time.Duration) Cache {
	be := backend.NewMemoryBackend(name, expiration, cleanupInterval)
	return NewBaseCache(disabled, retrieveFunc, be)
}

func NewRedisCache(name string, disabled bool, retrieveFunc RetrieveFunc,
	cli *redis.Client, expiration time.Duration) Cache {
	be := backend.NewRedisBackend(name, cli, expiration)
	return NewBaseCache(disabled, retrieveFunc, be)
}
