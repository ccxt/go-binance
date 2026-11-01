package options

import (
	"context"
	"encoding/json"
	"net/http"
)

type Ticker struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	Open               string `json:"open"`
	High               string `json:"high"`
	Low                string `json:"low"`
	Volume             string `json:"volume"`
	Amount             string `json:"amount"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstTradeId       int    `json:"firstTradeId"`
	TradeCount         int    `json:"tradeCount"`
	StrikePrice        string `json:"strikePrice"`
	ExercisePrice      string `json:"exercisePrice"`
}

// TickerService list recent trades in orderbook
type TickerService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *TickerService) Symbol(symbol string) *TickerService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *TickerService) Do(ctx context.Context, opts ...RequestOption) (res []*Ticker, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/ticker",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Ticker{}, err
	}
	res = make([]*Ticker, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Ticker{}, err
	}
	return res, nil
}
