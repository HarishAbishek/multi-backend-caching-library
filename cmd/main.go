package main

import (
	"log"
	"multi-backend-caching-library/internal/api"
	"multi-backend-caching-library/internal/cache"
	"multi-backend-caching-library/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	c, err := cache.NewCache(cfg)
	if err != nil {
		log.Fatalf("Failed to create cache: %v", err)
	}

	r := gin.Default()
	api.RegisterRoutes(r, c)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
