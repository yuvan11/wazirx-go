package wazirx_client

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

func TestPingService_FetchPingStatus(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *PingService
		args    args
		wantRes string
		WantErr bool
	}{
		{
			name: `Ping Service Test Positive`,
			s: &PingService{C: &Client{
				ApiKey:     "",
				Secretkey:  "",
				BaseURL:    api.BASE_URL,
				HTTPClient: http.Client{},
				do: func(req *http.Request) (*http.Response, error) {
					return *&req.Response, nil
				},
			}},
			args: args{
				ctx: context.TODO(),
			},
			wantRes: "{ }",
			WantErr: false,
		},
		{
			name: `Ping Service Test Negative(removing BASEURL)`,
			s: &PingService{C: &Client{
				ApiKey:     "",
				Secretkey:  "",
				BaseURL:    "",
				HTTPClient: http.Client{},
				do: func(req *http.Request) (*http.Response, error) {
					return *&req.Response, nil
				},
			}},
			args: args{
				ctx: context.TODO(),
			},
			wantRes: "",
			WantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes, gotErr := tt.s.FetchPingStatus(tt.args.ctx); gotRes != tt.wantRes && gotErr != nil {
				t.Errorf("PingService.FetchPingStatus() = %v, want %v", gotRes, tt.wantRes)
			}
			if gotRes, gotErr := tt.s.FetchPingStatus(tt.args.ctx); gotRes == tt.wantRes && gotErr == nil {
				t.Errorf("PingService.FetchPingStatus() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestServerTimeService_FetchServerTime(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *ServerTimeService
		args    args
		want    *ServerTime
		wantErr bool
	}{
		{
			name: "Server Time Test +ve",
			s: &ServerTimeService{
				C: &Client{
					ApiKey:     "",
					Secretkey:  "",
					BaseURL:    api.TIME,
					HTTPClient: http.Client{},
					do: func(req *http.Request) (*http.Response, error) {
						return *&req.Response, nil
					},
				},
			},
			args: args{
				ctx: context.TODO(),
			},
			want: &ServerTime{
				ServerTime: time.Now().Unix(),
			},
			wantErr: false,
		},
		{
			name: "Server Time Test -ve(Removing BASE URL)",
			s: &ServerTimeService{
				C: &Client{
					ApiKey:     "",
					Secretkey:  "",
					BaseURL:    "",
					HTTPClient: http.Client{},
					do: func(req *http.Request) (*http.Response, error) {
						return *&req.Response, nil
					},
				},
			},
			args: args{
				ctx: context.TODO(),
			},
			want: &ServerTime{
				ServerTime: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := tt.s.FetchServerTime(tt.args.ctx)
			if gotErr == nil {
				if err := gotErr != nil; err != tt.wantErr {
					t.Errorf("ServerTimeService.FetchServerTime() = %v, want %v", err, tt.wantErr)
				}
			}
			if gotErr != nil {
				if err := gotErr == nil; err == tt.wantErr {
					t.Errorf("ServerTimeService.FetchServerTime() = %v, want %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestSystemStatusService_FetchSystemStatus(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *SystemStatusService
		args    args
		want    *SystemStatus
		wantErr bool
	}{
		{
			// TODO: Add test cases.
			name: `System Status Test`,
			s: &SystemStatusService{
				C: &Client{
					ApiKey:     "",
					Secretkey:  "",
					BaseURL:    api.BASE_URL,
					HTTPClient: http.Client{},
					do: func(req *http.Request) (*http.Response, error) {
						return *&req.Response, nil
					},
				}},
			args: args{
				ctx: context.TODO(),
			},
			want: &SystemStatus{
				Status:  "normal",
				Message: "System is running normally.",
			},
			wantErr: false,
		},
		{
			// TODO: Add test cases.
			name: `System Status Test Negative`,
			s: &SystemStatusService{
				C: &Client{
					ApiKey:     "",
					Secretkey:  "",
					BaseURL:    "",
					HTTPClient: http.Client{},
					do: func(req *http.Request) (*http.Response, error) {
						return *&req.Response, nil
					},
				}},
			args: args{
				ctx: context.TODO(),
			},
			want: &SystemStatus{
				Status:  "",
				Message: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.s.FetchSystemStatus(tt.args.ctx)
			if gotErr == nil {
				if err := gotErr != nil; err != tt.wantErr {
					t.Errorf("SystemStatusService.FetchSystemStatus() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("SystemStatusService.FetchSystemStatus() = %v, want %v", got, tt.want)
				}
			} else {
				if err := gotErr == nil; err == tt.wantErr {
					t.Errorf("SystemStatusService.FetchSystemStatus() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}

		})
	}
}

func TestExchangeInfoService_FetchExchangeInfo(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *ExchangeInfoService
		args    args
		wantErr bool
	}{
		{
			name: "Exchange Info Test +ve",
			s: &ExchangeInfoService{
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
			wantErr: false,
		},
		{
			name: "Exchange Info Test -ve(Removing base URL)",
			s: &ExchangeInfoService{
				C: &Client{
					ApiKey:     "",
					Secretkey:  "",
					BaseURL:    "",
					HTTPClient: http.Client{},
					do: func(req *http.Request) (*http.Response, error) {
						return &http.Response{}, nil
					},
				},
			},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := tt.s.FetchExchangeInfo(tt.args.ctx)
			if gotErr == nil {
				if err := gotErr != nil; err != tt.wantErr {
					t.Errorf("ExchangeInfoService.FetchExchangeInfo() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if gotErr != nil {
				if err := gotErr == nil; err == tt.wantErr {
					t.Errorf("ExchangeInfoService.FetchExchangeInfo() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

		})
	}
}
