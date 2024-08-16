package cache

import (
	"testing"
	"time"
)

func TestCache_SetAndGet(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Set("key1", "value1", time.Second*10)
	cache.Set("key2", "value2", time.Second*10)

	val, err := cache.Get("key1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if val != "value1" {
		t.Fatalf("expected value1, got %v", val)
	}

	// Add one more item to trigger eviction
	cache.Set("key3", "value3", time.Second*10)
	_, err = cache.Get("key2")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestCache_Expiration(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Set("key1", "value1", time.Millisecond*10)
	time.Sleep(time.Millisecond * 20)

	_, err := cache.Get("key1")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
