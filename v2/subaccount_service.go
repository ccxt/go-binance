package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// TransferToSubAccountService transfer to subaccount
type TransferToSubAccountService struct {
	c       *Client
	toEmail string
	asset   string
	amount  string
}

// ToEmail set toEmail
func (s *TransferToSubAccountService) ToEmail(toEmail string) *TransferToSubAccountService {
	s.toEmail = toEmail
	return s
}

// Asset set asset
func (s *TransferToSubAccountService) Asset(asset string) *TransferToSubAccountService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *TransferToSubAccountService) Amount(amount string) *TransferToSubAccountService {
	s.amount = amount
	return s
}

func (s *TransferToSubAccountService) transferToSubaccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"toEmail": s.toEmail,
		"asset":   s.asset,
		"amount":  s.amount,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *TransferToSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *TransferToSubAccountResponse, err error) {
	data, err := s.transferToSubaccount(ctx, "/sapi/v1/sub-account/transfer/subToSub", opts...)
	if err != nil {
		return nil, err
	}
	res = &TransferToSubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TransferToSubAccountResponse define transfer to subaccount response
type TransferToSubAccountResponse struct {
	TxnID int64 `json:"txnId"`
}

type SubaccountDepositAddressService struct {
	c       *Client
	email   string
	coin    string
	network string
}

// Email set email
func (s *SubaccountDepositAddressService) Email(email string) *SubaccountDepositAddressService {
	s.email = email
	return s
}

// Coin set coin
func (s *SubaccountDepositAddressService) Coin(coin string) *SubaccountDepositAddressService {
	s.coin = coin
	return s
}

// Network set network
func (s *SubaccountDepositAddressService) Network(network string) *SubaccountDepositAddressService {
	s.network = network
	return s
}

func (s *SubaccountDepositAddressService) subaccountDepositAddress(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email":   s.email,
		"coin":    s.coin,
		"network": s.network,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountDepositAddressResponse, err error) {
	data, err := s.subaccountDepositAddress(ctx, "/sapi/v1/capital/deposit/subAddress", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountDepositAddressResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubaccountDepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	URL     string `json:"url"`
}

type SubaccountAssetsService struct {
	c     *Client
	email string
}

// Email set email
func (s *SubaccountAssetsService) Email(email string) *SubaccountAssetsService {
	s.email = email
	return s
}

func (s *SubaccountAssetsService) subaccountAssets(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email": s.email,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountAssetsService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountAssetsResponse, err error) {
	data, err := s.subaccountAssets(ctx, "/sapi/v3/sub-account/assets", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountAssetsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubaccountAssetsResponse Query Sub-account Assets response
type SubaccountAssetsResponse struct {
	Balances []AssetBalance `json:"balances"`
}

type AssetBalance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free"`
	Locked float64 `json:"locked"`
}

type SubaccountSpotSummaryService struct {
	c     *Client
	email *string
	page  *int32
	size  *int32
}

// Email set email
func (s *SubaccountSpotSummaryService) Email(email string) *SubaccountSpotSummaryService {
	s.email = &email
	return s
}

func (s *SubaccountSpotSummaryService) Page(page int32) *SubaccountSpotSummaryService {
	s.page = &page
	return s
}

func (s *SubaccountSpotSummaryService) Size(size int32) *SubaccountSpotSummaryService {
	s.size = &size
	return s
}

func (s *SubaccountSpotSummaryService) subaccountSpotSummary(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	if s.size != nil {
		r.setParam("size", *s.size)
	}

	if s.page != nil {
		r.setParam("page", *s.page)
	}

	if s.email != nil {
		r.setParam("email", *s.email)
	}
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountSpotSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountSpotSummaryResponse, err error) {
	data, err := s.subaccountSpotSummary(ctx, "/sapi/v1/sub-account/spotSummary", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountSpotSummaryResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubaccountSpotSummaryResponse Query Sub-account Spot Assets Summary response
type SubaccountSpotSummaryResponse struct {
	TotalCount                int64                       `json:"totalCount"`
	MasterAccountTotalAsset   string                      `json:"masterAccountTotalAsset"`
	SpotSubUserAssetBtcVoList []SpotSubUserAssetBtcVoList `json:"spotSubUserAssetBtcVoList"`
}

type SpotSubUserAssetBtcVoList struct {
	Email      string `json:"email"`
	TotalAsset string `json:"totalAsset"`
}

// SubAccountListService Query Sub-account List (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#query-sub-account-list-for-master-account
type SubAccountListService struct {
	c           *Client
	email       *string
	isFreeze    bool
	page, limit int
}

func (s *SubAccountListService) Email(v string) *SubAccountListService {
	s.email = &v
	return s
}

func (s *SubAccountListService) IsFreeze(v bool) *SubAccountListService {
	s.isFreeze = v
	return s
}

func (s *SubAccountListService) Page(v int) *SubAccountListService {
	s.page = v
	return s
}

func (s *SubAccountListService) Limit(v int) *SubAccountListService {
	s.limit = v
	return s
}

func (s *SubAccountListService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountList, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/list",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	if s.isFreeze {
		r.setParam("isFreeze", "true")
	} else {
		r.setParam("isFreeze", "false")
	}
	if s.page > 0 {
		r.setParam("page", s.page)
	}
	if s.limit > 200 {
		r.setParam("limit", 200)
	} else if s.limit <= 0 {
		r.setParam("limit", 10)
	} else {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountList)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountList struct {
	SubAccounts []SubAccount `json:"subAccounts"`
}

type SubAccount struct {
	SubUserID                   string `json:"subUserId"`
	Email                       string `json:"email"`
	Remark                      string `json:"remark"`
	IsFreeze                    bool   `json:"isFreeze"`
	CreateTime                  uint64 `json:"createTime"`
	IsManagedSubAccount         bool   `json:"isManagedSubAccount"`
	IsAssetManagementSubAccount bool   `json:"isAssetManagementSubAccount"`
}

// ManagedSubAccountDepositService
// Deposit Assets Into The Managed Sub-account（For Investor Master Account）
// https://developers.binance.com/docs/sub_account/managed-sub-account
type ManagedSubAccountDepositService struct {
	c       *Client
	toEmail string
	asset   string
	amount  string
}

func (s *ManagedSubAccountDepositService) ToEmail(email string) *ManagedSubAccountDepositService {
	s.toEmail = email
	return s
}

func (s *ManagedSubAccountDepositService) Asset(asset string) *ManagedSubAccountDepositService {
	s.asset = asset
	return s
}

func (s *ManagedSubAccountDepositService) Amount(amount string) *ManagedSubAccountDepositService {
	s.amount = amount
	return s
}

type ManagedSubAccountDepositResponse struct {
	ID int64 `json:"tranId"`
}

// Do send request
func (s *ManagedSubAccountDepositService) Do(ctx context.Context, opts ...RequestOption) (*ManagedSubAccountDepositResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/managed-subaccount/deposit",
		secType:  secTypeSigned,
	}

	r.setParam("toEmail", s.toEmail)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &ManagedSubAccountDepositResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ManagedSubAccountWithdrawalService
// Withdrawal Assets From The Managed Sub-account（For Investor Master Account）
// https://developers.binance.com/docs/sub_account/managed-sub-account/Withdrawl-Assets-From-The-Managed-Sub-account
type ManagedSubAccountWithdrawalService struct {
	c            *Client
	fromEmail    string
	asset        string
	amount       string
	transferDate int64 // Withdrawals is automatically occur on the transfer date(UTC0). If a date is not selected, the withdrawal occurs right now
}

func (s *ManagedSubAccountWithdrawalService) FromEmail(email string) *ManagedSubAccountWithdrawalService {
	s.fromEmail = email
	return s
}

func (s *ManagedSubAccountWithdrawalService) Asset(asset string) *ManagedSubAccountWithdrawalService {
	s.asset = asset
	return s
}

func (s *ManagedSubAccountWithdrawalService) Amount(amount string) *ManagedSubAccountWithdrawalService {
	s.amount = amount
	return s
}

func (s *ManagedSubAccountWithdrawalService) TransferDate(val int64) *ManagedSubAccountWithdrawalService {
	s.transferDate = val
	return s
}

type ManagedSubAccountWithdrawalResponse struct {
	ID int64 `json:"tranId"`
}

// Do send request
func (s *ManagedSubAccountWithdrawalService) Do(ctx context.Context, opts ...RequestOption) (*ManagedSubAccountWithdrawalResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/managed-subaccount/withdraw",
		secType:  secTypeSigned,
	}

	r.setParam("fromEmail", s.fromEmail)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	if s.transferDate > 0 {
		r.setParam("transferDate", s.transferDate)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &ManagedSubAccountWithdrawalResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ManagedSubAccountAssetsService
// Query Managed Sub-account Asset Details（For Investor Master Account）
// https://binance-docs.github.io/apidocs/spot/en/#query-managed-sub-account-asset-details-for-investor-master-account
type ManagedSubAccountAssetsService struct {
	c     *Client
	email string
}

func (s *ManagedSubAccountAssetsService) Email(email string) *ManagedSubAccountAssetsService {
	s.email = email
	return s
}

type ManagedSubAccountAsset struct {
	Coin             string `json:"coin"`
	Name             string `json:"name"`
	TotalBalance     string `json:"totalBalance"`
	AvailableBalance string `json:"availableBalance"`
	InOrder          string `json:"inOrder"`
	BtcValue         string `json:"btcValue"`
}

func (s *ManagedSubAccountAssetsService) Do(ctx context.Context, opts ...RequestOption) ([]*ManagedSubAccountAsset, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/managed-subaccount/asset",
		secType:  secTypeSigned,
	}

	r.setParam("email", s.email)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*ManagedSubAccountAsset, 0)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// SubAccountFuturesAccountService Get Detail on Sub-account's Futures Account (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#get-detail-on-sub-account-39-s-futures-account-for-master-account
type SubAccountFuturesAccountService struct {
	c     *Client
	email *string
}

func (s *SubAccountFuturesAccountService) Email(v string) *SubAccountFuturesAccountService {
	s.email = &v
	return s
}

func (s *SubAccountFuturesAccountService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesAccount, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/futures/account",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountFuturesAccount struct {
	Email                       string                          `json:"email"`
	Asset                       string                          `json:"asset"`
	Assets                      []SubAccountFuturesAccountAsset `json:"assets"`
	CanDeposit                  bool                            `json:"canDeposit"`
	CanTrade                    bool                            `json:"canTrade"`
	CanWithdraw                 bool                            `json:"canWithdraw"`
	FeeTier                     int                             `json:"feeTier"`
	MaxWithdrawAmount           string                          `json:"maxWithdrawAmount"`
	TotalInitialMargin          string                          `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string                          `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string                          `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string                          `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string                          `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string                          `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string                          `json:"totalWalletBalance"`
	UpdateTime                  int64                           `json:"updateTime"`
}

type SubAccountFuturesAccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}

// SubAccountFuturesSummaryV1Service Get Summary of Sub-account's Futures Account (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#get-summary-of-sub-account-39-s-futures-account-for-master-account
type SubAccountFuturesSummaryV1Service struct {
	c *Client
}

func (s *SubAccountFuturesSummaryV1Service) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesSummaryV1, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/futures/accountSummary",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesSummaryV1)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountFuturesSummaryCommon struct {
	Asset                       string `json:"asset"`
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
}

type SubAccountFuturesSummaryV1 struct {
	SubAccountFuturesSummaryCommon
	SubAccountList []SubAccountFuturesSummaryV1SubAccountList `json:"subAccountList"`
}

type SubAccountFuturesSummaryV1SubAccountList struct {
	Email string `json:"email"`
	SubAccountFuturesSummaryCommon
}

// SubAccountFuturesTransferV1Service Futures Transfer for Sub-account (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#futures-transfer-for-sub-account-for-master-account
type SubAccountFuturesTransferV1Service struct {
	c      *Client
	email  string
	asset  string
	amount float64
	/*
		1: transfer from subaccount's spot account to its USDT-margined futures account
		2: transfer from subaccount's USDT-margined futures account to its spot account
		3: transfer from subaccount's spot account to its COIN-margined futures account
		4:transfer from subaccount's COIN-margined futures account to its spot account
	*/
	transferType int
}

func (s *SubAccountFuturesTransferV1Service) Email(v string) *SubAccountFuturesTransferV1Service {
	s.email = v
	return s
}

func (s *SubAccountFuturesTransferV1Service) Asset(v string) *SubAccountFuturesTransferV1Service {
	s.asset = v
	return s
}

func (s *SubAccountFuturesTransferV1Service) Amount(v float64) *SubAccountFuturesTransferV1Service {
	s.amount = v
	return s
}

func (s *SubAccountFuturesTransferV1Service) TransferType(v int) *SubAccountFuturesTransferV1Service {
	s.transferType = v
	return s
}

func (s *SubAccountFuturesTransferV1Service) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesTransferResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/futures/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"email":  s.email,
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountFuturesTransferResponse struct {
	// seems api doc bug, return `tranId` as int64 actually in production environment
	TranID int64 `json:"tranId"`
}

// Sub-account Transfer History (For Sub-account)
// https://binance-docs.github.io/apidocs/spot/en/#sub-account-transfer-history-for-sub-account
type SubAccountTransferHistoryService struct {
	c                 *Client
	asset             *string
	transferType      *SubAccountTransferType
	startTime         *int64
	endTime           *int64
	limit             *int
	returnFailHistory *bool
}

func (s *SubAccountTransferHistoryService) Asset(v string) *SubAccountTransferHistoryService {
	s.asset = &v
	return s
}

func (s *SubAccountTransferHistoryService) TransferType(v SubAccountTransferType) *SubAccountTransferHistoryService {
	s.transferType = &v
	return s
}

func (s *SubAccountTransferHistoryService) StartTime(v int64) *SubAccountTransferHistoryService {
	s.startTime = &v
	return s
}

func (s *SubAccountTransferHistoryService) EndTime(v int64) *SubAccountTransferHistoryService {
	s.endTime = &v
	return s
}

func (s *SubAccountTransferHistoryService) Limit(v int) *SubAccountTransferHistoryService {
	s.limit = &v
	return s
}

func (s *SubAccountTransferHistoryService) ReturnFailHistory(v bool) *SubAccountTransferHistoryService {
	s.returnFailHistory = &v
	return s
}

func (s *SubAccountTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccountTransferHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/transfer/subUserHistory",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.transferType != nil {
		r.setParam("type", *s.transferType)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.returnFailHistory != nil {
		r.setParam("returnFailHistory", *s.returnFailHistory)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccountTransferHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountTransferHistory struct {
	CounterParty    string                 `json:"counterParty"`
	Email           string                 `json:"email"`
	Type            SubAccountTransferType `json:"type"`
	Asset           string                 `json:"asset"`
	Qty             string                 `json:"qty"`
	FromAccountType AccountType            `json:"fromAccountType"`
	ToAccountType   AccountType            `json:"toAccountType"`
	Status          string                 `json:"status"`
	TranID          int64                  `json:"tranId"`
	Time            int64                  `json:"time"`
}

// Create virtual sub-account
type CreateVirtualSubAccountService struct {
	c                *Client
	subAccountString string
	recvWindow       *int64
}

type CreateVirtualSubAccountResponse struct {
	Email string `json:"email"`
}

func (s *CreateVirtualSubAccountService) SubAccountString(subAccountString string) *CreateVirtualSubAccountService {
	s.subAccountString = subAccountString
	return s
}

func (s *CreateVirtualSubAccountService) RecvWindow(recvWindow int64) *CreateVirtualSubAccountService {
	s.recvWindow = &recvWindow
	return s
}

func (s *CreateVirtualSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *CreateVirtualSubAccountResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/virtualSubAccount",
		secType:  secTypeSigned,
	}
	r.setFormParam("subAccountString", s.subAccountString)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateVirtualSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Sub-Account spot transfer history
type SubAccountSpotTransferHistoryService struct {
	c          *Client
	fromEmail  *string // Sub-account email
	toEmail    *string // true or false
	startTime  *uint64
	endTime    *uint64
	page       *int32 // default 1
	limit      *int32 // default 1, max 200
	recvWindow *int64
}

type SubAccountSpotTransfer struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Asset  string `json:"asset"`
	Qty    string `json:"qty"`
	Status string `json:"status"`
	TranId uint64 `json:"tranId"`
	Time   uint64 `json:"time"`
}

func (s *SubAccountSpotTransferHistoryService) FromEmail(fromEmail string) *SubAccountSpotTransferHistoryService {
	s.fromEmail = &fromEmail
	return s
}

func (s *SubAccountSpotTransferHistoryService) ToEmail(toEmail string) *SubAccountSpotTransferHistoryService {
	s.toEmail = &toEmail
	return s
}

func (s *SubAccountSpotTransferHistoryService) StartTime(startTime uint64) *SubAccountSpotTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *SubAccountSpotTransferHistoryService) EndTime(endTime uint64) *SubAccountSpotTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *SubAccountSpotTransferHistoryService) Page(page int32) *SubAccountSpotTransferHistoryService {
	s.page = &page
	return s
}

func (s *SubAccountSpotTransferHistoryService) Limit(limit int32) *SubAccountSpotTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *SubAccountSpotTransferHistoryService) RecvWindow(recvWindow int64) *SubAccountSpotTransferHistoryService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountSpotTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccountSpotTransfer, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/sub/transfer/history",
		secType:  secTypeSigned,
	}
	if s.fromEmail != nil {
		r.setParam("fromEmail", *s.fromEmail)
	}
	if s.toEmail != nil {
		r.setParam("toEmail", *s.toEmail)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccountSpotTransfer, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Sub-Account futures transfer history
type SubAccountFuturesTransferHistoryService struct {
	c           *Client
	email       string // Sub-account email
	futuresType int64  // 1: usdt 2: coin
	startTime   *int64
	endTime     *int64
	page        *int32 // default 1
	limit       *int32 // default 50, max 500
	recvWindow  *int64
}

type SubAccountFuturesTransferHistoryResponse struct {
	Success     bool                         `json:"success"`
	FuturesType int32                        `json:"futuresType"`
	Transfers   []*SubAccountFuturesTransfer `json:"transfers"`
}

type SubAccountFuturesTransfer struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Asset  string `json:"asset"`
	Qty    string `json:"qty"`
	Status string `json:"status"`
	TranId uint64 `json:"tranId"`
	Time   uint64 `json:"time"`
}

func (s *SubAccountFuturesTransferHistoryService) Email(email string) *SubAccountFuturesTransferHistoryService {
	s.email = email
	return s
}

func (s *SubAccountFuturesTransferHistoryService) FuturesType(futuresType int64) *SubAccountFuturesTransferHistoryService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccountFuturesTransferHistoryService) StartTime(startTime int64) *SubAccountFuturesTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *SubAccountFuturesTransferHistoryService) EndTime(endTime int64) *SubAccountFuturesTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *SubAccountFuturesTransferHistoryService) Page(page int32) *SubAccountFuturesTransferHistoryService {
	s.page = &page
	return s
}

func (s *SubAccountFuturesTransferHistoryService) Limit(limit int32) *SubAccountFuturesTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *SubAccountFuturesTransferHistoryService) RecvWindow(recvWindow int64) *SubAccountFuturesTransferHistoryService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountFuturesTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesTransferHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/futures/internalTransfer",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("futuresType", s.futuresType)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesTransferHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Execute sub account futures balance transfer
type SubAccountFuturesInternalTransferService struct {
	c           *Client
	fromEmail   string // sender email
	toEmail     string // receiver email
	futuresType int64  // 1: usdt 2: coin
	asset       string
	amount      string // decimal format
	recvWindow  *int64
}

type SubAccountFuturesInternalTransferResponse struct {
	Success bool   `json:"success"`
	TxnId   string `json:"txnId"`
}

func (s *SubAccountFuturesInternalTransferService) FromEmail(fromEmail string) *SubAccountFuturesInternalTransferService {
	s.fromEmail = fromEmail
	return s
}

func (s *SubAccountFuturesInternalTransferService) ToEmail(toEmail string) *SubAccountFuturesInternalTransferService {
	s.toEmail = toEmail
	return s
}

func (s *SubAccountFuturesInternalTransferService) FuturesType(futuresType int64) *SubAccountFuturesInternalTransferService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccountFuturesInternalTransferService) Asset(asset string) *SubAccountFuturesInternalTransferService {
	s.asset = asset
	return s
}

func (s *SubAccountFuturesInternalTransferService) Amount(amount string) *SubAccountFuturesInternalTransferService {
	s.amount = amount
	return s
}

func (s *SubAccountFuturesInternalTransferService) RecvWindow(recvWindow int64) *SubAccountFuturesInternalTransferService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountFuturesInternalTransferService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesInternalTransferResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/futures/internalTransfer",
		secType:  secTypeSigned,
	}
	r.setFormParam("fromEmail", s.fromEmail)
	r.setFormParam("toEmail", s.toEmail)
	r.setFormParam("futuresType", s.futuresType)
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesInternalTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get sub account deposit record
type SubAccountDepositRecordService struct {
	c          *Client
	email      string // sub-account email
	coin       *string
	status     *int32 // 0(0:pending,6: credited but cannot withdraw,7:Wrong Deposit,8:Waiting User confirm,1:success)
	startTime  *int64
	endTime    *int64
	limit      *int // sub-account email
	offset     *int // default 0
	txId       *string
	recvWindow *int64
}

type SubAccountDepositRecord struct {
	Id            string `json:"id"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int32  `json:"status"`
	Address       string `json:"address"`
	TxId          string `json:"txId"`
	AddressTag    string `json:"addressTag"`
	InsertTime    int64  `json:"insertTime"`
	TransferType  int32  `json:"transferType"`
	ConfirmTimes  string `json:"confirmTimes"`
	UnlockConfirm int32  `json:"unlockConfirm"`
	WalletType    int32  `json:"walletType"`
}

func (s *SubAccountDepositRecordService) Email(email string) *SubAccountDepositRecordService {
	s.email = email
	return s
}

func (s *SubAccountDepositRecordService) Coin(coin string) *SubAccountDepositRecordService {
	s.coin = &coin
	return s
}

func (s *SubAccountDepositRecordService) Status(status int32) *SubAccountDepositRecordService {
	s.status = &status
	return s
}

func (s *SubAccountDepositRecordService) StartTime(startTime int64) *SubAccountDepositRecordService {
	s.startTime = &startTime
	return s
}

func (s *SubAccountDepositRecordService) EndTime(endTime int64) *SubAccountDepositRecordService {
	s.endTime = &endTime
	return s
}

func (s *SubAccountDepositRecordService) Limit(limit int) *SubAccountDepositRecordService {
	s.limit = &limit
	return s
}

func (s *SubAccountDepositRecordService) Offset(offset int) *SubAccountDepositRecordService {
	s.offset = &offset
	return s
}

func (s *SubAccountDepositRecordService) TxId(txId string) *SubAccountDepositRecordService {
	s.txId = &txId
	return s
}

func (s *SubAccountDepositRecordService) RecvWindow(recvWindow int64) *SubAccountDepositRecordService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountDepositRecordService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccountDepositRecord, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/subHisrec",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.txId != nil {
		r.setParam("txId", *s.txId)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccountDepositRecord, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get sub account margin futures status
type SubAccountMarginFuturesStatusService struct {
	c          *Client
	email      *string // sub-account email
	recvWindow *int64
}

type SubAccountMarginFuturesStatus struct {
	Email            string `json:"email"`
	IsSubUserEnabled bool   `json:"isSubUserEnabled"`
	IsUserActive     bool   `json:"isUserActive"`
	InsertTime       int64  `json:"insertTime"`
	IsMarginEnabled  bool   `json:"isMarginEnabled"`
	IsFutureEnabled  bool   `json:"isFutureEnabled"`
	Mobile           int64  `json:"mobile"`
}

func (s *SubAccountMarginFuturesStatusService) Email(email string) *SubAccountMarginFuturesStatusService {
	s.email = &email
	return s
}

func (s *SubAccountMarginFuturesStatusService) RecvWindow(recvWindow int64) *SubAccountMarginFuturesStatusService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountMarginFuturesStatusService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccountMarginFuturesStatus, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/status",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccountMarginFuturesStatus, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// sub account margin enable
type SubAccountMarginEnableService struct {
	c          *Client
	email      string // sub-account email
	recvWindow *int64
}

type SubAccountMarginEnableResponse struct {
	Email           string `json:"email"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
}

func (s *SubAccountMarginEnableService) Email(email string) *SubAccountMarginEnableService {
	s.email = email
	return s
}

func (s *SubAccountMarginEnableService) RecvWindow(recvWindow int64) *SubAccountMarginEnableService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountMarginEnableService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountMarginEnableResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/margin/enable",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountMarginEnableResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get sub-account margin account detail
type SubAccountMarginAccountInfoService struct {
	c          *Client
	email      string // sub-account email
	recvWindow *int64
}

type SubAccountMarginAccountInfo struct {
	Email                 string               `json:"email"`
	MarginLevel           string               `json:"marginLevel"`
	TotalAssetOfBtc       string               `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc   string               `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc    string               `json:"totalNetAssetOfBtc"`
	MarginTradeCoeffVo    *MarginTradeCoeffVo  `json:"marginTradeCoeffVo"`
	MarginUserAssetVoList []*MarginUserAssetVo `json:"marginUserAssetVoList"`
}

type MarginTradeCoeffVo struct {
	ForceLiquidationBar string `json:"forceLiquidationBar"`
	MarginCallBar       string `json:"marginCallBar"`
	NormalBar           string `json:"normalBar"`
}

type MarginUserAssetVo struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

func (s *SubAccountMarginAccountInfoService) Email(email string) *SubAccountMarginAccountInfoService {
	s.email = email
	return s
}

func (s *SubAccountMarginAccountInfoService) RecvWindow(recvWindow int64) *SubAccountMarginAccountInfoService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountMarginAccountInfoService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountMarginAccountInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/margin/account",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountMarginAccountInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get sub-account margin account summary
type SubAccountMarginAccountSummaryService struct {
	c          *Client
	recvWindow *int64
}

type SubAccountMarginAccountSummary struct {
	TotalAssetOfBtc     string              `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string              `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string              `json:"totalNetAssetOfBtc"`
	SubAccountList      []*MarginSubAccount `json:"subAccountList"`
}

type MarginSubAccount struct {
	Email               string `json:"email"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
}

func (s *SubAccountMarginAccountSummaryService) RecvWindow(recvWindow int64) *SubAccountMarginAccountSummaryService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountMarginAccountSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountMarginAccountSummary, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/margin/accountSummary",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountMarginAccountSummary)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type SubAccountFuturesEnableService struct {
	c          *Client
	email      string // sub-account email
	recvWindow *int64
}

type SubAccountFuturesEnableResponse struct {
	Email            string `json:"email"`
	IsFuturesEnabled bool   `json:"isFuturesEnabled"`
}

func (s *SubAccountFuturesEnableService) Email(email string) *SubAccountFuturesEnableService {
	s.email = email
	return s
}

func (s *SubAccountFuturesEnableService) RecvWindow(recvWindow int64) *SubAccountFuturesEnableService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountFuturesEnableService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesEnableResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/futures/enable",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesEnableResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get the target sub-account futures account detail, v2 interface.
type SubAccountFuturesAccountV2Service struct {
	c           *Client
	email       string // sub-account email
	futuresType int32  // 1:USDT Margined Futures, 2:COIN Margined Futures
	recvWindow  *int64
}

type SubAccountFuturesAccountV2ServiceResponse struct {
	FutureAccountResp   *SubAccountFuturesAccountV2  `json:"futureAccountResp"`   // set while futuresType=1(USDT margined)
	DeliveryAccountResp *SubAccountDeliveryAccountV2 `json:"deliveryAccountResp"` // set while futuresType=2(COIN margined)
}

type SubAccountFuturesAccountV2 struct {
	Email                       string          `json:"email"`
	Asset                       string          `json:"asset"`
	Assets                      []*FuturesAsset `json:"assets"`
	CanDeposit                  bool            `json:"canDeposit"`
	CanTrade                    bool            `json:"canTrade"`
	CanWithdraw                 bool            `json:"canWithdraw"`
	FeeTier                     int32           `json:"feeTier"`
	MaxWithdrawAmount           string          `json:"maxWithdrawAmount"`
	TotalInitialMargin          string          `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string          `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string          `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string          `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string          `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string          `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string          `json:"totalWalletBalance"`
	UpdateTime                  int64           `json:"updateTime"`
}

type SubAccountDeliveryAccountV2 struct {
	Email       string          `json:"email"`
	Assets      []*FuturesAsset `json:"assets"`
	CanDeposit  bool            `json:"canDeposit"`
	CanTrade    bool            `json:"canTrade"`
	CanWithdraw bool            `json:"canWithdraw"`
	FeeTier     int32           `json:"feeTier"`
	UpdateTime  int64           `json:"updateTime"`
}

type FuturesAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}

func (s *SubAccountFuturesAccountV2Service) Email(email string) *SubAccountFuturesAccountV2Service {
	s.email = email
	return s
}

func (s *SubAccountFuturesAccountV2Service) FuturesType(futuresType int32) *SubAccountFuturesAccountV2Service {
	s.futuresType = futuresType
	return s
}

func (s *SubAccountFuturesAccountV2Service) RecvWindow(recvWindow int64) *SubAccountFuturesAccountV2Service {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountFuturesAccountV2Service) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesAccountV2ServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v2/sub-account/futures/account",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("futuresType", s.futuresType)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesAccountV2ServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get sub-account futures account summary, include U-M and C-M, v2 interface
type SubAccountFuturesAccountSummaryService struct {
	c           *Client
	futuresType int32  // 1:USDT Margined Futures, 2:COIN Margined Futures
	page        *int32 // default 1
	limit       *int32 // default 10, max 20
	recvWindow  *int64
}

type SubAccountFuturesAccountSummaryServiceResponse struct {
	FutureAccountSummaryResp   *SubAccountFuturesAccountSummary  `json:"futureAccountSummaryResp"`   // set while futuresType=1
	DeliveryAccountSummaryResp *SubAccountDeliveryAccountSummary `json:"deliveryAccountSummaryResp"` // set while futuresType=2
}

type SubAccountFuturesAccountSummary struct {
	TotalInitialMargin          string               `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string               `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string               `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string               `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string               `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string               `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string               `json:"totalWalletBalance"`
	Asset                       string               `json:"asset"`
	SubAccountList              []*FuturesSubAccount `json:"subAccountList"`
}

type FuturesSubAccount struct {
	Email                       string `json:"email"`
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
	Asset                       string `json:"asset"`
}

type SubAccountDeliveryAccountSummary struct {
	TotalMarginBalanceOfBTC    string                `json:"totalMarginBalanceOfBTC"`
	TotalUnrealizedProfitOfBTC string                `json:"totalUnrealizedProfitOfBTC"`
	TotalWalletBalanceOfBTC    string                `json:"totalWalletBalanceOfBTC"`
	Asset                      string                `json:"asset"`
	SubAccountList             []*DeliverySubAccount `json:"subAccountList"`
}

type DeliverySubAccount struct {
	Email                 string `json:"email"`
	TotalMarginBalance    string `json:"totalMarginBalance"`
	TotalUnrealizedProfit string `json:"totalUnrealizedProfit"`
	TotalWalletBalance    string `json:"totalWalletBalance"`
	Asset                 string `json:"asset"`
}

func (s *SubAccountFuturesAccountSummaryService) FuturesType(futuresType int32) *SubAccountFuturesAccountSummaryService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccountFuturesAccountSummaryService) Page(page int32) *SubAccountFuturesAccountSummaryService {
	s.page = &page
	return s
}

func (s *SubAccountFuturesAccountSummaryService) Limit(limit int32) *SubAccountFuturesAccountSummaryService {
	s.limit = &limit
	return s
}

func (s *SubAccountFuturesAccountSummaryService) RecvWindow(recvWindow int64) *SubAccountFuturesAccountSummaryService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountFuturesAccountSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesAccountSummaryServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v2/sub-account/futures/accountSummary",
		secType:  secTypeSigned,
	}
	r.setParam("futuresType", s.futuresType)
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesAccountSummaryServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get target sub-account futures position information, include U-M and C-M, v2 interface.
type SubAccountFuturesPositionsService struct {
	c           *Client
	email       string
	futuresType int32 // 1:USDT Margined Futures, 2:COIN Margined Futures
	recvWindow  *int64
}

type SubAccountFuturesPositionsServiceResponse struct {
	FuturePositionRiskVos   []*SubAccountFuturesPosition  `json:"futurePositionRiskVos"`   // set while futuresType=1
	DeliveryPositionRiskVos []*SubAccountDeliveryPosition `json:"deliveryPositionRiskVos"` // set while futuresType=2
}

type SubAccountFuturesPosition struct {
	EntryPrice       string `json:"entryPrice"`
	Leverage         string `json:"leverage"`
	MaxNotional      string `json:"maxNotional"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmount   string `json:"positionAmount"`
	Symbol           string `json:"symbol"`
	UnrealizedProfit string `json:"unrealizedProfit"`
}

type SubAccountDeliveryPosition struct {
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	Leverage         string `json:"leverage"`
	Isolated         string `json:"isolated"`
	IsolatedWallet   string `json:"isolatedWallet"`
	IsolatedMargin   string `json:"isolatedMargin"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	PositionSide     string `json:"positionSide"`
	PositionAmount   string `json:"positionAmount"`
	Symbol           string `json:"symbol"`
	UnrealizedProfit string `json:"unrealizedProfit"`
}

func (s *SubAccountFuturesPositionsService) Email(email string) *SubAccountFuturesPositionsService {
	s.email = email
	return s
}

func (s *SubAccountFuturesPositionsService) FuturesType(futuresType int32) *SubAccountFuturesPositionsService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccountFuturesPositionsService) RecvWindow(recvWindow int64) *SubAccountFuturesPositionsService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountFuturesPositionsService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesPositionsServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v2/sub-account/futures/positionRisk",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("futuresType", s.futuresType)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesPositionsServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// execute sub-account margin account transfer
type SubAccountMarginTransferService struct {
	c            *Client
	email        string
	asset        string
	amount       string // decimal format
	transferType int32  // 1: Transfer from the spot account of the sub account to its leverage account; 2: Transfer from the leverage account of the sub account to its spot account
	recvWindow   *int64
}

type SubAccountMarginTransferResponse struct {
	TxnId string `json:"txnId"`
}

func (s *SubAccountMarginTransferService) Email(email string) *SubAccountMarginTransferService {
	s.email = email
	return s
}

func (s *SubAccountMarginTransferService) Asset(asset string) *SubAccountMarginTransferService {
	s.asset = asset
	return s
}

func (s *SubAccountMarginTransferService) Amount(amount string) *SubAccountMarginTransferService {
	s.amount = amount
	return s
}

func (s *SubAccountMarginTransferService) TransferType(transferType int32) *SubAccountMarginTransferService {
	s.transferType = transferType
	return s
}

func (s *SubAccountMarginTransferService) RecvWindow(recvWindow int64) *SubAccountMarginTransferService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountMarginTransferService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountMarginTransferResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/margin/transfer",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	r.setFormParam("type", s.transferType)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountMarginTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// sub-account transfer balance to master-account
type SubAccountTransferSubToMasterService struct {
	c          *Client
	asset      string
	amount     string // decimal format
	recvWindow *int64
}

type SubAccountTransferSubToMasterResponse struct {
	TxnId string `json:"txnId"`
}

func (s *SubAccountTransferSubToMasterService) Asset(asset string) *SubAccountTransferSubToMasterService {
	s.asset = asset
	return s
}

func (s *SubAccountTransferSubToMasterService) Amount(amount string) *SubAccountTransferSubToMasterService {
	s.amount = amount
	return s
}

func (s *SubAccountTransferSubToMasterService) RecvWindow(recvWindow int64) *SubAccountTransferSubToMasterService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountTransferSubToMasterService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountTransferSubToMasterResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/transfer/subToMaster",
		secType:  secTypeSigned,
	}
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountTransferSubToMasterResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Universal transfer of master and sub accounts
type SubAccountUniversalTransferService struct {
	c               *Client
	fromEmail       *string
	toEmail         *string
	fromAccountType string  // "SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	toAccountType   string  // "SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	clientTranId    *string // Non repeatable
	symbol          *string // Only used under ISOLATEd_MARGIN type
	asset           string
	amount          string // decimal format
	recvWindow      *int64
}

type SubAccountUniversalTransferResponse struct {
	TranId       int64  `json:"tranId"`
	ClientTranId string `json:"clientTranId"`
}

func (s *SubAccountUniversalTransferService) FromEmail(fromEmail string) *SubAccountUniversalTransferService {
	s.fromEmail = &fromEmail
	return s
}

func (s *SubAccountUniversalTransferService) ToEmail(toEmail string) *SubAccountUniversalTransferService {
	s.toEmail = &toEmail
	return s
}

func (s *SubAccountUniversalTransferService) FromAccountType(fromAccountType string) *SubAccountUniversalTransferService {
	s.fromAccountType = fromAccountType
	return s
}

func (s *SubAccountUniversalTransferService) ToAccountType(toAccountType string) *SubAccountUniversalTransferService {
	s.toAccountType = toAccountType
	return s
}

func (s *SubAccountUniversalTransferService) ClientTranId(clientTranId string) *SubAccountUniversalTransferService {
	s.clientTranId = &clientTranId
	return s
}

func (s *SubAccountUniversalTransferService) Symbol(symbol string) *SubAccountUniversalTransferService {
	s.symbol = &symbol
	return s
}

func (s *SubAccountUniversalTransferService) Asset(asset string) *SubAccountUniversalTransferService {
	s.asset = asset
	return s
}

func (s *SubAccountUniversalTransferService) Amount(amount string) *SubAccountUniversalTransferService {
	s.amount = amount
	return s
}

func (s *SubAccountUniversalTransferService) RecvWindow(recvWindow int64) *SubAccountUniversalTransferService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountUniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountUniversalTransferResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}
	if s.fromEmail != nil {
		r.setFormParam("fromEmail", *s.fromEmail)
	}
	if s.toEmail != nil {
		r.setFormParam("toEmail", *s.toEmail)
	}
	r.setFormParam("fromAccountType", s.fromAccountType)
	r.setFormParam("toAccountType", s.toAccountType)
	if s.clientTranId != nil {
		r.setFormParam("clientTranId", *s.clientTranId)
	}
	if s.symbol != nil {
		r.setFormParam("symbol", *s.symbol)
	}
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountUniversalTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query the universal transfer history of sub and master accounts
type SubAccUniversalTransferHistoryService struct {
	c            *Client
	fromEmail    *string
	toEmail      *string
	clientTranId *string // Non repeatable
	startTime    *int64
	endTime      *int64
	page         *int32 // default 1
	limit        *int32 // default 500, max 500
	recvWindow   *int64
}

type SubAccountUniversalTransferHistoryServiceResponse struct {
	Result     []*SubAccountUniversalTransferRecord `json:"result"`
	TotalCount int64                                `json:"totalCount"`
}

type SubAccountUniversalTransferRecord struct {
	TranId          int64  `json:"tranId"`
	FromEmail       string `json:"fromEmail"`
	ToEmail         string `json:"toEmail"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	CreateTimeStamp int64  `json:"createTimeStamp"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	ClientTranId    string `json:"clientTranId"`
}

func (s *SubAccUniversalTransferHistoryService) FromEmail(fromEmail string) *SubAccUniversalTransferHistoryService {
	s.fromEmail = &fromEmail
	return s
}

func (s *SubAccUniversalTransferHistoryService) ToEmail(toEmail string) *SubAccUniversalTransferHistoryService {
	s.toEmail = &toEmail
	return s
}

func (s *SubAccUniversalTransferHistoryService) ClientTranId(clientTranId string) *SubAccUniversalTransferHistoryService {
	s.clientTranId = &clientTranId
	return s
}

func (s *SubAccUniversalTransferHistoryService) StartTime(startTime int64) *SubAccUniversalTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *SubAccUniversalTransferHistoryService) EndTime(endTime int64) *SubAccUniversalTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *SubAccUniversalTransferHistoryService) Page(page int32) *SubAccUniversalTransferHistoryService {
	s.page = &page
	return s
}

func (s *SubAccUniversalTransferHistoryService) Limit(limit int32) *SubAccUniversalTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *SubAccUniversalTransferHistoryService) RecvWindow(recvWindow int64) *SubAccUniversalTransferHistoryService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccUniversalTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountUniversalTransferHistoryServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}
	if s.fromEmail != nil {
		r.setParam("fromEmail", *s.fromEmail)
	}
	if s.toEmail != nil {
		r.setParam("toEmail", *s.toEmail)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountUniversalTransferHistoryServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Binance Leveraged Tokens enable
type SubAccountBlvtEnableService struct {
	c          *Client
	email      string // sub-account email
	enableBlvt bool   // Only true for now
	recvWindow *int64
}

type SubAccountBlvtEnableServiceResponse struct {
	Email      string `json:"email"`
	EnableBlvt bool   `json:"enableBlvt"`
}

func (s *SubAccountBlvtEnableService) Email(email string) *SubAccountBlvtEnableService {
	s.email = email
	return s
}

func (s *SubAccountBlvtEnableService) EnableBlvt(enableBlvt bool) *SubAccountBlvtEnableService {
	s.enableBlvt = enableBlvt
	return s
}

func (s *SubAccountBlvtEnableService) RecvWindow(recvWindow int64) *SubAccountBlvtEnableService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountBlvtEnableService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountBlvtEnableServiceResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/blvt/enable",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	r.setFormParam("enableBlvt", s.enableBlvt)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountBlvtEnableServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// query sub-account api ip restriction
type SubAccountApiIpRestrictionService struct {
	c                *Client
	email            string // sub-account email
	subAccountApiKey string
	recvWindow       *int64
}

type SubAccountApiIpRestrictServiceResponse struct {
	IpRestrict string   `json:"ipRestrict"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}

func (s *SubAccountApiIpRestrictionService) Email(email string) *SubAccountApiIpRestrictionService {
	s.email = email
	return s
}

func (s *SubAccountApiIpRestrictionService) SubAccountApiKey(subAccountApiKey string) *SubAccountApiIpRestrictionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *SubAccountApiIpRestrictionService) RecvWindow(recvWindow int64) *SubAccountApiIpRestrictionService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountApiIpRestrictionService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountApiIpRestrictServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/subAccountApi/ipRestriction",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("subAccountApiKey", s.subAccountApiKey)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountApiIpRestrictServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// delete sub-account ip restriction
type SubAccountApiDeleteIpRestrictionService struct {
	c                *Client
	email            string // sub-account email
	subAccountApiKey string
	ipAddress        *string // Can be deleted in batches, separated by commas
	recvWindow       *int64
}

type SubAccountApiDeleteIpRestrictServiceResponse struct {
	IpRestrict string   `json:"ipRestrict"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}

func (s *SubAccountApiDeleteIpRestrictionService) Email(email string) *SubAccountApiDeleteIpRestrictionService {
	s.email = email
	return s
}

func (s *SubAccountApiDeleteIpRestrictionService) SubAccountApiKey(subAccountApiKey string) *SubAccountApiDeleteIpRestrictionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *SubAccountApiDeleteIpRestrictionService) IpAddress(ipAddress string) *SubAccountApiDeleteIpRestrictionService {
	s.ipAddress = &ipAddress
	return s
}

func (s *SubAccountApiDeleteIpRestrictionService) RecvWindow(recvWindow int64) *SubAccountApiDeleteIpRestrictionService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountApiDeleteIpRestrictionService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountApiDeleteIpRestrictServiceResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/sub-account/subAccountApi/ipRestriction/ipList",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	r.setFormParam("subAccountApiKey", s.subAccountApiKey)
	if s.ipAddress != nil {
		r.setFormParam("ipAddress", *s.ipAddress)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountApiDeleteIpRestrictServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// add sub-account ip restriction
type SubAccountApiAddIpRestrictionService struct {
	c                *Client
	email            string // sub-account email
	subAccountApiKey string
	status           string  // IP restriction status. 1 or not filled in (null)=IP not restricted. 2=Access is limited to trusted IPs only.
	ipAddress        *string // IP can be filled in bulk, separated by commas
	recvWindow       *int64
}

type SubAccountApiAddIpRestrictServiceResponse struct {
	IpRestrict string   `json:"ipRestrict"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}

func (s *SubAccountApiAddIpRestrictionService) Email(email string) *SubAccountApiAddIpRestrictionService {
	s.email = email
	return s
}

func (s *SubAccountApiAddIpRestrictionService) SubAccountApiKey(subAccountApiKey string) *SubAccountApiAddIpRestrictionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *SubAccountApiAddIpRestrictionService) Status(status string) *SubAccountApiAddIpRestrictionService {
	s.status = status
	return s
}

func (s *SubAccountApiAddIpRestrictionService) IpAddress(ipAddress string) *SubAccountApiAddIpRestrictionService {
	s.ipAddress = &ipAddress
	return s
}

func (s *SubAccountApiAddIpRestrictionService) RecvWindow(recvWindow int64) *SubAccountApiAddIpRestrictionService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountApiAddIpRestrictionService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountApiAddIpRestrictServiceResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/subAccountApi/ipRestriction",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	r.setFormParam("subAccountApiKey", s.subAccountApiKey)
	r.setFormParam("status", s.status)
	if s.ipAddress != nil {
		r.setFormParam("ipAddress", *s.ipAddress)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountApiAddIpRestrictServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ManagedSubAccountWithdrawService struct {
	c            *Client
	fromEmail    string
	asset        string
	amount       string // decimal format
	transferDate *int64 // The withdrawal will automatically occur on the selected date (UTC0). If no date is selected, the withdrawal will take effect immediately
	recvWindow   *int64
}

type ManagedSubAccountWithdrawServiceResponse struct {
	TranId int64 `json:"tranId"`
}

func (s *ManagedSubAccountWithdrawService) FromEmail(fromEmail string) *ManagedSubAccountWithdrawService {
	s.fromEmail = fromEmail
	return s
}

func (s *ManagedSubAccountWithdrawService) Asset(asset string) *ManagedSubAccountWithdrawService {
	s.asset = asset
	return s
}

func (s *ManagedSubAccountWithdrawService) Amount(amount string) *ManagedSubAccountWithdrawService {
	s.amount = amount
	return s
}

func (s *ManagedSubAccountWithdrawService) TransferDate(transferDate int64) *ManagedSubAccountWithdrawService {
	s.transferDate = &transferDate
	return s
}

func (s *ManagedSubAccountWithdrawService) RecvWindow(recvWindow int64) *ManagedSubAccountWithdrawService {
	s.recvWindow = &recvWindow
	return s
}

func (s *ManagedSubAccountWithdrawService) Do(ctx context.Context, opts ...RequestOption) (res *ManagedSubAccountWithdrawServiceResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/managed-subaccount/withdraw",
		secType:  secTypeSigned,
	}
	r.setFormParam("fromEmail", s.fromEmail)
	r.setFormParam("asset", s.asset)
	r.setFormParam("amount", s.amount)
	if s.transferDate != nil {
		r.setFormParam("transferDate", *s.transferDate)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountWithdrawServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query asset snapshot of managed-sub account
type ManagedSubAccountSnapshotService struct {
	c          *Client
	email      string
	accType    string // SPOT (spot), MARGIN (full position), FUTURES (U-based contract)
	startTime  *int64
	endTime    *int64
	limit      *int32 // min 7, max 30, default 7
	recvWindow *int64
}

type ManagedSubAccountSnapshotServiceResponse struct {
	Code        int64         `json:"code"`
	Msg         string        `json:"msg"`
	SnapshotVos []*SnapshotVo `json:"snapshotVos"`
}

type SnapshotVo struct {
	Data       SnapshotVoData `json:"data"`
	Type       string         `json:"type"`
	UpdateTime int64          `json:"updateTime"`
}

type SnapshotVoData struct {
	Balances            []*SnapShotSpotBalance `json:"balances"`            // set while SnapshotVo.type=spot
	TotalAssetOfBtc     string                 `json:"totalAssetOfBtc"`     // set while SnapshotVo.type is one of spot and margin
	MarginLevel         string                 `json:"marginLevel"`         // set while SnapshotVo.type=margin
	TotalLiabilityOfBtc string                 `json:"totalLiabilityOfBtc"` // set while SnapshotVo.type=margin
	TotalNetAssetOfBtc  string                 `json:"totalNetAssetOfBtc"`  // set while SnapshotVo.type=margin
	UserAssets          []*MarginUserAsset     `json:"userAssets"`          // set while SnapshotVo.type=margin
	Assets              []*FuturesUserAsset    `json:"assets"`              // set while SnapshotVo.type=futures
	Position            []*FuturesUserPosition `json:"position"`            // set while SnapshotVo.type=futures
}

type SnapShotSpotBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type MarginUserAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

type FuturesUserAsset struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	WalletBalance string `json:"walletBalance"`
}

type FuturesUserPosition struct {
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
}

func (s *ManagedSubAccountSnapshotService) Email(email string) *ManagedSubAccountSnapshotService {
	s.email = email
	return s
}

func (s *ManagedSubAccountSnapshotService) AccType(accType string) *ManagedSubAccountSnapshotService {
	s.accType = accType
	return s
}

func (s *ManagedSubAccountSnapshotService) StartTime(startTime int64) *ManagedSubAccountSnapshotService {
	s.startTime = &startTime
	return s
}

func (s *ManagedSubAccountSnapshotService) EndTime(endTime int64) *ManagedSubAccountSnapshotService {
	s.endTime = &endTime
	return s
}

func (s *ManagedSubAccountSnapshotService) Limit(limit int32) *ManagedSubAccountSnapshotService {
	s.limit = &limit
	return s
}

func (s *ManagedSubAccountSnapshotService) RecvWindow(recvWindow int64) *ManagedSubAccountSnapshotService {
	s.recvWindow = &recvWindow
	return s
}

func (s *ManagedSubAccountSnapshotService) Do(ctx context.Context, opts ...RequestOption) (res *ManagedSubAccountSnapshotServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/accountSnapshot",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("type", s.accType)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountSnapshotServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// managed-sub account query transfer log, this interface is for investor
type ManagedSubAccountQueryTransferLogForInvestorService struct {
	c                           *Client
	email                       string
	startTime                   int64
	endTime                     int64 // The start and end time intervals cannot exceed six months
	page                        int32
	limit                       int32   // max 500
	transfers                   *string // from/to
	transferFunctionAccountType *string // SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE
}

type ManagedSubAccountQueryTransferLogForInvestorServiceResponse struct {
	ManagerSubTransferHistoryVos []*ManagedSubTransferHistoryVo `json:"managerSubTransferHistoryVos"`
	Count                        int32                          `json:"count"`
}

type ManagedSubTransferHistoryVo struct {
	FromEmail       string `json:"fromEmail"`
	FromAccountType string `json:"fromAccountType"`
	ToEmail         string `json:"toEmail"`
	ToAccountType   string `json:"toAccountType"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	ScheduledData   int64  `json:"scheduledData"`
	CreateTime      int64  `json:"createTime"`
	Status          string `json:"status"`
	TranId          int64  `json:"tranId"`
}

func (s *ManagedSubAccountQueryTransferLogForInvestorService) Email(email string) *ManagedSubAccountQueryTransferLogForInvestorService {
	s.email = email
	return s
}

func (s *ManagedSubAccountQueryTransferLogForInvestorService) StartTime(startTime int64) *ManagedSubAccountQueryTransferLogForInvestorService {
	s.startTime = startTime
	return s
}

func (s *ManagedSubAccountQueryTransferLogForInvestorService) EndTime(endTime int64) *ManagedSubAccountQueryTransferLogForInvestorService {
	s.endTime = endTime
	return s
}

func (s *ManagedSubAccountQueryTransferLogForInvestorService) Page(page int32) *ManagedSubAccountQueryTransferLogForInvestorService {
	s.page = page
	return s
}

func (s *ManagedSubAccountQueryTransferLogForInvestorService) Limit(limit int32) *ManagedSubAccountQueryTransferLogForInvestorService {
	s.limit = limit
	return s
}

func (s *ManagedSubAccountQueryTransferLogForInvestorService) Transfers(transfers string) *ManagedSubAccountQueryTransferLogForInvestorService {
	s.transfers = &transfers
	return s
}

func (s *ManagedSubAccountQueryTransferLogForInvestorService) TransferFunctionAccountType(transferFunctionAccountType string) *ManagedSubAccountQueryTransferLogForInvestorService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *ManagedSubAccountQueryTransferLogForInvestorService) Do(ctx context.Context, opts ...RequestOption) (res *ManagedSubAccountQueryTransferLogForInvestorServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/queryTransLogForInvestor",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	r.setParam("page", s.page)
	r.setParam("limit", s.limit)
	if s.transfers != nil {
		r.setParam("transfers", *s.transfers)
	}
	if s.transferFunctionAccountType != nil {
		r.setParam("transferFunctionAccountType", *s.transferFunctionAccountType)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountQueryTransferLogForInvestorServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ManagedSubAccountQueryTransferLogForTradeParentService struct {
	c                           *Client
	email                       string
	startTime                   int64
	endTime                     int64 // The start and end time intervals cannot exceed six months
	page                        int32
	limit                       int32   // max 500
	transfers                   *string // from/to
	transferFunctionAccountType *string // SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE
}

type ManagedSubAccountQueryTransferLogForTradeParentServiceResponse struct {
	ManagerSubTransferHistoryVos []*ManagedSubTransferHistoryVo `json:"managerSubTransferHistoryVos"`
	Count                        int32                          `json:"count"`
}

func (s *ManagedSubAccountQueryTransferLogForTradeParentService) Email(email string) *ManagedSubAccountQueryTransferLogForTradeParentService {
	s.email = email
	return s
}

func (s *ManagedSubAccountQueryTransferLogForTradeParentService) StartTime(startTime int64) *ManagedSubAccountQueryTransferLogForTradeParentService {
	s.startTime = startTime
	return s
}

func (s *ManagedSubAccountQueryTransferLogForTradeParentService) EndTime(endTime int64) *ManagedSubAccountQueryTransferLogForTradeParentService {
	s.endTime = endTime
	return s
}

func (s *ManagedSubAccountQueryTransferLogForTradeParentService) Page(page int32) *ManagedSubAccountQueryTransferLogForTradeParentService {
	s.page = page
	return s
}

func (s *ManagedSubAccountQueryTransferLogForTradeParentService) Limit(limit int32) *ManagedSubAccountQueryTransferLogForTradeParentService {
	s.limit = limit
	return s
}

func (s *ManagedSubAccountQueryTransferLogForTradeParentService) Transfers(transfers string) *ManagedSubAccountQueryTransferLogForTradeParentService {
	s.transfers = &transfers
	return s
}

func (s *ManagedSubAccountQueryTransferLogForTradeParentService) TransferFunctionAccountType(transferFunctionAccountType string) *ManagedSubAccountQueryTransferLogForTradeParentService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *ManagedSubAccountQueryTransferLogForTradeParentService) Do(ctx context.Context, opts ...RequestOption) (res *ManagedSubAccountQueryTransferLogForTradeParentServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/queryTransLogForTradeParent",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	r.setParam("page", s.page)
	r.setParam("limit", s.limit)
	if s.transfers != nil {
		r.setParam("transfers", *s.transfers)
	}
	if s.transferFunctionAccountType != nil {
		r.setParam("transferFunctionAccountType", *s.transferFunctionAccountType)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountQueryTransferLogForTradeParentServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Investor account inquiry custody account futures assets
type ManagedSubAccountQueryFuturesAssetService struct {
	c     *Client
	email string
}

type ManagedSubAccountQueryFuturesAssetServiceResponse struct {
	Code        int32                             `json:"code"`
	Message     string                            `json:"message"`
	SnapshotVos []*ManagedSubFuturesAccountSnapVo `json:"snapshotVos"`
}

type ManagedSubFuturesAccountSnapVo struct {
	Type       string                              `json:"type"`
	UpdateTime int64                               `json:"updateTime"`
	Data       *ManagedSubFuturesAccountSnapVoData `json:"data"`
}

type ManagedSubFuturesAccountSnapVoData struct {
	Assets   []*ManagedSubFuturesAccountSnapVoDataAsset    `json:"assets"`
	Position []*ManagedSubFuturesAccountSnapVoDataPosition `json:"position"`
}

type ManagedSubFuturesAccountSnapVoDataAsset struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	WalletBalance string `json:"walletBalance"`
}

type ManagedSubFuturesAccountSnapVoDataPosition struct {
	Symbol      string `json:"symbol"`
	EntryPrice  string `json:"entryPrice"`
	MarkPrice   string `json:"markPrice"`
	PositionAmt string `json:"positionAmt"`
}

func (s *ManagedSubAccountQueryFuturesAssetService) Email(email string) *ManagedSubAccountQueryFuturesAssetService {
	s.email = email
	return s
}

func (s *ManagedSubAccountQueryFuturesAssetService) Do(ctx context.Context, opts ...RequestOption) (res *ManagedSubAccountQueryFuturesAssetServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/fetch-future-asset",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountQueryFuturesAssetServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Investor account inquiry for leveraged assets in custodial accounts
type ManagedSubAccountQueryMarginAssetService struct {
	c     *Client
	email string
}

type ManagedSubAccountQueryMarginAssetServiceResponse struct {
	MarginLevel         string                          `json:"marginLevel"`
	TotalAssetOfBtc     string                          `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string                          `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string                          `json:"totalNetAssetOfBtc"`
	UserAssets          []*ManagedSubAccountMarginAsset `json:"userAssets"`
}

type ManagedSubAccountMarginAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

func (s *ManagedSubAccountQueryMarginAssetService) Email(email string) *ManagedSubAccountQueryMarginAssetService {
	s.email = email
	return s
}

func (s *ManagedSubAccountQueryMarginAssetService) Do(ctx context.Context, opts ...RequestOption) (res *ManagedSubAccountQueryMarginAssetServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/marginAsset",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountQueryMarginAssetServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query sub account assets, v4 interface.
type SubAccountAssetService struct {
	c          *Client
	email      string
	recvWindow *int64
}

type SubAccountAssetServiceResponse struct {
	Balances []*SubAccountAssetBalance `json:"balances"`
}

type SubAccountAssetBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

func (s *SubAccountAssetService) Email(email string) *SubAccountAssetService {
	s.email = email
	return s
}

func (s *SubAccountAssetService) RecvWindow(recvWindow int64) *SubAccountAssetService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountAssetService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountAssetServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v4/sub-account/assets",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountAssetServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query the list of managed-accounts
type ManagedSubAccountInfoService struct {
	c          *Client
	email      string
	page       *int32 // default 1
	limit      *int32 // default 20, max 20
	recvWindow *int64
}

type ManagedSubAccountInfoServiceResponse struct {
	Total                    int32                          `json:"total"`
	ManagerSubUserInfoVoList []*ManagedSubAccountUserInfoVo `json:"managerSubUserInfoVoList"`
}

type ManagedSubAccountUserInfoVo struct {
	RootUserId               int64  `json:"rootUserId"`
	ManagersubUserId         int64  `json:"managersubUserId"`
	BindParentUserId         int64  `json:"bindParentUserId"`
	Email                    string `json:"email"`
	InsertTimeStamp          int64  `json:"insertTimeStamp"`
	BindParentEmail          string `json:"bindParentEmail"`
	IsSubUserEnabled         bool   `json:"isSubUserEnabled"`
	IsUserActive             bool   `json:"isUserActive"`
	IsMarginEnabled          bool   `json:"isMarginEnabled"`
	IsFutureEnabled          bool   `json:"isFutureEnabled"`
	IsSignedLVTRiskAgreement bool   `json:"isSignedLVTRiskAgreement"`
}

func (s *ManagedSubAccountInfoService) Email(email string) *ManagedSubAccountInfoService {
	s.email = email
	return s
}

func (s *ManagedSubAccountInfoService) Page(page int32) *ManagedSubAccountInfoService {
	s.page = &page
	return s
}

func (s *ManagedSubAccountInfoService) Limit(limit int32) *ManagedSubAccountInfoService {
	s.limit = &limit
	return s
}

func (s *ManagedSubAccountInfoService) RecvWindow(recvWindow int64) *ManagedSubAccountInfoService {
	s.recvWindow = &recvWindow
	return s
}

func (s *ManagedSubAccountInfoService) Do(ctx context.Context, opts ...RequestOption) (res *ManagedSubAccountInfoServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/info",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountInfoServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query sub account transaction volume statistics list
type SubAccountTransactionStatisticsService struct {
	c          *Client
	email      string
	recvWindow *int64
}

type SubAccountTransactionStatisticServiceResponse struct {
	Recent30BtcTotal         string         `json:"recent30BtcTotal"`
	Recent30BtcFuturesTotal  string         `json:"recent30BtcFuturesTotal"`
	Recent30BtcMarginTotal   string         `json:"recent30BtcMarginTotal"`
	Recent30BusdTotal        string         `json:"recent30BusdTotal"`
	Recent30BusdFuturesTotal string         `json:"recent30BusdFuturesTotal"`
	Recent30BusdMarginTotal  string         `json:"recent30BusdMarginTotal"`
	TradeInfoVos             []*TradeInfoVo `json:"tradeInfoVos"`
}

type TradeInfoVo struct {
	UserId      int64   `json:"userId"`
	Btc         float64 `json:"btc"`
	BtcFutures  float64 `json:"btcFutures"`
	BtcMargin   float64 `json:"btcMargin"`
	Busd        float64 `json:"busd"`
	BusdFutures float64 `json:"busdFutures"`
	BusdMargin  float64 `json:"busdMargin"`
	Date        int64   `json:"date"`
}

func (s *SubAccountTransactionStatisticsService) Email(email string) *SubAccountTransactionStatisticsService {
	s.email = email
	return s
}

func (s *SubAccountTransactionStatisticsService) RecvWindow(recvWindow int64) *SubAccountTransactionStatisticsService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountTransactionStatisticsService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountTransactionStatisticServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/transaction-statistics",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountTransactionStatisticServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Obtain the recharge address for the custody account
type ManagedSubAccountDepositAddressService struct {
	c          *Client
	email      string
	coin       string
	network    *string // The network can be obtained from GET /sapi/v1/capital/deposit/address. When the network is not transmitted, return the default network of the coin
	recvWindow *int64
}

type ManagedSubAccountDepositAddressServiceResponse struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

func (s *ManagedSubAccountDepositAddressService) Email(email string) *ManagedSubAccountDepositAddressService {
	s.email = email
	return s
}

func (s *ManagedSubAccountDepositAddressService) Coin(coin string) *ManagedSubAccountDepositAddressService {
	s.coin = coin
	return s
}

func (s *ManagedSubAccountDepositAddressService) Network(network string) *ManagedSubAccountDepositAddressService {
	s.network = &network
	return s
}

func (s *ManagedSubAccountDepositAddressService) RecvWindow(recvWindow int64) *ManagedSubAccountDepositAddressService {
	s.recvWindow = &recvWindow
	return s
}

func (s *ManagedSubAccountDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (res *ManagedSubAccountDepositAddressServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/deposit/address",
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("coin", s.coin)
	if s.network != nil {
		r.setParam("network", *s.network)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountDepositAddressServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type SubAccountOptionsEnableService struct {
	c          *Client
	email      string
	recvWindow *int64
}

type SubAccountOptionsEnableServiceResponse struct {
	Email             string `json:"email"`
	IsEOptionsEnabled bool   `json:"isEOptionsEnabled"`
}

func (s *SubAccountOptionsEnableService) Email(email string) *SubAccountOptionsEnableService {
	s.email = email
	return s
}

func (s *SubAccountOptionsEnableService) RecvWindow(recvWindow int64) *SubAccountOptionsEnableService {
	s.recvWindow = &recvWindow
	return s
}

func (s *SubAccountOptionsEnableService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountOptionsEnableServiceResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/eoptions/enable",
		secType:  secTypeSigned,
	}
	r.setFormParam("email", s.email)
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountOptionsEnableServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query transfer records of managed-sub accounts
type ManagedSubAccountQueryTransferLogService struct {
	c                           *Client
	startTime                   int64
	endTime                     int64 // End time (start time end time interval cannot exceed six months)
	page                        int32
	limit                       int32   // max 500
	transfers                   *string // from/to
	transferFunctionAccountType *string // SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE
	recvWindow                  *int64
}

type ManagedSubAccountQueryTransferLogServiceResponse struct {
	ManagerSubTransferHistoryVos []*ManagedSubTransferHistoryVo `json:"managerSubTransferHistoryVos"`
	Count                        int32                          `json:"count"`
}

func (s *ManagedSubAccountQueryTransferLogService) StartTime(startTime int64) *ManagedSubAccountQueryTransferLogService {
	s.startTime = startTime
	return s
}

func (s *ManagedSubAccountQueryTransferLogService) EndTime(endTime int64) *ManagedSubAccountQueryTransferLogService {
	s.endTime = endTime
	return s
}

func (s *ManagedSubAccountQueryTransferLogService) Page(page int32) *ManagedSubAccountQueryTransferLogService {
	s.page = page
	return s
}

func (s *ManagedSubAccountQueryTransferLogService) Limit(limit int32) *ManagedSubAccountQueryTransferLogService {
	s.limit = limit
	return s
}

func (s *ManagedSubAccountQueryTransferLogService) Transfers(transfers string) *ManagedSubAccountQueryTransferLogService {
	s.transfers = &transfers
	return s
}

func (s *ManagedSubAccountQueryTransferLogService) TransferFunctionAccountType(transferFunctionAccountType string) *ManagedSubAccountQueryTransferLogService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *ManagedSubAccountQueryTransferLogService) RecvWindow(recvWindow int64) *ManagedSubAccountQueryTransferLogService {
	s.recvWindow = &recvWindow
	return s
}

func (s *ManagedSubAccountQueryTransferLogService) Do(ctx context.Context, opts ...RequestOption) (res *ManagedSubAccountQueryTransferLogServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/managed-subaccount/query-trans-log",
		secType:  secTypeSigned,
	}
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	r.setParam("page", s.page)
	r.setParam("limit", s.limit)
	if s.transfers != nil {
		r.setParam("transfers", *s.transfers)
	}
	if s.transferFunctionAccountType != nil {
		r.setParam("transferFunctionAccountType", *s.transferFunctionAccountType)
	}
	if s.recvWindow != nil {
		r.recvWindow = *s.recvWindow
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountQueryTransferLogServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
