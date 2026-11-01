package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginCancelOrderService service to cancel margin account orders
type MarginCancelOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	newClientOrderID  *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *MarginCancelOrderService) Symbol(symbol string) *MarginCancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *MarginCancelOrderService) OrderID(orderID int64) *MarginCancelOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderId
func (s *MarginCancelOrderService) OrigClientOrderID(origClientOrderID string) *MarginCancelOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// NewClientOrderID set newClientOrderId
func (s *MarginCancelOrderService) NewClientOrderID(newClientOrderID string) *MarginCancelOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *MarginCancelOrderService) RecvWindow(recvWindow int64) *MarginCancelOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginCancelOrderService) Do(ctx context.Context) (*MarginCancelOrderResponse, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/margin/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.newClientOrderID != nil {
		r.setParam("newClientOrderId", *s.newClientOrderID)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(MarginCancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginCancelOrderResponse define cancel order response
type MarginCancelOrderResponse struct {
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	OrigClientOrderID   string `json:"origClientOrderId"`
	ClientOrderID       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}
