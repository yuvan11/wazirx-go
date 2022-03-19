package wazirx_client

import (
	"context"
	"encoding/json"

	"net/http"

	helpers "github.com/yuvan11/wazirx-go/wazirx-client/Helpers"
	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

type FundInfoService struct {
	C *Client
}

type FundInfo struct {
	Asset       string `json:"asset"`
	Free        string `json:"free"`
	Locked      string `json:"locked"`
	ReservedFee string `json:"reservedFee"`
}

func (s *FundInfoService) FetchFundsInfo(ctx context.Context) (*[]FundInfo, error) {

	r := &Request{
		Method:     http.MethodGet,
		Endpoint:   api.FUNDS,
		SecType:    SecTypeSigned,
		RecvWindow: 10000,
		TimeStamp:  helpers.CurrentTimestamp(),
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		return &[]FundInfo{}, err
	}

	result := new([]FundInfo)

	if err := json.Unmarshal(data, &result); err != nil {
		return &[]FundInfo{}, err
	}

	return result, nil

}
