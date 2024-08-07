package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	CacheType        string
	MaxSize          int

}

func LoadConfig() *Config {
	return &Config{
		CacheType:        getEnv("CACHE_TYPE", "inmemory"),
		MaxSize:          getEnvAsInt("CACHE_MAX_SIZE", 100),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.Split(value, sep)
	}
	return defaultValue
}
