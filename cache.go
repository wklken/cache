package cache

import (
	"time"

	"github.com/wklken/cache/backend"
)

func NewCache(name string, disabled bool, retrieveFunc RetrieveFunc,
	expiration time.Duration) Cache {
	be := backend.NewMemoryBackend(name, expiration)
	return NewBaseCache(disabled, retrieveFunc, be)
}
