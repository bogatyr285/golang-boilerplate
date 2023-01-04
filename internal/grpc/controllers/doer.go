//go:generate mockgen -source=./doer.go -destination=../../../mocks/doer_svc_mock.go -package=mocks
package controllers

import (
	"context"

	pb "github.com/bogatyr285/golang-boilerplate/api/v1/doer"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DoerService interface {
	Do(ctx context.Context, input string) (string, error)
}

type DoerController struct {
	doerService DoerService
	logger      *zap.Logger
	*pb.UnimplementedDoerAPIServer
}

func NewDoerController(doerService DoerService, logger *zap.Logger) *DoerController {
	return &DoerController{doerService: doerService, logger: logger.Named("doer-grpc-ctrl")}
}

func (c *DoerController) DoAwesome(ctx context.Context, in *pb.DoAwesomeRequest) (*pb.DoAwesomeResponse, error) {
	ctx, span := trace.StartSpan(ctx, "v1.doer.grpc.doawesome")
	defer span.End()
	if err := in.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error")
	}

	r, err := c.doerService.Do(ctx, in.Input)
	if err != nil {
		// to log/trace full err with details, in response just common message
		span.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return nil, status.Errorf(codes.Internal, "couldn't do awesome. we're working to fix it")
	}

	return &pb.DoAwesomeResponse{Msg: r}, nil
}
