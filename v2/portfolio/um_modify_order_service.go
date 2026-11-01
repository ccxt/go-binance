package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMModifyOrderService service to modify UM orders
type UMModifyOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	side              string
	quantity          string
	price             *string
	priceMatch        *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *UMModifyOrderService) Symbol(symbol string) *UMModifyOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *UMModifyOrderService) OrderID(orderID int64) *UMModifyOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderId
func (s *UMModifyOrderService) OrigClientOrderID(origClientOrderID string) *UMModifyOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Side set side
func (s *UMModifyOrderService) Side(side SideType) *UMModifyOrderService {
	s.side = string(side)
	return s
}

// Quantity set quantity
func (s *UMModifyOrderService) Quantity(quantity string) *UMModifyOrderService {
	s.quantity = quantity
	return s
}

// Price set price
func (s *UMModifyOrderService) Price(price string) *UMModifyOrderService {
	s.price = &price
	return s
}

// PriceMatch set priceMatch
func (s *UMModifyOrderService) PriceMatch(priceMatch string) *UMModifyOrderService {
	s.priceMatch = &priceMatch
	return s
}

// RecvWindow set recvWindow
func (s *UMModifyOrderService) RecvWindow(recvWindow int64) *UMModifyOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMModifyOrderService) Do(ctx context.Context) (*UMModifyOrderResponse, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	r.setParam("quantity", s.quantity)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.price != nil {
		r.setParam("price", *s.price)
	}
	if s.priceMatch != nil {
		r.setParam("priceMatch", *s.priceMatch)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(UMModifyOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMModifyOrderResponse define modify order response
type UMModifyOrderResponse struct {
	OrderID                 int64  `json:"orderId"`
	Symbol                  string `json:"symbol"`
	Status                  string `json:"status"`
	ClientOrderID           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	AvgPrice                string `json:"avgPrice"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumQty                  string `json:"cumQty"`
	CumQuote                string `json:"cumQuote"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	OrigType                string `json:"origType"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
	UpdateTime              int64  `json:"updateTime"`
	PriceMatch              string `json:"priceMatch"`
}
