package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMOpenConditionalOrderService service to get current UM open conditional order
type UMOpenConditionalOrderService struct {
	c                   *Client
	symbol              string
	strategyID          *int64
	newClientStrategyID *string
	recvWindow          *int64
}

// Symbol set symbol
func (s *UMOpenConditionalOrderService) Symbol(symbol string) *UMOpenConditionalOrderService {
	s.symbol = symbol
	return s
}

// StrategyID set strategyId
func (s *UMOpenConditionalOrderService) StrategyID(strategyID int64) *UMOpenConditionalOrderService {
	s.strategyID = &strategyID
	return s
}

// NewClientStrategyID set newClientStrategyId
func (s *UMOpenConditionalOrderService) NewClientStrategyID(newClientStrategyID string) *UMOpenConditionalOrderService {
	s.newClientStrategyID = &newClientStrategyID
	return s
}

// RecvWindow set recvWindow
func (s *UMOpenConditionalOrderService) RecvWindow(recvWindow int64) *UMOpenConditionalOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMOpenConditionalOrderService) Do(ctx context.Context) (*UMOpenConditionalOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/conditional/openOrder",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.strategyID != nil {
		r.setParam("strategyId", *s.strategyID)
	}
	if s.newClientStrategyID != nil {
		r.setParam("newClientStrategyId", *s.newClientStrategyID)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(UMOpenConditionalOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMOpenConditionalOrderResponse define open conditional order response
type UMOpenConditionalOrderResponse struct {
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
	BookTime                int64  `json:"bookTime"`
	UpdateTime              int64  `json:"updateTime"`
	TimeInForce             string `json:"timeInForce"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
	PriceMatch              string `json:"priceMatch"`
}
