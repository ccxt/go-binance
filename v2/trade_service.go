package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListTradesService list trades
type ListTradesService struct {
	c         *Client
	symbol    string
	startTime *int64
	endTime   *int64
	limit     *int
	fromID    *int64
	orderId   *int64
}

// Symbol set symbol
func (s *ListTradesService) Symbol(symbol string) *ListTradesService {
	s.symbol = symbol
	return s
}

// StartTime set starttime
func (s *ListTradesService) StartTime(startTime int64) *ListTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListTradesService) EndTime(endTime int64) *ListTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListTradesService) Limit(limit int) *ListTradesService {
	s.limit = &limit
	return s
}

// FromID set fromID
func (s *ListTradesService) FromID(fromID int64) *ListTradesService {
	s.fromID = &fromID
	return s
}

// OrderId set OrderId
func (s *ListTradesService) OrderId(OrderId int64) *ListTradesService {
	s.orderId = &OrderId
	return s
}

// Do send request
func (s *ListTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*TradeV3, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/myTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*TradeV3{}, err
	}
	res = make([]*TradeV3, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*TradeV3{}, err
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
func (s *HistoricalTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*Trade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/historicalTrades",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*Trade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// Trade define trade info
type Trade struct {
	ID            int64  `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"qty"`
	QuoteQuantity string `json:"quoteQty"`
	Time          int64  `json:"time"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
	IsBestMatch   bool   `json:"isBestMatch"`
	IsIsolated    bool   `json:"isIsolated"`
}

// TradeV3 define v3 trade info
type TradeV3 struct {
	ID              int64  `json:"id"`
	Symbol          string `json:"symbol"`
	OrderID         int64  `json:"orderId"`
	OrderListId     int64  `json:"orderListId"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	QuoteQuantity   string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
	IsIsolated      bool   `json:"isIsolated"`
}

// AggTradesService list aggregate trades
type AggTradesService struct {
	c         *Client
	symbol    string
	fromID    *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *AggTradesService) Symbol(symbol string) *AggTradesService {
	s.symbol = symbol
	return s
}

// FromID set fromID
func (s *AggTradesService) FromID(fromID int64) *AggTradesService {
	s.fromID = &fromID
	return s
}

// StartTime set startTime
func (s *AggTradesService) StartTime(startTime int64) *AggTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *AggTradesService) EndTime(endTime int64) *AggTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *AggTradesService) Limit(limit int) *AggTradesService {
	s.limit = &limit
	return s
}

// Do send request
func (s *AggTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*AggTrade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/aggTrades",
	}
	r.setParam("symbol", s.symbol)
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
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*AggTrade{}, err
	}
	res = make([]*AggTrade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*AggTrade{}, err
	}
	return res, nil
}

// AggTrade define aggregate trade info
type AggTrade struct {
	AggTradeID       int64  `json:"a"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int64  `json:"f"`
	LastTradeID      int64  `json:"l"`
	Timestamp        int64  `json:"T"`
	IsBuyerMaker     bool   `json:"m"`
	IsBestPriceMatch bool   `json:"M"`
}

// RecentTradesService list recent trades
type RecentTradesService struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *RecentTradesService) Symbol(symbol string) *RecentTradesService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *RecentTradesService) Limit(limit int) *RecentTradesService {
	s.limit = &limit
	return s
}

// Do send request
func (s *RecentTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*Trade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v1/trades",
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
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

// RecentTradesV3Service list recent trades v3
type RecentTradesV3Service struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *RecentTradesV3Service) Symbol(symbol string) *RecentTradesV3Service {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *RecentTradesV3Service) Limit(limit int) *RecentTradesV3Service {
	s.limit = &limit
	return s
}

// Do send request
func (s *RecentTradesV3Service) Do(ctx context.Context, opts ...RequestOption) (res []*Trade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/trades",
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
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
