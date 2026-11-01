package portfolio

import (
	"context"
	"encoding/json"
)

// UMAccountConfigService get UM futures account configuration
type UMAccountConfigService struct {
	c *Client
}

// UMAccountConfig define UM futures account configuration
type UMAccountConfig struct {
	FeeTier           int   `json:"feeTier"`          // Account commission tier
	CanTrade          bool  `json:"canTrade"`         // If can trade
	CanDeposit        bool  `json:"canDeposit"`       // If can transfer in asset
	CanWithdraw       bool  `json:"canWithdraw"`      // If can transfer out asset
	DualSidePosition  bool  `json:"dualSidePosition"` // If dual side position is enabled
	UpdateTime        int64 `json:"updateTime"`       // Reserved property
	MultiAssetsMargin bool  `json:"multiAssetsMargin"`
	TradeGroupId      int   `json:"tradeGroupId"`
}

// Do send request
func (s *UMAccountConfigService) Do(ctx context.Context) (*UMAccountConfig, error) {
	r := &request{
		method:   "GET",
		endpoint: "/papi/v1/um/accountConfig",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(UMAccountConfig)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
