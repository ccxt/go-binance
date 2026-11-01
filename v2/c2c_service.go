package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// C2CTradeHistoryService retrieve c2c trade history
type C2CTradeHistoryService struct {
	c              *Client
	tradeType      SideType
	startTimestamp *int64
	endTimestamp   *int64
	page           *int32
	rows           *int32
}

// TransactionType set transaction type
func (s *C2CTradeHistoryService) TradeType(tradeType SideType) *C2CTradeHistoryService {
	s.tradeType = tradeType
	return s
}

// BeginTime set beginTime
func (s *C2CTradeHistoryService) StartTimestamp(startTimestamp int64) *C2CTradeHistoryService {
	s.startTimestamp = &startTimestamp
	return s
}

// EndTime set endTime
func (s *C2CTradeHistoryService) EndTime(endTimestamp int64) *C2CTradeHistoryService {
	s.endTimestamp = &endTimestamp
	return s
}

// Page set page
func (s *C2CTradeHistoryService) Page(page int32) *C2CTradeHistoryService {
	s.page = &page
	return s
}

// Rows set rows
func (s *C2CTradeHistoryService) Rows(rows int32) *C2CTradeHistoryService {
	s.rows = &rows
	return s
}

// Do send request
func (s *C2CTradeHistoryService) Do(ctx context.Context, opts ...RequestOption) (*C2CTradeHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/c2c/orderMatch/listUserOrderHistory",
		secType:  secTypeSigned,
	}
	r.setParam("tradeType", s.tradeType)
	if s.startTimestamp != nil {
		r.setParam("startTimestamp", *s.startTimestamp)
	}
	if s.endTimestamp != nil {
		r.setParam("endTime", *s.endTimestamp)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.rows != nil {
		r.setParam("rows", *s.rows)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := C2CTradeHistory{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// C2CTradeHistory response
type C2CTradeHistory struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    []C2CRecord `json:"data"`
	Total   int64       `json:"total"`
	Success bool        `json:"success"`
}

// C2CRecord a record of c2c
type C2CRecord struct {
	OrderNumber         string `json:"orderNumber"`
	AdvNo               string `json:"advNo"`
	TradeType           string `json:"tradeType"`
	Asset               string `json:"asset"`
	Fiat                string `json:"fiat"`
	FiatSymbol          string `json:"fiatSymbol"`
	Amount              string `json:"amount"`
	TotalPrice          string `json:"totalPrice"`
	UnitPrice           string `json:"unitPrice"`
	OrderStatus         string `json:"orderStatus"`
	CreateTime          int64  `json:"createTime"`
	Commission          string `json:"commission"`
	CounterPartNickName string `json:"counterPartNickName"`
	AdvertisementRole   string `json:"advertisementRole"`
}
