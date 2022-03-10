package wazirx_client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	helpers "wazirx-go/wazirx-client/Helpers"
	api "wazirx-go/wazirx-client/endpoints"
)

type HistoricalTrades struct {
	ID            int64  `json:"id"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
	Price         string `json:"price"`
	Quantity      string `json:"qty"`
	QuoteQuantity string `json:"quoteQty"`
	Time          int64  `json:"time"`
}

type HistoricalTradesService struct {
	C      *Client
	Symbol string
	Limit  int32
}

func (s *HistoricalTradesService) SetSymbol(symbol string) *HistoricalTradesService {
	s.Symbol = symbol
	return s
}

func (s *HistoricalTradesService) SetLimit(limit int32) *HistoricalTradesService {
	s.Limit = limit
	return s
}

func (s *HistoricalTradesService) DisplayHistoricalTrade(ctx context.Context) *[]HistoricalTrades {

	r := &Request{
		Method:     http.MethodGet,
		Endpoint:   api.HISTORICAL_TRADES,
		SecType:    SecTypeSigned,
		RecvWindow: 10000,
		TimeStamp:  helpers.CurrentTimestamp(),
	}

	if s.Limit > 0 {
		r.SetParam("limit", s.Limit)
	}

	if s.Symbol != "" {
		r.SetParam("symbol", s.Symbol)
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		fmt.Println(err)
	}

	result := new([]HistoricalTrades)

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}
