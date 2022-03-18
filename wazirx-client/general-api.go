package wazirx_client

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

// PING SERVICE

type PingService struct {
	C *Client
}

func (s *PingService) FetchPingStatus(ctx context.Context) (res string, err error) {

	r := &Request{
		Method:   http.MethodGet,
		Endpoint: api.PING,
		SecType:  SecTypeNone,
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		return "", err
	}
	res = string(data)

	return res, nil
}

// SERVER TIME

type ServerTimeService struct {
	C *Client
}
type ServerTime struct {
	ServerTime int64 `json:"serverTime"`
}

// SYSTEM STATUS

type SystemStatusService struct {
	C *Client
}

type SystemStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (s *SystemStatusService) FetchSystemStatus(ctx context.Context) (*SystemStatus, error) {

	r := &Request{
		Method:   http.MethodGet,
		Endpoint: api.SYSTEM_STATUS,
		SecType:  SecTypeNone,
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		return &SystemStatus{}, err
	}

	result := new(SystemStatus)

	if err := json.Unmarshal(data, &result); err != nil {
		return &SystemStatus{}, err
	}

	return result, nil
}

func (s *ServerTimeService) FetchServerTime(ctx context.Context) (*ServerTime, error) {

	r := &Request{
		Method:   http.MethodGet,
		Endpoint: api.TIME,
		SecType:  SecTypeNone,
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		fmt.Println(err)
	}

	result := new(ServerTime)

	if err := json.Unmarshal(data, &result); err != nil {
		return &ServerTime{}, err
	}

	return result, nil
}

//  EXCHANGE INFORMATION

type Filters struct {
	FilterType string `json:"filterType"`
	MinPrice   string `json:"minPrice"`
	TickSize   string `json:"tickSize"`
}

type Symbols struct {
	Symbol               string    `json:"symbol"`
	Status               string    `json:"status"`
	BaseAsset            string    `json:"baseAsset"`
	QuoteAsset           string    `json:"quoteAsset"`
	BaseAssetPrecision   int32     `json:"baseAssetPrecision"`
	QuoteAssetPrecision  int32     `json:"quoteAssetPrecision"`
	OrderTypes           []string  `json:"orderTypes"`
	IsSpotTradingAllowed bool      `json:"isSpotTradingAllowed"`
	Filters              []Filters `json:"filters"`
}

type ExchangeInfo struct {
	TimeZone   string    `json:"timezone"`
	ServerTIme int64     `json:"serverTime"`
	Symbols    []Symbols `json:"symbols"`
}

type ExchangeInfoService struct {
	C *Client
}

func (s *ExchangeInfoService) FetchExchangeInfo(ctx context.Context) *ExchangeInfo {

	r := &Request{
		Method:   http.MethodGet,
		Endpoint: api.EXCHAGE_INFO,
		SecType:  SecTypeNone,
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		fmt.Println(err)
	}

	result := new(ExchangeInfo)

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Can not unmarshal JSON", err)
	}

	return result
}
