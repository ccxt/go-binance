package binance

import (
	"context"
	"fmt"
	"net/http"
)

type UiKlinesService struct {
	c         *Client
	symbol    string
	interval  string
	startTime *uint64
	endTime   *uint64
	timeZone  *string // default: 0(utc)
	limit     *uint32 // default 500, max 1000
}

type UiKline struct {
	OpenTime                 uint64
	Open                     string
	High                     string
	Low                      string
	Close                    string
	Volume                   string
	CloseTime                uint64
	QuoteVolume              string
	TradeNum                 uint64
	TakerBuyBaseAssetVolume  string
	TakerBuyQuoteAssetVolume string
}

func (s *UiKlinesService) Symbol(symbol string) *UiKlinesService {
	s.symbol = symbol
	return s
}

func (s *UiKlinesService) Interval(interval string) *UiKlinesService {
	s.interval = interval
	return s
}

func (s *UiKlinesService) StartTime(startTime uint64) *UiKlinesService {
	s.startTime = &startTime
	return s
}

func (s *UiKlinesService) EndTime(endTime uint64) *UiKlinesService {
	s.endTime = &endTime
	return s
}

func (s *UiKlinesService) TimeZone(timeZone string) *UiKlinesService {
	s.timeZone = &timeZone
	return s
}

func (s *UiKlinesService) Limit(limit uint32) *UiKlinesService {
	s.limit = &limit
	return s
}

func (s *UiKlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*UiKline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/uiKlines",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.timeZone != nil {
		r.setParam("timeZone", *s.timeZone)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	j, err := newJSON(data)
	if err != nil {
		return []*UiKline{}, err
	}
	num := len(j.MustArray())
	res = make([]*UiKline, num)
	for i := 0; i < num; i++ {
		item := j.GetIndex(i)
		if len(item.MustArray()) < 12 {
			err = fmt.Errorf("invalid UiKline response")
			return []*UiKline{}, err
		}
		res[i] = &UiKline{
			OpenTime:                 item.GetIndex(0).MustUint64(),
			Open:                     item.GetIndex(1).MustString(),
			High:                     item.GetIndex(2).MustString(),
			Low:                      item.GetIndex(3).MustString(),
			Close:                    item.GetIndex(4).MustString(),
			Volume:                   item.GetIndex(5).MustString(),
			CloseTime:                item.GetIndex(6).MustUint64(),
			QuoteVolume:              item.GetIndex(7).MustString(),
			TradeNum:                 item.GetIndex(8).MustUint64(),
			TakerBuyBaseAssetVolume:  item.GetIndex(9).MustString(),
			TakerBuyQuoteAssetVolume: item.GetIndex(10).MustString(),
		}
	}
	return res, nil
}
