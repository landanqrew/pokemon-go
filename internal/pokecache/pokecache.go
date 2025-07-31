package pokecache

import (
	"sync"
	"time"

	"github.com/landanqrew/pokemon-go/internal/web"
)

type Cache struct {
	Entries map[string]CacheEntry `json:"entries"`
	mu      sync.RWMutex
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.Entries[key]; !ok {
		c.Entries[key] = CacheEntry{
			CreatedAt: time.Now(),
			Value:     value,
		}
	}

}

func (c *Cache) Get(key string) ([]byte, error) {
	c.mu.RLock()
	entry, ok := c.Entries[key]
	c.mu.RUnlock()

	if ok {
		return entry.Value, nil
	}

	value, err := web.GetResponseBytesBaseUrl(key)
	if err != nil {
		return nil, err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// Double-check in case another goroutine added to cache during the network call time
	if entry, exists := c.Entries[key]; exists {
		return entry.Value, nil
	}

	c.Entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Value:     value,
	}
	return value, nil
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.Entries {
		if entry.CreatedAt.Before(time.Now().Add(interval * -1)) {
			delete(c.Entries, key)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries: make(map[string]CacheEntry),
	}

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			c.reapLoop(interval)
		}
	}()

	return c
}

type CacheEntry struct {
	CreatedAt time.Time `json:"created_at"`
	Value     []byte    `json:"value"`
}
