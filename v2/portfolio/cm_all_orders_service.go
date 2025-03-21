package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMAllOrdersService service to get all CM orders
type CMAllOrdersService struct {
	c          *Client
	symbol     *string
	pair       *string
	orderID    *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *CMAllOrdersService) Symbol(symbol string) *CMAllOrdersService {
	s.symbol = &symbol
	return s
}

// Pair set pair
func (s *CMAllOrdersService) Pair(pair string) *CMAllOrdersService {
	s.pair = &pair
	return s
}

// OrderID set orderID
func (s *CMAllOrdersService) OrderID(orderID int64) *CMAllOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set startTime
func (s *CMAllOrdersService) StartTime(startTime int64) *CMAllOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *CMAllOrdersService) EndTime(endTime int64) *CMAllOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *CMAllOrdersService) Limit(limit int) *CMAllOrdersService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *CMAllOrdersService) RecvWindow(recvWindow int64) *CMAllOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMAllOrdersService) Do(ctx context.Context) ([]*CMAllOrdersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/allOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.pair != nil {
		r.setParam("pair", *s.pair)
	}
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
	var res []*CMAllOrdersResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMAllOrdersResponse define all orders response
type CMAllOrdersResponse struct {
	AvgPrice      string `json:"avgPrice"`
	ClientOrderID string `json:"clientOrderId"`
	CumBase       string `json:"cumBase"`
	ExecutedQty   string `json:"executedQty"`
	OrderID       int64  `json:"orderId"`
	OrigQty       string `json:"origQty"`
	OrigType      string `json:"origType"`
	Price         string `json:"price"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	Status        string `json:"status"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	Time          int64  `json:"time"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	UpdateTime    int64  `json:"updateTime"`
}
