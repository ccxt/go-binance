package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMAllOrdersService service to get all UM orders
type UMAllOrdersService struct {
	c          *Client
	symbol     string
	orderID    *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *UMAllOrdersService) Symbol(symbol string) *UMAllOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *UMAllOrdersService) OrderID(orderID int64) *UMAllOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set startTime
func (s *UMAllOrdersService) StartTime(startTime int64) *UMAllOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UMAllOrdersService) EndTime(endTime int64) *UMAllOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *UMAllOrdersService) Limit(limit int) *UMAllOrdersService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *UMAllOrdersService) RecvWindow(recvWindow int64) *UMAllOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMAllOrdersService) Do(ctx context.Context) ([]*UMAllOrdersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/allOrders",
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
	var res []*UMAllOrdersResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMAllOrdersResponse define all orders response
type UMAllOrdersResponse struct {
	AvgPrice                string `json:"avgPrice"`
	ClientOrderID           string `json:"clientOrderId"`
	CumQuote                string `json:"cumQuote"`
	ExecutedQty             string `json:"executedQty"`
	OrderID                 int64  `json:"orderId"`
	OrigQty                 string `json:"origQty"`
	OrigType                string `json:"origType"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	Status                  string `json:"status"`
	Symbol                  string `json:"symbol"`
	Time                    int64  `json:"time"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	UpdateTime              int64  `json:"updateTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
	PriceMatch              string `json:"priceMatch"`
}
