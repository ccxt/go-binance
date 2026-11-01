package portfolio

import (
	"context"
	"encoding/json"
)

// UMAccountDetailV2Service get UM account detail v2
type UMAccountDetailV2Service struct {
	c *Client
}

// UMAccountDetailV2 define UM account detail v2
type UMAccountDetailV2 struct {
	Assets    []*UMAssetV2    `json:"assets"`
	Positions []*UMPositionV2 `json:"positions"`
}

// UMAssetV2 define UM asset detail v2
type UMAssetV2 struct {
	Asset                  string `json:"asset"`                  // Asset name
	CrossWalletBalance     string `json:"crossWalletBalance"`     // Wallet balance
	CrossUnPnl             string `json:"crossUnPnl"`             // Unrealized profit
	MaintMargin            string `json:"maintMargin"`            // Maintenance margin required
	InitialMargin          string `json:"initialMargin"`          // Total initial margin required
	PositionInitialMargin  string `json:"positionInitialMargin"`  // Initial margin required for positions
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // Initial margin required for open orders
	UpdateTime             int64  `json:"updateTime"`             // Last update time
}

// UMPositionV2 define UM position detail v2
type UMPositionV2 struct {
	Symbol           string `json:"symbol"`           // Symbol name
	InitialMargin    string `json:"initialMargin"`    // Initial margin required
	MaintMargin      string `json:"maintMargin"`      // Maintenance margin required
	UnrealizedProfit string `json:"unrealizedProfit"` // Unrealized profit
	PositionSide     string `json:"positionSide"`     // Position side
	PositionAmt      string `json:"positionAmt"`      // Position amount
	UpdateTime       int64  `json:"updateTime"`       // Last update time
	Notional         string `json:"notional"`         // Notional value
}

// Do send request
func (s *UMAccountDetailV2Service) Do(ctx context.Context) (*UMAccountDetailV2, error) {
	r := &request{
		method:   "GET",
		endpoint: "/papi/v2/um/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(UMAccountDetailV2)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
