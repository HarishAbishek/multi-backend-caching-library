package test

import (
	"multi-backend-caching-library/internal/cache"
	"multi-backend-caching-library/internal/config"
	"testing"
	"time"
)

func TestCacheIntegration(t *testing.T) {
	cfg := &config.Config{
		CacheType: "inmemory",
		MaxSize:   2,
	}
	c, err := cache.NewCache(cfg)
	if err != nil {
		t.Fatalf("could not create cache: %v", err)
	}

	c.Set("key1", "value1", time.Second*10)
	val, err := c.Get("key1")
	if err != nil {
		t.Fatalf("could not get value from cache: %v", err)
	}

	if val != "value1" {
		t.Fatalf("expected value1, got %v", val)
	}
}
