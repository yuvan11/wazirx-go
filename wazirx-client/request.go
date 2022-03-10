package wazirx_client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	SecTypeNone SecType = iota
	SecTypeAPIKey
	SecTypeSigned // if the 'timestamp' parameter is required
)

type SecType int

type Params map[string]interface{}

type Request struct {
	Method     string
	Endpoint   string
	Query      url.Values
	Form       url.Values
	RecvWindow int64
	TimeStamp  int64
	SecType    SecType
	Header     http.Header
	Body       io.Reader
	FullURL    string
}

// addParam add single param with key/value to Query string
func (r *Request) AddParam(key string, value interface{}) *Request {
	if r.Query == nil {
		r.Query = url.Values{}
	}
	r.Query.Add(key, fmt.Sprintf("%v", value))
	return r
}

// setParam set single param with key/value to Query string
func (r *Request) SetParam(key string, value interface{}) *Request {
	if r.Query == nil {
		r.Query = url.Values{}
	}
	r.Query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setParams set multiple Params with key/values to Query string
func (r *Request) SetParams(m Params) *Request {
	for k, v := range m {
		r.SetParam(k, v)
	}
	return r
}

// setFormParam set single Form param with key/value to Request Form body
func (r *Request) SetFormParam(key string, value interface{}) *Request {
	if r.Form == nil {
		r.Form = url.Values{}
	}
	r.Form.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setFormParams set multiple Form Params with key/values to Request Form body
func (r *Request) SetFormParams(m Params) *Request {
	for k, v := range m {
		r.SetFormParam(k, v)
	}
	return r
}

func (r *Request) ValidateReq() (err error) {

	if r.Query == nil {
		r.Query = url.Values{}
	}
	if r.Form == nil {
		r.Form = url.Values{}
	}

	return nil
}

// RequestOption define option type for Request
type RequestOption func(*Request)

// WithRecvWindow set RecvWindow param for the Request
func WithRecvWindow(RecvWindow int64) RequestOption {
	return func(r *Request) {
		r.RecvWindow = RecvWindow
	}
}

// WithHeader set or add a Header value to the Request
func WithHeader(key, value string, replace bool) RequestOption {
	return func(r *Request) {
		if r.Header == nil {
			r.Header = http.Header{}
		}
		if replace {
			r.Header.Set(key, value)
		} else {
			r.Header.Add(key, value)
		}
	}
}

// WithHeaders set or replace the Headers of the Request
func WithHeaders(Header http.Header) RequestOption {
	return func(r *Request) {
		r.Header = Header.Clone()
	}
}
