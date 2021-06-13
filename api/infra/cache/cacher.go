package cache

import (
	"io"
	"time"
)

type Cacher interface {
	Set(hash string, cache io.Reader, expiration time.Duration) error
	Get(hash string) (cache io.Reader, err error)
	Ping() error
}
