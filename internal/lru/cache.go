package lru

import (
	"sync"
	"time"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=Cache

type Cache interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, bool)
}

type MyCache struct {
	sync.RWMutex
	cleanupInterval time.Duration
	items           map[string]Item
}

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

func New(cleanupInterval time.Duration) Cache {
	items := make(map[string]Item)

	cache := MyCache{
		items:           items,
		cleanupInterval: cleanupInterval,
	}

	if cleanupInterval > 0 {
		cache.StartGC()
	}

	return &cache
}

func (c *MyCache) Set(key string, value interface{}, duration time.Duration) {
	var expiration int64

	expiration = time.Now().Add(duration).UnixNano()

	c.Lock()
	defer c.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}
}

func (c *MyCache) Get(key string) (interface{}, bool) {

	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}

	return item.Value, true
}

func (c *MyCache) StartGC() {
	go c.GC()
}

func (c *MyCache) GC() {
	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

func (c *MyCache) expiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

func (c *MyCache) clearItems(keys []string) {
	c.Lock()
	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
