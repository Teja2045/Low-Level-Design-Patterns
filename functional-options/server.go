package functionalOptions

type Server struct {
	Opts Opts
}

func NewServer(options ...OptFunc) *Server {
	opts := DefaultOpts()
	for _, option := range options {
		option(&opts)
	}

	return &Server{
		Opts: opts,
	}
}
