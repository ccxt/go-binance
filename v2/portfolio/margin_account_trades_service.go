package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginAccountTradesService service to get margin account trade list
type MarginAccountTradesService struct {
	c          *Client
	symbol     string
	orderID    *int64
	startTime  *int64
	endTime    *int64
	fromID     *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *MarginAccountTradesService) Symbol(symbol string) *MarginAccountTradesService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *MarginAccountTradesService) OrderID(orderID int64) *MarginAccountTradesService {
	s.orderID = &orderID
	return s
}

// StartTime set startTime
func (s *MarginAccountTradesService) StartTime(startTime int64) *MarginAccountTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginAccountTradesService) EndTime(endTime int64) *MarginAccountTradesService {
	s.endTime = &endTime
	return s
}

// FromID set fromID
func (s *MarginAccountTradesService) FromID(fromID int64) *MarginAccountTradesService {
	s.fromID = &fromID
	return s
}

// Limit set limit
func (s *MarginAccountTradesService) Limit(limit int) *MarginAccountTradesService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *MarginAccountTradesService) RecvWindow(recvWindow int64) *MarginAccountTradesService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginAccountTradesService) Do(ctx context.Context) ([]*MarginTrade, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/myTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
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
	var res []*MarginTrade
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginTrade define margin trade info
type MarginTrade struct {
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	ID              int64  `json:"id"`
	IsBestMatch     bool   `json:"isBestMatch"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	OrderID         int64  `json:"orderId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Symbol          string `json:"symbol"`
	Time            int64  `json:"time"`
}
