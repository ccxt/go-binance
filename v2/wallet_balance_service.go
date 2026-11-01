package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// WalletBalanceService fetches all user wallet balance.
//
// See https://developers.binance.com/docs/wallet/asset/query-user-wallet-balance
type WalletBalanceService struct {
	c          *Client
	quoteAsset *string
}

// QuoteAsset sets the quoteAsset parameter.
func (s *WalletBalanceService) QuoteAsset(asset string) *WalletBalanceService {
	s.quoteAsset = &asset
	return s
}

// Do sends the request.
func (s *WalletBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*WalletBalance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/wallet/balance",
		secType:  secTypeSigned,
	}
	if s.quoteAsset != nil {
		r.setParam("quoteAsset", *s.quoteAsset)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*WalletBalance, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// WalletBalanceResponse defines the response of WalletBalanceService
type WalletBalance struct {
	Activate   bool   `json:"activate"`
	Balance    string `json:"balance"`
	WalletName string `json:"walletName"`
}
