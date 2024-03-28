package cache

import (
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
)

var _ ICache = (*Client)(nil)

type Client struct {
	rdb *redis.Client
}

func (c Client) Set(key string, value interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) Get(key string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func NewRedisClient(dsn string, opt *option) *Client {
	// todo implement me

	// generate redis options
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

func NewRedisClientWithPassword(address string, password string, db int, opt *option) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     password,
		DB:           db,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
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

	return rOpt, nil
}
