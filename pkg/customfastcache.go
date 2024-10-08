package cachetest

import (
	"sync"
)

// This is my very simple cache. I was surprised how well it works.
type CustomFastCache struct {
	sync.RWMutex
	items      map[string][]byte
	totalsizeB int
	maxsizeB   int
}

func NewCustomFastCache(maxsizeMB int) *CustomFastCache {
	return &CustomFastCache{
		items:      make(map[string][]byte),
		totalsizeB: 0,
		maxsizeB:   maxsizeMB * 1024 * 1024,
	}
}

func (c *CustomFastCache) Set(key string, value []byte) {
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

func (c *CustomFastCache) Get(key string) ([]byte, bool) {
	c.RLock()
	defer c.RUnlock()
	item, found := c.items[key]

	return item, found
}

func (c *CustomFastCache) Del(key string) bool {
	c.Lock()
	defer c.Unlock()
	v, found := c.items[key]
	if found {
		c.totalsizeB -= len(v)
		delete(c.items, key)

	}
	return found
}
