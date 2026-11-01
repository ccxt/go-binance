package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// RateLimitService get rate limits
type RateLimitService struct {
	c *Client
}

// Do send request
func (s *RateLimitService) Do(ctx context.Context, opts ...RequestOption) (res []*RateLimitFull, err error) {
	res = make([]*RateLimitFull, 0)
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/rateLimit/order",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

type RateLimitFull struct {
	RateLimitType RateLimitType     `json:"rateLimitType"`
	Interval      RateLimitInterval `json:"interval"`
	IntervalNum   int               `json:"intervalNum"`
	Limit         int               `json:"limit"`
	Count         int               `json:"count"`
}
