package options

import (
	"context"
	"encoding/json"
	"net/http"
)

// AccountService create order
type AccountService struct {
	c *Client
}

// Do send request
func (s *AccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/account",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(Account)
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, err
	}
	return res, nil
}

type Asset struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	Equity        string `json:"equity"`
	Available     string `json:"available"`
	Locked        string `json:"locked"`
	UnrealizedPNL string `json:"unrealizedPNL"`
}

type Greek struct {
	Underlying string `json:"underlying"`
	Delta      string `json:"delta"`
	Gamma      string `json:"gamma"`
	Theta      string `json:"theta"`
	Vega       string `json:"vega"`
}

// Account define create order response
type Account struct {
	Asset     []*Asset `json:"asset"`
	Greek     []*Greek `json:"greek"`
	RiskLevel string   `json:"riskLevel"`
	Time      uint64   `json:"time"`
}
