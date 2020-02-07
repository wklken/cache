package backend

import "time"

type Backend interface {
    Set(key string, value interface{}, duration time.Duration)
    Get(key string) (interface{}, bool)
}
