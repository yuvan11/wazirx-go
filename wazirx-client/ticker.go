package wazirx_client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

type TickerService struct {
	C      *Client
	Symbol string
}

type Ticker struct {
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

func (s *TickerService) SetSymbol(symbol string) *TickerService {
	s.Symbol = symbol
	return s
}

func (s *TickerService) TickerData(ctx context.Context) *Ticker {

	r := &Request{
		Method:   http.MethodGet,
		Endpoint: api.TICKER,
		SecType:  SecTypeNone,
	}

	if s.Symbol != "" {
		r.SetParam("symbol", s.Symbol)
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		fmt.Println(err)
	}

	result := new(Ticker)

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result

}
