package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMarginMaxWithdrawService get margin max withdrawable amount
type GetMarginMaxWithdrawService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMarginMaxWithdrawService) Asset(asset string) *GetMarginMaxWithdrawService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMarginMaxWithdrawService) Do(ctx context.Context, opts ...RequestOption) (res *MaxWithdraw, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/maxWithdraw",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxWithdraw)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MaxWithdraw define margin max withdrawable amount info
type MaxWithdraw struct {
	Amount string `json:"amount"` // max withdrawable amount
}
