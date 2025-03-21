package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMCancelConditionalOrderService service to cancel UM conditional orders
type UMCancelConditionalOrderService struct {
	c                   *Client
	symbol              string
	strategyID          *int64
	newClientStrategyID *string
	recvWindow          *int64
}

// Symbol set symbol
func (s *UMCancelConditionalOrderService) Symbol(symbol string) *UMCancelConditionalOrderService {
	s.symbol = symbol
	return s
}

// StrategyID set strategyId
func (s *UMCancelConditionalOrderService) StrategyID(strategyID int64) *UMCancelConditionalOrderService {
	s.strategyID = &strategyID
	return s
}

// NewClientStrategyID set newClientStrategyId
func (s *UMCancelConditionalOrderService) NewClientStrategyID(newClientStrategyID string) *UMCancelConditionalOrderService {
	s.newClientStrategyID = &newClientStrategyID
	return s
}

// RecvWindow set recvWindow
func (s *UMCancelConditionalOrderService) RecvWindow(recvWindow int64) *UMCancelConditionalOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMCancelConditionalOrderService) Do(ctx context.Context) (*UMCancelConditionalOrderResponse, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/conditional/order",
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
	res := new(UMCancelConditionalOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMCancelConditionalOrderResponse define cancel conditional order response
type UMCancelConditionalOrderResponse struct {
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
	TimeInForce             string `json:"timeInForce"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	BookTime                int64  `json:"bookTime"`
	UpdateTime              int64  `json:"updateTime"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
	PriceMatch              string `json:"priceMatch"`
}
