package options

import (
	"context"
	"encoding/json"
	"net/http"
)

type Trade struct {
	Id       uint64 `json:"id"`
	TradeId  int    `json:"tradeId"`
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	Qty      string `json:"qty"`
	QuoteQty string `json:"quoteQty"`
	Side     int    `json:"side"`
	Time     uint64 `json:"time"`
}

// TradesService list recent trades in orderbook
type TradesService struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *TradesService) Symbol(symbol string) *TradesService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *TradesService) Limit(limit int) *TradesService {
	s.limit = &limit
	return s
}

// Do send request
func (s *TradesService) Do(ctx context.Context, opts ...RequestOption) (res []*Trade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/trades",
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Trade{}, err
	}
	res = make([]*Trade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Trade{}, err
	}
	return res, nil
}

// HistoricalTradesService trades
type HistoricalTradesService struct {
	c      *Client
	symbol string
	limit  *int
	fromID *int64
}

// Symbol set symbol
func (s *HistoricalTradesService) Symbol(symbol string) *HistoricalTradesService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *HistoricalTradesService) Limit(limit int) *HistoricalTradesService {
	s.limit = &limit
	return s
}

// FromID set fromID
func (s *HistoricalTradesService) FromID(fromID int64) *HistoricalTradesService {
	s.fromID = &fromID
	return s
}

// Do send request
func (s *HistoricalTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*HistoricalTrade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/historicalTrades",
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*HistoricalTrade{}, err
	}
	res = make([]*HistoricalTrade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*HistoricalTrade{}, err
	}
	return res, nil
}

// HistoricalTrade define historical trade info
type HistoricalTrade struct {
	Id       uint64 `json:"id"`
	TradeId  int    `json:"tradeId"`
	Price    string `json:"price"`
	Qty      string `json:"qty"`
	QuoteQty string `json:"quoteQty"`
	Side     int    `json:"side"`
	Time     uint64 `json:"time"`
}
