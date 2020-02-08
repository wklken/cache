package backend

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTTLCache(t *testing.T) {
	c := newTTLCache(5*time.Second, 10*time.Second)
	assert.NotNil(t, c)
}

func TestMemoryBackend(t *testing.T) {
	be := NewMemoryBackend("test", 5*time.Second)
	assert.NotNil(t, be)

	_, found := be.Get("not_exists")
	assert.False(t, found)

	be.Set("hello", "world", time.Duration(0))
	value, found := be.Get("hello")
	assert.True(t, found)
	assert.Equal(t, "world", value)
}
