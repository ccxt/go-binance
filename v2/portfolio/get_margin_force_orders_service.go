package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMarginForceOrdersService service to get user's margin force orders
type GetMarginForceOrdersService struct {
	c          *Client
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

// StartTime set startTime
func (s *GetMarginForceOrdersService) StartTime(startTime int64) *GetMarginForceOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetMarginForceOrdersService) EndTime(endTime int64) *GetMarginForceOrdersService {
	s.endTime = &endTime
	return s
}

// Current set current page
func (s *GetMarginForceOrdersService) Current(current int64) *GetMarginForceOrdersService {
	s.current = &current
	return s
}

// Size set page size
func (s *GetMarginForceOrdersService) Size(size int64) *GetMarginForceOrdersService {
	s.size = &size
	return s
}

// RecvWindow set recvWindow
func (s *GetMarginForceOrdersService) RecvWindow(recvWindow int64) *GetMarginForceOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetMarginForceOrdersService) Do(ctx context.Context) (*MarginForceOrdersResponse, error) {
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
