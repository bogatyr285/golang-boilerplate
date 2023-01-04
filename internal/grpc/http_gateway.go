package grpc

import (
	"context"
	"net"
	"net/http"

	pb "github.com/bogatyr285/golang-boilerplate/api/v1/doer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

// traceIDHeader - header with traceID will be added to each request to simplify user issues debugging
const traceIDHeader = "X-Trace-ID"

// Gateway - proxy which tranforms HTTP requests to GRPC
type Gateway struct {
	mux        http.Handler
	httpGwAddr string
	logger     *zap.Logger
}

func NewGateway(ctx context.Context, grpcAddr, httpGwAddr string, logger *zap.Logger) (*Gateway, error) {
	gwMux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
		}),
	)

	traceHeaderDecoratedMux := tracing()(gwMux)

	traceDecoratedMux := &ochttp.Handler{Handler: traceHeaderDecoratedMux}

	err := pb.RegisterDoerAPIHandlerFromEndpoint(
		context.Background(),
		gwMux,
		grpcAddr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(new(ocgrpc.ClientHandler)),
		},
	)
	if err != nil {
		return nil, err
	}

	return &Gateway{
		mux:        traceDecoratedMux,
		httpGwAddr: httpGwAddr,
		logger:     logger.Named("grpc/http-gateway"),
	}, nil
}

func (g *Gateway) Start() (func() error, error) {
	hserver := http.Server{
		Handler: g.mux,
	}

	g.logger.Info("starting", zap.String("addr", g.httpGwAddr))
	l, err := net.Listen("tcp", g.httpGwAddr)
	if err != nil {
		return nil, err
	}

	go func() {
		err = hserver.Serve(l)
		if err != nil {
			g.logger.Error("http/grpc gateway server", zap.Error(err))
		}
	}()

	return func() error {
		g.logger.Info("shutting down")
		return hserver.Close()
	}, nil
}

/* tracing middleware */
type tracingOption func(*ochttp.Handler)

func tracing(options ...tracingOption) func(http.Handler) http.Handler {
	handler := &ochttp.Handler{
		FormatSpanName: func(r *http.Request) string {
			return "HTTP " + r.Method + " " + r.URL.Path
		},
	}
	for _, option := range options {
		option(handler)
	}

	return func(next http.Handler) http.Handler {
		handler.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if span := trace.FromContext(r.Context()); span != nil {
				if sc := span.SpanContext(); sc.IsSampled() {
					w.Header().Set(traceIDHeader, sc.TraceID.String())
				}
			}
			next.ServeHTTP(w, r)
		})
		return handler
	}
}
