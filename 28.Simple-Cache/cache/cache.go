package cache

import (
    "sync"
    "time"
)

type Item struct {
    Value      interface{}
    Expiration int64
}

type Cache struct {
    items map[string]Item
    mut    sync.Mutex
}

func NewCache() *Cache {
    return &Cache{
        items: make(map[string]Item),
    }
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
    c.mut.Lock()
    defer c.mut.Unlock()
    expiration := time.Now().Add(duration).UnixNano()
    c.items[key] = Item{Value: value, Expiration: expiration}
    go func() {
        time.Sleep(duration)
        c.Delete(key)
    }()
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mut.Lock()
    defer c.mut.Unlock()
    item, found := c.items[key]
    if !found {
        return nil, false
    }
    if time.Now().UnixNano() > item.Expiration {
        delete(c.items, key)
        return nil, false
    }
    return item.Value, true
}

func (c *Cache) Delete(key string) {
    c.mut.Lock()
    defer c.mut.Unlock()
    delete(c.items, key)
}
