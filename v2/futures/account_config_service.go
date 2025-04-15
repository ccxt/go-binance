package futures

import (
	"context"
	"encoding/json"
)

// AccountConfigService get futures account configuration
type AccountConfigService struct {
	c *Client
}

// AccountConfig define futures account configuration
type AccountConfig struct {
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
func (s *AccountConfigService) Do(ctx context.Context) (*AccountConfig, error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/accountConfig",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(AccountConfig)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
