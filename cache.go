package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Cache .
type Cache struct {
	redisclient *redis.Client
	conf        *Config
}

// NewCache create a new cache
func NewCache(r *redis.Client, c *Config) *Cache {
	var conf2 = Config{
		DefaultTTL:  time.Second * 5,
		CachePrefix: "",
	}
	if c != nil {
		conf2.CachePrefix = c.CachePrefix
		conf2.DefaultTTL = c.DefaultTTL
	}
	return &Cache{
		redisclient: r,
		conf:        &conf2,
	}
}

// Set set value
func (c *Cache) Set(ctx context.Context, key, value string, vs ...interface{}) error {
	cachekey := fmt.Sprintf("%s%s", c.conf.CachePrefix, key)
	var ttl = c.conf.DefaultTTL
	for _, v := range vs {
		switch v.(type) {
		case KeyTTL:
			ttl, _ = v.(KeyTTL)
		}
	}
	_, err := c.redisclient.Set(ctx, cachekey, value, ttl).Result()
	return err
}

// Get get value
func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	cachekey := fmt.Sprintf("%s%s", c.conf.CachePrefix, key)
	result, err := c.redisclient.Get(ctx, cachekey).Result()
	if errors.Is(err, redis.Nil) {
		return result, ErrorNil
	}
	return result, err
}

// Remove remove value
func (c *Cache) Remove(ctx context.Context, key string) (int64, error) {
	cachekey := fmt.Sprintf("%s%s", c.conf.CachePrefix, key)
	return c.redisclient.Del(ctx, cachekey).Result()
}
