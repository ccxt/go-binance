package futures

import (
	"context"
	"fmt"
	"net/http"
)

type LvtKlinesService struct {
	c         *Client
	symbol    string // for example, BTCDOWN
	interval  string
	startTime *uint64
	endTime   *uint64
	limit     *uint32
}

type LvtKline struct {
	OpenTime      uint64
	Open          string
	High          string
	Low           string
	Close         string
	CloseLeverage string
	CloseTime     uint64
}

func (s *LvtKlinesService) Symbol(symbol string) *LvtKlinesService {
	s.symbol = symbol
	return s
}

func (s *LvtKlinesService) Interval(interval string) *LvtKlinesService {
	s.interval = interval
	return s
}

func (s *LvtKlinesService) StartTime(startTime uint64) *LvtKlinesService {
	s.startTime = &startTime
	return s
}

func (s *LvtKlinesService) EndTime(endTime uint64) *LvtKlinesService {
	s.endTime = &endTime
	return s
}

func (s *LvtKlinesService) Limit(limit uint32) *LvtKlinesService {
	s.limit = &limit
	return s
}

func (s *LvtKlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*LvtKline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/v1/lvtKlines",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	j, err := newJSON(data)
	if err != nil {
		return []*LvtKline{}, err
	}
	num := len(j.MustArray())
	res = make([]*LvtKline, num)
	for i := 0; i < num; i++ {
		item := j.GetIndex(i)
		if len(item.MustArray()) < 12 {
			err = fmt.Errorf("invalid LvtKline response")
			return []*LvtKline{}, err
		}
		res[i] = &LvtKline{
			OpenTime:      item.GetIndex(0).MustUint64(),
			Open:          item.GetIndex(1).MustString(),
			High:          item.GetIndex(2).MustString(),
			Low:           item.GetIndex(3).MustString(),
			Close:         item.GetIndex(4).MustString(),
			CloseLeverage: item.GetIndex(5).MustString(),
			CloseTime:     item.GetIndex(6).MustUint64(),
		}
	}
	return res, nil
}
