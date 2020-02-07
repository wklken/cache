package key

import "fmt"

type Int64Key struct {
	key int64
}

func NewInt64Key(key int64) *Int64Key {
	return &Int64Key{
		key: key,
	}
}

func (k Int64Key) Key() string {
	return fmt.Sprintf("%d", k.key)
}
