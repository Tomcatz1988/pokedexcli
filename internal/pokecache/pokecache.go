package pokecache

import (
	"fmt"
	"sync"
	"time"
)


type Cache struct {
	entries map[string]cacheEntry
	mux sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


func NewCache (interval time.Duration) Cache {
	entries := make(map[string]cacheEntry)
	mux := sync.RWMutex{}
	cache := Cache{
		entries,
		mux,
	}
	go cache.reapLoop(interval)
	return cache
}


func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	toCull := make([]string, 0)
	for tick := range(ticker.C) {
		cache.mux.RLock()
		for key, entry := range(cache.entries) {
			elapsed := tick.Sub(entry.createdAt)
			if elapsed > interval {
				toCull = append(toCull, key)
			}
		}
		cache.mux.RUnlock()
		cache.mux.Lock()
		for _, key := range (toCull) {
			delete(cache.entries, key)
		}
		cache.mux.Unlock()
		toCull = toCull[:0]
	}
}


func (cache *Cache) Add(key string, val []byte) {
	cache.mux.Lock()
	cache.entries[key] = cacheEntry {
		time.Now(),
		val,
	}
	cache.mux.Unlock()
	fmt.Println("result cached")
}

func (cache *Cache) Get(key string) (val []byte, exists bool) {
	cache.mux.RLock()
	entry, exists := cache.entries[key]
	if exists {
		val = entry.val
	} 
	cache.mux.RUnlock()
	return val, exists
}
