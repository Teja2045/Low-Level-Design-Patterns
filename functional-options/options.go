package functionalOptions

type OptFunc func(opts *Opts)

type Opts struct {
	id               string
	maxConns         int
	tls              bool
	timeoutInSeconds int
}

func DefaultOpts() Opts {
	return Opts{
		id:               "default",
		maxConns:         1,
		timeoutInSeconds: 10,
		tls:              false,
	}
}

func WithCustomID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

func WithMaxConns(maxConns int) OptFunc {
	return func(opts *Opts) {
		opts.maxConns = maxConns
	}
}

func WithTimoutInSeconds(timeout int) OptFunc {
	return func(opts *Opts) {
		opts.timeoutInSeconds = timeout
	}
}

func WithTLS(tls bool) OptFunc {
	return func(opts *Opts) {
		opts.tls = tls
	}
}
