package binance

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

type TradingDayTickerService struct {
	c          *Client
	symbol     *string
	symbols    []string
	timeZone   *string // Hours and minutes (e.g. -1:00, 05:45); Only hours (e.g. 0, 8, 4);
	tickerType *string // [FULL, MINT]. default FULL
}

type TradingDayTicker struct {
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
	OpenTime           uint64 `json:"openTime"`
	CloseTime          uint64 `json:"closeTime"`
	FirstId            uint64 `json:"firstId"`
	LastId             uint64 `json:"lastId"`
	Count              uint64 `json:"count"`
}

func (s *TradingDayTickerService) Symbol(symbol string) *TradingDayTickerService {
	s.symbol = &symbol
	return s
}

func (s *TradingDayTickerService) Symbols(symbols []string) *TradingDayTickerService {
	s.symbols = symbols
	return s
}

func (s *TradingDayTickerService) TimeZone(timeZone string) *TradingDayTickerService {
	s.timeZone = &timeZone
	return s
}

func (s *TradingDayTickerService) TickerType(tickerType string) *TradingDayTickerService {
	s.tickerType = &tickerType
	return s
}

func (s *TradingDayTickerService) Do(ctx context.Context, opts ...RequestOption) (res []*TradingDayTicker, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/tradingDay",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.symbols != nil {
		r.setParam("symbols", s.symbols)
	}
	if s.timeZone != nil {
		r.setParam("timeZone", *s.timeZone)
	}
	if s.tickerType != nil {
		r.setParam("type", *s.tickerType)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	res = make([]*TradingDayTicker, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
