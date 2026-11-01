package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMarginInterestHistoryService get margin borrow/loan interest history
type GetMarginInterestHistoryService struct {
	c         *Client
	asset     *string
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
	archived  *bool
}

// Asset set asset
func (s *GetMarginInterestHistoryService) Asset(asset string) *GetMarginInterestHistoryService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *GetMarginInterestHistoryService) StartTime(startTime int64) *GetMarginInterestHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetMarginInterestHistoryService) EndTime(endTime int64) *GetMarginInterestHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current page
func (s *GetMarginInterestHistoryService) Current(current int64) *GetMarginInterestHistoryService {
	s.current = &current
	return s
}

// Size set page size
func (s *GetMarginInterestHistoryService) Size(size int64) *GetMarginInterestHistoryService {
	s.size = &size
	return s
}

// Archived set archived
func (s *GetMarginInterestHistoryService) Archived(archived bool) *GetMarginInterestHistoryService {
	s.archived = &archived
	return s
}

// Do send request
func (s *GetMarginInterestHistoryService) Do(ctx context.Context, opts ...RequestOption) (*MarginInterestHistoryResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/marginInterestHistory",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
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
	if s.archived != nil {
		r.setParam("archived", *s.archived)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(MarginInterestHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginInterestHistoryResponse define margin interest history response
type MarginInterestHistoryResponse struct {
	Rows  []MarginInterest `json:"rows"`
	Total int64            `json:"total"`
}

// MarginInterest define margin interest info
type MarginInterest struct {
	TxID                int64  `json:"txId"`
	InterestAccuredTime int64  `json:"interestAccuredTime"`
	Asset               string `json:"asset"`
	RawAsset            string `json:"rawAsset"`
	Principal           string `json:"principal"`
	Interest            string `json:"interest"`
	InterestRate        string `json:"interestRate"`
	Type                string `json:"type"` // PERIODIC/ON_BORROW/PERIODIC_CONVERTED/ON_BORROW_CONVERTED/PORTFOLIO
}
