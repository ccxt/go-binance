package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMCancelOrderService service to cancel UM orders
type UMCancelOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *UMCancelOrderService) Symbol(symbol string) *UMCancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *UMCancelOrderService) OrderID(orderID int64) *UMCancelOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *UMCancelOrderService) OrigClientOrderID(origClientOrderID string) *UMCancelOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *UMCancelOrderService) RecvWindow(recvWindow int64) *UMCancelOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMCancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *UMOrder, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/order",
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

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
