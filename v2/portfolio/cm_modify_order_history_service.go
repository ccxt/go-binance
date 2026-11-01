package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMModifyOrderHistoryService service to get CM order modification history
type CMModifyOrderHistoryService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	startTime         *int64
	endTime           *int64
	limit             *int
	recvWindow        *int64
}

// Symbol set symbol
func (s *CMModifyOrderHistoryService) Symbol(symbol string) *CMModifyOrderHistoryService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CMModifyOrderHistoryService) OrderID(orderID int64) *CMModifyOrderHistoryService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CMModifyOrderHistoryService) OrigClientOrderID(origClientOrderID string) *CMModifyOrderHistoryService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// StartTime set startTime
func (s *CMModifyOrderHistoryService) StartTime(startTime int64) *CMModifyOrderHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *CMModifyOrderHistoryService) EndTime(endTime int64) *CMModifyOrderHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *CMModifyOrderHistoryService) Limit(limit int) *CMModifyOrderHistoryService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *CMModifyOrderHistoryService) RecvWindow(recvWindow int64) *CMModifyOrderHistoryService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMModifyOrderHistoryService) Do(ctx context.Context) ([]*CMModifyOrderHistoryResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/orderAmendment",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
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
	var res []*CMModifyOrderHistoryResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMModifyOrderHistoryResponse define modify order history response
type CMModifyOrderHistoryResponse struct {
	AmendmentID   int64     `json:"amendmentId"`
	Symbol        string    `json:"symbol"`
	Pair          string    `json:"pair"`
	OrderID       int64     `json:"orderId"`
	ClientOrderID string    `json:"clientOrderId"`
	Time          int64     `json:"time"`
	Amendment     Amendment `json:"amendment"`
}
