package inmemdb

import "time"

type Inmem interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
}
