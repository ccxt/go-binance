package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetCMPositionRiskService get CM position risk information
type GetCMPositionRiskService struct {
	c           *Client
	marginAsset *string
	pair        *string
}

// MarginAsset set margin asset
func (s *GetCMPositionRiskService) MarginAsset(marginAsset string) *GetCMPositionRiskService {
	s.marginAsset = &marginAsset
	return s
}

// Pair set trading pair
func (s *GetCMPositionRiskService) Pair(pair string) *GetCMPositionRiskService {
	s.pair = &pair
	return s
}

// Do send request
func (s *GetCMPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*CMPosition, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/positionRisk",
		secType:  secTypeSigned,
	}
	if s.marginAsset != nil {
		r.setParam("marginAsset", *s.marginAsset)
	}
	if s.pair != nil {
		r.setParam("pair", *s.pair)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*CMPosition{}, err
	}
	res = make([]*CMPosition, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CMPosition{}, err
	}
	return res, nil
}
