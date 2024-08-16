package api

import (
	"bytes"
	"encoding/json"
	"multi-backend-caching-library/internal/cache"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCacheHandler_SetAndGet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cache := cache.NewLRUCache(2)

	r := gin.Default()
	RegisterRoutes(r, cache)

	// Test setting a value
	body := map[string]interface{}{
		"key":   "key1",
		"value": "value1",
		"ttl":   100,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/cache", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %v", w.Code)
	}

	// Test getting the value
	req, _ = http.NewRequest(http.MethodGet, "/cache/key1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %v", w.Code)
	}

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp["value"] != "value1" {
		t.Fatalf("expected value1, got %v", resp["value"])
	}
}

func TestCacheHandler_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cache := cache.NewLRUCache(2)

	r := gin.Default()
	RegisterRoutes(r, cache)

	// Test setting a value
	body := map[string]interface{}{
		"key":   "key2",
		"value": "value2",
		"ttl":   1003,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/cache", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %v", w.Code)
	}

	// Test deleting the value
	req, _ = http.NewRequest(http.MethodDelete, "/cache/key2", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %v", w.Code)
	}

	// Test getting the value should return 404
	req, _ = http.NewRequest(http.MethodGet, "/cache/key2", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %v", w.Code)
	}
}
