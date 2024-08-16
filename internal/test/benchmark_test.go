package test

import (
	"testing"
	"time"
	"multi-backend-caching-library/internal/cache"
)

func BenchmarkLRUCache_Set(b *testing.B) {
	cache := cache.NewLRUCache(1000)
	for i := 0; i < b.N; i++ {
		cache.Set("key", "value", time.Second*10)
	}
}

func BenchmarkLRUCache_Get(b *testing.B) {
	cache := cache.NewLRUCache(1000)
	cache.Set("key", "value", time.Second*10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get("key")
	}
}
