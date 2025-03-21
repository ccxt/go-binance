package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetNegativeBalanceExchangeRecordService get user negative balance auto exchange record
type GetNegativeBalanceExchangeRecordService struct {
	c          *Client
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

// StartTime set startTime
func (s *GetNegativeBalanceExchangeRecordService) StartTime(startTime int64) *GetNegativeBalanceExchangeRecordService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetNegativeBalanceExchangeRecordService) EndTime(endTime int64) *GetNegativeBalanceExchangeRecordService {
	s.endTime = &endTime
	return s
}

// RecvWindow set recvWindow
func (s *GetNegativeBalanceExchangeRecordService) RecvWindow(recvWindow int64) *GetNegativeBalanceExchangeRecordService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetNegativeBalanceExchangeRecordService) Do(ctx context.Context, opts ...RequestOption) (*NegativeBalanceExchangeRecord, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/portfolio/negative-balance-exchange-record",
		secType:  secTypeSigned,
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(NegativeBalanceExchangeRecord)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NegativeBalanceExchangeRecord define negative balance exchange record response
type NegativeBalanceExchangeRecord struct {
	Total int64                      `json:"total"`
	Rows  []*NegativeBalanceExchange `json:"rows"`
}

// NegativeBalanceExchange define negative balance exchange info
type NegativeBalanceExchange struct {
	StartTime int64                            `json:"startTime"`
	EndTime   int64                            `json:"endTime"`
	Details   []*NegativeBalanceExchangeDetail `json:"details"`
}

// NegativeBalanceExchangeDetail define negative balance exchange detail
type NegativeBalanceExchangeDetail struct {
	Asset                string  `json:"asset"`
	NegativeBalance      float64 `json:"negativeBalance"`
	NegativeMaxThreshold float64 `json:"negativeMaxThreshold"`
}
