package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginForceOrdersService service to get user's margin force orders
type MarginForceOrdersService struct {
	c          *Client
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

// StartTime set startTime
func (s *MarginForceOrdersService) StartTime(startTime int64) *MarginForceOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginForceOrdersService) EndTime(endTime int64) *MarginForceOrdersService {
	s.endTime = &endTime
	return s
}

// Current set current page
func (s *MarginForceOrdersService) Current(current int64) *MarginForceOrdersService {
	s.current = &current
	return s
}

// Size set page size
func (s *MarginForceOrdersService) Size(size int64) *MarginForceOrdersService {
	s.size = &size
	return s
}

// RecvWindow set recvWindow
func (s *MarginForceOrdersService) RecvWindow(recvWindow int64) *MarginForceOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginForceOrdersService) Do(ctx context.Context) (*MarginForceOrdersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/forceOrders",
		secType:  secTypeSigned,
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(MarginForceOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginForceOrder define margin force order
type MarginForceOrder struct {
	AvgPrice    string `json:"avgPrice"`
	ExecutedQty string `json:"executedQty"`
	OrderID     int64  `json:"orderId"`
	Price       string `json:"price"`
	Qty         string `json:"qty"`
	Side        string `json:"side"`
	Symbol      string `json:"symbol"`
	TimeInForce string `json:"timeInForce"`
	UpdatedTime int64  `json:"updatedTime"`
}

// MarginForceOrdersResponse define margin force orders response
type MarginForceOrdersResponse struct {
	Rows  []*MarginForceOrder `json:"rows"`
	Total int64               `json:"total"`
}
