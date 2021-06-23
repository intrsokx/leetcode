package httputil

import (
	"net/http"
)

type HttpUtil struct {
	http.Client
}

//在new一个新的中间件时，采用闭包+多个可选参数的方式来实现中间件的灵活性和扩展性
func NewHttpUtil(opts ...Option) *HttpUtil {
	cfg := &option{}
	for _, opt := range opts {
		opt(cfg)
	}

	var client http.Client
	client = http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	if cfg.proxy != nil {
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(cfg.proxy),
		}
	}
	if cfg.timeout != 0 {
		client.Timeout = cfg.timeout
	}

	return &HttpUtil{client}
}
