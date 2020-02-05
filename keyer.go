package cache

import (
	"fmt"
)

type Keyer interface {
	Key() string
}

type StringKey struct {
	key string
}

func NewStringKey(key string) StringKey {
	return StringKey{
		key: key,
	}
}

func (s *StringKey) Key() string {
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

func (k *Int64Key) Key() string {
	return fmt.Sprintf("%d", k.key)
}

type ExampleKey struct {
	Field1 string
	Field2 int64
}

func NewExampleKey(field1 string, field2 int64) ExampleKey {
	return ExampleKey{
		Field1: field1,
		Field2: field2,
	}
}

func (k *ExampleKey) Key() string {
	return fmt.Sprintf("%s:%d", k.Field1, k.Field2)
}
