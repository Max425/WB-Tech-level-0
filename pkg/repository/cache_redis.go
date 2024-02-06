package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{client: client}
}

func (r *RedisStore) Set(ctx context.Context, key int, value []byte, lifetime time.Duration) error {
	return r.client.Set(ctx, strconv.Itoa(key), value, lifetime).Err()

}

func (r *RedisStore) Get(ctx context.Context, key int) ([]byte, error) {
	return r.client.Get(ctx, strconv.Itoa(key)).Bytes()
}

func (r *RedisStore) Delete(ctx context.Context, key int) error {
	return r.client.Del(ctx, strconv.Itoa(key)).Err()
}
