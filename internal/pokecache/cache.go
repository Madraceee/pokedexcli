package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		store: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return &cache
}

// Cache Methods

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.store[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, present := c.store[key]
	return val.val, present
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mutex.Lock()
		for k, v := range c.store {
			now := time.Now().UTC()
			if v.createdAt.Before(now.Add(-interval)) {
				delete(c.store, k)
			}
		}
		c.mutex.Unlock()
	}
}
