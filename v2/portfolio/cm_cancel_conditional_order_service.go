package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMCancelConditionalOrderService service to cancel CM conditional orders
type CMCancelConditionalOrderService struct {
	c                   *Client
	symbol              string
	strategyID          *int64
	newClientStrategyID *string
	recvWindow          *int64
}

// Symbol set symbol
func (s *CMCancelConditionalOrderService) Symbol(symbol string) *CMCancelConditionalOrderService {
	s.symbol = symbol
	return s
}

// StrategyID set strategyId
func (s *CMCancelConditionalOrderService) StrategyID(strategyID int64) *CMCancelConditionalOrderService {
	s.strategyID = &strategyID
	return s
}

// NewClientStrategyID set newClientStrategyId
func (s *CMCancelConditionalOrderService) NewClientStrategyID(newClientStrategyID string) *CMCancelConditionalOrderService {
	s.newClientStrategyID = &newClientStrategyID
	return s
}

// RecvWindow set recvWindow
func (s *CMCancelConditionalOrderService) RecvWindow(recvWindow int64) *CMCancelConditionalOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMCancelConditionalOrderService) Do(ctx context.Context) (*CMCancelConditionalOrderResponse, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/cm/conditional/order",
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
	res := new(CMCancelConditionalOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMCancelConditionalOrderResponse define cancel conditional order response
type CMCancelConditionalOrderResponse struct {
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
	TimeInForce         string `json:"timeInForce"`
	ActivatePrice       string `json:"activatePrice"`
	PriceRate           string `json:"priceRate"`
	BookTime            int64  `json:"bookTime"`
	UpdateTime          int64  `json:"updateTime"`
	WorkingType         string `json:"workingType"`
	PriceProtect        bool   `json:"priceProtect"`
}
