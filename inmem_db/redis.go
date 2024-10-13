package inmemdb

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisClient{
		client: rdb,
	}
}

func (rdc *RedisClient) Set(key string, value interface{}, exp time.Duration) error {
	err := rdc.client.Set(context.Background(), key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rdc *RedisClient) Get(key string) (string, error) {
	val, err := rdc.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, err
}
