package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetCMCommissionRateService get user commission rate for CM
type GetCMCommissionRateService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetCMCommissionRateService) Symbol(symbol string) *GetCMCommissionRateService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetCMCommissionRateService) Do(ctx context.Context, opts ...RequestOption) (*CommissionRate, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/commissionRate",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(CommissionRate)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
