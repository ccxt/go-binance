package binance

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// MarginTransferService transfer between spot account and margin account
type MarginTransferService struct {
	c            *Client
	asset        string
	amount       string
	transferType int
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginTransferService) Asset(asset string) *MarginTransferService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginTransferService) Amount(amount string) *MarginTransferService {
	s.amount = amount
	return s
}

// Type 1: transfer from main account to margin account 2: transfer from margin account to main account
func (s *MarginTransferService) Type(transferType MarginTransferType) *MarginTransferService {
	s.transferType = int(transferType)
	return s
}

// Do send request
func (s *MarginTransferService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TransactionResponse define transaction response
type TransactionResponse struct {
	TranID int64 `json:"tranId"`
}

// MarginLoanService apply for a loan
// Deprecated: use MarginBorrowRepayService instead
type MarginLoanService struct {
	c          *Client
	asset      string
	amount     string
	isIsolated bool
	symbol     *string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginLoanService) Asset(asset string) *MarginLoanService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginLoanService) Amount(amount string) *MarginLoanService {
	s.amount = amount
	return s
}

// IsIsolated is for isolated margin or not, "TRUE", "FALSE"，default "FALSE"
func (s *MarginLoanService) IsIsolated(isIsolated bool) *MarginLoanService {
	s.isIsolated = isIsolated
	return s
}

// Symbol set isolated symbol
func (s *MarginLoanService) Symbol(symbol string) *MarginLoanService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginLoanService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/loan",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.setFormParams(m)
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginRepayService repay loan for margin account
// Deprecated: use MarginBorrowRepayService instead
type MarginRepayService struct {
	c          *Client
	asset      string
	amount     string
	isIsolated bool
	symbol     *string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginRepayService) Asset(asset string) *MarginRepayService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginRepayService) Amount(amount string) *MarginRepayService {
	s.amount = amount
	return s
}

// IsIsolated is for isolated margin or not, "TRUE", "FALSE"，default "FALSE"
func (s *MarginRepayService) IsIsolated(isIsolated bool) *MarginRepayService {
	s.isIsolated = isIsolated
	return s
}

// Symbol set isolated symbol
func (s *MarginRepayService) Symbol(symbol string) *MarginRepayService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginRepayService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/repay",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.setFormParams(m)
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginBorrowRepayService borrow/repay for margin account
type MarginBorrowRepayService struct {
	c          *Client
	asset      string                       // Mandatory:YES
	amount     string                       // Mandatory:YES
	isIsolated bool                         // Mandatory:YES, TRUE for Isolated Margin, FALSE for Cross Margin, Default FALSE
	symbol     string                       // Mandatory:YES, Only for Isolated margin
	_type      MarginAccountBorrowRepayType // Mandatory:YES, BORROW or REPAY
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginBorrowRepayService) Asset(asset string) *MarginBorrowRepayService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginBorrowRepayService) Amount(amount string) *MarginBorrowRepayService {
	s.amount = amount
	return s
}

// IsIsolated is for isolated margin or not, "TRUE", "FALSE"，default "FALSE"
func (s *MarginBorrowRepayService) IsIsolated(isIsolated bool) *MarginBorrowRepayService {
	s.isIsolated = isIsolated
	return s
}

// Symbol set isolated symbol
func (s *MarginBorrowRepayService) Symbol(symbol string) *MarginBorrowRepayService {
	s.symbol = symbol
	return s
}

func (s *MarginBorrowRepayService) Type(marginBorrowRepayType MarginAccountBorrowRepayType) *MarginBorrowRepayService {
	s._type = marginBorrowRepayType
	return s
}

// Do send request
func (s *MarginBorrowRepayService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/borrow-repay",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":      s.asset,
		"isIsolated": s.isIsolated,
		"symbol":     s.symbol,
		"amount":     s.amount,
		"type":       string(s._type),
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListMarginBorrowRepayService Query borrow/repay records in Margin account
// 1. txId or startTime must be sent. txId takes precedence. Response in descending order
// 2. If an asset is sent, data within 30 days before endTime; If an asset is not sent, data within 7 days before endTime
// 3. If neither startTime nor endTime is sent, the recent 7-day data will be returned.
// 4. startTime set as endTime - 7days by default, endTime set as current time by default
type ListMarginBorrowRepayService struct {
	c              *Client
	asset          *string                      // Mandatory:NO
	isolatedSymbol *string                      // Mandatory:NO, Symbol in Isolated Margin
	txId           *int64                       // Mandatory:NO, Transaction id
	startTime      *int64                       // Mandatory:NO, milliseconds
	endTime        *int64                       // Mandatory:NO, milliseconds
	current        *int64                       // Mandatory:NO, Current querying page. Start from 1. Default:1
	size           *int64                       // Mandatory:NO, Default:10 Max:100
	_type          MarginAccountBorrowRepayType // Mandatory:YES, BORROW or REPAY
}

// Asset set asset being transferred, e.g., BTC
func (s *ListMarginBorrowRepayService) Asset(asset string) *ListMarginBorrowRepayService {
	s.asset = &asset
	return s
}

func (s *ListMarginBorrowRepayService) IsolatedSymbol(isolatedSymbol string) *ListMarginBorrowRepayService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

func (s *ListMarginBorrowRepayService) TxId(txId int64) *ListMarginBorrowRepayService {
	s.txId = &txId
	return s
}

func (s *ListMarginBorrowRepayService) StartTime(startTime int64) *ListMarginBorrowRepayService {
	s.startTime = &startTime
	return s
}

func (s *ListMarginBorrowRepayService) EndTime(endTime int64) *ListMarginBorrowRepayService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginBorrowRepayService) Current(current int64) *ListMarginBorrowRepayService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginBorrowRepayService) Size(size int64) *ListMarginBorrowRepayService {
	s.size = &size
	return s
}

func (s *ListMarginBorrowRepayService) Type(marginBorrowRepayType MarginAccountBorrowRepayType) *ListMarginBorrowRepayService {
	s._type = marginBorrowRepayType
	return s
}

// Do send request
func (s *ListMarginBorrowRepayService) Do(ctx context.Context, opts ...RequestOption) (res *MarginBorrowRepayResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/borrow-repay",
		secType:  secTypeSigned,
	}
	m := params{
		//"asset":          s.asset,
		//"isolatedSymbol": s.isolatedSymbol,
		//"txId":           s.txId,
		//"startTime":      s.startTime,
		//"endTime":        s.endTime,
		//"current":        s.current,
		//"size":           s.size,
		"type": string(s._type),
	}
	r.setParams(m)
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.isolatedSymbol != nil {
		r.setParam("isolatedSymbol", *s.isolatedSymbol)
	}
	if s.txId != nil {
		r.setParam("txId", *s.txId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}

	res = new(MarginBorrowRepayResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type MarginBorrowRepayResponse struct {
	Rows  []MarginBorrowRepay `json:"rows"`
	Total int64               `json:"total"`
}

type MarginBorrowRepay struct {
	Type           string `json:"type"`           // AUTO,MANUAL for Cross Margin Borrow; MANUAL，AUTO，BNB_AUTO_REPAY，AUTO_BORROW_REPAY, POINT_AUTO_REPAY for Cross Margin Repay; AUTO，MANUAL for Isolated Margin Borrow/Repay;
	IsolatedSymbol string `json:"isolatedSymbol"` // "BNBUSDT"  isolated symbol, will not be returned for crossed margin
	Amount         string `json:"amount"`
	Asset          string `json:"asset"`
	Interest       string `json:"interest"`  // Interest repaid
	Principal      string `json:"principal"` // Principal repaid
	Status         string `json:"status"`    //one of PENDING (pending execution), CONFIRMED (successfully execution), FAILED (execution failed, nothing happened to your account);
	Timestamp      int64  `json:"timestamp"`
	TxID           int64  `json:"txId"`

	ClientTag string `json:"clientTag"` // This field is not in the document
}

// ListMarginLoansService list loan record
// Deprecated: use ListMarginBorrowRepayService instead
type ListMarginLoansService struct {
	c         *Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListMarginLoansService) Asset(asset string) *ListMarginLoansService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *ListMarginLoansService) TxID(txID int64) *ListMarginLoansService {
	s.txID = &txID
	return s
}

// StartTime set start time
func (s *ListMarginLoansService) StartTime(startTime int64) *ListMarginLoansService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *ListMarginLoansService) EndTime(endTime int64) *ListMarginLoansService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginLoansService) Current(current int64) *ListMarginLoansService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginLoansService) Size(size int64) *ListMarginLoansService {
	s.size = &size
	return s
}

// Do send request
func (s *ListMarginLoansService) Do(ctx context.Context, opts ...RequestOption) (res *MarginLoanResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/loan",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.txID != nil {
		r.setParam("txId", *s.txID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginLoanResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginLoanResponse define margin loan response
type MarginLoanResponse struct {
	Rows  []MarginLoan `json:"rows"`
	Total int64        `json:"total"`
}

// MarginLoan define margin loan
type MarginLoan struct {
	Asset     string               `json:"asset"`
	Principal string               `json:"principal"`
	Timestamp int64                `json:"timestamp"`
	Status    MarginLoanStatusType `json:"status"`
}

// ListMarginRepaysService list repay record
// Deprecated: use ListMarginBorrowRepayService instead
type ListMarginRepaysService struct {
	c         *Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListMarginRepaysService) Asset(asset string) *ListMarginRepaysService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *ListMarginRepaysService) TxID(txID int64) *ListMarginRepaysService {
	s.txID = &txID
	return s
}

// StartTime set start time
func (s *ListMarginRepaysService) StartTime(startTime int64) *ListMarginRepaysService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *ListMarginRepaysService) EndTime(endTime int64) *ListMarginRepaysService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginRepaysService) Current(current int64) *ListMarginRepaysService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginRepaysService) Size(size int64) *ListMarginRepaysService {
	s.size = &size
	return s
}

// Do send request
func (s *ListMarginRepaysService) Do(ctx context.Context, opts ...RequestOption) (res *MarginRepayResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/repay",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.txID != nil {
		r.setParam("txId", *s.txID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginRepayResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginRepayResponse define margin repay response
type MarginRepayResponse struct {
	Rows  []MarginRepay `json:"rows"`
	Total int64         `json:"total"`
}

// MarginRepay define margin repay
type MarginRepay struct {
	Asset     string                `json:"asset"`
	Amount    string                `json:"amount"`
	Interest  string                `json:"interest"`
	Principal string                `json:"principal"`
	Timestamp int64                 `json:"timestamp"`
	Status    MarginRepayStatusType `json:"status"`
	TxID      int64                 `json:"txId"`
}

// GetIsolatedMarginAccountService gets isolated margin account info
type GetIsolatedMarginAccountService struct {
	c *Client

	symbols []string
}

// Symbols set symbols to the isolated margin account
func (s *GetIsolatedMarginAccountService) Symbols(symbols ...string) *GetIsolatedMarginAccountService {
	s.symbols = symbols
	return s
}

// Do send request
func (s *GetIsolatedMarginAccountService) Do(ctx context.Context, opts ...RequestOption) (res *IsolatedMarginAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/isolated/account",
		secType:  secTypeSigned,
	}

	if len(s.symbols) > 0 {
		r.setParam("symbols", strings.Join(s.symbols, ","))
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(IsolatedMarginAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IsolatedMarginAccount defines isolated user assets of margin account
type IsolatedMarginAccount struct {
	TotalAssetOfBTC     string                `json:"totalAssetOfBtc"`
	TotalLiabilityOfBTC string                `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBTC  string                `json:"totalNetAssetOfBtc"`
	Assets              []IsolatedMarginAsset `json:"assets"`
}

// IsolatedMarginAsset defines isolated margin asset information, like margin level, liquidation price... etc
type IsolatedMarginAsset struct {
	Symbol     string            `json:"symbol"`
	QuoteAsset IsolatedUserAsset `json:"quoteAsset"`
	BaseAsset  IsolatedUserAsset `json:"baseAsset"`

	IsolatedCreated   bool   `json:"isolatedCreated"`
	Enabled           bool   `json:"enabled"`
	MarginLevel       string `json:"marginLevel"`
	MarginLevelStatus string `json:"marginLevelStatus"`
	MarginRatio       string `json:"marginRatio"`
	IndexPrice        string `json:"indexPrice"`
	LiquidatePrice    string `json:"liquidatePrice"`
	LiquidateRate     string `json:"liquidateRate"`
	TradeEnabled      bool   `json:"tradeEnabled"`
}

// IsolatedUserAsset defines isolated user assets of the margin account
type IsolatedUserAsset struct {
	Asset         string `json:"asset"`
	Borrowed      string `json:"borrowed"`
	Free          string `json:"free"`
	Interest      string `json:"interest"`
	Locked        string `json:"locked"`
	NetAsset      string `json:"netAsset"`
	NetAssetOfBtc string `json:"netAssetOfBtc"`

	BorrowEnabled bool   `json:"borrowEnabled"`
	RepayEnabled  bool   `json:"repayEnabled"`
	TotalAsset    string `json:"totalAsset"`
}

// GetMarginAccountService get margin account info
type GetMarginAccountService struct {
	c *Client
}

// Do send request
func (s *GetMarginAccountService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginAccount define margin account info
type MarginAccount struct {
	Created                    bool        `json:"created"`
	BorrowEnabled              bool        `json:"borrowEnabled"`
	MarginLevel                string      `json:"marginLevel"`
	CollateralMarginLevel      string      `json:"collateralMarginLevel"`
	TotalAssetOfBTC            string      `json:"totalAssetOfBtc"`
	TotalLiabilityOfBTC        string      `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBTC         string      `json:"totalNetAssetOfBtc"`
	TotalCollateralValueInUSDT string      `json:"totalCollateralValueInUSDT"`
	TradeEnabled               bool        `json:"tradeEnabled"`
	TransferInEnabled          bool        `json:"transferInEnabled"`
	TransferOutEnabled         bool        `json:"transferOutEnabled"`
	AccountType                string      `json:"accountType"`
	UserAssets                 []UserAsset `json:"userAssets"`
}

// UserAsset define user assets of margin account
type UserAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// GetMarginAssetService get margin asset info
type GetMarginAssetService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMarginAssetService) Asset(asset string) *GetMarginAssetService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMarginAssetService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAsset, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/asset",
		secType:  secTypeAPIKey,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginAsset)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginAsset define margin asset info
type MarginAsset struct {
	FullName      string `json:"assetFullName"`
	Name          string `json:"assetName"`
	Borrowable    bool   `json:"isBorrowable"`
	Mortgageable  bool   `json:"isMortgageable"`
	UserMinBorrow string `json:"userMinBorrow"`
	UserMinRepay  string `json:"userMinRepay"`
}

// GetMarginPairService get margin pair info
type GetMarginPairService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetMarginPairService) Symbol(symbol string) *GetMarginPairService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetMarginPairService) Do(ctx context.Context, opts ...RequestOption) (res *MarginPair, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/pair",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginPair)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginPair define margin pair info
type MarginPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// GetMarginAllPairsService get margin pair info
type GetMarginAllPairsService struct {
	c *Client
}

// Do send request
func (s *GetMarginAllPairsService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAllPair, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/allPairs",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAllPair{}, err
	}
	res = make([]*MarginAllPair, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAllPair{}, err
	}
	return res, nil
}

// MarginAllPair define margin pair info
type MarginAllPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// GetMarginPriceIndexService get margin price index
type GetMarginPriceIndexService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetMarginPriceIndexService) Symbol(symbol string) *GetMarginPriceIndexService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetMarginPriceIndexService) Do(ctx context.Context, opts ...RequestOption) (res *MarginPriceIndex, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/priceIndex",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginPriceIndex)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginPriceIndex define margin price index
type MarginPriceIndex struct {
	CalcTime int64  `json:"calcTime"`
	Price    string `json:"price"`
	Symbol   string `json:"symbol"`
}

// ListMarginTradesService list trades
type ListMarginTradesService struct {
	c          *Client
	symbol     string
	startTime  *int64
	endTime    *int64
	limit      *int
	fromID     *int64
	isIsolated bool
}

// Symbol set symbol
func (s *ListMarginTradesService) Symbol(symbol string) *ListMarginTradesService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *ListMarginTradesService) IsIsolated(isIsolated bool) *ListMarginTradesService {
	s.isIsolated = isIsolated
	return s
}

// StartTime set starttime
func (s *ListMarginTradesService) StartTime(startTime int64) *ListMarginTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListMarginTradesService) EndTime(endTime int64) *ListMarginTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListMarginTradesService) Limit(limit int) *ListMarginTradesService {
	s.limit = &limit
	return s
}

// FromID set fromID
func (s *ListMarginTradesService) FromID(fromID int64) *ListMarginTradesService {
	s.fromID = &fromID
	return s
}

// Do send request
func (s *ListMarginTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*TradeV3, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/myTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*TradeV3{}, err
	}
	res = make([]*TradeV3, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*TradeV3{}, err
	}
	return res, nil
}

// GetMaxBorrowableService get max borrowable of asset
// https://developers.binance.com/docs/margin_trading/borrow-and-repay/Query-Max-Borrow
type GetMaxBorrowableService struct {
	c              *Client
	asset          string
	isolatedSymbol string
}

// Asset set asset
func (s *GetMaxBorrowableService) Asset(asset string) *GetMaxBorrowableService {
	s.asset = asset
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *GetMaxBorrowableService) IsolatedSymbol(isolatedSymbol string) *GetMaxBorrowableService {
	s.isolatedSymbol = isolatedSymbol
	return s
}

// Do send request
func (s *GetMaxBorrowableService) Do(ctx context.Context, opts ...RequestOption) (res *MaxBorrowable, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/maxBorrowable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.isolatedSymbol != "" {
		r.setParam("isolatedSymbol", s.isolatedSymbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxBorrowable)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MaxBorrowable define max borrowable response
type MaxBorrowable struct {
	Amount      string `json:"amount"`      // account's currently max borrowable amount with sufficient system availability
	BorrowLimit string `json:"borrowLimit"` // max borrowable amount limited by the account level
}

// GetMaxTransferableService get max transferable of asset
type GetMaxTransferableService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMaxTransferableService) Asset(asset string) *GetMaxTransferableService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMaxTransferableService) Do(ctx context.Context, opts ...RequestOption) (res *MaxTransferable, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/maxTransferable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxTransferable)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MaxTransferable define max transferable response
type MaxTransferable struct {
	Amount string `json:"amount"`
}

// StartIsolatedMarginUserStreamService create listen key for margin user stream service
type StartIsolatedMarginUserStreamService struct {
	c      *Client
	symbol string
}

// Symbol sets the user stream to isolated margin user stream
func (s *StartIsolatedMarginUserStreamService) Symbol(symbol string) *StartIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *StartIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/userDataStream/isolated",
		secType:  secTypeAPIKey,
	}

	r.setFormParam("symbol", s.symbol)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveIsolatedMarginUserStreamService updates listen key for isolated margin user data stream
type KeepaliveIsolatedMarginUserStreamService struct {
	c         *Client
	listenKey string
	symbol    string
}

// Symbol set symbol to the isolated margin keepalive request
func (s *KeepaliveIsolatedMarginUserStreamService) Symbol(symbol string) *KeepaliveIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// ListenKey set listen key
func (s *KeepaliveIsolatedMarginUserStreamService) ListenKey(listenKey string) *KeepaliveIsolatedMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/sapi/v1/userDataStream/isolated",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	r.setFormParam("symbol", s.symbol)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseIsolatedMarginUserStreamService delete listen key
type CloseIsolatedMarginUserStreamService struct {
	c         *Client
	listenKey string

	symbol string
}

// ListenKey set listen key
func (s *CloseIsolatedMarginUserStreamService) ListenKey(listenKey string) *CloseIsolatedMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Symbol set symbol to the isolated margin user stream close request
func (s *CloseIsolatedMarginUserStreamService) Symbol(symbol string) *CloseIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CloseIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/userDataStream/isolated",
		secType:  secTypeAPIKey,
	}

	r.setFormParam("listenKey", s.listenKey)
	r.setFormParam("symbol", s.symbol)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// StartMarginUserStreamService create listen key for margin user stream service
type StartMarginUserStreamService struct {
	c *Client
}

// Do send request
func (s *StartMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveMarginUserStreamService update listen key
type KeepaliveMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveMarginUserStreamService) ListenKey(listenKey string) *KeepaliveMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseMarginUserStreamService delete listen key
type CloseMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseMarginUserStreamService) ListenKey(listenKey string) *CloseMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}

	r.setFormParam("listenKey", s.listenKey)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// GetAllMarginAssetsService get margin pair info
type GetAllMarginAssetsService struct {
	c *Client
}

// Do send request
func (s *GetAllMarginAssetsService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAsset, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/allAssets",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAsset{}, err
	}
	res = make([]*MarginAsset, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAsset{}, err
	}
	return res, nil
}

// GetIsolatedMarginAllPairsService get isolated margin pair info
type GetIsolatedMarginAllPairsService struct {
	c *Client
}

// Do send request
func (s *GetIsolatedMarginAllPairsService) Do(ctx context.Context, opts ...RequestOption) (res []*IsolatedMarginAllPair, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/isolated/allPairs",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*IsolatedMarginAllPair{}, err
	}
	res = make([]*IsolatedMarginAllPair, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*IsolatedMarginAllPair{}, err
	}
	return res, nil
}

// IsolatedMarginAllPair define isolated margin pair info
type IsolatedMarginAllPair struct {
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// IsolatedMarginTransferService transfer assets between spot and isolated margin.
type IsolatedMarginTransferService struct {
	c *Client

	asset     string
	symbol    string
	transFrom AccountType
	transTo   AccountType
	amount    string
}

func (s *IsolatedMarginTransferService) Symbol(symbol string) *IsolatedMarginTransferService {
	s.symbol = symbol
	return s
}

func (s *IsolatedMarginTransferService) Asset(asset string) *IsolatedMarginTransferService {
	s.asset = asset
	return s
}

// TransFrom supports account types: "SPOT", "ISOLATED_MARGIN"
func (s *IsolatedMarginTransferService) TransFrom(transFrom AccountType) *IsolatedMarginTransferService {
	s.transFrom = transFrom
	return s
}

// TransTo supports account types: "SPOT", "ISOLATED_MARGIN"
func (s *IsolatedMarginTransferService) TransTo(transTo AccountType) *IsolatedMarginTransferService {
	s.transTo = transTo
	return s
}

func (s *IsolatedMarginTransferService) Amount(amount string) *IsolatedMarginTransferService {
	s.amount = amount
	return s
}

// Do send request
func (s *IsolatedMarginTransferService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/isolated/transfer",
		secType:  secTypeSigned,
	}
	r.setFormParam("asset", s.asset)
	r.setFormParam("symbol", s.symbol)
	r.setFormParam("transFrom", s.transFrom)
	r.setFormParam("transTo", s.transTo)
	r.setFormParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(TransactionResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
