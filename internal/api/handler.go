package api

import (
	"log"
	"net/http"
	"time"

	"multi-backend-caching-library/internal/cache"

	"github.com/gin-gonic/gin"
)

type CacheHandler struct {
	cache cache.Cache
}

func RegisterRoutes(r *gin.Engine, c cache.Cache) {
	handler := &CacheHandler{cache: c}
	r.GET("/cache/:key", handler.getCache)
	r.POST("/cache", handler.setCache)
	r.PUT("/cache", handler.updateCache)
	r.DELETE("/cache/:key", handler.deleteCache)
}

func (h *CacheHandler) getCache(c *gin.Context) {
	key := c.Param("key")
	log.Printf("Retrieving key: %s", key)
	value, err := h.cache.Get(key)
	if err != nil {
		log.Printf("Error retrieving key: %s, error: %v", key, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "cache miss"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}

func (h *CacheHandler) setCache(c *gin.Context) {
	var req struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
		TTL   int64       `json:"ttl"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	log.Printf("Setting key: %s with value: %v and TTL: %d", req.Key, req.Value, req.TTL)
	err := h.cache.Set(req.Key, req.Value, time.Duration(req.TTL)*time.Second)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to set cache"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "cache set successfully"})
}

func (h *CacheHandler) updateCache(c *gin.Context) {
	var req struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
		TTL   int64       `json:"ttl"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	log.Printf("Updating key: %s with value: %v and TTL: %d", req.Key, req.Value, req.TTL)
	err := h.cache.Set(req.Key, req.Value, time.Duration(req.TTL)*time.Second)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update cache"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "cache updated successfully"})
}

func (h *CacheHandler) deleteCache(c *gin.Context) {
	key := c.Param("key")
	log.Printf("Deleting key: %s", key)
	err := h.cache.Delete(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete cache"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "cache deleted successfully"})
}
