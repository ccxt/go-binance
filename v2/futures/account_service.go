package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetBalanceService get account balance
type GetBalanceService struct {
	c *Client
}

// Do send request
//
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Account-Balance-V3
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*Balance, rateLimits map[string]string, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v3/balance",
		secType:  secTypeSigned,
	}
	data, rateLimits, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Balance{}, nil, err
	}
	res = make([]*Balance, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Balance{}, nil, err
	}
	return res, rateLimits, nil
}

// Balance define user balance of your account
type Balance struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
	MarginAvailable    bool   `json:"marginAvailable"`
	UpdateTime         int64  `json:"updateTime"`
}

// GetAccountService get account info
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, rateLimits map[string]string, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/account",
		secType:  secTypeSigned,
	}
	data, rateLimits, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, nil, err
	}
	return res, rateLimits, nil
}

// Account define account info
type Account struct {
	Assets                      []*AccountAsset    `json:"assets"`
	FeeTier                     int                `json:"feeTier"`
	CanTrade                    bool               `json:"canTrade"`
	CanDeposit                  bool               `json:"canDeposit"`
	CanWithdraw                 bool               `json:"canWithdraw"`
	UpdateTime                  int64              `json:"updateTime"`
	MultiAssetsMargin           bool               `json:"multiAssetsMargin"`
	TotalInitialMargin          string             `json:"totalInitialMargin"`
	TotalMaintMargin            string             `json:"totalMaintMargin"`
	TotalWalletBalance          string             `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string             `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string             `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string             `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string             `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string             `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string             `json:"totalCrossUnPnl"`
	AvailableBalance            string             `json:"availableBalance"`
	MaxWithdrawAmount           string             `json:"maxWithdrawAmount"`
	Positions                   []*AccountPosition `json:"positions"`
}

// AccountAsset define account asset
type AccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}

// AccountPosition define account position
type AccountPosition struct {
	Isolated               bool             `json:"isolated"`
	Leverage               string           `json:"leverage"`
	InitialMargin          string           `json:"initialMargin"`
	MaintMargin            string           `json:"maintMargin"`
	OpenOrderInitialMargin string           `json:"openOrderInitialMargin"`
	PositionInitialMargin  string           `json:"positionInitialMargin"`
	Symbol                 string           `json:"symbol"`
	UnrealizedProfit       string           `json:"unrealizedProfit"`
	EntryPrice             string           `json:"entryPrice"`
	MaxNotional            string           `json:"maxNotional"`
	PositionSide           PositionSideType `json:"positionSide"`
	PositionAmt            string           `json:"positionAmt"`
	Notional               string           `json:"notional"`
	BidNotional            string           `json:"bidNotional"`
	AskNotional            string           `json:"askNotional"`
	IsolatedWallet         string           `json:"isolatedWallet"`
	UpdateTime             int64            `json:"updateTime"`
}

// GetAccountV3Service get account info
type GetAccountV3Service struct {
	c *Client
}

// Do send request
func (s *GetAccountV3Service) Do(ctx context.Context, opts ...RequestOption) (res *AccountV3, rateLimits map[string]string, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v3/account",
		secType:  secTypeSigned,
	}
	data, rateLimits, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, nil, err
	}
	res = new(AccountV3)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, nil, err
	}
	return res, rateLimits, nil
}

// AccountV3 define account info
type AccountV3 struct {
	TotalInitialMargin          string               `json:"totalInitialMargin"`
	TotalMaintMargin            string               `json:"totalMaintMargin"`
	TotalWalletBalance          string               `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string               `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string               `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string               `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string               `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string               `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string               `json:"totalCrossUnPnl"`
	AvailableBalance            string               `json:"availableBalance"`
	MaxWithdrawAmount           string               `json:"maxWithdrawAmount"`
	Assets                      []*AccountAssetV3    `json:"assets"`
	Positions                   []*AccountPositionV3 `json:"positions"`
}

// AccountAssetV3 define account asset
type AccountAssetV3 struct {
	Asset                  string `json:"asset"`
	WalletBalance          string `json:"walletBalance"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	MarginBalance          string `json:"marginBalance"`
	MaintMargin            string `json:"maintMargin"`
	InitialMargin          string `json:"initialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}

// AccountPositionV3 define account position
type AccountPositionV3 struct {
	Symbol           string `json:"symbol"`
	PositionSide     string `json:"positionSide"`
	PositionAmt      string `json:"positionAmt"`
	UnrealizedProfit string `json:"unrealizedProfit"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Notional         string `json:"notional"`
	IsolatedWallet   string `json:"isolatedWallet"`
	InitialMargin    string `json:"initialMargin"`
	MaintMargin      string `json:"maintMargin"`
	UpdateTime       int64  `json:"updateTime"`
}
