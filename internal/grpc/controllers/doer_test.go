package controllers_test

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/bogatyr285/golang-boilerplate/api/v1/doer"
	"github.com/bogatyr285/golang-boilerplate/internal/grpc/controllers"
	"github.com/bogatyr285/golang-boilerplate/mocks"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

func TestDoAwesome(t *testing.T) {
	logger := zap.NewExample()
	doerServiceMock := mocks.NewMockDoerService(gomock.NewController(t))

	tests := []struct {
		name    string
		prepare func(m *mocks.MockDoerService)
		req     *pb.DoAwesomeRequest
		wantErr bool
	}{
		{
			name: "should process sucessfully",
			prepare: func(m *mocks.MockDoerService) {
				m.EXPECT().Do(gomock.Any(), gomock.Any()).Return("ok", nil)
			},
			req: &pb.DoAwesomeRequest{
				Input: "input_ok",
			},
			wantErr: false,
		},
		{
			name:    "should return validation error",
			prepare: nil,
			req: &pb.DoAwesomeRequest{
				Input: "in", // too short input
			},
			wantErr: true,
		},
		{
			name: "should return internal error",
			prepare: func(m *mocks.MockDoerService) {
				m.EXPECT().Do(gomock.Any(), gomock.Any()).Return("", fmt.Errorf("oh, no. service error"))
			},
			req: &pb.DoAwesomeRequest{
				Input: "input",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			if tt.prepare != nil {
				tt.prepare(doerServiceMock)
			}

			doerController := controllers.NewDoerController(doerServiceMock, logger)
			ctx := context.Background()
			if _, err := doerController.DoAwesome(ctx, tt.req); (err != nil) != tt.wantErr {
				t.Errorf("DoAwesome() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
