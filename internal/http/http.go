package http

import (
	"net"
	"net/http"

	"go.uber.org/zap"
)

const swaggerDefaultEndpoint = `/swagger.json`

type HTTPServer struct {
	mux    *http.ServeMux
	h      *http.Server
	addr   string
	logger *zap.Logger
}

// NewHTTPServer - http server which can be used for instumental things like:
// metrics, buildifo, debug
func NewHTTPServer(address string, logger *zap.Logger) (*HTTPServer, error) {
	hserver := &http.Server{}
	mux := http.NewServeMux()

	return &HTTPServer{
		mux:    mux,
		h:      hserver,
		addr:   address,
		logger: logger.Named("http-server"),
	}, nil
}

func (s *HTTPServer) Handle(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

func (s *HTTPServer) AddSwagger(swaggerBytes []byte) {
	s.mux.HandleFunc(swaggerDefaultEndpoint, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(swaggerBytes)
	})
}

func (s *HTTPServer) Start() (func() error, error) {
	s.logger.Info("starting", zap.String("addr", s.addr))
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return nil, err
	}
	s.h.Handler = s.mux

	// och := &ochttp.Handler{
	// 	Handler: s.h.Handler,
	// }
	// s.Handle("/", och)

	go func() {
		err = s.h.Serve(l)
		if err != nil {
			s.logger.Error("http server", zap.Error(err))
		}
	}()

	return func() error {
		s.logger.Info("shutting down")
		return s.h.Close()
	}, nil

}
