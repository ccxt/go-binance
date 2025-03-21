package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMarginLoanService query margin loan record
type GetMarginLoanService struct {
	c         *Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
	archived  *bool
}

// Asset set asset
func (s *GetMarginLoanService) Asset(asset string) *GetMarginLoanService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *GetMarginLoanService) TxID(txID int64) *GetMarginLoanService {
	s.txID = &txID
	return s
}

// StartTime set startTime
func (s *GetMarginLoanService) StartTime(startTime int64) *GetMarginLoanService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetMarginLoanService) EndTime(endTime int64) *GetMarginLoanService {
	s.endTime = &endTime
	return s
}

// Current set current page
func (s *GetMarginLoanService) Current(current int64) *GetMarginLoanService {
	s.current = &current
	return s
}

// Size set page size
func (s *GetMarginLoanService) Size(size int64) *GetMarginLoanService {
	s.size = &size
	return s
}

// Archived set archived
func (s *GetMarginLoanService) Archived(archived bool) *GetMarginLoanService {
	s.archived = &archived
	return s
}

// Do send request
func (s *GetMarginLoanService) Do(ctx context.Context, opts ...RequestOption) (*GetMarginLoanResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/marginLoan",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.txID != nil {
		r.setParam("txId", *s.txID)
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
	res := new(GetMarginLoanResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMarginLoanResponse define margin loan response
type GetMarginLoanResponse struct {
	Rows  []MarginLoan `json:"rows"`
	Total int64        `json:"total"`
}

// MarginLoan define margin loan info
type MarginLoan struct {
	TxID      int64  `json:"txId"`
	Asset     string `json:"asset"`
	Principal string `json:"principal"`
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"` // PENDING/CONFIRMED/FAILED
}
