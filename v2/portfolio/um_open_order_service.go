package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMOpenOrderService service to get current UM open order
type UMOpenOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *UMOpenOrderService) Symbol(symbol string) *UMOpenOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *UMOpenOrderService) OrderID(orderID int64) *UMOpenOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderId
func (s *UMOpenOrderService) OrigClientOrderID(origClientOrderID string) *UMOpenOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *UMOpenOrderService) RecvWindow(recvWindow int64) *UMOpenOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMOpenOrderService) Do(ctx context.Context) (*UMOpenOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/openOrder",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(UMOpenOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMOpenOrderResponse define open order response
type UMOpenOrderResponse struct {
	AvgPrice                string `json:"avgPrice"`
	ClientOrderID           string `json:"clientOrderId"`
	CumQuote                string `json:"cumQuote"`
	ExecutedQty             string `json:"executedQty"`
	OrderID                 int64  `json:"orderId"`
	OrigQty                 string `json:"origQty"`
	OrigType                string `json:"origType"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	Status                  string `json:"status"`
	Symbol                  string `json:"symbol"`
	Time                    int64  `json:"time"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	UpdateTime              int64  `json:"updateTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
	PriceMatch              string `json:"priceMatch"`
}
