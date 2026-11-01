package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMConditionalOrderHistoryService service to get CM conditional order history
type CMConditionalOrderHistoryService struct {
	c                   *Client
	symbol              string
	strategyID          *int64
	newClientStrategyID *string
	recvWindow          *int64
}

// Symbol set symbol
func (s *CMConditionalOrderHistoryService) Symbol(symbol string) *CMConditionalOrderHistoryService {
	s.symbol = symbol
	return s
}

// StrategyID set strategyId
func (s *CMConditionalOrderHistoryService) StrategyID(strategyID int64) *CMConditionalOrderHistoryService {
	s.strategyID = &strategyID
	return s
}

// NewClientStrategyID set newClientStrategyId
func (s *CMConditionalOrderHistoryService) NewClientStrategyID(newClientStrategyID string) *CMConditionalOrderHistoryService {
	s.newClientStrategyID = &newClientStrategyID
	return s
}

// RecvWindow set recvWindow
func (s *CMConditionalOrderHistoryService) RecvWindow(recvWindow int64) *CMConditionalOrderHistoryService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMConditionalOrderHistoryService) Do(ctx context.Context) (*CMConditionalOrderHistoryResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/conditional/orderHistory",
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
	res := new(CMConditionalOrderHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMConditionalOrderHistoryResponse define conditional order history response
type CMConditionalOrderHistoryResponse struct {
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
	WorkingType         string `json:"workingType"`
	PriceProtect        bool   `json:"priceProtect"`
	PriceMatch          string `json:"priceMatch"`
}
