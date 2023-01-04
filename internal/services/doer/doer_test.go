package doer_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/bogatyr285/golang-boilerplate/internal/models"
	"github.com/bogatyr285/golang-boilerplate/internal/services/doer"
	"github.com/bogatyr285/golang-boilerplate/mocks"
	"github.com/golang/mock/gomock"
)

func TestDo(t *testing.T) {
	doerMock := mocks.NewMockSomethingFetcher(gomock.NewController(t))

	type args struct {
		input string
	}

	tests := []struct {
		name    string
		prepare func(m *mocks.MockSomethingFetcher)
		args    args
		wantErr bool
	}{
		{
			name: "should do everything right",
			prepare: func(m *mocks.MockSomethingFetcher) {
				m.EXPECT().GetSomething(gomock.Any(), gomock.Any()).Return([]*models.Something{}, nil)
			},
			args:    args{input: "test_input"},
			wantErr: false,
		},
		{
			name: "should fail because db err",
			prepare: func(m *mocks.MockSomethingFetcher) {
				m.EXPECT().GetSomething(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("oh, no. DB error"))
			},
			args:    args{input: "test_input"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			if tt.prepare != nil {
				tt.prepare(doerMock)
			}

			doerService := doer.NewDoer(doerMock)
			ctx := context.Background()
			if _, err := doerService.Do(ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
