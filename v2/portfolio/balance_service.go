package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetBalanceService get account balance
type GetBalanceService struct {
	c     *Client
	asset *string
}

// Asset set asset
func (s *GetBalanceService) Asset(asset string) *GetBalanceService {
	s.asset = &asset
	return s
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*Balance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/balance",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Balance{}, err
	}

	if s.asset != nil {
		// Single balance response when asset is specified
		singleRes := new(Balance)
		err = json.Unmarshal(data, singleRes)
		if err != nil {
			return nil, err
		}
		return []*Balance{singleRes}, nil
	}

	// Array of balances when no asset is specified
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Balance{}, err
	}
	return res, nil
}

// Balance define user balance of your account
type Balance struct {
	Asset               string `json:"asset"`
	TotalWalletBalance  string `json:"totalWalletBalance"`
	CrossMarginAsset    string `json:"crossMarginAsset"`
	CrossMarginBorrowed string `json:"crossMarginBorrowed"`
	CrossMarginFree     string `json:"crossMarginFree"`
	CrossMarginInterest string `json:"crossMarginInterest"`
	CrossMarginLocked   string `json:"crossMarginLocked"`
	UMWalletBalance     string `json:"umWalletBalance"`
	UMUnrealizedPNL     string `json:"umUnrealizedPNL"`
	CMWalletBalance     string `json:"cmWalletBalance"`
	CMUnrealizedPNL     string `json:"cmUnrealizedPNL"`
	UpdateTime          int64  `json:"updateTime"`
	NegativeBalance     string `json:"negativeBalance"`
}
