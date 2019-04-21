package bitly

import (
	"sync"
)

type cache struct {
	mu   sync.RWMutex
	URLs map[string]string
}

func newCache() *cache {
	return &cache{
		URLs: make(map[string]string),
	}
}

func (c *cache) add(longURL string, shortURL string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.URLs[longURL] = shortURL
}

func (c *cache) get(longURL string) (string, bool) {
	c.mu.RLock()
	shortURL, found := c.URLs[longURL]
	c.mu.RUnlock()
	return shortURL, found
}
