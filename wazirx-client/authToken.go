package wazirx_client

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"
	helpers "wazirx-go/wazirx-client/Helpers"
	api "wazirx-go/wazirx-client/endpoints"
)

type AuthInfoService struct {
	C *Client
}

type AuthInfo struct {
	AuthToken        string `json:"auth_key"`
	Timeout_Duration int32  `json:"timeout_duration"`
}

func (s *AuthInfoService) CreateAuthToken(ctx context.Context) *AuthInfo {

	r := &Request{
		Method:     http.MethodPost,
		Endpoint:   api.CREATE_AUTH_TOKEN,
		SecType:    SecTypeSigned,
		RecvWindow: 10000,
		TimeStamp:  helpers.CurrentTimestamp(),
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		fmt.Println(err)
	}

	result := new(AuthInfo)

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result

}
