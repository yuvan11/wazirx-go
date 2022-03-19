package wazirx_client

import (
	"context"
	"encoding/json"
	"net/http"

	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

type RecentTradesService struct {
	C      *Client
	Symbol string
	Limit  int32
}

type RecentTrades struct {
	ID            int64  `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"qty"`
	QuoteQuantity string `json:"quoteQty"`
	Time          int64  `json:"time"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
}

func (s *RecentTradesService) SetSymbol(symbol string) *RecentTradesService {
	s.Symbol = symbol
	return s
}

func (s *RecentTradesService) SetLimit(limit int32) *RecentTradesService {
	s.Limit = limit
	return s
}

func (s *RecentTradesService) DisplayRecentMarketTrade(ctx context.Context) (*[]RecentTrades, error) {

	r := &Request{
		Method:   http.MethodGet,
		Endpoint: api.MARKET_TRADES,
		SecType:  SecTypeNone,
	}

	if s.Limit > 0 {
		r.SetParam("limit", s.Limit)
	}

	if s.Symbol != "" {
		r.SetParam("symbol", s.Symbol)
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		return &[]RecentTrades{}, err
	}

	result := new([]RecentTrades)

	if err := json.Unmarshal(data, &result); err != nil {
		return &[]RecentTrades{}, err
	}

	return result, nil
}
