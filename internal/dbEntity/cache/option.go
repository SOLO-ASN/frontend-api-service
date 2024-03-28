package cache

import "time"

type OptionFn func(*option)

type option struct {
	EnableTrace  bool
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func defaultOption() *option {
	return &option{
		EnableTrace:  false,
		DialTimeout:  30 * time.Second,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (o *option) apply(opts ...OptionFn) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithEnableTrace(enable bool) OptionFn {
	return func(opt *option) {
		opt.EnableTrace = enable
	}
}

func WithDialTimeout(timeout time.Duration) OptionFn {
	return func(opt *option) {
		opt.DialTimeout = timeout
	}
}

func WithReadTimeout(timeout time.Duration) OptionFn {
	return func(opt *option) {
		opt.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) OptionFn {
	return func(opt *option) {
		opt.WriteTimeout = timeout
	}
}
