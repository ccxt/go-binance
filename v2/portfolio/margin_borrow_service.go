package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMarginMaxBorrowService get margin max borrowable amount
type GetMarginMaxBorrowService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMarginMaxBorrowService) Asset(asset string) *GetMarginMaxBorrowService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMarginMaxBorrowService) Do(ctx context.Context, opts ...RequestOption) (res *MaxBorrow, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/maxBorrowable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxBorrow)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MaxBorrow define margin max borrowable amount info
type MaxBorrow struct {
	Amount      string `json:"amount"`      // account's currently max borrowable amount with sufficient system availability
	BorrowLimit string `json:"borrowLimit"` // max borrowable amount limited by the account level
}
