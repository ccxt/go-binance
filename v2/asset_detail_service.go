package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAssetDetailService fetches all asset detail.
//
// See https://binance-docs.github.io/apidocs/spot/en/#asset-detail-user_data
type GetAssetDetailService struct {
	c     *Client
	asset *string
}

// Asset sets the asset parameter.
func (s *GetAssetDetailService) Asset(asset string) *GetAssetDetailService {
	s.asset = &asset
	return s
}

// Do sends the request.
func (s *GetAssetDetailService) Do(ctx context.Context) (res map[string]AssetDetail, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/assetDetail",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make(map[string]AssetDetail)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

type GetAllCoinsInfoService struct {
	c *Client
}

// Do send request
func (s *GetAllCoinsInfoService) Do(ctx context.Context) (res []*CoinInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/config/getall",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return []*CoinInfo{}, err
	}
	res = make([]*CoinInfo, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CoinInfo{}, err
	}
	return res, nil
}

// AssetDetail represents the detail of an asset
type AssetDetail struct {
	MinWithdrawAmount string `json:"minWithdrawAmount"`
	DepositStatus     bool   `json:"depositStatus"`
	WithdrawFee       string `json:"withdrawFee"`
	WithdrawStatus    bool   `json:"withdrawStatus"`
	DepositTip        string `json:"depositTip"`
}

type CoinInfo struct {
	Coin              string    `json:"coin"`
	DepositAllEnable  bool      `json:"depositAllEnable"`
	Free              string    `json:"free"`
	Freeze            string    `json:"freeze"`
	Ipoable           string    `json:"ipoable"`
	Ipoing            string    `json:"ipoing"`
	IsLegalMoney      bool      `json:"isLegalMoney"`
	Locked            string    `json:"locked"`
	Name              string    `json:"name"`
	NetworkList       []Network `json:"networkList"`
	Storage           string    `json:"storage"`
	Trading           bool      `json:"trading"`
	WithdrawAllEnable bool      `json:"withdrawAllEnable"`
	Withdrawing       string    `json:"withdrawing"`
}

type Network struct {
	AddressRegex            string `json:"addressRegex"`
	Coin                    string `json:"coin"`
	DepositDesc             string `json:"depositDesc,omitempty"` // 仅在充值关闭时返回
	DepositEnable           bool   `json:"depositEnable"`
	IsDefault               bool   `json:"isDefault"`
	MemoRegex               string `json:"memoRegex"`
	MinConfirm              int    `json:"minConfirm"` // 上账所需的最小确认数
	Name                    string `json:"name"`
	Network                 string `json:"network"`
	ResetAddressStatus      bool   `json:"resetAddressStatus"`
	SpecialTips             string `json:"specialTips"`
	UnLockConfirm           int    `json:"unLockConfirm"`          // 解锁需要的确认数
	WithdrawDesc            string `json:"withdrawDesc,omitempty"` // 仅在提现关闭时返回
	WithdrawEnable          bool   `json:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
	WithdrawMax             string `json:"withdrawMax"`
	WithdrawMin             string `json:"withdrawMin"`
	SameAddress             bool   `json:"sameAddress"` // 是否需要memo
	EstimatedArrivalTime    int    `json:"estimatedArrivalTime"`
	Busy                    bool   `json:"busy"`
	ContractAddressUrl      string `json:"contractAddressUrl"`
	ContractAddress         string `json:"contractAddress"`
}

// GetUserAssetService Get user assets
// See https://binance-docs.github.io/apidocs/spot/en/#user-asset-user_data
type GetUserAssetService struct {
	c                *Client
	asset            *string
	needBtcValuation bool
}

func (s *GetUserAssetService) Asset(asset string) *GetUserAssetService {
	s.asset = &asset
	return s
}

func (s *GetUserAssetService) NeedBtcValuation(val bool) *GetUserAssetService {
	s.needBtcValuation = val
	return s
}

type UserAssetRecord struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}

func (s *GetUserAssetService) Do(ctx context.Context) (res []UserAssetRecord, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v3/asset/getUserAsset",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.needBtcValuation {
		r.setParam("needBtcValuation", s.needBtcValuation)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	return
}

// GetFundingAssetService funding wallet
// See https://developers.binance.com/docs/wallet/asset/funding-wallet
type GetFundingAssetService struct {
	c                *Client
	asset            *string
	needBtcValuation *string
}

// Asset set asset
func (s *GetFundingAssetService) Asset(asset string) *GetFundingAssetService {
	s.asset = &asset
	return s
}

// NeedBtcValuation set needBtcValuation
func (s *GetFundingAssetService) NeedBtcValuation(needBtcValuation string) *GetFundingAssetService {
	s.needBtcValuation = &needBtcValuation
	return s
}

// FundingAsset define response of GetFundingAssetService
type FundingAsset struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	BtcValuation string `json:"btcValuation"`
}

func (s *GetFundingAssetService) Do(ctx context.Context) (res []FundingAsset, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/asset/get-funding-asset",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.needBtcValuation != nil {
		r.setParam("needBtcValuation", *s.needBtcValuation)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]FundingAsset, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
