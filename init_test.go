package cache

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"

	"github.com/wklken/cache/key"
)

func retrieveOK(k key.Key) (interface{}, error) {
	return "", nil
}

func TestNewCache(t *testing.T) {
	expiration := 5 * time.Minute
	cleanupInterval := 6 * time.Minute

	c := NewCache("test", false, retrieveOK, expiration, cleanupInterval)
	assert.NotNil(t, c)
}

func TestNewRedisCache(t *testing.T) {
	expiration := 5 * time.Minute

	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	cli := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	c := NewRedisCache("test", false, retrieveOK, cli, expiration)
	assert.NotNil(t, c)
}
