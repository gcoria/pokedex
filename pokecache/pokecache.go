package pokecache

import (
	"time"
)

type Cache struct {
	cache map[string]cachEntry
}

type cachEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]cachEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cachEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Gett(key string) ([]byte, bool) {
	cache, ok := c.cache[key]
	return cache.val, ok
}
