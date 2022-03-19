package wazirx_client

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

func TestTickersService_DisplayTickersData(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *TickersService
		args    args
		want    *[]Tickers
		wantErr bool
	}{

		{
			name: "Tickers Service Test",
			s: &TickersService{
				C: &Client{
					ApiKey:     "",
					Secretkey:  "",
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
			want: &[]Tickers{
				{
					Symbol:     "btcinr",
					BaseAsset:  "btc",
					QuoteAsset: "inr",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DisplayTickersData(tt.args.ctx)

			if (err != nil) != tt.wantErr {
				t.Errorf("TickersService.DisplayTickersData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, a := range *got {
				for _, d := range *tt.want {
					if a.Symbol == d.Symbol {
						if !reflect.DeepEqual(a.Symbol, d.Symbol) {
							t.Errorf("TickersService.DisplayTickersData() = %v, want %v", a.Symbol, d.Symbol)
							return
						}
					}

				}

			}

		})
	}
}
