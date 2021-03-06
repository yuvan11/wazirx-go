package wazirx_client

import (
	"context"
	"encoding/json"

	"net/http"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

type TickersService struct {
	C *Client
}

type Tickers struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
	OpenPrice  string `json:"openPrice"`
	LowPrice   string `json:"lowPrice"`
	HighPrice  string `json:"highPrice"`
	LastPrice  string `json:"lastPrice"`
	Volume     string `json:"volume"`
	BidPrice   string `json:"bidPrice"`
	AskPrice   string `json:"askPrice"`
	At         int64  `json:"at"`
}

func (s *TickersService) DisplayTickersData(ctx context.Context) (*[]Tickers, error) {

	r := &Request{
		Method:   http.MethodGet,
		Endpoint: api.TICKERS,
		SecType:  SecTypeNone,
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		return &[]Tickers{}, err
	}

	result := new([]Tickers)

	if err := json.Unmarshal(data, &result); err != nil {
		return &[]Tickers{}, err
	}

	return result, nil
}
