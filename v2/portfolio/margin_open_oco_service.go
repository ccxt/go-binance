package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginOpenOCOService service to get margin account's open OCO orders
type MarginOpenOCOService struct {
	c          *Client
	recvWindow *int64
}

// RecvWindow set recvWindow
func (s *MarginOpenOCOService) RecvWindow(recvWindow int64) *MarginOpenOCOService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginOpenOCOService) Do(ctx context.Context) ([]*MarginOCOResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/openOrderList",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*MarginOCOResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginOCOOrderDetail define margin OCO order detail
type MarginOCOOrderDetail struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}
