package options

import (
	"context"
	"encoding/json"
	"net/http"
)

type ExerciseHistory struct {
	Symbol          string `json:"symbol"`
	StrikePrice     string `json:"strikePrice"`
	RealStrikePrice string `json:"realStrikePrice"`
	ExpiryDate      uint64 `json:"expiryDate"`
	StrikeResult    string `json:"strikeResult"`
}

// underlying: Spot trading pairs such as BTCUSDT
type ExerciseHistoryService struct {
	c          *Client
	underlying *string
	startTime  *uint64
	endTime    *uint64
	limit      *uint64
}

// Underlying set underlying
func (s *ExerciseHistoryService) Underlying(underlying string) *ExerciseHistoryService {
	s.underlying = &underlying
	return s
}

func (s *ExerciseHistoryService) StartTime(startTime uint64) *ExerciseHistoryService {
	s.startTime = &startTime
	return s
}

func (s *ExerciseHistoryService) EndTime(endTime uint64) *ExerciseHistoryService {
	s.endTime = &endTime
	return s
}

func (s *ExerciseHistoryService) Limit(limit uint64) *ExerciseHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ExerciseHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*ExerciseHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/exerciseHistory",
	}
	if s.underlying != nil {
		r.setParam("underlying", s.underlying)
	}
	if s.startTime != nil {
		r.setParam("startTime", s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", s.limit)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*ExerciseHistory{}, err
	}
	res = make([]*ExerciseHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*ExerciseHistory{}, err
	}
	return res, nil
}
