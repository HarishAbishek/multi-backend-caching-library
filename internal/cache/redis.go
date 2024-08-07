package cache

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	err := r.client.Set(r.ctx, key, value, ttl).Err()
	return err
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	log.Printf("Error in Get key: %s, error: %v", key, err)
	if err == redis.Nil {
		return nil, errors.New("cache miss")
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *RedisCache) Delete(key string) error {
	err := r.client.Del(r.ctx, key).Err()
	return err
}
