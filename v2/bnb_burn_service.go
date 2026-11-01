package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetBNBBurnService get BNB Burn on spot trade and margin interest
type GetBNBBurnService struct {
	c *Client
}

// Do send request
func (s *GetBNBBurnService) Do(ctx context.Context, opts ...RequestOption) (*BNBBurn, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bnbBurn",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := BNBBurn{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// BNBBurn response
type BNBBurn struct {
	SpotBNBBurn     bool `json:"spotBNBBurn"`
	InterestBNBBurn bool `json:"interestBNBBurn"`
}

// ToggleBNBBurnService toggle BNB Burn on spot trade and margin interest
type ToggleBNBBurnService struct {
	c *Client

	spotBNBBurn     *bool
	interestBNBBurn *bool
}

// SpotBNBBurn sets the spot bnb burn parameter.
func (s *ToggleBNBBurnService) SpotBNBBurn(v bool) *ToggleBNBBurnService {
	s.spotBNBBurn = &v
	return s
}

// InterestBNBBurn sets the interest BNB burn parameter
func (s *ToggleBNBBurnService) InterestBNBBurn(v bool) *ToggleBNBBurnService {
	s.interestBNBBurn = &v
	return s
}

// Do send request
func (s *ToggleBNBBurnService) Do(ctx context.Context, opts ...RequestOption) (*BNBBurn, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/bnbBurn",
		secType:  secTypeSigned,
	}
	if s.spotBNBBurn != nil {
		r.setParam("spotBNBBurn", *s.spotBNBBurn)
	}
	if s.interestBNBBurn != nil {
		r.setParam("interestBNBBurn", *s.interestBNBBurn)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := BNBBurn{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
