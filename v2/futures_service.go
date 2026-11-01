package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// FuturesTransferService transfer asset between spot account and futures account
type FuturesTransferService struct {
	c            *Client
	asset        string
	amount       string
	transferType int
}

// Asset set asset being transferred, e.g., BTC
func (s *FuturesTransferService) Asset(asset string) *FuturesTransferService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *FuturesTransferService) Amount(amount string) *FuturesTransferService {
	s.amount = amount
	return s
}

// Type 1: transfer from spot account to futures account 2: transfer from futures account to spot account
func (s *FuturesTransferService) Type(transferType FuturesTransferType) *FuturesTransferService {
	s.transferType = int(transferType)
	return s
}

// Do send request
func (s *FuturesTransferService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/futures/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListFuturesTransferService list futures transfer
type ListFuturesTransferService struct {
	c         *Client
	asset     *string
	startTime int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListFuturesTransferService) Asset(asset string) *ListFuturesTransferService {
	s.asset = &asset
	return s
}

// StartTime set start time
func (s *ListFuturesTransferService) StartTime(startTime int64) *ListFuturesTransferService {
	s.startTime = startTime
	return s
}

// EndTime set end time
func (s *ListFuturesTransferService) EndTime(endTime int64) *ListFuturesTransferService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListFuturesTransferService) Current(current int64) *ListFuturesTransferService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListFuturesTransferService) Size(size int64) *ListFuturesTransferService {
	s.size = &size
	return s
}

// Do send request
func (s *ListFuturesTransferService) Do(ctx context.Context, opts ...RequestOption) (res *FuturesTransferHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/futures/transfer",
		secType:  secTypeSigned,
	}
	r.setParams(params{
		"startTime": s.startTime,
	})
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FuturesTransferHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FuturesTransferHistory define futures transfer history
type FuturesTransferHistory struct {
	Rows  []FuturesTransfer `json:"rows"`
	Total int64             `json:"total"`
}

// FuturesTransfer define futures transfer history item
type FuturesTransfer struct {
	Asset     string                    `json:"asset"`
	TranID    int64                     `json:"tranId"`
	Amount    string                    `json:"amount"`
	Type      int64                     `json:"type"`
	Timestamp int64                     `json:"timestamp"`
	Status    FuturesTransferStatusType `json:"status"`
}

type FuturesOrderBookHistoryService struct {
	c         *Client
	symbol    string
	dataType  string
	startTime int64
	endTime   int64
}

func (s *FuturesOrderBookHistoryService) Symbol(symbol string) *FuturesOrderBookHistoryService {
	s.symbol = symbol
	return s
}

func (s *FuturesOrderBookHistoryService) DataType(dataType FuturesOrderBookHistoryDataType) *FuturesOrderBookHistoryService {
	s.dataType = string(dataType)
	return s
}

func (s *FuturesOrderBookHistoryService) StartTime(startTime int64) *FuturesOrderBookHistoryService {
	s.startTime = startTime
	return s
}

func (s *FuturesOrderBookHistoryService) EndTime(endTime int64) *FuturesOrderBookHistoryService {
	s.endTime = endTime
	return s
}

func (s *FuturesOrderBookHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *FuturesOrderBookHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/futures/histDataLink",
		secType:  secTypeSigned,
	}
	r.setParams(params{
		"symbol":    s.symbol,
		"dataType":  s.dataType,
		"startTime": s.startTime,
		"endTime":   s.endTime,
	})
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FuturesOrderBookHistory)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type FuturesOrderBookHistoryItem struct {
	Day string `json:"day"`
	Url string `json:"url"`
}

type FuturesOrderBookHistory struct {
	Data []*FuturesOrderBookHistoryItem `json:"data"`
}
