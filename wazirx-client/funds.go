package wazirx_client

import (
	"context"
	"encoding/json"
	"fmt"

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

func (s *FundInfoService) FetchFundsInfo(ctx context.Context) *[]FundInfo {

	r := &Request{
		Method:     http.MethodGet,
		Endpoint:   api.FUNDS,
		SecType:    SecTypeSigned,
		RecvWindow: 10000,
		TimeStamp:  helpers.CurrentTimestamp(),
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		fmt.Println(err)
	}

	result := new([]FundInfo)

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result

}
