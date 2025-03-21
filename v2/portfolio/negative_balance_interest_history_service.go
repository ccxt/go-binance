package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetNegativeBalanceInterestHistoryService get portfolio margin negative balance interest history
type GetNegativeBalanceInterestHistoryService struct {
	c         *Client
	asset     *string
	startTime *int64
	endTime   *int64
	size      *int64
}

// Asset set asset
func (s *GetNegativeBalanceInterestHistoryService) Asset(asset string) *GetNegativeBalanceInterestHistoryService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *GetNegativeBalanceInterestHistoryService) StartTime(startTime int64) *GetNegativeBalanceInterestHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetNegativeBalanceInterestHistoryService) EndTime(endTime int64) *GetNegativeBalanceInterestHistoryService {
	s.endTime = &endTime
	return s
}

// Size set size
func (s *GetNegativeBalanceInterestHistoryService) Size(size int64) *GetNegativeBalanceInterestHistoryService {
	s.size = &size
	return s
}

// Do send request
func (s *GetNegativeBalanceInterestHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*NegativeBalanceInterest, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/portfolio/interest-history",
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
	if s.size != nil {
		r.setParam("size", *s.size)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*NegativeBalanceInterest{}, err
	}
	res = make([]*NegativeBalanceInterest, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*NegativeBalanceInterest{}, err
	}
	return res, nil
}

// NegativeBalanceInterest define negative balance interest info
type NegativeBalanceInterest struct {
	Asset               string `json:"asset"`
	Interest            string `json:"interest"` // interest amount
	InterestAccuredTime int64  `json:"interestAccuredTime"`
	InterestRate        string `json:"interestRate"` // daily interest rate
	Principal           string `json:"principal"`
}
