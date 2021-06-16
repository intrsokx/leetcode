package ratelimiter

type option struct {
	qps  int64
	conf *RateConf
}

type RateConf struct {
}

type Option func(opt *option)

func WithQps(qps int64) Option {
	return func(opt *option) {
		opt.qps = qps
	}
}

func WithRateConf(conf *RateConf) Option {
	return func(opt *option) {
		opt.conf = conf
	}
}
