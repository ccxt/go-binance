package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMCancelAllOrdersService service to cancel all open CM orders
type CMCancelAllOrdersService struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

// Symbol set symbol
func (s *CMCancelAllOrdersService) Symbol(symbol string) *CMCancelAllOrdersService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *CMCancelAllOrdersService) RecvWindow(recvWindow int64) *CMCancelAllOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMCancelAllOrdersService) Do(ctx context.Context) (*CMCancelAllOrdersResponse, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/cm/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(CMCancelAllOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMCancelAllOrdersResponse define cancel all orders response
type CMCancelAllOrdersResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
