package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMCancelAllOrdersService service to cancel all active UM orders on a symbol
type UMCancelAllOrdersService struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

// Symbol set symbol
func (s *UMCancelAllOrdersService) Symbol(symbol string) *UMCancelAllOrdersService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *UMCancelAllOrdersService) RecvWindow(recvWindow int64) *UMCancelAllOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMCancelAllOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *UMCancelAllOrdersResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/allOpenOrders",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMCancelAllOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMCancelAllOrdersResponse defines cancel all orders response
type UMCancelAllOrdersResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
