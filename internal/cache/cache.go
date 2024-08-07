package cache

import (
	"errors"
	"multi-backend-caching-library/internal/config"
	"time"
)

type Cache interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

func NewCache(cfg *config.Config) (Cache, error) {

	switch cfg.CacheType {
	case "inmemory":
		return NewLRUCache(cfg.MaxSize), nil
	case "redis":
		return NewRedisCache(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB), nil
	default:
		return nil, errors.New("unsupported cache type")
	}
}
