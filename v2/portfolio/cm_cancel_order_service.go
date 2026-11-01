package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMCancelOrderService service to cancel CM orders
type CMCancelOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *CMCancelOrderService) Symbol(symbol string) *CMCancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CMCancelOrderService) OrderID(orderID int64) *CMCancelOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderId
func (s *CMCancelOrderService) OrigClientOrderID(origClientOrderID string) *CMCancelOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *CMCancelOrderService) RecvWindow(recvWindow int64) *CMCancelOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMCancelOrderService) Do(ctx context.Context) (*CMCancelOrderResponse, error) {
	r := &request{
		method:   http.MethodDelete,
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
	res := new(CMCancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMCancelOrderResponse define cancel order response
type CMCancelOrderResponse struct {
	AvgPrice      string `json:"avgPrice"`
	ClientOrderID string `json:"clientOrderId"`
	CumQty        string `json:"cumQty"`
	CumBase       string `json:"cumBase"`
	ExecutedQty   string `json:"executedQty"`
	OrderID       int64  `json:"orderId"`
	OrigQty       string `json:"origQty"`
	Price         string `json:"price"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	Status        string `json:"status"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	UpdateTime    int64  `json:"updateTime"`
}
