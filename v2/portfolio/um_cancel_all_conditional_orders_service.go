package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMCancelAllConditionalOrdersService service to cancel all open UM conditional orders
type UMCancelAllConditionalOrdersService struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

// Symbol set symbol
func (s *UMCancelAllConditionalOrdersService) Symbol(symbol string) *UMCancelAllConditionalOrdersService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *UMCancelAllConditionalOrdersService) RecvWindow(recvWindow int64) *UMCancelAllConditionalOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMCancelAllConditionalOrdersService) Do(ctx context.Context) (*UMCancelAllConditionalOrdersResponse, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/conditional/allOpenOrders",
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
	res := new(UMCancelAllConditionalOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMCancelAllConditionalOrdersResponse define cancel all conditional orders response
type UMCancelAllConditionalOrdersResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
