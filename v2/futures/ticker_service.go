package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// ListBookTickersService list best price/qty on the order book for a symbol or symbols
type ListBookTickersService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *ListBookTickersService) Symbol(symbol string) *ListBookTickersService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListBookTickersService) Do(ctx context.Context, opts ...RequestOption) (res []*BookTicker, rateLimits map[string]string, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/ticker/bookTicker",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, rateLimits, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return []*BookTicker{}, nil, err
	}
	res = make([]*BookTicker, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*BookTicker{}, nil, err
	}
	return res, rateLimits, nil
}

// BookTicker define book ticker info
type BookTicker struct {
	Symbol       string `json:"symbol"`
	BidPrice     string `json:"bidPrice"`
	BidQuantity  string `json:"bidQty"`
	AskPrice     string `json:"askPrice"`
	AskQuantity  string `json:"askQty"`
	Time         int64  `json:"time"`
	LastUpdateId int64  `json:"lastUpdateId"`
}

// ListPricesService list latest price for a symbol or symbols
type ListPricesService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *ListPricesService) Symbol(symbol string) *ListPricesService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListPricesService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolPrice, rateLimits map[string]string, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/ticker/price",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, rateLimits, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*SymbolPrice{}, nil, err
	}
	data = common.ToJSONList(data)
	res = make([]*SymbolPrice, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*SymbolPrice{}, nil, err
	}
	return res, rateLimits, nil
}

// SymbolPrice define symbol and price pair
type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// ListPriceChangeStatsService show stats of price change in last 24 hours for all symbols
type ListPriceChangeStatsService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *ListPriceChangeStatsService) Symbol(symbol string) *ListPriceChangeStatsService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListPriceChangeStatsService) Do(ctx context.Context, opts ...RequestOption) (res []*PriceChangeStats, rateLimits map[string]string, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/ticker/24hr",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, rateLimits, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, nil, err
	}
	data = common.ToJSONList(data)
	res = make([]*PriceChangeStats, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, nil, err
	}
	return res, rateLimits, nil
}

// PriceChangeStats define price change stats
type PriceChangeStats struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQuantity       string `json:"lastQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}
