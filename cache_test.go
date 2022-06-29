package cache

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestCache(t *testing.T) {
	r := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	c := NewCache(r, &Config{
		DefaultTTL:  time.Second * 10,
		CachePrefix: "prefix",
	})
	ctx := context.Background()
	key := "test"
	value := "value"
	_, err := c.Get(ctx, key)
	if !errors.Is(err, ErrorNil) {
		t.Fatal("key exists")
	}
	err = c.Set(ctx, key, value, time.Second*5)
	if err != nil {
		t.Fatal(err)
	}
	res, err := c.Get(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	if res != value {
		t.Fatal("not equal")
	}
}
