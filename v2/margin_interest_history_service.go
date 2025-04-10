package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginInterestHistoryService fetches the margin interest history
type MarginInterestHistoryService struct {
	c              *Client
	asset          *string
	isolatedSymbol *string
	startTime      *int64
	endTime        *int64
	current        *int64
	size           *int64
}

// Asset sets the asset parameter.
func (s *MarginInterestHistoryService) Asset(asset string) *MarginInterestHistoryService {
	s.asset = &asset
	return s
}

// IsolatedSymbol sets the isolatedSymbol parameter.
func (s *MarginInterestHistoryService) IsolatedSymbol(symbol string) *MarginInterestHistoryService {
	s.isolatedSymbol = &symbol
	return s
}

// StartTime sets the startTime parameter.
func (s *MarginInterestHistoryService) StartTime(startTime int64) *MarginInterestHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
func (s *MarginInterestHistoryService) EndTime(endTime int64) *MarginInterestHistoryService {
	s.endTime = &endTime
	return s
}

// Current sets the current parameter.
func (s *MarginInterestHistoryService) Current(current int64) *MarginInterestHistoryService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *MarginInterestHistoryService) Size(size int64) *MarginInterestHistoryService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *MarginInterestHistoryService) Do(ctx context.Context) (*MarginInterestHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/interestHistory",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.isolatedSymbol != nil {
		r.setParam("isolatedSymbol", *s.isolatedSymbol)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
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
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(MarginInterestHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginInterestHistory represents the response
type MarginInterestHistory struct {
	Rows  []MarginInterestHistoryRow `json:"rows"`
	Total int64                      `json:"total"`
}

type MarginInterestHistoryRow struct {
	TxId                int64  `json:"txId"`
	InterestAccuredTime int64  `json:"interestAccuredTime"`
	Asset               string `json:"asset"`
	RawAsset            string `json:"rawAsset,omitempty"`
	Principal           string `json:"principal"`
	Interest            string `json:"interest"`
	InterestRate        string `json:"interestRate"`
	Type                string `json:"type"`
	IsolatedSymbol      string `json:"isolatedSymbol,omitempty"`
}
