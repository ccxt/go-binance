package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMAccountTradesService service to get UM account trade list
type UMAccountTradesService struct {
	c          *Client
	symbol     string
	startTime  *int64
	endTime    *int64
	fromID     *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *UMAccountTradesService) Symbol(symbol string) *UMAccountTradesService {
	s.symbol = symbol
	return s
}

// StartTime set startTime
func (s *UMAccountTradesService) StartTime(startTime int64) *UMAccountTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UMAccountTradesService) EndTime(endTime int64) *UMAccountTradesService {
	s.endTime = &endTime
	return s
}

// FromID set fromID
func (s *UMAccountTradesService) FromID(fromID int64) *UMAccountTradesService {
	s.fromID = &fromID
	return s
}

// Limit set limit
func (s *UMAccountTradesService) Limit(limit int) *UMAccountTradesService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *UMAccountTradesService) RecvWindow(recvWindow int64) *UMAccountTradesService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMAccountTradesService) Do(ctx context.Context) ([]*UMAccountTrade, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/userTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*UMAccountTrade
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
