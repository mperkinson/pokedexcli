// Package pokecache
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.ReapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		value:     value,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	return entry.value, true
}

func (c *Cache) Reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	timeout := time.Now().UTC().Add(-interval)

	for k, v := range c.cache {
		if v.createdAt.Before(timeout) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.Reap(interval)
	}
}
