package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginInterestRateHistoryService fetches the margin interest history
type MarginInterestRateHistoryService struct {
	c         *Client
	asset     *string
	vipLevel  *int64
	startTime *int64
	endTime   *int64
}

// Asset sets the asset parameter.
func (s *MarginInterestRateHistoryService) Asset(asset string) *MarginInterestRateHistoryService {
	s.asset = &asset
	return s
}

// VipLevel sets the current parameter.
func (s *MarginInterestRateHistoryService) VipLevel(vipLevel int64) *MarginInterestRateHistoryService {
	s.vipLevel = &vipLevel
	return s
}

// StartTime sets the startTime parameter.
func (s *MarginInterestRateHistoryService) StartTime(startTime int64) *MarginInterestRateHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
func (s *MarginInterestRateHistoryService) EndTime(endTime int64) *MarginInterestRateHistoryService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *MarginInterestRateHistoryService) Do(ctx context.Context) (*MarginInterestRateHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/interestRateHistory",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.vipLevel != nil {
		r.setParam("vipLevel", *s.vipLevel)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(MarginInterestRateHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginInterestRateHistory represents the response
type MarginInterestRateHistory []MarginInterestRateHistoryElement

type MarginInterestRateHistoryElement struct {
	Timestamp         int64  `json:"timestamp"`
	VipLevel          int64  `json:"vipLevel"`
	Asset             string `json:"asset"`
	DailyInterestRate string `json:"dailyInterestRate"`
}
