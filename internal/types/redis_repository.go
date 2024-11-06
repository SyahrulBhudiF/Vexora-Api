package types

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository struct {
	redis  *redis.Client
	prefix string
	ctx    context.Context
}

func NewRedisRepository(redis *redis.Client, prefix string) *RedisRepository {
	return &RedisRepository{
		redis,
		fmt.Sprintf("%s:", prefix),
		context.Background(),
	}
}

func (r *RedisRepository) Set(key string, value interface{}, expiration time.Duration) error {
	key = r.prefix + key
	return r.redis.Set(r.ctx, key, value, expiration).Err()
}

func (r *RedisRepository) Get(key string) (string, error) {
	key = r.prefix + key
	return r.redis.Get(r.ctx, key).Result()
}

func (r *RedisRepository) Delete(key string) error {
	key = r.prefix + key
	return r.redis.Del(r.ctx, key).Err()
}

func (r *RedisRepository) Exists(key string) (bool, error) {
	key = r.prefix + key
	exists, err := r.redis.Exists(r.ctx, key).Result()
	return exists > 0, err
}
