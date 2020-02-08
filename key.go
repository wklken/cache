package cache

import "fmt"

// StringKey
type StringKey struct {
	key string
}

func NewStringKey(key string) StringKey {
	return StringKey{
		key: key,
	}
}

func (s StringKey) Key() string {
	return s.key
}

type Int64Key struct {
	key int64
}

func NewInt64Key(key int64) Int64Key {
	return Int64Key{
		key: key,
	}
}

func (k Int64Key) Key() string {
	return fmt.Sprintf("%d", k.key)
}
