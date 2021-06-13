package redis

import (
	"bytes"
	"io"
	"io/ioutil"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func (r *Redis) Set(hash string, cache io.Reader, expiration time.Duration) error {
	content, err := ioutil.ReadAll(cache)
	if err != nil {
		return err
	}
	cmd := r.client.Set(hash, content, expiration)
	return cmd.Err()
}

func (r *Redis) Get(hash string) (cache io.Reader, err error) {

	cmd := r.client.Get(hash)
	if cmd.Err() != nil {
		return nil, err
	}
	response, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(response), nil
}

func NewRedis(configOpt ConfigOpt) (*Redis, error) {
	var config Config
	configOpt(&config)
	client := redis.NewClient(
		&redis.Options{
			Addr: config.DNS(),
			DB:   config.Database,
		})
	rd := &Redis{
		client: client,
	}
	return rd, rd.client.Ping().Err()
}
