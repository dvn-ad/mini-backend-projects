package cache

import (
	"net/http"
	"sync"
)

type CachedResponse struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

type CacheSystem struct {
	lock    sync.Mutex
	storage map[string]CachedResponse
}

func NewCacheSystem() *CacheSystem{
	return &CacheSystem{
		storage: make(map[string]CachedResponse),
	}
}

func (c *CacheSystem) Get(key string) (CachedResponse, bool) {
	c.lock.Lock()

	defer c.lock.Unlock()

	item, exists := c.storage[key]
	
	return item, exists
}

func (c *CacheSystem) Set(key string, response CachedResponse) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.storage[key] = response
}

func (c *CacheSystem) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.storage = make(map[string]CachedResponse)
}
