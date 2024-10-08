package cachetest

import (
	"sync"
)

// This is my very simple cache. I was surprised how well it works.
type CustomCache struct {
	sync.RWMutex
	items      map[string][]byte
	totalsizeB int
	maxsizeB   int
	hits       int
	misses     int
}

func NewCustomSimpleCache(maxsizeMB int) *CustomCache {
	return &CustomCache{
		items:      make(map[string][]byte),
		totalsizeB: 0,
		maxsizeB:   maxsizeMB * 1024 * 1024,
		hits:       0,
		misses:     0,
	}
}

func (c *CustomCache) Set(key string, value []byte) {
	c.Lock()
	defer c.Unlock()

	// No entries are allowed that are larger than the maxsize.
	if len(value) >= c.maxsizeB {
		return
	}

	// If the key already exists, remove its size from the total
	if oldValue, exists := c.items[key]; exists {
		c.totalsizeB -= len(oldValue)
	}

	// If adding the new value would exceed the max size, remove items until it fits
	for {
		if (len(value) + c.totalsizeB) < c.maxsizeB {
			break
		}
		for k, v := range c.items {
			delete(c.items, k)
			c.totalsizeB -= len(v)
			break
		}
	}

	// Add the new item
	c.items[key] = value
	c.totalsizeB += len(value)

}

func (c *CustomCache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()
	item, found := c.items[key]
	if found {
		c.hits += 1
	} else {
		c.misses += 1
	}

	return item, found
}

func (c *CustomCache) Del(key string) bool {
	c.Lock()
	defer c.Unlock()
	v, found := c.items[key]
	if found {
		c.hits += 1
		c.totalsizeB -= len(v)
		delete(c.items, key)

	} else {
		c.misses += 1
	}
	return found
}
