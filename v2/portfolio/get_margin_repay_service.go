package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMarginRepayService query margin repay record
type GetMarginRepayService struct {
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
func (s *GetMarginRepayService) Asset(asset string) *GetMarginRepayService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *GetMarginRepayService) TxID(txID int64) *GetMarginRepayService {
	s.txID = &txID
	return s
}

// StartTime set startTime
func (s *GetMarginRepayService) StartTime(startTime int64) *GetMarginRepayService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetMarginRepayService) EndTime(endTime int64) *GetMarginRepayService {
	s.endTime = &endTime
	return s
}

// Current set current page
func (s *GetMarginRepayService) Current(current int64) *GetMarginRepayService {
	s.current = &current
	return s
}

// Size set page size
func (s *GetMarginRepayService) Size(size int64) *GetMarginRepayService {
	s.size = &size
	return s
}

// Archived set archived
func (s *GetMarginRepayService) Archived(archived bool) *GetMarginRepayService {
	s.archived = &archived
	return s
}

// Do send request
func (s *GetMarginRepayService) Do(ctx context.Context, opts ...RequestOption) (*GetMarginRepayResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/repayLoan",
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
	res := new(GetMarginRepayResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMarginRepayResponse define margin repay response
type GetMarginRepayResponse struct {
	Rows  []MarginRepay `json:"rows"`
	Total int64         `json:"total"`
}

// MarginRepay define margin repay info
type MarginRepay struct {
	Amount    string `json:"amount"` // Total amount repaid
	Asset     string `json:"asset"`
	Interest  string `json:"interest"`  // Interest repaid
	Principal string `json:"principal"` // Principal repaid
	Status    string `json:"status"`    // PENDING/CONFIRMED/FAILED
	Timestamp int64  `json:"timestamp"`
	TxID      int64  `json:"txId"`
}
