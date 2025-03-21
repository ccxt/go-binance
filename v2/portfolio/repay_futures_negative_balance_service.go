package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// RepayFuturesNegativeBalanceService repay futures negative balance
type RepayFuturesNegativeBalanceService struct {
	c *Client
}

// Do send request
func (s *RepayFuturesNegativeBalanceService) Do(ctx context.Context, opts ...RequestOption) (*SuccessResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/repay-futures-negative-balance",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(SuccessResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
