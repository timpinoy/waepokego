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
	cache map[string]cacheEntry
	mutex *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}
	go cache.reaploop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reaploop(loopDuration time.Duration) {
	ticker := time.NewTicker(loopDuration)
	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			for key, entry := range c.cache {
				if time.Since(entry.createdAt) > loopDuration {
					delete(c.cache, key)
				}
			}
			c.mutex.Unlock()
		}
	}
}
