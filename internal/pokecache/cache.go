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
	mu    *sync.RWMutex
	cache map[string]cacheEntry
}

type Cache interface {
	Get(key string) ([]byte, bool)
	Put(key string, value []byte) error
	reapLoop(interval time.Duration)
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
	cache := &entriesCache{
		cache: make(map[string]cacheEntry),
		mu:    mu,
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *entriesCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for d := range ticker.C {
		for k, v := range c.cache {
			diff := d.Sub(v.createdAt)
			if diff.Microseconds() >= interval.Microseconds() {
				delete(c.cache, k)
			}
		}
	}
}
