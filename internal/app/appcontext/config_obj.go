package appcontext

import (
	"time"
)

// CacheItem defines object for cache item collection (cache keys & TTLs)
type CacheItem struct {
	// UserByID                 CacheData
}

// CacheData represent data for each cache item
type CacheData struct {
	Key string
	TTL time.Duration
}
