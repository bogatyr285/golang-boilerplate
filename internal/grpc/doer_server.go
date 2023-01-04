package grpc

import (
	"context"

	pb "github.com/bogatyr285/golang-boilerplate/api/v1/doer"
	"go.uber.org/zap"
)

type DoerGRPCServer struct {
	dpCtrl pb.DoerAPIServer
	logger *zap.Logger
	*pb.UnimplementedDoerAPIServer
}

func NewDoerGRPCServer(dpCtrl pb.DoerAPIServer, logger *zap.Logger) *DoerGRPCServer {
	return &DoerGRPCServer{dpCtrl: dpCtrl, logger: logger}
}

func (c *DoerGRPCServer) DoAwesome(ctx context.Context, in *pb.DoAwesomeRequest) (*pb.DoAwesomeResponse, error) {
	return c.dpCtrl.DoAwesome(ctx, in)
}
