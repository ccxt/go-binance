package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// PayTradeHistoryService retrieves the pay transaction history
type PayTradeHistoryService struct {
	c              *Client
	startTimestamp *int64
	endTimestamp   *int64
	limit          *int32
}

// StartTimestamp set startTimestamp
func (s *PayTradeHistoryService) StartTimestamp(startTimestamp int64) *PayTradeHistoryService {
	s.startTimestamp = &startTimestamp
	return s
}

// EndTimestamp set endTimestamp
func (s *PayTradeHistoryService) EndTimestamp(endTimestamp int64) *PayTradeHistoryService {
	s.endTimestamp = &endTimestamp
	return s
}

// Rows set rows
func (s *PayTradeHistoryService) Limit(limit int32) *PayTradeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *PayTradeHistoryService) Do(ctx context.Context, opts ...RequestOption) (*PayTradeHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/pay/transactions",
		secType:  secTypeSigned,
	}
	if s.startTimestamp != nil {
		r.setParam("startTimestamp", *s.startTimestamp)
	}
	if s.endTimestamp != nil {
		r.setParam("endTimestamp", *s.endTimestamp)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := PayTradeHistory{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type PayTradeHistory struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    []PayTradeItem `json:"data"`
	Success bool           `json:"success"`
}

type PayTradeItem struct {
	OrderType       string        `json:"orderType"`
	TransactionID   string        `json:"transactionId"`
	TransactionTime int64         `json:"transactionTime"`
	Amount          string        `json:"amount"`
	Currency        string        `json:"currency"`
	PayerInfo       *PayerInfo    `json:"payerInfo"`
	ReceiverInfo    *ReceiverInfo `json:"receiverInfo"`
	FundsDetail     []FundsDetail `json:"fundsDetail"`
}

type FundsDetail struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}
type PayerInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Email       string `json:"email"`
	BinanceId   int    `json:"binanceId"`
	AccountId   int    `json:"accountId"`
	CountryCode int    `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
	MobileCode  string `json:"mobileCode"`
	UnmaskData  bool   `json:"unmaskData"`
}
type ReceiverInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Email       string `json:"email,omitempty"`
	BinanceId   int    `json:"binanceId,omitempty"`
	AccountId   int    `json:"accountId"`
	CountryCode int    `json:"countryCode,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	MobileCode  string `json:"mobileCode,omitempty"`
	UnmaskData  bool   `json:"unmaskData"`
	Extend      struct {
		PhoneOrEmailChanged bool `json:"phoneOrEmailChanged"`
	} `json:"extend,omitempty"`
}
