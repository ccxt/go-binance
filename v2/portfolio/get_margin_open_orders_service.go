package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMarginOpenOrdersService service to get current margin open orders
type GetMarginOpenOrdersService struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}

// Symbol set symbol
func (s *GetMarginOpenOrdersService) Symbol(symbol string) *GetMarginOpenOrdersService {
	s.symbol = &symbol
	return s
}

// RecvWindow set recvWindow
func (s *GetMarginOpenOrdersService) RecvWindow(recvWindow int64) *GetMarginOpenOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetMarginOpenOrdersService) Do(ctx context.Context) ([]*MarginOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*MarginOrder
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
