package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	data     map[string]cacheEntry
	interval time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	_, exists := c.data[key]
	defer c.mu.Unlock()
	if exists == false {
		c.data[key] = cacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exists := c.data[key]
	defer c.mu.Unlock()
	return entry.val, exists
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.mu.Lock()
		for i, v := range c.data {
			if time.Since(v.createdAt) > c.interval {
				delete(c.data, i)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop()

	return &cache
}
