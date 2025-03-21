package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMForceOrdersService service to get user's UM force orders
type UMForceOrdersService struct {
	c             *Client
	symbol        *string
	autoCloseType *string
	startTime     *int64
	endTime       *int64
	limit         *int
	recvWindow    *int64
}

// Symbol set symbol
func (s *UMForceOrdersService) Symbol(symbol string) *UMForceOrdersService {
	s.symbol = &symbol
	return s
}

// AutoCloseType set autoCloseType
func (s *UMForceOrdersService) AutoCloseType(autoCloseType string) *UMForceOrdersService {
	s.autoCloseType = &autoCloseType
	return s
}

// StartTime set startTime
func (s *UMForceOrdersService) StartTime(startTime int64) *UMForceOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UMForceOrdersService) EndTime(endTime int64) *UMForceOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *UMForceOrdersService) Limit(limit int) *UMForceOrdersService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *UMForceOrdersService) RecvWindow(recvWindow int64) *UMForceOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMForceOrdersService) Do(ctx context.Context) ([]*UMForceOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/forceOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.autoCloseType != nil {
		r.setParam("autoCloseType", *s.autoCloseType)
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
	var res []*UMForceOrderResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMForceOrderResponse define force order response
type UMForceOrderResponse struct {
	OrderID       int64  `json:"orderId"`
	Symbol        string `json:"symbol"`
	Status        string `json:"status"`
	ClientOrderID string `json:"clientOrderId"`
	Price         string `json:"price"`
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"`
	ExecutedQty   string `json:"executedQty"`
	CumQuote      string `json:"cumQuote"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	OrigType      string `json:"origType"`
	Time          int64  `json:"time"`
	UpdateTime    int64  `json:"updateTime"`
}
