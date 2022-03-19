package wazirx_client

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

func TestAccountInfoService_FetchAccountInfo(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *AccountInfoService
		args    args
		want    *AccountInfo
		wantErr bool
	}{
		{
			name: "Account Info Service Test",
			s: &AccountInfoService{
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
			want: &AccountInfo{
				AccountType: "default",
				CanTrade:    true,
				CanWithdraw: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FetchAccountInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountInfoService.FetchAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.AccountType, tt.want.AccountType) {
				t.Errorf("AccountInfoService.FetchAccountInfo() = %v, want %v", got.AccountType, tt.want.AccountType)
			}
		})
	}
}
