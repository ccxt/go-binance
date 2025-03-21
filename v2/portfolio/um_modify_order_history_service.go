package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMModifyOrderHistoryService service to get UM order modification history
type UMModifyOrderHistoryService struct {
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
func (s *UMModifyOrderHistoryService) Symbol(symbol string) *UMModifyOrderHistoryService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *UMModifyOrderHistoryService) OrderID(orderID int64) *UMModifyOrderHistoryService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *UMModifyOrderHistoryService) OrigClientOrderID(origClientOrderID string) *UMModifyOrderHistoryService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// StartTime set startTime
func (s *UMModifyOrderHistoryService) StartTime(startTime int64) *UMModifyOrderHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UMModifyOrderHistoryService) EndTime(endTime int64) *UMModifyOrderHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *UMModifyOrderHistoryService) Limit(limit int) *UMModifyOrderHistoryService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *UMModifyOrderHistoryService) RecvWindow(recvWindow int64) *UMModifyOrderHistoryService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMModifyOrderHistoryService) Do(ctx context.Context) ([]*UMModifyOrderHistoryResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/orderAmendment",
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
	var res []*UMModifyOrderHistoryResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AmendmentDetail define amendment detail
type AmendmentDetail struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

// Amendment define amendment information
type Amendment struct {
	Price   AmendmentDetail `json:"price"`
	OrigQty AmendmentDetail `json:"origQty"`
	Count   int             `json:"count"`
}

// UMModifyOrderHistoryResponse define modify order history response
type UMModifyOrderHistoryResponse struct {
	AmendmentID   int64     `json:"amendmentId"`
	Symbol        string    `json:"symbol"`
	Pair          string    `json:"pair"`
	OrderID       int64     `json:"orderId"`
	ClientOrderID string    `json:"clientOrderId"`
	Time          int64     `json:"time"`
	Amendment     Amendment `json:"amendment"`
	PriceMatch    string    `json:"priceMatch"`
}
