package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetFundingInfoService struct {
	c *Client
}

type FundingInfo struct {
	Symbol                   string `json:"symbol"`
	AdjustedFundingRateCap   string `json:"adjustedFundingRateCap"`
	AdjustedFundingRateFloor string `json:"adjustedFundingRateFloor"`
	FundingIntervalHours     int    `json:"fundingIntervalHours"`
	Disclaimer               bool   `json:"disclaimer"`
}

func (s *GetFundingInfoService) Do(ctx context.Context, opts ...RequestOption) (fundingInfo []*FundingInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/fundingInfo",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	fundingInfo = make([]*FundingInfo, 0)
	err = json.Unmarshal(data, &fundingInfo)
	if err != nil {
		return nil, err
	}
	return fundingInfo, nil
}
