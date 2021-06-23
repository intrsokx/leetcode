package httputil

import (
	"net/url"
	"time"
)

type option struct {
	timeout time.Duration
	proxy   *url.URL
}

type Option func(opt *option)

func WithProxy(proxy *url.URL) Option {
	return func(opt *option) {
		opt.proxy = proxy
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opt *option) {
		opt.timeout = timeout
	}
}
