package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMAccountTradeService service to get CM account trade list
type CMAccountTradeService struct {
	c          *Client
	symbol     *string
	pair       *string
	startTime  *int64
	endTime    *int64
	fromID     *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *CMAccountTradeService) Symbol(symbol string) *CMAccountTradeService {
	s.symbol = &symbol
	return s
}

// Pair set pair
func (s *CMAccountTradeService) Pair(pair string) *CMAccountTradeService {
	s.pair = &pair
	return s
}

// StartTime set startTime
func (s *CMAccountTradeService) StartTime(startTime int64) *CMAccountTradeService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *CMAccountTradeService) EndTime(endTime int64) *CMAccountTradeService {
	s.endTime = &endTime
	return s
}

// FromID set fromId
func (s *CMAccountTradeService) FromID(fromID int64) *CMAccountTradeService {
	s.fromID = &fromID
	return s
}

// Limit set limit
func (s *CMAccountTradeService) Limit(limit int) *CMAccountTradeService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *CMAccountTradeService) RecvWindow(recvWindow int64) *CMAccountTradeService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMAccountTradeService) Do(ctx context.Context) ([]*CMAccountTrade, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/userTrades",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.pair != nil {
		r.setParam("pair", *s.pair)
	}
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
	var res []*CMAccountTrade
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
