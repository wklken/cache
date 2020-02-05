package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// func TestNewCache(t *testing.T) {
//     c := NewCache("test")
//     assert.NotNil(t, c)
// }

func TestPrivateNewTTLCache(t *testing.T) {
	c := createTTLCache(5*time.Second, 10*time.Second)
	assert.NotNil(t, c)

	c = createTTLCache(5*time.Second, 0)
	assert.NotNil(t, c)
}

func TestNewTTLCache(t *testing.T) {
	c := NewTTLCache("test", 5*time.Second, 10*time.Second)
	assert.NotNil(t, c)
}

func TestNewTTLCacheWithEvictedFunc(t *testing.T) {
	f := func(key string, value interface{}) {
	}
	c := NewTTLCacheWithEvictedFunc("test", 5*time.Second, 10*time.Second, f)
	assert.NotNil(t, c)
}
