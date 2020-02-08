package cache

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"

)

func retrieveOK(k Key) (interface{}, error) {
	return "ok", nil
}

func TestNewCache(t *testing.T) {
	expiration := 5 * time.Minute

	c := NewCache("test", false, retrieveOK, expiration)
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
