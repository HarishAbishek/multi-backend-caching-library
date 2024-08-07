# Multi-Backend Caching Library in Go

## Project Overview
This project aims to develop a robust caching library in Go that supports multiple backends. The library features an in-memory cache with an LRU (Least Recently Used) eviction policy and integrates with external caching solutions such as Redis or Memcached. It provides an easy-to-use API for cache operations and includes cache invalidation and expiration policies.

## Objectives
1. Develop an in-memory cache with LRU eviction policy.
2. Integrate with external caches like Redis or Memcached.
3. Design an intuitive API for setting, getting, updating, and deleting cache  entries.
4. Implement cache invalidation and expiration policies.
5. Ensure high performance and scalability of the caching library.

## Project Deliverables
1. **Core Library**:
    - In-memory caching implementation with LRU eviction.
    - Redis or Memcached integration modules.
    - Unified API for cache operations.
    - Configuration options for cache invalidation and expiration.
2. **Documentation**:
    - Comprehensive API documentation.
    - Usage guides and examples.
    - Best practices for integrating the library into Go applications.

## Technical Details
### In-Memory Cache with LRU Eviction
- Use a doubly linked list combined with a hash map for O(1) access and eviction.
- Implement a configurable maximum size for the cache.

### Unified API
- Provide methods like `Set(key string, value interface{}, ttl time.Duration) error`, `Get(key string) (interface{}, error)`, `Update(key string, value interface{}, ttl time.Duration) error`, and `Delete(key string) error`.
- Support both synchronous and asynchronous operations.

### Cache Invalidation and Expiration
- Implement TTL (Time-To-Live) for cache entries.
- Provide methods for manual invalidation.

## Installation and Usage
1. Clone the repository: `git clone https://github.com/yourusername/multi-backend-caching-library.git`
2. Navigate to the project directory: `cd multi-backend-caching-library`
3. Build the project: `go build -o main ./cmd/main.go`
4. Run the application: `./main`
5. Use the provided API endpoints to interact with the cache.

## Docker Setup
1. Build and start the containers: `docker-compose up --build`
2. The application will be available at `http://localhost:8080`

## Environment Variables
- `CACHE_TYPE`: Type of cache to use (`inmemory`, `redis`).
- `CACHE_MAX_SIZE`: Maximum size for the in-memory cache.
- `REDIS_ADDR`: Address of the Redis server.
- `REDIS_PASSWORD`: Password for the Redis server.
- `REDIS_DB`: Redis database number.

## API Endpoints
- `GET /cache/:key`: Retrieve a value from the cache.
- `POST /cache`: Set a value in the cache.
- `PUT /cache`: Update a value in the cache.
- `DELETE /cache/:key`: Delete a value from the cache.

## Testing
- Run unit tests: `go test ./...`
- Run integration tests: `go test -tags=integration ./test`

## License
This project is licensed under the MIT License.
