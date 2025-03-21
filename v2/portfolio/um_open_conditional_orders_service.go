package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMOpenConditionalOrdersService service to get all current UM open conditional orders
type UMOpenConditionalOrdersService struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}

// Symbol set symbol
func (s *UMOpenConditionalOrdersService) Symbol(symbol string) *UMOpenConditionalOrdersService {
	s.symbol = &symbol
	return s
}

// RecvWindow set recvWindow
func (s *UMOpenConditionalOrdersService) RecvWindow(recvWindow int64) *UMOpenConditionalOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMOpenConditionalOrdersService) Do(ctx context.Context) ([]*UMOpenConditionalOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/conditional/openOrders",
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
	var res []*UMOpenConditionalOrderResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
