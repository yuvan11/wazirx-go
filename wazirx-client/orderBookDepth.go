package wazirx_client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	helpers "github.com/yuvan11/wazirx-go/wazirx-client/Helpers"
	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

type OrderBookDepthService struct {
	C      *Client
	Symbol string
	Limit  int32
}

/*
type OrderPrice struct {
	PriceValue string
	Quantity   string
}
*/

type OrderBookDepth struct {
	TimeStamp int64      `json:"timestamp"`
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
}

func (s *OrderBookDepthService) SetSymbol(symbol string) *OrderBookDepthService {
	s.Symbol = symbol
	return s
}

func (s *OrderBookDepthService) SetLimit(limit int32) *OrderBookDepthService {
	s.Limit = limit
	return s
}

func (s *OrderBookDepthService) DisplayOrdersBookDepth(ctx context.Context) *OrderBookDepth {

	r := &Request{
		Method:     http.MethodGet,
		Endpoint:   api.MARKET_DEPTH,
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

	//var result map[string]interface{}
	result := new(OrderBookDepth)

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}
