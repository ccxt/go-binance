package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMPositionRiskService get UM position risk information
type GetUMPositionRiskService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *GetUMPositionRiskService) Symbol(symbol string) *GetUMPositionRiskService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *GetUMPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*UMPosition, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*UMPosition{}, err
	}
	res = make([]*UMPosition, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*UMPosition{}, err
	}
	return res, nil
}
