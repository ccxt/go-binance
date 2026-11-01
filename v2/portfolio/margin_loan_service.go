package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginLoanService service to borrow margin loan
type MarginLoanService struct {
	c      *Client
	asset  string
	amount string
}

// Asset set asset
func (s *MarginLoanService) Asset(asset string) *MarginLoanService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *MarginLoanService) Amount(amount string) *MarginLoanService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginLoanService) Do(ctx context.Context, opts ...RequestOption) (res *MarginLoanResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/marginLoan",
		secType:  secTypeSigned,
	}

	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginLoanResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginLoanResponse define margin loan response
type MarginLoanResponse struct {
	TranID int64 `json:"tranId"`
}
