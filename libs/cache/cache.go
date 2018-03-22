package cache

import (
	"sync"
	"time"
)

type ICache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, ttl time.Duration)
	Remove(key string)
	Keys() []string
}

var _ ICache = &cache{}

type cache struct {
	data map[string]cacheItem
	mu sync.RWMutex
	tick time.Ticker
}

func NewCache(garbageTick time.Duration) ICache {
	c := &cache{}

	return c
}

func (c *cache) Get(key string) (interface{}, bool)  {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if v, ok := c.data[key]; ok && !v.Expired() {
		return v.item, true
	} else {
		return nil, false
	}
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheItem{
		item: value,
		exp: time.Now().Add(ttl).UnixNano(),
	}
}

func (c *cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

func (c *cache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]string, 0, len(c.data))
	for key, _ := range c.data {
		result = append(result, key)
	}

	return result
}

type cacheItem struct {
	item interface{}
	exp  int64
}

func (ci cacheItem) Expired() bool {
	if ci.exp == 0 {
		return false
	}

	return time.Now().UnixNano() > ci.exp
}