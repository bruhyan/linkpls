package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type RedisClient struct {
	client *redis.Client
}

func New(addr string, password string, db int, protocol int) *RedisClient {
	return &RedisClient{
		client: redis.NewClient(
			&redis.Options{
				Addr:     addr,
				Password: password,
				DB:       db,
				Protocol: protocol,
			},
		),
	}
}

func (r *RedisClient) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *RedisClient) Set(key, value string) error {
	return r.client.Set(context.Background(), key, value, time.Hour).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *RedisClient) Incr(key string) (int64, error) {
	return r.client.Incr(context.Background(), key).Result()
}
