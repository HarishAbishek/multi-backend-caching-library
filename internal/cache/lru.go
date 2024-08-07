package cache

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type lruItem struct {
	key        string
	value      interface{}
	expiration int64
}

type LRUCache struct {
	mu         sync.Mutex
	items      map[string]*list.Element
	evictionList *list.List
	maxSize    int
}

func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		items:      make(map[string]*list.Element),
		evictionList: list.New(),
		maxSize:    maxSize,
	}
}

func (c *LRUCache) Set(key string, value interface{}, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		c.evictionList.MoveToFront(item)
		item.Value.(*lruItem).value = value
		item.Value.(*lruItem).expiration = time.Now().Add(ttl).UnixNano()
		return nil
	}

	if c.evictionList.Len() >= c.maxSize {
		c.evict()
	}

	item := &lruItem{
		key:        key,
		value:      value,
		expiration: time.Now().Add(ttl).UnixNano(),
	}
	entry := c.evictionList.PushFront(item)
	c.items[key] = entry
	return nil
}

func (c *LRUCache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		if item.Value.(*lruItem).expiration > time.Now().UnixNano() {
			c.evictionList.MoveToFront(item)
			return item.Value.(*lruItem).value, nil
		}
		c.remove(item)
	}
	return nil, errors.New("cache miss")
}

func (c *LRUCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		c.remove(item)
		return nil
	}
	return errors.New("key not found")
}

func (c *LRUCache) evict() {
	item := c.evictionList.Back()
	if item != nil {
		c.remove(item)
	}
}

func (c *LRUCache) remove(item *list.Element) {
	c.evictionList.Remove(item)
	delete(c.items, item.Value.(*lruItem).key)
}
