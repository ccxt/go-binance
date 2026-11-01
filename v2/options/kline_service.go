package options

import (
	"context"
	"fmt"
	"net/http"
)

// KlinesService list klines
type KlinesService struct {
	c         *Client
	symbol    string
	interval  string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *KlinesService) Symbol(symbol string) *KlinesService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *KlinesService) Interval(interval string) *KlinesService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *KlinesService) Limit(limit int) *KlinesService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *KlinesService) StartTime(startTime int64) *KlinesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *KlinesService) EndTime(endTime int64) *KlinesService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *KlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/klines",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Kline{}, err
	}
	j, err := newJSON(data)
	if err != nil {
		return []*Kline{}, err
	}
	num := len(j.MustArray())
	res = make([]*Kline, num)
	for i := 0; i < num; i++ {
		item := j.GetIndex(i)
		if len(item.MustMap()) < 12 {
			err = fmt.Errorf("invalid kline response")
			return []*Kline{}, err
		}
		res[i] = &Kline{
			OpenTime:    item.Get("openTime").MustInt64(),
			Open:        item.Get("open").MustString(),
			High:        item.Get("high").MustString(),
			Low:         item.Get("low").MustString(),
			Close:       item.Get("close").MustString(),
			CloseTime:   item.Get("closeTime").MustInt64(),
			Amount:      item.Get("amount").MustString(),
			TakerAmount: item.Get("takerAmount").MustString(),
			Volume:      item.Get("volume").MustString(),
			TakerVolume: item.Get("takerVolume").MustString(),
			Interval:    item.Get("interval").MustString(),
			TradeCount:  item.Get("tradeCount").MustInt64(),
		}
	}
	return res, nil
}

// Kline define kline info
type Kline struct {
	Open        string `json:"open"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Close       string `json:"close"`
	Volume      string `json:"volume"`
	Amount      string `json:"amount"`
	Interval    string `json:"interval"`
	TradeCount  int64  `json:"tradeCount"`
	TakerVolume string `json:"takerVolume"`
	TakerAmount string `json:"takerAmount"`
	OpenTime    int64  `json:"openTime"`
	CloseTime   int64  `json:"closeTime"`
}
