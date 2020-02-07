package backend

import (
	"fmt"
	"time"

	gordscache "github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v4"
)

type RedisBackend struct {
	name string

	codec *gordscache.Codec

	defaultExpiration time.Duration
}

func (c *RedisBackend) Set(key string, value interface{}, duration time.Duration) {
	rKey := fmt.Sprintf("%s:%s", c.name, key)

	if duration == time.Duration(0) {
		duration = c.defaultExpiration
	}

	err := c.codec.Set(&gordscache.Item{
		Key:        rKey,
		Object:     value,
		Expiration: duration,
	})

	if err != nil {
		log.WithError(err).Errorf("set key into redis fail [cache=%s, key=%s]", c.name, key)
	}
}

func (c *RedisBackend) Get(key string) (interface{}, bool) {
	rKey := fmt.Sprintf("%s:%s", c.name, key)

	var value interface{}
	err := c.codec.Get(rKey, &value)

	if err != nil {
		return nil, false
	}

	return value, true
}

func NewRedisBackend(name string, cli *redis.Client, expiration time.Duration) *RedisBackend {
	return &RedisBackend{
		name: name,
		codec: &gordscache.Codec{
			Redis: cli,

			Marshal:   msgpack.Marshal,
			Unmarshal: msgpack.Unmarshal,
		},
		defaultExpiration: expiration,
	}
}
