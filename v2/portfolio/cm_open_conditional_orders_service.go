package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMOpenConditionalOrdersService service to get all current CM open conditional orders
type CMOpenConditionalOrdersService struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}

// Symbol set symbol
func (s *CMOpenConditionalOrdersService) Symbol(symbol string) *CMOpenConditionalOrdersService {
	s.symbol = &symbol
	return s
}

// RecvWindow set recvWindow
func (s *CMOpenConditionalOrdersService) RecvWindow(recvWindow int64) *CMOpenConditionalOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMOpenConditionalOrdersService) Do(ctx context.Context) ([]*CMOpenConditionalOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/conditional/openOrders",
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
	var res []*CMOpenConditionalOrderResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
