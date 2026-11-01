package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMConditionalOrdersService service to get all CM conditional orders
type CMConditionalOrdersService struct {
	c          *Client
	symbol     *string
	strategyID *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *CMConditionalOrdersService) Symbol(symbol string) *CMConditionalOrdersService {
	s.symbol = &symbol
	return s
}

// StrategyID set strategyId
func (s *CMConditionalOrdersService) StrategyID(strategyID int64) *CMConditionalOrdersService {
	s.strategyID = &strategyID
	return s
}

// StartTime set startTime
func (s *CMConditionalOrdersService) StartTime(startTime int64) *CMConditionalOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *CMConditionalOrdersService) EndTime(endTime int64) *CMConditionalOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *CMConditionalOrdersService) Limit(limit int) *CMConditionalOrdersService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *CMConditionalOrdersService) RecvWindow(recvWindow int64) *CMConditionalOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMConditionalOrdersService) Do(ctx context.Context) ([]*CMConditionalOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/conditional/allOrders",
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
	var res []*CMConditionalOrderResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMConditionalOrderResponse define conditional order response
type CMConditionalOrderResponse struct {
	NewClientStrategyID string `json:"newClientStrategyId"`
	StrategyID          int64  `json:"strategyId"`
	StrategyStatus      string `json:"strategyStatus"`
	StrategyType        string `json:"strategyType"`
	OrigQty             string `json:"origQty"`
	Price               string `json:"price"`
	ReduceOnly          bool   `json:"reduceOnly"`
	Side                string `json:"side"`
	PositionSide        string `json:"positionSide"`
	StopPrice           string `json:"stopPrice"`
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	Status              string `json:"status"`
	BookTime            int64  `json:"bookTime"`
	UpdateTime          int64  `json:"updateTime"`
	TriggerTime         int64  `json:"triggerTime"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	ActivatePrice       string `json:"activatePrice"`
	PriceRate           string `json:"priceRate"`
}
