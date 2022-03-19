package wazirx_client

import (
	"context"
	"net/http"
	"testing"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

func TestAuthInfoService_CreateAuthToken(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *AuthInfoService
		args    args
		wantErr bool
	}{
		{
			name: "AuthInfo Service Test",
			s: &AuthInfoService{
				C: &Client{
					ApiKey:     api.API_KEY,
					Secretkey:  api.SECRET_KEY,
					BaseURL:    api.BASE_URL,
					HTTPClient: http.Client{},
					do: func(req *http.Request) (*http.Response, error) {
						return &http.Response{}, nil
					},
				},
			},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.CreateAuthToken(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthInfoService.CreateAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
