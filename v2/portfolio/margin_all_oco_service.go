package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginAllOCOService service to get all OCO orders for a margin account
type MarginAllOCOService struct {
	c          *Client
	fromID     *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
}

// FromID set fromID
func (s *MarginAllOCOService) FromID(fromID int64) *MarginAllOCOService {
	s.fromID = &fromID
	return s
}

// StartTime set startTime
func (s *MarginAllOCOService) StartTime(startTime int64) *MarginAllOCOService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginAllOCOService) EndTime(endTime int64) *MarginAllOCOService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *MarginAllOCOService) Limit(limit int) *MarginAllOCOService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *MarginAllOCOService) RecvWindow(recvWindow int64) *MarginAllOCOService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginAllOCOService) Do(ctx context.Context) ([]*MarginOCOResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/allOrderList",
		secType:  secTypeSigned,
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
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

// MarginOCODetail defines detail of an OCO order
type MarginOCODetail struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}
