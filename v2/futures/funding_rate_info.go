package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// FundingRateInfoService gets funding rate info
type FundingRateInfoService struct {
	c *Client
}

// Do sends request
func (s *FundingRateInfoService) Do(ctx context.Context, opts ...RequestOption) (res []*FundingRateInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/fundingInfo",
		secType:  secTypeNone,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return []*FundingRateInfo{}, err
	}
	res = make([]*FundingRateInfo, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*FundingRateInfo{}, err
	}
	return res, nil
}

// FundingRateInfo defines funding rate info for symbols
type FundingRateInfo struct {
	Symbol                   string `json:"symbol"`
	AdjustedFundingRateCap   string `json:"adjustedFundingRateCap"`
	AdjustedFundingRateFloor string `json:"adjustedFundingRateFloor"`
	FundingIntervalHours     int64  `json:"fundingIntervalHours"`
}
