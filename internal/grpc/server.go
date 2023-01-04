package grpc

import (
	"context"
	"fmt"
	"net"
	"runtime"
	"time"

	pb "github.com/bogatyr285/golang-boilerplate/api/v1/doer"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_middlewarerecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_middlewaretags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	// GRPCDefaultGracefulStopTimeout - period to wait result of grpc.GracefulStop
	// after call grpc.Stop
	GRPCDefaultGracefulStopTimeout = 5 * time.Second
)

// GRPC - structure describes gRPC props
type Server struct {
	grpcAddr            string
	grpcSrv             *grpc.Server
	listener            net.Listener
	gracefulStopTimeout time.Duration

	logger *zap.Logger
}

func NewServer(
	grpcAddr string,
	doerCtrl pb.DoerAPIServer,
	logger *zap.Logger,
) (*Server, error) {
	logger = logger.Named("grpc-server")
	netListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return nil, err
	}
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpczap.StreamServerInterceptor(logger),
			grpc_middlewaretags.StreamServerInterceptor(),
			grpc_middlewarerecovery.StreamServerInterceptor(),
			StreamValidationInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryServerLogTraceIDDecorator(logger),
			grpczap.UnaryServerInterceptor(logger),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_middlewaretags.UnaryServerInterceptor(),
			grpc_middlewarerecovery.UnaryServerInterceptor(grpc_middlewarerecovery.WithRecoveryHandler(onPanicStackLogger(logger))),
			UnaryValidationInterceptor(),
		)),
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	}

	grpcSrv := grpc.NewServer(opts...)
	pb.RegisterDoerAPIServer(grpcSrv, doerCtrl)

	// register health check service
	healthService := NewHealthChecker(logger)
	grpc_health_v1.RegisterHealthServer(grpcSrv, healthService)

	// Register reflection service on gRPC server. can be a flag
	reflection.Register(grpcSrv)

	//trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	grpc_prometheus.Register(grpcSrv)

	server := &Server{
		grpcAddr:            grpcAddr,
		listener:            netListener,
		grpcSrv:             grpcSrv,
		gracefulStopTimeout: GRPCDefaultGracefulStopTimeout,
		logger:              logger,
	}

	return server, nil
}

func (s *Server) Run() (func() error, error) {
	s.logger.Info("starting", zap.String("grpcAddr", s.grpcAddr))

	go func() {
		err := s.grpcSrv.Serve(s.listener)
		if err == grpc.ErrServerStopped {
			s.logger.Error("grpc server", zap.Error(err))
		}
	}()

	return s.close, nil
}

// stop - gracefully stop server & listeners
func (s *Server) close() error {
	s.logger.Info("gracefully stopping....", zap.String("grpcAddr", s.grpcAddr))

	stopped := make(chan struct{})
	go func() {
		s.grpcSrv.GracefulStop()
		close(stopped)
	}()

	t := time.NewTimer(s.gracefulStopTimeout)
	select {
	case <-t.C:
		s.logger.Info("ungracefully stopping....", zap.String("grpcAddr", s.grpcAddr))
		s.grpcSrv.Stop()
	case <-stopped:
		t.Stop()
	}
	s.logger.Info("stopped", zap.String("grpcAddr", s.grpcAddr))
	return nil
}

/* util middlewares */
func withTraceID(ctx context.Context, logger *zap.Logger) context.Context {
	rootTraceID := ""
	rootSpan := trace.FromContext(ctx)
	if rootSpan != nil {
		rootTraceID = rootSpan.SpanContext().TraceID.String()
	}

	ctxzap.AddFields(ctx, zap.String("trace_id", rootTraceID))

	return ctx
}

// adds trace_id to logger output
func unaryServerLogTraceIDDecorator(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(withTraceID(ctx, logger), req)
	}
}

// add stack if panic happened
func onPanicStackLogger(logger *zap.Logger) func(p interface{}) error {
	return func(p interface{}) error {
		const maxStacksize = 2 * 1024

		stack := make([]byte, maxStacksize)
		stack = stack[:runtime.Stack(stack, true)]
		// keep a multiline stack
		logger.Error("panic happened", zap.Any("error", p), zap.ByteString("stack", stack))

		return fmt.Errorf("internal error. we've sent group to handle it. %v", p)
	}
}
