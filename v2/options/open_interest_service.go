package options

import (
	"context"
	"encoding/json"
	"net/http"
)

type OpenInterest struct {
	Symbol             string `json:"symbol"`
	SumOpenInterest    string `json:"sumOpenInterest"`
	SumOpenInterestUsd string `json:"sumOpenInterestUsd"`
	Timestamp          string `json:"timestamp"`
}

// underlying: Spot trading pairs such as BTCUSDT
type OpenInterestService struct {
	c               *Client
	underlyingAsset string //Target assets, such as ETH or BTC
	expiration      string //Maturity date, such as 221225
}

// Underlying set underlying
func (s *OpenInterestService) UnderlyingAsset(underlyingAsset string) *OpenInterestService {
	s.underlyingAsset = underlyingAsset
	return s
}

func (s *OpenInterestService) Expiration(expiration string) *OpenInterestService {
	s.expiration = expiration
	return s
}

// Do send request
func (s *OpenInterestService) Do(ctx context.Context, opts ...RequestOption) (res []*OpenInterest, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/openInterest",
	}
	r.setParam("underlyingAsset", s.underlyingAsset)
	r.setParam("expiration", s.expiration)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*OpenInterest{}, err
	}
	res = make([]*OpenInterest, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*OpenInterest{}, err
	}
	return res, nil
}
