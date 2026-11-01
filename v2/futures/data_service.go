package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type DeliveryPriceService struct {
	c    *Client
	pair string
}

type DeliveryPrice struct {
	DeliveryTime  uint64  `json:"deliveryTime"`  // deliveryTime
	DeliveryPrice float64 `json:"deliveryPrice"` // deliveryPrice
}

func (s *DeliveryPriceService) Pair(pair string) *DeliveryPriceService {
	s.pair = pair
	return s
}

func (s *DeliveryPriceService) Do(ctx context.Context, opts ...RequestOption) (res []*DeliveryPrice, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/delivery-price",
	}
	r.setParam("pair", s.pair)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*DeliveryPrice, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type TopLongShortAccountRatioService struct {
	c         *Client
	symbol    string
	period    string  // "5m","15m","30m","1h","2h","4h","6h","12h","1d"
	limit     *uint32 // default 30, max 500
	startTime *uint64
	endTime   *uint64
}

type TopLongShortAccountRatio struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      uint64 `json:"timestamp"`
}

func (s *TopLongShortAccountRatioService) Symbol(symbol string) *TopLongShortAccountRatioService {
	s.symbol = symbol
	return s
}

func (s *TopLongShortAccountRatioService) Period(period string) *TopLongShortAccountRatioService {
	s.period = period
	return s
}

func (s *TopLongShortAccountRatioService) Limit(limit uint32) *TopLongShortAccountRatioService {
	s.limit = &limit
	return s
}

func (s *TopLongShortAccountRatioService) StartTime(startTime uint64) *TopLongShortAccountRatioService {
	s.startTime = &startTime
	return s
}

func (s *TopLongShortAccountRatioService) EndTime(endTime uint64) *TopLongShortAccountRatioService {
	s.endTime = &endTime
	return s
}

func (s *TopLongShortAccountRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*TopLongShortAccountRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/topLongShortAccountRatio",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)
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
		return nil, err
	}
	res = make([]*TopLongShortAccountRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type TopLongShortPositionRatioService struct {
	c         *Client
	symbol    string
	period    string  // "5m","15m","30m","1h","2h","4h","6h","12h","1d"
	limit     *uint32 // default 30, max 500
	startTime *uint64
	endTime   *uint64
}

type TopLongShortPositionRatio struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      uint64 `json:"timestamp"`
}

func (s *TopLongShortPositionRatioService) Symbol(symbol string) *TopLongShortPositionRatioService {
	s.symbol = symbol
	return s
}

func (s *TopLongShortPositionRatioService) Period(period string) *TopLongShortPositionRatioService {
	s.period = period
	return s
}

func (s *TopLongShortPositionRatioService) Limit(limit uint32) *TopLongShortPositionRatioService {
	s.limit = &limit
	return s
}

func (s *TopLongShortPositionRatioService) StartTime(startTime uint64) *TopLongShortPositionRatioService {
	s.startTime = &startTime
	return s
}

func (s *TopLongShortPositionRatioService) EndTime(endTime uint64) *TopLongShortPositionRatioService {
	s.endTime = &endTime
	return s
}

func (s *TopLongShortPositionRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*TopLongShortPositionRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/topLongShortPositionRatio",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)
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
		return nil, err
	}
	res = make([]*TopLongShortPositionRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type TakerLongShortRatioService struct {
	c         *Client
	symbol    string
	period    string  // "5m","15m","30m","1h","2h","4h","6h","12h","1d"
	limit     *uint32 // default 30, max 500
	startTime *uint64
	endTime   *uint64
}

type TakerLongShortRatio struct {
	BuySellRatio string `json:"buySellRatio"`
	BuyVol       string `json:"buyVol"`
	SellVol      string `json:"sellVol"`
	Timestamp    uint64 `json:"timestamp"`
}

func (s *TakerLongShortRatioService) Symbol(symbol string) *TakerLongShortRatioService {
	s.symbol = symbol
	return s
}

func (s *TakerLongShortRatioService) Period(period string) *TakerLongShortRatioService {
	s.period = period
	return s
}

func (s *TakerLongShortRatioService) Limit(limit uint32) *TakerLongShortRatioService {
	s.limit = &limit
	return s
}

func (s *TakerLongShortRatioService) StartTime(startTime uint64) *TakerLongShortRatioService {
	s.startTime = &startTime
	return s
}

func (s *TakerLongShortRatioService) EndTime(endTime uint64) *TakerLongShortRatioService {
	s.endTime = &endTime
	return s
}

func (s *TakerLongShortRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*TakerLongShortRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/takerlongshortRatio",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)
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
		return nil, err
	}
	res = make([]*TakerLongShortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type BasisService struct {
	c            *Client
	pair         string // for example, BTCUSDT
	contractType string // CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
	period       string // "5m","15m","30m","1h","2h","4h","6h","12h","1d"
	limit        uint32 // default 30, max 500
	startTime    *uint64
	endTime      *uint64
}

type Basis struct {
	Pair                string `json:"pair"`
	IndexPrice          string `json:"indexPrice"`
	ContractType        string `json:"contractType"`
	BasisRate           string `json:"basisRate"`
	FuturesPrice        string `json:"futuresPrice"`
	AnnualizedBasisRate string `json:"annualizedBasisRate"`
	Basis               string `json:"basis"`
	Timestamp           uint64 `json:"timestamp"`
}

func (s *BasisService) Pair(pair string) *BasisService {
	s.pair = pair
	return s
}

func (s *BasisService) ContractType(contractType string) *BasisService {
	s.contractType = contractType
	return s
}

func (s *BasisService) Period(period string) *BasisService {
	s.period = period
	return s
}

func (s *BasisService) Limit(limit uint32) *BasisService {
	s.limit = limit
	return s
}

func (s *BasisService) StartTime(startTime uint64) *BasisService {
	s.startTime = &startTime
	return s
}

func (s *BasisService) EndTime(endTime uint64) *BasisService {
	s.endTime = &endTime
	return s
}

func (s *BasisService) Do(ctx context.Context, opts ...RequestOption) (res []*Basis, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/basis",
	}
	r.setParam("pair", s.pair)
	r.setParam("contractType", s.contractType)
	r.setParam("period", s.period)
	r.setParam("limit", s.limit)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*Basis, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
