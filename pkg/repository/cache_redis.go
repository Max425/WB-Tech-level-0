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

func (r *RedisStore) Set(ctx context.Context, key int, value interface{}, lifetime time.Duration) error {
	return r.client.Set(ctx, strconv.Itoa(key), value, lifetime).Err()

}

func (r *RedisStore) Get(ctx context.Context, key int) (interface{}, error) {
	return r.client.Get(ctx, strconv.Itoa(key)).Result()
}

func (r *RedisStore) Delete(ctx context.Context, key int) error {
	return r.client.Del(ctx, strconv.Itoa(key)).Err()
}
