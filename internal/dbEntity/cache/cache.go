package cache

import (
	"errors"
	"github.com/go-redis/redis/v8"
)

var (
	// ErrCacheGenerate returns when cache generate error occurs.
	ErrCacheGenerate = errors.New("cache generate error")

	// ErrCacheTypeMismatch returns when cache type mismatch occurs.
	ErrCacheTypeMismatch = errors.New("cache type mismatch")

	// ErrCacheNotFound returns cache not found, or use redis.Nil??
	ErrCacheNotFound = errors.New("cache not found")
)

type ICache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type Cache struct {
	CacheType string
	Rdb       *redis.Client

	// todo add memory cache
	// MemoryDb
}

func Init(dsn string, cType string, opts ...OptionFn) (ICache, error) {
	opt := defaultOption()
	opt.apply(opts...)

	if cType == "redis" {
		rdb := NewRedisClient(dsn, opt)
		if rdb == nil {
			return nil, ErrCacheGenerate
		}
		return rdb, nil
	} else if cType == "memory" {
		// todo implement me
		panic("implement me")
	} else {
		return nil, ErrCacheTypeMismatch
	}
}
