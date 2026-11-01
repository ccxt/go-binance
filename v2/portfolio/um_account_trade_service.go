package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMAccountTradeService service to get UM account trade list
type UMAccountTradeService struct {
	c          *Client
	symbol     string
	startTime  *int64
	endTime    *int64
	fromID     *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *UMAccountTradeService) Symbol(symbol string) *UMAccountTradeService {
	s.symbol = symbol
	return s
}

// StartTime set startTime
func (s *UMAccountTradeService) StartTime(startTime int64) *UMAccountTradeService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UMAccountTradeService) EndTime(endTime int64) *UMAccountTradeService {
	s.endTime = &endTime
	return s
}

// FromID set fromId
func (s *UMAccountTradeService) FromID(fromID int64) *UMAccountTradeService {
	s.fromID = &fromID
	return s
}

// Limit set limit
func (s *UMAccountTradeService) Limit(limit int) *UMAccountTradeService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *UMAccountTradeService) RecvWindow(recvWindow int64) *UMAccountTradeService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMAccountTradeService) Do(ctx context.Context) ([]*UMAccountTrade, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/userTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
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
	var res []*UMAccountTrade
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMAccountTrade define UM account trade
type UMAccountTrade struct {
	Symbol          string `json:"symbol"`
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	RealizedPnl     string `json:"realizedPnl"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	Buyer           bool   `json:"buyer"`
	Maker           bool   `json:"maker"`
	PositionSide    string `json:"positionSide"`
}
