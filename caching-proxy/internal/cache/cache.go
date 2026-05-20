package cache

import (
	"sync"
)

type CachedResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

func NewCacheSystem() *CacheSystem{
	return &CacheSystem{
		storage: make(map[string]CachedResponse),
	}
}
type CacheSystem struct {
	lock    sync.Mutex
	storage map[string]CachedResponse
}

func (c *CacheSystem) Get(key string) (CachedResponse, bool) {
	c.lock.Lock()

	defer c.lock.Unlock()

	item, exists := c.storage[key]
	if !exists {
		return CachedResponse{}, false
	}
	return item, true
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
