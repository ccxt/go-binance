package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMCancelAllConditionalOrdersService service to cancel all open CM conditional orders
type CMCancelAllConditionalOrdersService struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

// Symbol set symbol
func (s *CMCancelAllConditionalOrdersService) Symbol(symbol string) *CMCancelAllConditionalOrdersService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *CMCancelAllConditionalOrdersService) RecvWindow(recvWindow int64) *CMCancelAllConditionalOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMCancelAllConditionalOrdersService) Do(ctx context.Context) (*CMCancelAllConditionalOrdersResponse, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/cm/conditional/allOpenOrders",
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
	res := new(CMCancelAllConditionalOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMCancelAllConditionalOrdersResponse define cancel all conditional orders response
type CMCancelAllConditionalOrdersResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
