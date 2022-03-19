package wazirx_client

import (
	"context"
	"net/http"
	"testing"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

func TestHistoricalTradesService_DisplayHistoricalTrade(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *HistoricalTradesService
		args    args
		wantErr bool
	}{
		{
			name: "Historical Trade Service",
			s: &HistoricalTradesService{
				C: &Client{
					ApiKey:     api.API_KEY,
					Secretkey:  api.SECRET_KEY,
					BaseURL:    api.BASE_URL,
					HTTPClient: http.Client{},
					do: func(req *http.Request) (*http.Response, error) {
						return &http.Response{}, nil
					},
				},
				Symbol: "manausdt",
				Limit:  10,
			},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DisplayHistoricalTrade(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("HistoricalTradesService.DisplayHistoricalTrade() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if int32(len(*got)) != (tt.s.Limit) {
				t.Errorf("HistoricalTradesService.DisplayHistoricalTrade() = %v, want %v", int32(len(*got)), tt.s.Limit)
			}
		})
	}
}
