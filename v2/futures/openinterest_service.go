package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetOpenInterestService get present open interest of a specific symbol.
type GetOpenInterestService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetOpenInterestService) Symbol(symbol string) *GetOpenInterestService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetOpenInterestService) Do(ctx context.Context, opts ...RequestOption) (res *OpenInterest, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openInterest",
	}
	r.setParam("symbol", s.symbol)
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(OpenInterest)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type OpenInterest struct {
	OpenInterest string `json:"openInterest"`
	Symbol       string `json:"symbol"`
	Time         int64  `json:"time"`
}

// OpenInterestStatisticsService list open history data of a symbol.
type OpenInterestStatisticsService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *OpenInterestStatisticsService) Symbol(symbol string) *OpenInterestStatisticsService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *OpenInterestStatisticsService) Period(period string) *OpenInterestStatisticsService {
	s.period = period
	return s
}

// Limit set limit
func (s *OpenInterestStatisticsService) Limit(limit int) *OpenInterestStatisticsService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *OpenInterestStatisticsService) StartTime(startTime int64) *OpenInterestStatisticsService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *OpenInterestStatisticsService) EndTime(endTime int64) *OpenInterestStatisticsService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *OpenInterestStatisticsService) Do(ctx context.Context, opts ...RequestOption) (res []*OpenInterestStatistic, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/openInterestHist",
	}

	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)

	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*OpenInterestStatistic{}, err
	}

	res = make([]*OpenInterestStatistic, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*OpenInterestStatistic{}, err
	}

	return res, nil
}

type OpenInterestStatistic struct {
	Symbol               string `json:"symbol"`
	SumOpenInterest      string `json:"sumOpenInterest"`
	SumOpenInterestValue string `json:"sumOpenInterestValue"`
	Timestamp            int64  `json:"timestamp"`
	CMCCirculatingSupply string `json:"CMCCirculatingSupply"`
}
