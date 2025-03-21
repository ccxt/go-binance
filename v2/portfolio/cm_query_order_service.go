package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMQueryOrderService service to query CM orders
type CMQueryOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *CMQueryOrderService) Symbol(symbol string) *CMQueryOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CMQueryOrderService) OrderID(orderID int64) *CMQueryOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderId
func (s *CMQueryOrderService) OrigClientOrderID(origClientOrderID string) *CMQueryOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *CMQueryOrderService) RecvWindow(recvWindow int64) *CMQueryOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMQueryOrderService) Do(ctx context.Context) (*CMQueryOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/order",
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
	res := new(CMQueryOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMQueryOrderResponse define query order response
type CMQueryOrderResponse struct {
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
	Status        string `json:"status"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	PositionSide  string `json:"positionSide"`
	Time          int64  `json:"time"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	UpdateTime    int64  `json:"updateTime"`
}
