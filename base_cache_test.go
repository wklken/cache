package cache

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/wklken/cache/backend"
)

func retrieveTest(k Key) (interface{}, error) {
	kStr := k.Key()
	switch kStr {
	case "a":
		return "1", nil
	case "b":
		return "2", nil
	case "error":
		return nil, errors.New("error")
	case "bool":
		return true, nil
	case "time":
		return time.Time{}, nil
	default:
		return "", nil
	}
}

func TestNewBaseCache(t *testing.T) {
	expiration := 5 * time.Minute

	be := backend.NewMemoryBackend("test", expiration)

	c := NewBaseCache(false, retrieveTest, be)

	// get from cache
	aKey := NewStringKey("a")
	x, err := c.Get(aKey)
	assert.NoError(t, err)
	assert.Equal(t, "1", x.(string))

	x, err = c.Get(aKey)
	assert.NoError(t, err)
	assert.Equal(t, "1", x.(string))

	// get string
	x, err = c.GetString(aKey)
	assert.NoError(t, err)
	assert.Equal(t, "1", x)

	// get bool
	boolKey := NewStringKey("bool")
	x, err = c.GetBool(boolKey)
	assert.NoError(t, err)
	assert.Equal(t, true, x.(bool))
	// get time
	timeKey := NewStringKey("time")
	x, err = c.GetTime(timeKey)
	assert.NoError(t, err)
	assert.IsType(t, time.Time{}, x)

	// get fail
	errorKey := NewStringKey("error")
	x, err = c.Get(errorKey)
	assert.Error(t, err)
	assert.Nil(t, x)

	err1 := err

	// get fail twice
	x, err = c.Get(errorKey)
	assert.Error(t, err)
	assert.Nil(t, x)

	err2 := err

	// the error should be the same
	assert.Equal(t, err1, err2)

	x, err = c.GetString(errorKey)
	assert.Error(t, err)
	assert.Equal(t, "", x)

	// disabled=true
	// c = NewCache("test", true, retrieveOK, expiration, cleanupInterval)
	c = NewBaseCache(true, retrieveTest, be)
	assert.NotNil(t, c)
	x, err = c.Get(aKey)
	assert.NoError(t, err)
	assert.Equal(t, "1", x.(string))

	_, err = c.GetString(timeKey)
	assert.Error(t, err)

	_, err = c.GetBool(aKey)
	assert.Error(t, err)

	_, err = c.GetTime(aKey)
	assert.Error(t, err)

	// TODO: add emptyCache here
}
