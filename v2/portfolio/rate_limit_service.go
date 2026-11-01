package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetRateLimitService get user rate limit
type GetRateLimitService struct {
	c          *Client
	recvWindow *int64
}

// RecvWindow set recvWindow
func (s *GetRateLimitService) RecvWindow(recvWindow int64) *GetRateLimitService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetRateLimitService) Do(ctx context.Context, opts ...RequestOption) ([]*RateLimit, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/rateLimit/order",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := make([]*RateLimit, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RateLimit define rate limit info
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int64  `json:"intervalNum"`
	Limit         int64  `json:"limit"`
}
