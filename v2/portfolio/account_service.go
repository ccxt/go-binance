package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAccountService get account information
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Account define account info
type Account struct {
	UniMMR                   string `json:"uniMMR"`        // Portfolio margin account maintenance margin rate
	AccountEquity            string `json:"accountEquity"` // Account equity, in USD value
	ActualEquity             string `json:"actualEquity"`  // Account equity without collateral rate, in USD value
	AccountInitialMargin     string `json:"accountInitialMargin"`
	AccountMaintMargin       string `json:"accountMaintMargin"`       // Portfolio margin account maintenance margin, unit：USD
	AccountStatus            string `json:"accountStatus"`            // Portfolio margin account status
	VirtualMaxWithdrawAmount string `json:"virtualMaxWithdrawAmount"` // Portfolio margin maximum amount for transfer out in USD
	TotalAvailableBalance    string `json:"totalAvailableBalance"`
	TotalMarginOpenLoss      string `json:"totalMarginOpenLoss"` // in USD margin open order
	UpdateTime               int64  `json:"updateTime"`          // last update time
}
