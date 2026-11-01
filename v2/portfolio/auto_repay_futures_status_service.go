package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAutoRepayFuturesStatusService get auto-repay-futures status
type GetAutoRepayFuturesStatusService struct {
	c *Client
}

// Do send request
func (s *GetAutoRepayFuturesStatusService) Do(ctx context.Context, opts ...RequestOption) (*AutoRepayFuturesStatus, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/repay-futures-switch",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(AutoRepayFuturesStatus)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AutoRepayFuturesStatus define auto repay futures status
type AutoRepayFuturesStatus struct {
	AutoRepay bool `json:"autoRepay"` // true for turn on the auto-repay futures; false for turn off
}
