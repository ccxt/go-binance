package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMarginAllOrdersService service to get all margin account orders
type GetMarginAllOrdersService struct {
	c          *Client
	symbol     string
	orderID    *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *GetMarginAllOrdersService) Symbol(symbol string) *GetMarginAllOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetMarginAllOrdersService) OrderID(orderID int64) *GetMarginAllOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set startTime
func (s *GetMarginAllOrdersService) StartTime(startTime int64) *GetMarginAllOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetMarginAllOrdersService) EndTime(endTime int64) *GetMarginAllOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetMarginAllOrdersService) Limit(limit int) *GetMarginAllOrdersService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *GetMarginAllOrdersService) RecvWindow(recvWindow int64) *GetMarginAllOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetMarginAllOrdersService) Do(ctx context.Context) ([]*MarginOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/allOrders",
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
	var res []*MarginOrder
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
