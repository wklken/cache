package cache

import (
	"time"
)

type Key interface {
	Key() string
}

type RetrieveFunc func(key Key) (interface{}, error)

type Cache interface {
	Get(key Key) (interface{}, error)

	GetString(key Key) (string, error)
	GetBool(key Key) (bool, error)
	GetInt64(key Key) (int64, error)
	GetTime(key Key) (time.Time, error)

	Delete(key Key) error
	Exists(key Key) bool

	DirectGet(key Key) (interface{}, bool)

	Disabled() bool
}
