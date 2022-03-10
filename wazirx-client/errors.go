package wazirx_client

import "fmt"

type CustomAPIError struct {
	ErrCode    int64  `json:"errcode"`
	ErrMessage string `json:"errmessage"`
}

func (e CustomAPIError) Error() string {
	return fmt.Sprintf("<CustomAPIError> ErrCode=%d, ErrMessage=%s", e.ErrCode, e.ErrMessage)
}
