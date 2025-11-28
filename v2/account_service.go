package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAccountService get account info
type GetAccountService struct {
	c                *Client
	omitZeroBalances *bool
}

// OmitZeroBalances sets the omitZeroBalances parameter on the request.
// When set to true, the API will return the non-zero balances of an account.
func (s *GetAccountService) OmitZeroBalances(v bool) *GetAccountService {
	s.omitZeroBalances = &v
	return s
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/account",
		secType:  secTypeSigned,
	}
	if s.omitZeroBalances != nil {
		r.setParam("omitZeroBalances", *s.omitZeroBalances)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
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
	MakerCommission  int64           `json:"makerCommission"`
	TakerCommission  int64           `json:"takerCommission"`
	BuyerCommission  int64           `json:"buyerCommission"`
	SellerCommission int64           `json:"sellerCommission"`
	CommissionRates  CommissionRates `json:"commissionRates"`
	CanTrade         bool            `json:"canTrade"`
	CanWithdraw      bool            `json:"canWithdraw"`
	CanDeposit       bool            `json:"canDeposit"`
	UpdateTime       uint64          `json:"updateTime"`
	AccountType      string          `json:"accountType"`
	Balances         []Balance       `json:"balances"`
	Permissions      []string        `json:"permissions"`
	UID              int64           `json:"uid"`
}

// Balance define user balance of your account
type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// GetAccountSnapshotService all account orders; active, canceled, or filled
type GetAccountSnapshotService struct {
	c           *Client
	accountType string
	startTime   *int64
	endTime     *int64
	limit       *int
}

type CommissionRates struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

// Type set account type ("SPOT", "MARGIN", "FUTURES")
func (s *GetAccountSnapshotService) Type(accountType string) *GetAccountSnapshotService {
	s.accountType = accountType
	return s
}

// StartTime set starttime
func (s *GetAccountSnapshotService) StartTime(startTime int64) *GetAccountSnapshotService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *GetAccountSnapshotService) EndTime(endTime int64) *GetAccountSnapshotService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetAccountSnapshotService) Limit(limit int) *GetAccountSnapshotService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetAccountSnapshotService) Do(ctx context.Context, opts ...RequestOption) (res *Snapshot, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/accountSnapshot",
		secType:  secTypeSigned,
	}
	r.setParam("type", s.accountType)

	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Snapshot{}, err
	}
	res = new(Snapshot)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &Snapshot{}, err
	}
	return res, nil
}

// Snapshot define snapshot
type Snapshot struct {
	Code     int            `json:"code"`
	Msg      string         `json:"msg"`
	Snapshot []*SnapshotVos `json:"snapshotVos"`
}

// SnapshotVos define content of a snapshot
type SnapshotVos struct {
	Data       *SnapshotData `json:"data"`
	Type       string        `json:"type"`
	UpdateTime int64         `json:"updateTime"`
}

// SnapshotData define content of a snapshot
type SnapshotData struct {
	MarginLevel         string `json:"marginLevel"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`

	Balances   []*SnapshotBalances   `json:"balances"`
	UserAssets []*SnapshotUserAssets `json:"userAssets"`
	Assets     []*SnapshotAssets     `json:"assets"`
	Positions  []*SnapshotPositions  `json:"position"`
}

// SnapshotBalances define snapshot balances
type SnapshotBalances struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// SnapshotUserAssets define snapshot user assets
type SnapshotUserAssets struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// SnapshotAssets define snapshot assets
type SnapshotAssets struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	WalletBalance string `json:"walletBalance"`
}

// SnapshotPositions define snapshot positions
type SnapshotPositions struct {
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
}

// GetAPIKeyPermission get API Key permission info
type GetAPIKeyPermission struct {
	c *Client
}

// Do send request
func (s *GetAPIKeyPermission) Do(ctx context.Context, opts ...RequestOption) (res *APIKeyPermission, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/account/apiRestrictions",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(APIKeyPermission)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// APIKeyPermission define API key permission
type APIKeyPermission struct {
	IPRestrict                     bool   `json:"ipRestrict"`
	CreateTime                     uint64 `json:"createTime"`
	EnableWithdrawals              bool   `json:"enableWithdrawals"`
	EnableInternalTransfer         bool   `json:"enableInternalTransfer"`
	PermitsUniversalTransfer       bool   `json:"permitsUniversalTransfer"`
	EnableVanillaOptions           bool   `json:"enableVanillaOptions"`
	EnableReading                  bool   `json:"enableReading"`
	EnableFutures                  bool   `json:"enableFutures"`
	EnableMargin                   bool   `json:"enableMargin"`
	EnableSpotAndMarginTrading     bool   `json:"enableSpotAndMarginTrading"`
	TradingAuthorityExpirationTime uint64 `json:"tradingAuthorityExpirationTime"`
}

// GetAccountService get account info
// See https://developers.binance.com/docs/zh-CN/wallet/account
type GetSapiAccountInfoService struct {
	c *Client
}

type SapiAccountInfo struct {
	VipLevel                       int  `json:"vipLevel"`
	IsMarginEnabled                bool `json:"isMarginEnabled"`
	IsFutureEnabled                bool `json:"isFutureEnabled"`
	IsOptionsEnabled               bool `json:"isOptionsEnabled"`
	IsPortfolioMarginRetailEnabled bool `json:"isPortfolioMarginRetailEnabled"`
}

// Do send request
func (s *GetSapiAccountInfoService) Do(ctx context.Context, opts ...RequestOption) (res *SapiAccountInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/account/info",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SapiAccountInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
