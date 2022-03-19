package wazirx_client

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

func TestTickerService_TickerData(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *TickerService
		args    args
		want    *Ticker
		wantErr bool
	}{
		{
			name: "Ticker Data Service Test",
			s: &TickerService{
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
			},
			args: args{
				ctx: context.TODO(),
			},
			want: &Ticker{
				Symbol:     "manausdt",
				BaseAsset:  "mana",
				QuoteAsset: "usdt",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.TickerData(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TickerService.TickerData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Symbol, tt.want.Symbol) &&
				!reflect.DeepEqual(got.BaseAsset, tt.want.BaseAsset) &&
				!reflect.DeepEqual(got.QuoteAsset, tt.want.QuoteAsset) {
				t.Errorf("TickerService.TickerData() = %v, want %v", got, tt.want)
			}

		})
	}
}
