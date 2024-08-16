package config

import (
	"os"
	"testing"
)

func TestLoadConfig_Defaults(t *testing.T) {
	cfg := LoadConfig()

	if cfg.CacheType != "inmemory" {
		t.Errorf("expected CacheType to be 'inmemory', got '%s'", cfg.CacheType)
	}

	if cfg.MaxSize != 100 {
		t.Errorf("expected MaxSize to be 100, got %d", cfg.MaxSize)
	}

	if cfg.RedisAddr != "localhost:6379" {
		t.Errorf("expected RedisAddr to be 'localhost:6379', got '%s'", cfg.RedisAddr)
	}

	if cfg.RedisPassword != "" {
		t.Errorf("expected RedisPassword to be empty, got '%s'", cfg.RedisPassword)
	}

	if cfg.RedisDB != 0 {
		t.Errorf("expected RedisDB to be 0, got %d", cfg.RedisDB)
	}
}

func TestLoadConfig_EnvironmentVariables(t *testing.T) {
	os.Setenv("CACHE_TYPE", "redis")
	os.Setenv("CACHE_MAX_SIZE", "200")
	os.Setenv("REDIS_ADDR", "192.168.1.1:6379")
	os.Setenv("REDIS_PASSWORD", "mysecretpassword")
	os.Setenv("REDIS_DB", "1")

	cfg := LoadConfig()

	if cfg.CacheType != "redis" {
		t.Errorf("expected CacheType to be 'redis', got '%s'", cfg.CacheType)
	}

	if cfg.MaxSize != 200 {
		t.Errorf("expected MaxSize to be 200, got %d", cfg.MaxSize)
	}

	if cfg.RedisAddr != "192.168.1.1:6379" {
		t.Errorf("expected RedisAddr to be '192.168.1.1:6379', got '%s'", cfg.RedisAddr)
	}

	if cfg.RedisPassword != "mysecretpassword" {
		t.Errorf("expected RedisPassword to be 'mysecretpassword', got '%s'", cfg.RedisPassword)
	}

	if cfg.RedisDB != 1 {
		t.Errorf("expected RedisDB to be 1, got %d", cfg.RedisDB)
	}

	// Clean up environment variables after test
	os.Clearenv()
}

func TestLoadConfig_InvalidEnvValues(t *testing.T) {
	os.Setenv("CACHE_MAX_SIZE", "invalid")

	cfg := LoadConfig()

	if cfg.MaxSize != 100 { // Should fall back to default value
		t.Errorf("expected MaxSize to fall back to 100, got %d", cfg.MaxSize)
	}

	// Clean up environment variables after test
	os.Clearenv()
}

func BenchmarkLoadConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoadConfig()
	}
}
