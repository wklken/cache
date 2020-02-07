package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleKey(t *testing.T) {
	k := NewExampleKey("hello", 123)
	assert.NotNil(t, k)
	assert.Equal(t, "hello:123", k.Key())
}
