package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMOpenConditionalOrderService service to get current CM open conditional order
type CMOpenConditionalOrderService struct {
	c                   *Client
	symbol              string
	strategyID          *int64
	newClientStrategyID *string
	recvWindow          *int64
}

// Symbol set symbol
func (s *CMOpenConditionalOrderService) Symbol(symbol string) *CMOpenConditionalOrderService {
	s.symbol = symbol
	return s
}

// StrategyID set strategyId
func (s *CMOpenConditionalOrderService) StrategyID(strategyID int64) *CMOpenConditionalOrderService {
	s.strategyID = &strategyID
	return s
}

// NewClientStrategyID set newClientStrategyId
func (s *CMOpenConditionalOrderService) NewClientStrategyID(newClientStrategyID string) *CMOpenConditionalOrderService {
	s.newClientStrategyID = &newClientStrategyID
	return s
}

// RecvWindow set recvWindow
func (s *CMOpenConditionalOrderService) RecvWindow(recvWindow int64) *CMOpenConditionalOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMOpenConditionalOrderService) Do(ctx context.Context) (*CMOpenConditionalOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/conditional/openOrder",
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
	res := new(CMOpenConditionalOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMOpenConditionalOrderResponse define open conditional order response
type CMOpenConditionalOrderResponse struct {
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
	BookTime            int64  `json:"bookTime"`
	UpdateTime          int64  `json:"updateTime"`
	TimeInForce         string `json:"timeInForce"`
	ActivatePrice       string `json:"activatePrice"`
	PriceRate           string `json:"priceRate"`
}
