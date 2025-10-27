package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type entriesCache struct {
	interval time.Duration
	mu       *sync.RWMutex
	cache    map[string]cacheEntry
}

type Cache interface {
	Get(key string) ([]byte, bool)
	Put(key string, value []byte) error
}

func (c *entriesCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.cache[key]
	if !ok {
		return nil, ok
	}
	return v.val, ok
}
func (c *entriesCache) Put(key string, value []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	return nil
}

func NewCache(interval time.Duration) Cache {
	mu := &sync.RWMutex{}
	return &entriesCache{
		cache:    make(map[string]cacheEntry),
		mu:       mu,
		interval: interval,
	}
}
