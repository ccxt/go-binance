package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginRepayService service to repay margin loans
type MarginRepayService struct {
	c          *Client
	asset      string
	amount     string
	recvWindow *int64
}

// Asset set asset
func (s *MarginRepayService) Asset(asset string) *MarginRepayService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *MarginRepayService) Amount(amount string) *MarginRepayService {
	s.amount = amount
	return s
}

// RecvWindow set recvWindow
func (s *MarginRepayService) RecvWindow(recvWindow int64) *MarginRepayService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginRepayService) Do(ctx context.Context, opts ...RequestOption) (res *MarginRepayResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/repayLoan",
		secType:  secTypeSigned,
	}

	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginRepayResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginRepayResponse define margin repay response
type MarginRepayResponse struct {
	TranID int64 `json:"tranId"`
}
