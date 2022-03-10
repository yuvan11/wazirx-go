package wazirx_client

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"
	helpers "wazirx-go/wazirx-client/Helpers"
	api "wazirx-go/wazirx-client/endpoints"
)

type AccountInfoService struct {
	C *Client
}

type AccountInfo struct {
	AccountType string `json:"accountType"`
	CanTrade    bool   `json:"canTrade"`
	CanWithdraw bool   `json:"canWithdraw"`
	UpdateTime  int64  `json:"updateTime"`
}

func (s *AccountInfoService) FetchAccountInfo(ctx context.Context) *AccountInfo {

	r := &Request{
		Method:     http.MethodGet,
		Endpoint:   api.ACCOUNT,
		SecType:    SecTypeSigned,
		RecvWindow: 10000,
		TimeStamp:  helpers.CurrentTimestamp(),
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		fmt.Println(err)
	}

	result := new(AccountInfo)

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result

}
