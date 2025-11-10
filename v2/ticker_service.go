package binance

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

// buildRequest creates the API request for ListBookTickers
func (s *ListBookTickersService) buildRequest() *request {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/bookTicker",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	return r
}

// Symbol set symbol
func (s *ListBookTickersService) Symbol(symbol string) *ListBookTickersService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListBookTickersService) Do(ctx context.Context, opts ...RequestOption) (res []*BookTicker, err error) {
	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return []*BookTicker{}, err
	}
	res = make([]*BookTicker, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*BookTicker{}, err
	}
	return res, nil
}

// BookTicker define book ticker info
type BookTicker struct {
	Symbol      string `json:"symbol"`
	BidPrice    string `json:"bidPrice"`
	BidQuantity string `json:"bidQty"`
	AskPrice    string `json:"askPrice"`
	AskQuantity string `json:"askQty"`
}

// ListPricesService list latest price for a symbol or symbols
type ListPricesService struct {
	c       *Client
	symbol  *string
	symbols []string
}

// buildRequest creates the API request for ListPrices
func (s *ListPricesService) buildRequest() *request {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/price",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	} else if s.symbols != nil {
		s, _ := json.Marshal(s.symbols)
		r.setParam("symbols", string(s))
	}
	return r
}

// Symbol set symbol
func (s *ListPricesService) Symbol(symbol string) *ListPricesService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListPricesService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolPrice, err error) {
	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*SymbolPrice{}, err
	}
	data = common.ToJSONList(data)
	res = make([]*SymbolPrice, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*SymbolPrice{}, err
	}
	return res, nil
}

// SymbolPrice define symbol and price pair
type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// ListPriceChangeStatsService show stats of price change in last 24 hours for all symbols
type ListPriceChangeStatsService struct {
	c       *Client
	symbol  *string
	symbols []string
}

// buildRequest creates the API request for ListPriceChangeStats
func (s *ListPriceChangeStatsService) buildRequest() *request {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/24hr",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	} else if s.symbols != nil {
		r.setParam("symbols", s.symbols)
	}
	return r
}

// Symbol set symbol
func (s *ListPriceChangeStatsService) Symbol(symbol string) *ListPriceChangeStatsService {
	s.symbol = &symbol
	return s
}

// Symbols set symbols
func (s *ListPriceChangeStatsService) Symbols(symbols []string) *ListPriceChangeStatsService {
	s.symbols = symbols
	return s
}

// Symbols set symbols
func (s *ListPricesService) Symbols(symbols []string) *ListPricesService {
	s.symbols = symbols
	return s
}

// Do send request
func (s *ListPriceChangeStatsService) Do(ctx context.Context, opts ...RequestOption) (res []*PriceChangeStats, err error) {
	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	data = common.ToJSONList(data)
	res = make([]*PriceChangeStats, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PriceChangeStats define price change stats
type PriceChangeStats struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
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

// AveragePriceService show current average price for a symbol
type AveragePriceService struct {
	c      *Client
	symbol string
}

// buildRequest creates the API request for AveragePrice
func (s *AveragePriceService) buildRequest() *request {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/avgPrice",
	}
	r.setParam("symbol", s.symbol)
	return r
}

// Symbol set symbol
func (s *AveragePriceService) Symbol(symbol string) *AveragePriceService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *AveragePriceService) Do(ctx context.Context, opts ...RequestOption) (res *AvgPrice, err error) {
	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	res = new(AvgPrice)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AvgPrice define average price
type AvgPrice struct {
	Mins  int64  `json:"mins"`
	Price string `json:"price"`
}

type ListSymbolTickerService struct {
	c          *Client
	symbol     *string
	symbols    []string
	windowSize *string
}

// buildRequest creates the API request for ListSymbolTicker
func (s *ListSymbolTickerService) buildRequest() *request {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	} else if s.symbols != nil {
		s, _ := json.Marshal(s.symbols)
		r.setParam("symbols", string(s))
	}
	if s.windowSize != nil {
		r.setParam("windowSize", *s.windowSize)
	}
	return r
}

type SymbolTicker struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}

func (s *ListSymbolTickerService) Symbol(symbol string) *ListSymbolTickerService {
	s.symbol = &symbol
	return s
}

func (s *ListSymbolTickerService) Symbols(symbols []string) *ListSymbolTickerService {
	s.symbols = symbols
	return s
}

// Defaults to 1d if no parameter provided
//
// Supported windowSize values:
//
// - 1m,2m....59m for minutes
//
// - 1h, 2h....23h - for hours
//
// - 1d...7d - for days
//
// Units cannot be combined (e.g. 1d2h is not allowed).
//
// Reference: https://binance-docs.github.io/apidocs/spot/en/#rolling-window-price-change-statistics
func (s *ListSymbolTickerService) WindowSize(windowSize string) *ListSymbolTickerService {
	s.windowSize = &windowSize
	return s
}

func (s *ListSymbolTickerService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolTicker, err error) {
	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return []*SymbolTicker{}, err
	}
	res = make([]*SymbolTicker, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*SymbolTicker{}, err
	}
	return res, nil
}

// DoSBE sends the request with SBE encoding and returns the decoded response
func (s *ListBookTickersService) DoSBE(ctx context.Context, opts ...RequestOption) (res []*BookTicker, err error) {
	// Add SBE headers
	opts = append(opts, WithSBE(3, 1))

	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	// Decode SBE response using centralized decoder
	res = make([]*BookTicker, 0)
	if err := sbeDecoder.DecodeResponse(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// DoSBE sends the request with SBE encoding and returns the decoded response
func (s *ListPricesService) DoSBE(ctx context.Context, opts ...RequestOption) (res []*SymbolPrice, err error) {
	// Add SBE headers
	opts = append(opts, WithSBE(3, 1))

	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	// Decode SBE response using centralized decoder
	res = make([]*SymbolPrice, 0)
	if err := sbeDecoder.DecodeResponse(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// DoSBE sends the request with SBE encoding and returns the decoded response
func (s *ListPriceChangeStatsService) DoSBE(ctx context.Context, opts ...RequestOption) (res []*PriceChangeStats, err error) {
	// Add SBE headers
	opts = append(opts, WithSBE(3, 1))

	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	// Decode SBE response using centralized decoder
	res = make([]*PriceChangeStats, 0)
	if err := sbeDecoder.DecodeResponse(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// DoSBE sends the request with SBE encoding and returns the decoded response
func (s *AveragePriceService) DoSBE(ctx context.Context, opts ...RequestOption) (res *AvgPrice, err error) {
	// Add SBE headers
	opts = append(opts, WithSBE(3, 1))

	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	// Decode SBE response using centralized decoder
	res = &AvgPrice{}
	if err := sbeDecoder.DecodeResponse(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DoSBE sends the request with SBE encoding and returns the decoded response
func (s *ListSymbolTickerService) DoSBE(ctx context.Context, opts ...RequestOption) (res []*SymbolTicker, err error) {
	// Add SBE headers
	opts = append(opts, WithSBE(3, 1))

	r := s.buildRequest()
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	// Decode SBE response using centralized decoder
	res = make([]*SymbolTicker, 0)
	if err := sbeDecoder.DecodeResponse(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}
