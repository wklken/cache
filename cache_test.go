package cache

import (
	"testing"
	"time"

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
