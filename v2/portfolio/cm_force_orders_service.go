package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMForceOrdersService service to get user's CM force orders
type CMForceOrdersService struct {
	c             *Client
	symbol        *string
	autoCloseType *string
	startTime     *int64
	endTime       *int64
	limit         *int
	recvWindow    *int64
}

// Symbol set symbol
func (s *CMForceOrdersService) Symbol(symbol string) *CMForceOrdersService {
	s.symbol = &symbol
	return s
}

// AutoCloseType set autoCloseType
func (s *CMForceOrdersService) AutoCloseType(autoCloseType string) *CMForceOrdersService {
	s.autoCloseType = &autoCloseType
	return s
}

// StartTime set startTime
func (s *CMForceOrdersService) StartTime(startTime int64) *CMForceOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *CMForceOrdersService) EndTime(endTime int64) *CMForceOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *CMForceOrdersService) Limit(limit int) *CMForceOrdersService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *CMForceOrdersService) RecvWindow(recvWindow int64) *CMForceOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMForceOrdersService) Do(ctx context.Context) ([]*CMForceOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/forceOrders",
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
	var res []*CMForceOrderResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMForceOrderResponse define force order response
type CMForceOrderResponse struct {
	OrderID       int64  `json:"orderId"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	Status        string `json:"status"`
	ClientOrderID string `json:"clientOrderId"`
	Price         string `json:"price"`
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"`
	ExecutedQty   string `json:"executedQty"`
	CumBase       string `json:"cumBase"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	OrigType      string `json:"origType"`
	Time          int64  `json:"time"`
	UpdateTime    int64  `json:"updateTime"`
}
