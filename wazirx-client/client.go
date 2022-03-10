package wazirx_client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	api "wazirx-go/wazirx-client/endpoints"
)

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	ApiKey     string
	Secretkey  string
	BaseURL    string
	HTTPClient http.Client
	TimeOffset int64
	do         doFunc
}

func (c *Client) ParseRequest(r *Request) (err error) {

	client := &Client{}
	err = r.ValidateReq()

	if err != nil {
		return err
	}
	client.BaseURL = api.BASE_URL

	fullURL := fmt.Sprintf("%s%s", client.BaseURL, r.Endpoint)

	if r.RecvWindow > 0 {
		r.SetParam("recvWindow", r.RecvWindow)
	}
	if r.TimeStamp > 0 {
		r.SetParam("timestamp", r.TimeStamp)
	}

	queryString := r.Query.Encode()

	body := &bytes.Buffer{}

	bodyString := r.Form.Encode()

	header := http.Header{}

	if r.Header != nil {
		header = r.Header.Clone()
	}

	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}

	if r.SecType == SecTypeAPIKey || r.SecType == SecTypeSigned {
		client.ApiKey = api.API_KEY
		header.Set("X-Api-Key", client.ApiKey)
	}

	if r.SecType == SecTypeSigned {
		client.Secretkey = api.SECRET_KEY
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(client.Secretkey))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set("signature", fmt.Sprintf("%x", (mac.Sum(nil))))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}

	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}

	r.FullURL = fullURL
	r.Header = header
	r.Body = body

	return nil

}

func (c *Client) HandleAPI(ctx context.Context, r *Request) (data []byte, err error) {

	client := &Client{}
	err = c.ParseRequest(r)

	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest(r.Method, r.FullURL, r.Body)

	if err != nil {
		return []byte{}, err
	}

	req = req.WithContext(ctx)
	req.Header = r.Header

	f := client.do
	if f == nil {
		f = client.HTTPClient.Do
	}
	res, err := f(req)

	if err != nil {
		return []byte{}, err
	}

	data, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	}

	defer func() {
		closeErr := res.Body.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(CustomAPIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			log.Println("failed to unmarshal", e)
		}
		return nil, apiErr
	}
	return data, nil

}

// Public API Services init collection

// Ping service initiates
func (c *Client) NewPingService() *PingService {
	return &PingService{C: c}
}

// Server service initiates
func (c *Client) NewServerTimeService() *ServerTimeService {
	return &ServerTimeService{C: c}
}

// System status service initiates
func (c *Client) NewSystemStatusService() *SystemStatusService {
	return &SystemStatusService{C: c}
}

// Exchange status service initiates
func (c *Client) NewExchangeInfoService() *ExchangeInfoService {
	return &ExchangeInfoService{C: c}
}

// Tickers service initiates
func (c *Client) NewTickersService() *TickersService {
	return &TickersService{C: c}
}

// Ticker service initiates
func (c *Client) NewTickerService() *TickerService {
	return &TickerService{C: c}
}

// Market depth service initiates
func (c *Client) NewOrderBookDepthService() *OrderBookDepthService {
	return &OrderBookDepthService{C: c}
}

// Market trade service initiates
func (c *Client) NewRecentTradesService() *RecentTradesService {
	return &RecentTradesService{C: c}
}

// Historical trade service initiates
func (c *Client) NewHistoricalTradesService() *HistoricalTradesService {
	return &HistoricalTradesService{C: c}
}

// Account  API Services init collection

// Account display service initiates
func (c *Client) NewAccountInfoService() *AccountInfoService {
	return &AccountInfoService{C: c}
}

// Account funds service initiates
func (c *Client) NewFundInfoService() *FundInfoService {
	return &FundInfoService{C: c}
}

// WebSocket-Client  API Services init collection

// Generate auth token service initiates
func (c *Client) NewAuthInfoService() *AuthInfoService {
	return &AuthInfoService{C: c}
}
