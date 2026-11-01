package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// InterestHistoryService fetches the interest history
type InterestHistoryService struct {
	c           *Client
	lendingType LendingType
	asset       *string
	startTime   *int64
	endTime     *int64
	current     *int32
	size        *int32
}

// LendingType sets the lendingType parameter.
func (s *InterestHistoryService) LendingType(lendingType LendingType) *InterestHistoryService {
	s.lendingType = lendingType
	return s
}

// Asset sets the asset parameter.
func (s *InterestHistoryService) Asset(asset string) *InterestHistoryService {
	s.asset = &asset
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-30 days.
func (s *InterestHistoryService) StartTime(startTime int64) *InterestHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *InterestHistoryService) EndTime(endTime int64) *InterestHistoryService {
	s.endTime = &endTime
	return s
}

// Current sets the current parameter.
func (s *InterestHistoryService) Current(current int32) *InterestHistoryService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *InterestHistoryService) Size(size int32) *InterestHistoryService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *InterestHistoryService) Do(ctx context.Context) (*InterestHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/union/interestHistory",
		secType:  secTypeSigned,
	}
	r.setParam("lendingType", s.lendingType)
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(InterestHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// InterestHistory represents a response from InterestHistoryService.
type InterestHistory []InterestHistoryElement

type InterestHistoryElement struct {
	Asset       string      `json:"asset"`
	Interest    string      `json:"interest"`
	LendingType LendingType `json:"lendingType"`
	ProductName string      `json:"productName"`
	Time        int64       `json:"time"`
}
