package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginNextHourlyInterestRateService fetches the margin interest history
type MarginNextHourlyInterestRateService struct {
	c        *Client
	assets   *string
	isolated *bool
}

// Assets sets the assets parameter.
func (s *MarginNextHourlyInterestRateService) Assets(assets string) *MarginNextHourlyInterestRateService {
	s.assets = &assets
	return s
}

// Isolated sets the isolated parameter.
func (s *MarginNextHourlyInterestRateService) Isolated(isolated bool) *MarginNextHourlyInterestRateService {
	s.isolated = &isolated
	return s
}

// Do sends the request.
func (s *MarginNextHourlyInterestRateService) Do(ctx context.Context) (*MarginNextHourlyInterestRate, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/next-hourly-interest-rate",
		secType:  secTypeSigned,
	}
	if s.assets != nil {
		r.setParam("assets", *s.assets)
	}
	if s.isolated != nil {
		r.setParam("isolated", *s.isolated)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(MarginNextHourlyInterestRate)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginNextHourlyInterestRate represents the response
type MarginNextHourlyInterestRate []MarginNextHourlyInterestRateElement

type MarginNextHourlyInterestRateElement struct {
	Asset                  string `json:"asset"`
	NextHourlyInterestRate string `json:"nextHourlyInterestRate"`
}
