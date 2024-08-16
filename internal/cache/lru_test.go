package cache

import (
	"testing"
	"time"
)

func TestLRUCache_SetAndGet(t *testing.T) {
	cache := NewLRUCache(2)

	// Test setting and getting a value
	cache.Set("key1", "value1", time.Second*10)
	cache.Set("key2", "value2", time.Second*10)

	// Retrieve the first key
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

func TestLRUCache_Expiration(t *testing.T) {
	cache := NewLRUCache(2)

	// Test expiration
	cache.Set("key1", "value1", time.Millisecond*10)
	time.Sleep(time.Millisecond * 20)

	_, err := cache.Get("key1")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestLRUCache_Delete(t *testing.T) {
	cache := NewLRUCache(2)

	// Test deletion
	cache.Set("key1", "value1", time.Second*10)
	err := cache.Delete("key1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = cache.Get("key1")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestLRUCache_Eviction(t *testing.T) {
	cache := NewLRUCache(2)

	// Test eviction policy
	cache.Set("key1", "value1", time.Second*10)
	cache.Set("key2", "value2", time.Second*10)
	cache.Set("key3", "value3", time.Second*10)

	_, err := cache.Get("key1")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	val, err := cache.Get("key3")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if val != "value3" {
		t.Fatalf("expected value3, got %v", val)
	}
}
