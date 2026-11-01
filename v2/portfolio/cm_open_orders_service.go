package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMOpenOrdersService service to get all current CM open orders
type CMOpenOrdersService struct {
	c          *Client
	symbol     *string
	pair       *string
	recvWindow *int64
}

// Symbol set symbol
func (s *CMOpenOrdersService) Symbol(symbol string) *CMOpenOrdersService {
	s.symbol = &symbol
	return s
}

// Pair set pair
func (s *CMOpenOrdersService) Pair(pair string) *CMOpenOrdersService {
	s.pair = &pair
	return s
}

// RecvWindow set recvWindow
func (s *CMOpenOrdersService) RecvWindow(recvWindow int64) *CMOpenOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMOpenOrdersService) Do(ctx context.Context) ([]*CMOpenOrdersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.pair != nil {
		r.setParam("pair", *s.pair)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*CMOpenOrdersResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMOpenOrdersResponse define open orders response
type CMOpenOrdersResponse struct {
	AvgPrice      string `json:"avgPrice"`
	ClientOrderID string `json:"clientOrderId"`
	CumBase       string `json:"cumBase"`
	ExecutedQty   string `json:"executedQty"`
	OrderID       int64  `json:"orderId"`
	OrigQty       string `json:"origQty"`
	OrigType      string `json:"origType"`
	Price         string `json:"price"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	Status        string `json:"status"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	Time          int64  `json:"time"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	UpdateTime    int64  `json:"updateTime"`
}
