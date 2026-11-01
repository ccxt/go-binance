package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMModifyOrderService service to modify CM orders
type CMModifyOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	side              string
	quantity          string
	price             string
	recvWindow        *int64
}

// Symbol set symbol
func (s *CMModifyOrderService) Symbol(symbol string) *CMModifyOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CMModifyOrderService) OrderID(orderID int64) *CMModifyOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderId
func (s *CMModifyOrderService) OrigClientOrderID(origClientOrderID string) *CMModifyOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Side set side
func (s *CMModifyOrderService) Side(side SideType) *CMModifyOrderService {
	s.side = string(side)
	return s
}

// Quantity set quantity
func (s *CMModifyOrderService) Quantity(quantity string) *CMModifyOrderService {
	s.quantity = quantity
	return s
}

// Price set price
func (s *CMModifyOrderService) Price(price string) *CMModifyOrderService {
	s.price = price
	return s
}

// RecvWindow set recvWindow
func (s *CMModifyOrderService) RecvWindow(recvWindow int64) *CMModifyOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMModifyOrderService) Do(ctx context.Context) (*CMModifyOrderResponse, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/papi/v1/cm/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	r.setParam("quantity", s.quantity)
	r.setParam("price", s.price)
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
	res := new(CMModifyOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMModifyOrderResponse define modify order response
type CMModifyOrderResponse struct {
	OrderID       int64  `json:"orderId"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	Status        string `json:"status"`
	ClientOrderID string `json:"clientOrderId"`
	Price         string `json:"price"`
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"`
	ExecutedQty   string `json:"executedQty"`
	CumQty        string `json:"cumQty"`
	CumBase       string `json:"cumBase"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	OrigType      string `json:"origType"`
	UpdateTime    int64  `json:"updateTime"`
}
