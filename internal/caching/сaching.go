package caching

import (
	"sync"
	"time"
)

var cache sync.Map

func GetCache(key int) (interface{}, bool) {
	val, ok := cache.Load(key)
	return val, ok
}

func SetCache(key int, value interface{}, ttl time.Duration) {
	cache.Store(key, value)
	go func() {
		time.Sleep(ttl)
		cache.Delete(key)
	}()
}
