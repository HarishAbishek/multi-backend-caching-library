package cache

import (
	"testing"
	"time"
	"multi-backend-caching-library/internal/config"
)

func TestRedisCache_SetAndGet(t *testing.T) {
	cfg := &config.Config{
		RedisAddr:     "localhost:6379",
		RedisPassword: "",
		RedisDB:       0,
	}

	cache := NewRedisCache(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	// Test setting and getting a value
	err := cache.Set("key1", "value1", time.Second*10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	val, err := cache.Get("key1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if val != "value1" {
		t.Fatalf("expected value1, got %v", val)
	}
}

func TestRedisCache_Expiration(t *testing.T) {
	cfg := &config.Config{
		RedisAddr:     "localhost:6379",
		RedisPassword: "",
		RedisDB:       0,
	}

	cache := NewRedisCache(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	// Test expiration
	err := cache.Set("key1", "value1", time.Millisecond*10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	time.Sleep(time.Millisecond * 20)

	_, err = cache.Get("key1")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestRedisCache_Delete(t *testing.T) {
	cfg := &config.Config{
		RedisAddr:     "localhost:6379",
		RedisPassword: "",
		RedisDB:       0,
	}

	cache := NewRedisCache(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	// Test deletion
	err := cache.Set("key1", "value1", time.Second*10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = cache.Delete("key1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = cache.Get("key1")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
