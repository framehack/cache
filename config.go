package cache

import "time"

// Config cache config
type Config struct {
	DefaultTTL  time.Duration // default ttl
	CachePrefix string        // cache key prefix
}

// KeyTTL key ttl
type KeyTTL = time.Duration
