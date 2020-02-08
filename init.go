package cache

import (
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/wklken/cache/backend"
	"github.com/wklken/cache/types"
)

func NewCache(name string, disabled bool, retrieveFunc types.RetrieveFunc,
	expiration time.Duration, cleanupInterval time.Duration) types.Cache {
	be := backend.NewMemoryBackend(name, expiration, cleanupInterval)
	return types.NewBaseCache(disabled, retrieveFunc, be)
}

func NewRedisCache(name string, disabled bool, retrieveFunc types.RetrieveFunc,
	cli *redis.Client, expiration time.Duration) types.Cache {
	be := backend.NewRedisBackend(name, cli, expiration)
	return types.NewBaseCache(disabled, retrieveFunc, be)
}
