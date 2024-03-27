package server

import "time"

type HTTPOptionFunc func(options *httpOptions)

type httpOptions struct {
	readTimeout  time.Duration
	writeTimeout time.Duration

	// three modes, which are: debug, test, release
	mode string
}

func defaultHTTPOptions() *httpOptions {
	return &httpOptions{
		readTimeout:  60 * time.Second,
		writeTimeout: 60 * time.Second,
		mode:         "debug",
	}
}

func (ho *httpOptions) applyHTTPOptions(opts []HTTPOptionFunc) {
	for _, opt := range opts {
		opt(ho)
	}
}

func WithReadTimeout(timeout time.Duration) HTTPOptionFunc {
	return func(opt *httpOptions) {
		opt.readTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) HTTPOptionFunc {
	return func(opt *httpOptions) {
		opt.writeTimeout = timeout
	}
}

func WithMode(mode string) HTTPOptionFunc {
	return func(opt *httpOptions) {
		// todo: check mode
		opt.mode = mode
	}
}
