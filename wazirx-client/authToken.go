package wazirx_client

import (
	"context"
	"encoding/json"

	"net/http"

	helpers "github.com/yuvan11/wazirx-go/wazirx-client/Helpers"
	api "github.com/yuvan11/wazirx-go/wazirx-client/endpoints"
)

type AuthInfoService struct {
	C *Client
}

type AuthInfo struct {
	AuthToken        string `json:"auth_key"`
	Timeout_Duration int32  `json:"timeout_duration"`
}

func (s *AuthInfoService) CreateAuthToken(ctx context.Context) (*AuthInfo, error) {

	r := &Request{
		Method:     http.MethodPost,
		Endpoint:   api.CREATE_AUTH_TOKEN,
		SecType:    SecTypeSigned,
		RecvWindow: 10000,
		TimeStamp:  helpers.CurrentTimestamp(),
	}

	data, err := s.C.HandleAPI(ctx, r)

	if err != nil {
		return &AuthInfo{}, err
	}

	result := new(AuthInfo)

	if err := json.Unmarshal(data, &result); err != nil {
		return &AuthInfo{}, err
	}

	return result, nil

}
