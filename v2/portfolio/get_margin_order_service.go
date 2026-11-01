package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMarginOrderService service to get current margin open orders
type GetMarginOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *GetMarginOrderService) Symbol(symbol string) *GetMarginOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetMarginOrderService) OrderID(orderID int64) *GetMarginOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetMarginOrderService) OrigClientOrderID(origClientOrderID string) *GetMarginOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *GetMarginOrderService) RecvWindow(recvWindow int64) *GetMarginOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetMarginOrderService) Do(ctx context.Context) (*MarginOrder, error) {
	r := &request{
		method:   http.MethodGet,
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
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(MarginOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
