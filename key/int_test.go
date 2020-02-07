package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64Key(t *testing.T) {
	k := NewInt64Key(int64(123))
	assert.NotNil(t, k)
	assert.Equal(t, "123", k.Key())
}
