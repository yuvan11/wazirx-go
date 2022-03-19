package wazirx_client

import (
	"context"
	"net/http"
	"testing"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

func TestOrderBookDepthService_DisplayOrdersBookDepth(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		s    *OrderBookDepthService
		args args

		wantErr bool
	}{
		{
			name: "Depth Service Test",
			s: &OrderBookDepthService{
				C: &Client{
					ApiKey:     "",
					Secretkey:  "",
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
			got, err := tt.s.DisplayOrdersBookDepth(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderBookDepthService.DisplayOrdersBookDepth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if int32(len(got.Asks)) != tt.s.Limit {
				t.Errorf("OrderBookDepthService.DisplayOrdersBookDepth() = %v, want %v", len(got.Asks), tt.s.Limit)
			}
		})
	}
}
