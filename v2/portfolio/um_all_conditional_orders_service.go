package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMAllConditionalOrdersService service to get all UM conditional orders
type UMAllConditionalOrdersService struct {
	c          *Client
	symbol     *string
	strategyID *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *UMAllConditionalOrdersService) Symbol(symbol string) *UMAllConditionalOrdersService {
	s.symbol = &symbol
	return s
}

// StrategyID set strategyId
func (s *UMAllConditionalOrdersService) StrategyID(strategyID int64) *UMAllConditionalOrdersService {
	s.strategyID = &strategyID
	return s
}

// StartTime set startTime
func (s *UMAllConditionalOrdersService) StartTime(startTime int64) *UMAllConditionalOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UMAllConditionalOrdersService) EndTime(endTime int64) *UMAllConditionalOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *UMAllConditionalOrdersService) Limit(limit int) *UMAllConditionalOrdersService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *UMAllConditionalOrdersService) RecvWindow(recvWindow int64) *UMAllConditionalOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMAllConditionalOrdersService) Do(ctx context.Context) ([]*UMConditionalOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/conditional/allOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.strategyID != nil {
		r.setParam("strategyId", *s.strategyID)
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
	var res []*UMConditionalOrderResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMConditionalOrderResponse define conditional order response
type UMConditionalOrderResponse struct {
	NewClientStrategyID     string `json:"newClientStrategyId"`
	StrategyID              int64  `json:"strategyId"`
	StrategyStatus          string `json:"strategyStatus"`
	StrategyType            string `json:"strategyType"`
	OrigQty                 string `json:"origQty"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	StopPrice               string `json:"stopPrice"`
	Symbol                  string `json:"symbol"`
	OrderID                 int64  `json:"orderId"`
	Status                  string `json:"status"`
	BookTime                int64  `json:"bookTime"`
	UpdateTime              int64  `json:"updateTime"`
	TriggerTime             int64  `json:"triggerTime"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
	PriceMatch              string `json:"priceMatch"`
}
