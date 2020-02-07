package backend

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
)

func TestRedisBackend(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	cli := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	be := NewRedisBackend("test", cli, 5*time.Second)

	_, found := be.Get("not_exists")
	assert.False(t, found)

	be.Set("hello", "world", time.Duration(0))

	value, found := be.Get("hello")
	assert.True(t, found)
	assert.Equal(t, "world", value)
}
