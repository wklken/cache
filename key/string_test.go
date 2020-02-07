package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringKey(t *testing.T) {
	k := NewStringKey("hello")
	assert.NotNil(t, k)
	assert.Equal(t, "hello", k.Key())
}
