package cache

import (
	"context"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
)

var _ ICache = (*Client)(nil)

type Client struct {
	rdb *redis.Client
}

func (c Client) Set(key string, value interface{}) error {
	return c.rdb.Set(context.Background(), key, value, 0).Err()

}

func (c Client) Get(key string) (interface{}, error) {
	return c.rdb.Get(context.Background(), key).Result()
}

func NewRedisClient(dsn string, opts ...OptionFn) *Client {
	// generate redis options
	opt := defaultOption()
	opt.apply(opts...)

	o, err := genRedisOptions(dsn, opt)
	if err != nil {
		return nil
	}
	rdb := redis.NewClient(o)

	if opt.EnableTrace {
		rdb.AddHook(redisotel.TracingHook{})
	}

	return &Client{rdb: rdb}
}

func NewRedisClientWithPassword(address string, password string, db int, opts ...OptionFn) *Client {
	// generate redis options
	opt := defaultOption()
	opt.apply(opts...)

	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return &Client{rdb: rdb}
}

func genRedisOptions(dsn string, opt *option) (*redis.Options, error) {
	rOpt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}
	if opt.DialTimeout > 0 {
		rOpt.DialTimeout = opt.DialTimeout
	}

	if opt.ReadTimeout > 0 {
		rOpt.ReadTimeout = opt.ReadTimeout
	}

	if opt.WriteTimeout > 0 {
		rOpt.WriteTimeout = opt.WriteTimeout
	}

	rOpt.DB = opt.Db

	return rOpt, nil
}
