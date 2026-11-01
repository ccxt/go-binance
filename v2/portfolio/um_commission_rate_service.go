package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMCommissionRateService get user commission rate for UM
type GetUMCommissionRateService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetUMCommissionRateService) Symbol(symbol string) *GetUMCommissionRateService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetUMCommissionRateService) Do(ctx context.Context, opts ...RequestOption) (*CommissionRate, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/commissionRate",
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

// CommissionRate define commission rate info
type CommissionRate struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"` // 0.02%
	TakerCommissionRate string `json:"takerCommissionRate"` // 0.04%
}
