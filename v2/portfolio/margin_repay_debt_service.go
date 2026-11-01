package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginRepayDebtService service to repay margin debt
type MarginRepayDebtService struct {
	c                  *Client
	asset              string
	amount             *string
	specifyRepayAssets *string
	recvWindow         *int64
}

// Asset set asset
func (s *MarginRepayDebtService) Asset(asset string) *MarginRepayDebtService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *MarginRepayDebtService) Amount(amount string) *MarginRepayDebtService {
	s.amount = &amount
	return s
}

// SpecifyRepayAssets set specific assets to repay debt
func (s *MarginRepayDebtService) SpecifyRepayAssets(assets string) *MarginRepayDebtService {
	s.specifyRepayAssets = &assets
	return s
}

// RecvWindow set recvWindow
func (s *MarginRepayDebtService) RecvWindow(recvWindow int64) *MarginRepayDebtService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginRepayDebtService) Do(ctx context.Context) (*MarginRepayDebtResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/margin/repay-debt",
		secType:  secTypeSigned,
	}

	r.setParam("asset", s.asset)
	if s.amount != nil {
		r.setParam("amount", *s.amount)
	}
	if s.specifyRepayAssets != nil {
		r.setParam("specifyRepayAssets", *s.specifyRepayAssets)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(MarginRepayDebtResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginRepayDebtResponse represents the response from repaying margin debt
type MarginRepayDebtResponse struct {
	Amount             string   `json:"amount"`
	Asset              string   `json:"asset"`
	SpecifyRepayAssets []string `json:"specifyRepayAssets"`
	UpdateTime         int64    `json:"updateTime"`
	Success            bool     `json:"success"`
}
