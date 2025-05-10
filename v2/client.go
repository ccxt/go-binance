package binance

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/delivery"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/adshao/go-binance/v2/options"
)

// SideType define side type of order
type SideType string

// OrderType define order type
type OrderType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// NewOrderRespType define response JSON verbosity
type NewOrderRespType string

// OrderStatusType define order status type
type OrderStatusType string

// SymbolType define symbol type
type SymbolType string

// SymbolStatusType define symbol status type
type SymbolStatusType string

// SymbolFilterType define symbol filter type
type SymbolFilterType string

// UserDataEventType define spot user data event type
type UserDataEventType string

// MarginTransferType define margin transfer type
type MarginTransferType int

// MarginLoanStatusType define margin loan status type
type MarginLoanStatusType string

// MarginRepayStatusType define margin repay status type
type MarginRepayStatusType string

// FuturesTransferStatusType define futures transfer status type
type FuturesTransferStatusType string

// SideEffectType define side effect type for orders
type SideEffectType string

// FuturesTransferType define futures transfer type
type FuturesTransferType int

// TransactionType define transaction type
type TransactionType string

// LendingType define the type of lending (flexible saving, activity, ...)
type LendingType string

// StakingProduct define the staking product (locked staking, flexible defi staking, locked defi staking, ...)
type StakingProduct string

// StakingTransactionType define the staking transaction type (subscription, redemption, interest)
type StakingTransactionType string

// LiquidityOperationType define the type of adding/removing liquidity to a liquidity pool(COMBINATION, SINGLE)
type LiquidityOperationType string

// SwappingStatus define the status of swap when querying the swap history
type SwappingStatus int

// LiquidityRewardType define the type of reward we'd claim
type LiquidityRewardType int

// RewardClaimStatus define the status of claiming a reward
type RewardClaimStatus int

// RateLimitType define the rate limitation types
// see https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md#enum-definitions
type RateLimitType string

// RateLimitInterval define the rate limitation intervals
type RateLimitInterval string

// AccountType define the account types
type AccountType string

// SubAccountTransferType define the sub account transfer types
type SubAccountTransferType int

// UserUniversalTransferType define the user universal transfer types
type UserUniversalTransferType string

// UserUniversalTransferStatus define the user universal transfer status
type UserUniversalTransferStatusType string

// FuturesOrderBookHistoryDataType define the futures order book history data types
type FuturesOrderBookHistoryDataType string

// FutureAlgoType define future algo types
type FuturesAlgoType string

// FutureAlgoUrgencyType define future algo urgency type
type FuturesAlgoUrgencyType string

// FutureAlgoOrderStatusType define future algo order status
type FuturesAlgoOrderStatusType string

// Endpoints
var (
	BaseAPIMainURL    = "https://api.binance.com"
	BaseAPIMainUSURL  = "https://api.binance.us"
	BaseAPITestnetURL = "https://testnet.binance.vision"
)

// SelfTradePreventionMode define self trade prevention strategy
type SelfTradePreventionMode string

type MarginAccountBorrowRepayType string

// UseTestnet switch all the API endpoints from production to the testnet
var UseTestnet = false

// UseUSDomain switch all the API endpoints from production to the binance.us
var UseUSDomain = false

// Global enums
const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	OrderTypeLimit           OrderType = "LIMIT"
	OrderTypeMarket          OrderType = "MARKET"
	OrderTypeLimitMaker      OrderType = "LIMIT_MAKER"
	OrderTypeStopLoss        OrderType = "STOP_LOSS"
	OrderTypeStopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	OrderTypeTakeProfit      OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"

	TimeInForceTypeGTC TimeInForceType = "GTC"
	TimeInForceTypeIOC TimeInForceType = "IOC"
	TimeInForceTypeFOK TimeInForceType = "FOK"

	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeRESULT NewOrderRespType = "RESULT"
	NewOrderRespTypeFULL   NewOrderRespType = "FULL"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"
	OrderStatusTypePendingCancel   OrderStatusType = "PENDING_CANCEL"
	OrderStatusTypeRejected        OrderStatusType = "REJECTED"
	OrderStatusTypeExpired         OrderStatusType = "EXPIRED"
	OrderStatusExpiredInMatch      OrderStatusType = "EXPIRED_IN_MATCH" // STP Expired

	SymbolTypeSpot SymbolType = "SPOT"

	SymbolStatusTypePreTrading   SymbolStatusType = "PRE_TRADING"
	SymbolStatusTypeTrading      SymbolStatusType = "TRADING"
	SymbolStatusTypePostTrading  SymbolStatusType = "POST_TRADING"
	SymbolStatusTypeEndOfDay     SymbolStatusType = "END_OF_DAY"
	SymbolStatusTypeHalt         SymbolStatusType = "HALT"
	SymbolStatusTypeAuctionMatch SymbolStatusType = "AUCTION_MATCH"
	SymbolStatusTypeBreak        SymbolStatusType = "BREAK"

	SymbolFilterTypeLotSize            SymbolFilterType = "LOT_SIZE"
	SymbolFilterTypePriceFilter        SymbolFilterType = "PRICE_FILTER"
	SymbolFilterTypePercentPriceBySide SymbolFilterType = "PERCENT_PRICE_BY_SIDE"
	SymbolFilterTypeMinNotional        SymbolFilterType = "MIN_NOTIONAL"
	SymbolFilterTypeNotional           SymbolFilterType = "NOTIONAL"
	SymbolFilterTypeIcebergParts       SymbolFilterType = "ICEBERG_PARTS"
	SymbolFilterTypeMarketLotSize      SymbolFilterType = "MARKET_LOT_SIZE"
	SymbolFilterTypeMaxNumOrders       SymbolFilterType = "MAX_NUM_ORDERS"
	SymbolFilterTypeMaxNumAlgoOrders   SymbolFilterType = "MAX_NUM_ALGO_ORDERS"
	SymbolFilterTypeTrailingDelta      SymbolFilterType = "TRAILING_DELTA"

	UserDataEventTypeOutboundAccountPosition UserDataEventType = "outboundAccountPosition"
	UserDataEventTypeBalanceUpdate           UserDataEventType = "balanceUpdate"
	UserDataEventTypeExecutionReport         UserDataEventType = "executionReport"
	UserDataEventTypeListStatus              UserDataEventType = "ListStatus"

	MarginTransferTypeToMargin MarginTransferType = 1
	MarginTransferTypeToMain   MarginTransferType = 2

	FuturesTransferTypeToFutures       FuturesTransferType = 1
	FuturesTransferTypeToMain          FuturesTransferType = 2
	FuturesTransferTypeToFuturesCM     FuturesTransferType = 3
	FuturesTransferTypeFuturesCMToMain FuturesTransferType = 4

	MarginLoanStatusTypePending   MarginLoanStatusType = "PENDING"
	MarginLoanStatusTypeConfirmed MarginLoanStatusType = "CONFIRMED"
	MarginLoanStatusTypeFailed    MarginLoanStatusType = "FAILED"

	MarginRepayStatusTypePending   MarginRepayStatusType = "PENDING"
	MarginRepayStatusTypeConfirmed MarginRepayStatusType = "CONFIRMED"
	MarginRepayStatusTypeFailed    MarginRepayStatusType = "FAILED"

	FuturesTransferStatusTypePending   FuturesTransferStatusType = "PENDING"
	FuturesTransferStatusTypeConfirmed FuturesTransferStatusType = "CONFIRMED"
	FuturesTransferStatusTypeFailed    FuturesTransferStatusType = "FAILED"

	SideEffectTypeNoSideEffect    SideEffectType = "NO_SIDE_EFFECT"
	SideEffectTypeMarginBuy       SideEffectType = "MARGIN_BUY"
	SideEffectTypeAutoRepay       SideEffectType = "AUTO_REPAY"
	SideEffectTypeAutoBorrowRepay SideEffectType = "AUTO_BORROW_REPAY"

	TransactionTypeDeposit  TransactionType = "0"
	TransactionTypeWithdraw TransactionType = "1"
	TransactionTypeBuy      TransactionType = "0"
	TransactionTypeSell     TransactionType = "1"

	LendingTypeFlexible LendingType = "DAILY"
	LendingTypeFixed    LendingType = "CUSTOMIZED_FIXED"
	LendingTypeActivity LendingType = "ACTIVITY"

	LiquidityOperationTypeCombination LiquidityOperationType = "COMBINATION"
	LiquidityOperationTypeSingle      LiquidityOperationType = "SINGLE"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"

	StakingProductLockedStaking       = "STAKING"
	StakingProductFlexibleDeFiStaking = "F_DEFI"
	StakingProductLockedDeFiStaking   = "L_DEFI"

	StakingTransactionTypeSubscription = "SUBSCRIPTION"
	StakingTransactionTypeRedemption   = "REDEMPTION"
	StakingTransactionTypeInterest     = "INTEREST"

	SwappingStatusPending SwappingStatus = 0
	SwappingStatusDone    SwappingStatus = 1
	SwappingStatusFailed  SwappingStatus = 2

	RewardTypeTrading   LiquidityRewardType = 0
	RewardTypeLiquidity LiquidityRewardType = 1

	RewardClaimPending RewardClaimStatus = 0
	RewardClaimDone    RewardClaimStatus = 1

	RateLimitTypeRequestWeight RateLimitType = "REQUEST_WEIGHT"
	RateLimitTypeOrders        RateLimitType = "ORDERS"
	RateLimitTypeRawRequests   RateLimitType = "RAW_REQUESTS"

	RateLimitIntervalSecond RateLimitInterval = "SECOND"
	RateLimitIntervalMinute RateLimitInterval = "MINUTE"
	RateLimitIntervalDay    RateLimitInterval = "DAY"

	AccountTypeSpot           AccountType = "SPOT"
	AccountTypeMargin         AccountType = "MARGIN"
	AccountTypeIsolatedMargin AccountType = "ISOLATED_MARGIN"
	AccountTypeUSDTFuture     AccountType = "USDT_FUTURE"
	AccountTypeCoinFuture     AccountType = "COIN_FUTURE"

	SubAccountTransferTypeTransferIn  SubAccountTransferType = 1
	SubAccountTransferTypeTransferOut SubAccountTransferType = 2

	UserUniversalTransferTypeMainToUmFutures                UserUniversalTransferType = "MAIN_UMFUTURE"
	UserUniversalTransferTypeMainToCmFutures                UserUniversalTransferType = "MAIN_CMFUTURE"
	UserUniversalTransferTypeMainToMargin                   UserUniversalTransferType = "MAIN_MARGIN"
	UserUniversalTransferTypeUmFuturesToMain                UserUniversalTransferType = "UMFUTURE_MAIN"
	UserUniversalTransferTypeUmFuturesToMargin              UserUniversalTransferType = "UMFUTURE_MARGIN"
	UserUniversalTransferTypeCmFuturesToMain                UserUniversalTransferType = "CMFUTURE_MAIN"
	UserUniversalTransferTypeMarginToMain                   UserUniversalTransferType = "MARGIN_MAIN"
	UserUniversalTransferTypeMarginToUmFutures              UserUniversalTransferType = "MARGIN_UMFUTURE"
	UserUniversalTransferTypeMarginToCmFutures              UserUniversalTransferType = "MARGIN_CMFUTURE"
	UserUniversalTransferTypeCmFuturesToMargin              UserUniversalTransferType = "CMFUTURE_MARGIN"
	UserUniversalTransferTypeIsolatedMarginToMargin         UserUniversalTransferType = "ISOLATEDMARGIN_MARGIN"
	UserUniversalTransferTypeMarginToIsolatedMargin         UserUniversalTransferType = "MARGIN_ISOLATEDMARGIN"
	UserUniversalTransferTypeIsolatedMarginToIsolatedMargin UserUniversalTransferType = "ISOLATEDMARGIN_ISOLATEDMARGIN"
	UserUniversalTransferTypeMainToFunding                  UserUniversalTransferType = "MAIN_FUNDING"
	UserUniversalTransferTypeFundingToMain                  UserUniversalTransferType = "FUNDING_MAIN"
	UserUniversalTransferTypeFundingToUmFutures             UserUniversalTransferType = "FUNDING_UMFUTURE"
	UserUniversalTransferTypeUmFuturesToFunding             UserUniversalTransferType = "UMFUTURE_FUNDING"
	UserUniversalTransferTypeMarginToFunding                UserUniversalTransferType = "MARGIN_FUNDING"
	UserUniversalTransferTypeFundingToMargin                UserUniversalTransferType = "FUNDING_MARGIN"
	UserUniversalTransferTypeFundingToCmFutures             UserUniversalTransferType = "FUNDING_CMFUTURE"
	UserUniversalTransferTypeCmFuturesToFunding             UserUniversalTransferType = "CMFUTURE_FUNDING"
	UserUniversalTransferTypeMainToOption                   UserUniversalTransferType = "MAIN_OPTION"
	UserUniversalTransferTypeOptionToMain                   UserUniversalTransferType = "OPTION_MAIN"
	UserUniversalTransferTypeUmFuturesToOption              UserUniversalTransferType = "UMFUTURE_OPTION"
	UserUniversalTransferTypeOptionToUmFutures              UserUniversalTransferType = "OPTION_UMFUTURE"
	UserUniversalTransferTypeMarginToOption                 UserUniversalTransferType = "MARGIN_OPTION"
	UserUniversalTransferTypeOptionToMargin                 UserUniversalTransferType = "OPTION_MARGIN"
	UserUniversalTransferTypeFundingToOption                UserUniversalTransferType = "FUNDING_OPTION"
	UserUniversalTransferTypeOptionToFunding                UserUniversalTransferType = "OPTION_FUNDING"
	UserUniversalTransferTypeMainToPortfolioMargin          UserUniversalTransferType = "MAIN_PORTFOLIO_MARGIN"
	UserUniversalTransferTypePortfolioMarginToMain          UserUniversalTransferType = "PORTFOLIO_MARGIN_MAIN"
	UserUniversalTransferTypeMainToIsolatedMargin           UserUniversalTransferType = "MAIN_ISOLATED_MARGIN"
	UserUniversalTransferTypeIsolatedMarginToMain           UserUniversalTransferType = "ISOLATED_MARGIN_MAIN"

	UserUniversalTransferStatusTypePending   UserUniversalTransferStatusType = "PENDING"
	UserUniversalTransferStatusTypeConfirmed UserUniversalTransferStatusType = "CONFIRMED"
	UserUniversalTransferStatusTypeFailed    UserUniversalTransferStatusType = "FAILED"

	FuturesOrderBookHistoryDataTypeTDepth FuturesOrderBookHistoryDataType = "T_DEPTH"
	FuturesOrderBookHistoryDataTypeSDepth FuturesOrderBookHistoryDataType = "S_DEPTH"

	FuturesAlgoTypeVp   FuturesAlgoType = "VP"
	FuturesAlgoTypeTwap FuturesAlgoType = "TWAP"

	FuturesAlgoUrgencyTypeLow    FuturesAlgoUrgencyType = "LOW"
	FuturesAlgoUrgencyTypeMedium FuturesAlgoUrgencyType = "MEDIUM"
	FuturesAlgoUrgencyTypeHigh   FuturesAlgoUrgencyType = "HIGH"

	FuturesAlgoOrderStatusTypeWorking   FuturesAlgoOrderStatusType = "WORKING"
	FuturesAlgoOrderStatusTypeFinished  FuturesAlgoOrderStatusType = "FINISHED"
	FuturesAlgoOrderStatusTypeCancelled FuturesAlgoOrderStatusType = "CANCELLED"

	SelfTradePreventionModeNone        SelfTradePreventionMode = "NONE"
	SelfTradePreventionModeExpireTaker SelfTradePreventionMode = "EXPIRE_TAKER"
	SelfTradePreventionModeExpireBoth  SelfTradePreventionMode = "EXPIRE_BOTH"
	SelfTradePreventionModeExpireMaker SelfTradePreventionMode = "EXPIRE_MAKER"

	MarginAccountBorrow MarginAccountBorrowRepayType = "BORROW"
	MarginAccountRepay  MarginAccountBorrowRepayType = "REPAY"

	MarginAccountBorrowRepayStatusPending   string = "PENDING"
	MarginAccountBorrowRepayStatusConfirmed string = "CONFIRMED"
	MarginAccountBorrowRepayStatusFailed    string = "FAILED"
)

func currentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// getAPIEndpoint return the base endpoint of the Rest API according the UseTestnet flag
func getAPIEndpoint() string {
	if UseTestnet {
		return BaseAPITestnetURL
	}
	if UseUSDomain {
		return BaseAPIMainUSURL
	}
	return BaseAPIMainURL
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    common.KeyTypeHmac,
		BaseURL:    getAPIEndpoint(),
		UserAgent:  "Binance/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance-golang ", log.LstdFlags),
	}
}

// NewProxiedClient passing a proxy url
func NewProxiedClient(apiKey, secretKey, proxyUrl string) *Client {
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		log.Fatal(err)
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Client{
		APIKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
		BaseURL:   getAPIEndpoint(),
		UserAgent: "Binance/golang",
		HTTPClient: &http.Client{
			Transport: tr,
		},
		Logger: log.New(os.Stderr, "Binance-golang ", log.LstdFlags),
	}
}

// NewFuturesClient initialize client for futures API
func NewFuturesClient(apiKey, secretKey string) *futures.Client {
	return futures.NewClient(apiKey, secretKey)
}

// NewDeliveryClient initialize client for coin-M futures API
func NewDeliveryClient(apiKey, secretKey string) *delivery.Client {
	return delivery.NewClient(apiKey, secretKey)
}

// NewOptionsClient initialize client for options API
func NewOptionsClient(apiKey, secretKey string) *options.Client {
	return options.NewClient(apiKey, secretKey)
}

type doFunc func(req *http.Request) (*http.Response, error)

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	KeyType    string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc

	UsedWeight UsedWeight
	OrderCount OrderCount
}

type UsedWeight struct {
	Used   int64
	Used1M int64 // used in last 1 minute
}

type OrderCount struct {
	Count10s int64
	Count1d  int64
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	if r.recvWindow > 0 {
		r.setParam(recvWindowKey, r.recvWindow)
	}
	if r.secType == secTypeSigned {
		r.setParam(timestampKey, currentTimestamp()-c.TimeOffset)
	}
	queryString := r.query.Encode()
	// @ is a safe character and does not require escape, So replace it back.
	queryString = strings.ReplaceAll(queryString, "%40", "@")
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	// @ is a safe character and does not require escape, So replace it back.
	bodyString = strings.ReplaceAll(bodyString, "%40", "@")
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set("X-MBX-APIKEY", c.APIKey)
	}
	kt := c.KeyType
	if kt == "" {
		kt = common.KeyTypeHmac
	}
	sf, err := common.SignFunc(kt)
	if err != nil {
		return err
	}
	if r.secType == secTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		sign, err := sf(c.SecretKey, raw)
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set(signatureKey, *sign)
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s\n", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v\n", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}

	usedWeight := res.Header.Get("X-Mbx-Used-Weight")
	if usedWeight != "" {
		if used, err := strconv.ParseInt(usedWeight, 10, 64); err == nil {
			c.UsedWeight.Used = used
		}
	}
	usedWeight1M := res.Header.Get("X-Mbx-Used-Weight-1m")
	if usedWeight1M != "" {
		if used, err := strconv.ParseInt(usedWeight1M, 10, 64); err == nil {
			c.UsedWeight.Used1M = used
		}
	}

	orderCount10s := res.Header.Get("X-Mbx-Order-Count-10s")
	if orderCount10s != "" {
		if count, err := strconv.ParseInt(orderCount10s, 10, 64); err == nil {
			c.OrderCount.Count10s = count
		}
	}

	orderCount1d := res.Header.Get("X-Mbx-Order-Count-1d")
	if orderCount1d != "" {
		if count, err := strconv.ParseInt(orderCount1d, 10, 64); err == nil {
			c.OrderCount.Count1d = count
		}
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the returned error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v\n", res)
	c.debug("response body: %s\n", string(data))
	c.debug("response status code: %d\n", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(common.APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s\n", e)
		}
		if !apiErr.IsValid() {
			apiErr.Response = data
		}
		return nil, apiErr
	}
	return data, nil
}

// SetApiEndpoint set api Endpoint
func (c *Client) SetApiEndpoint(url string) *Client {
	c.BaseURL = url
	return c
}

// NewPingService init ping service
func (c *Client) NewPingService() *PingService {
	return &PingService{c: c}
}

// NewServerTimeService init server time service
func (c *Client) NewServerTimeService() *ServerTimeService {
	return &ServerTimeService{c: c}
}

// NewSetServerTimeService init set server time service
func (c *Client) NewSetServerTimeService() *SetServerTimeService {
	return &SetServerTimeService{c: c}
}

// NewDepthService init depth service
func (c *Client) NewDepthService() *DepthService {
	return &DepthService{c: c}
}

// NewAggTradesService init aggregate trades service
func (c *Client) NewAggTradesService() *AggTradesService {
	return &AggTradesService{c: c}
}

// NewRecentTradesService init recent trades service
func (c *Client) NewRecentTradesService() *RecentTradesService {
	return &RecentTradesService{c: c}
}

// NewKlinesService init klines service
func (c *Client) NewKlinesService() *KlinesService {
	return &KlinesService{c: c}
}

func (c *Client) NewUiKlinesService() *UiKlinesService {
	return &UiKlinesService{c: c}
}

// NewListPriceChangeStatsService init list prices change stats service
func (c *Client) NewListPriceChangeStatsService() *ListPriceChangeStatsService {
	return &ListPriceChangeStatsService{c: c}
}

// NewListPricesService init listing prices service
func (c *Client) NewListPricesService() *ListPricesService {
	return &ListPricesService{c: c}
}

func (c *Client) NewTradingDayTickerService() *TradingDayTickerService {
	return &TradingDayTickerService{c: c}
}

// NewListBookTickersService init listing booking tickers service
func (c *Client) NewListBookTickersService() *ListBookTickersService {
	return &ListBookTickersService{c: c}
}

// NewListSymbolTickerService init listing symbols tickers
func (c *Client) NewListSymbolTickerService() *ListSymbolTickerService {
	return &ListSymbolTickerService{c: c}
}

// NewCreateOrderService init creating order service
func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{c: c}
}

// NewCreateOCOService init creating OCO service
func (c *Client) NewCreateOCOService() *CreateOCOService {
	return &CreateOCOService{c: c}
}

// NewCancelOCOService init cancel OCO service
func (c *Client) NewCancelOCOService() *CancelOCOService {
	return &CancelOCOService{c: c}
}

// NewGetOrderService init get order service
func (c *Client) NewGetOrderService() *GetOrderService {
	return &GetOrderService{c: c}
}

// NewCancelOrderService init cancel order service
func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{c: c}
}

// NewCancelOpenOrdersService init cancel open orders service
func (c *Client) NewCancelOpenOrdersService() *CancelOpenOrdersService {
	return &CancelOpenOrdersService{c: c}
}

// NewListOpenOrdersService init list open orders service
func (c *Client) NewListOpenOrdersService() *ListOpenOrdersService {
	return &ListOpenOrdersService{c: c}
}

// NewListOpenOcoService init list open oco service
func (c *Client) NewListOpenOcoService() *ListOpenOcoService {
	return &ListOpenOcoService{c: c}
}

// NewListOrdersService init listing orders service
func (c *Client) NewListOrdersService() *ListOrdersService {
	return &ListOrdersService{c: c}
}

// NewGetAccountService init getting account service
func (c *Client) NewGetAccountService() *GetAccountService {
	return &GetAccountService{c: c}
}

// NewGetAPIKeyPermission init getting API key permission
func (c *Client) NewGetAPIKeyPermission() *GetAPIKeyPermission {
	return &GetAPIKeyPermission{c: c}
}

// NewSavingFlexibleProductPositionsService get flexible products positions (Savings)
func (c *Client) NewSavingFlexibleProductPositionsService() *SavingFlexibleProductPositionsService {
	return &SavingFlexibleProductPositionsService{c: c}
}

// NewSavingFixedProjectPositionsService get fixed project positions (Savings)
func (c *Client) NewSavingFixedProjectPositionsService() *SavingFixedProjectPositionsService {
	return &SavingFixedProjectPositionsService{c: c}
}

// NewListSavingsFlexibleProductsService get flexible products list (Savings)
func (c *Client) NewListSavingsFlexibleProductsService() *ListSavingsFlexibleProductsService {
	return &ListSavingsFlexibleProductsService{c: c}
}

// NewPurchaseSavingsFlexibleProductService purchase a flexible product (Savings)
func (c *Client) NewPurchaseSavingsFlexibleProductService() *PurchaseSavingsFlexibleProductService {
	return &PurchaseSavingsFlexibleProductService{c: c}
}

// NewRedeemSavingsFlexibleProductService redeem a flexible product (Savings)
func (c *Client) NewRedeemSavingsFlexibleProductService() *RedeemSavingsFlexibleProductService {
	return &RedeemSavingsFlexibleProductService{c: c}
}

// NewListSavingsFixedAndActivityProductsService get fixed and activity product list (Savings)
func (c *Client) NewListSavingsFixedAndActivityProductsService() *ListSavingsFixedAndActivityProductsService {
	return &ListSavingsFixedAndActivityProductsService{c: c}
}

// NewGetAccountSnapshotService init getting account snapshot service
func (c *Client) NewGetAccountSnapshotService() *GetAccountSnapshotService {
	return &GetAccountSnapshotService{c: c}
}

// NewListTradesService init listing trades service
func (c *Client) NewListTradesService() *ListTradesService {
	return &ListTradesService{c: c}
}

// NewHistoricalTradesService init listing trades service
func (c *Client) NewHistoricalTradesService() *HistoricalTradesService {
	return &HistoricalTradesService{c: c}
}

// NewListDepositsService init listing deposits service
func (c *Client) NewListDepositsService() *ListDepositsService {
	return &ListDepositsService{c: c}
}

// NewGetDepositAddressService init getting deposit address service
func (c *Client) NewGetDepositAddressService() *GetDepositsAddressService {
	return &GetDepositsAddressService{c: c}
}

// NewCreateWithdrawService init creating withdraw service
func (c *Client) NewCreateWithdrawService() *CreateWithdrawService {
	return &CreateWithdrawService{c: c}
}

// NewListWithdrawsService init listing withdraw service
func (c *Client) NewListWithdrawsService() *ListWithdrawsService {
	return &ListWithdrawsService{c: c}
}

// NewStartUserStreamService init starting user stream service
func (c *Client) NewStartUserStreamService() *StartUserStreamService {
	return &StartUserStreamService{c: c}
}

// NewKeepaliveUserStreamService init keep alive user stream service
func (c *Client) NewKeepaliveUserStreamService() *KeepaliveUserStreamService {
	return &KeepaliveUserStreamService{c: c}
}

// NewCloseUserStreamService init closing user stream service
func (c *Client) NewCloseUserStreamService() *CloseUserStreamService {
	return &CloseUserStreamService{c: c}
}

// NewExchangeInfoService init exchange info service
func (c *Client) NewExchangeInfoService() *ExchangeInfoService {
	return &ExchangeInfoService{c: c}
}

// NewRateLimitService init rate limit service
func (c *Client) NewRateLimitService() *RateLimitService {
	return &RateLimitService{c: c}
}

// NewGetAssetDetailService init get asset detail service
func (c *Client) NewGetAssetDetailService() *GetAssetDetailService {
	return &GetAssetDetailService{c: c}
}

func (c *Client) NewWalletBalanceService() *WalletBalanceService {
	return &WalletBalanceService{c: c}
}

// NewAveragePriceService init average price service
func (c *Client) NewAveragePriceService() *AveragePriceService {
	return &AveragePriceService{c: c}
}

// NewMarginTransferService init margin account transfer service
func (c *Client) NewMarginTransferService() *MarginTransferService {
	return &MarginTransferService{c: c}
}

// NewMarginLoanService init margin account loan service
// Deprecated: use NewMarginBorrowRepayService instead
func (c *Client) NewMarginLoanService() *MarginLoanService {
	return &MarginLoanService{c: c}
}

// NewMarginRepayService init margin account repay service
// Deprecated: use NewMarginBorrowRepayService instead
func (c *Client) NewMarginRepayService() *MarginRepayService {
	return &MarginRepayService{c: c}
}

// NewMarginBorrowRepayService init margin account borrow/repay service
func (c *Client) NewMarginBorrowRepayService() *MarginBorrowRepayService {
	return &MarginBorrowRepayService{c: c}
}

func (c *Client) NewListMarginBorrowRepayService() *ListMarginBorrowRepayService {
	return &ListMarginBorrowRepayService{c: c}
}

// NewCreateMarginOrderService init creating margin order service
func (c *Client) NewCreateMarginOrderService() *CreateMarginOrderService {
	return &CreateMarginOrderService{c: c}
}

// NewCancelMarginOrderService init cancel order service
func (c *Client) NewCancelMarginOrderService() *CancelMarginOrderService {
	return &CancelMarginOrderService{c: c}
}

// NewCreateMarginOCOService init creating margin order service
func (c *Client) NewCreateMarginOCOService() *CreateMarginOCOService {
	return &CreateMarginOCOService{c: c}
}

// NewCancelMarginOCOService init cancel order service
func (c *Client) NewCancelMarginOCOService() *CancelMarginOCOService {
	return &CancelMarginOCOService{c: c}
}

// NewGetMarginOrderService init get order service
func (c *Client) NewGetMarginOrderService() *GetMarginOrderService {
	return &GetMarginOrderService{c: c}
}

// NewListMarginLoansService init list margin loan service
// Deprecated: use NewListMarginBorrowRepayService instead
func (c *Client) NewListMarginLoansService() *ListMarginLoansService {
	return &ListMarginLoansService{c: c}
}

// NewListMarginRepaysService init list margin repay service
// Deprecated: use NewListMarginBorrowRepayService instead
func (c *Client) NewListMarginRepaysService() *ListMarginRepaysService {
	return &ListMarginRepaysService{c: c}
}

// NewGetMarginAccountService init get margin account service
func (c *Client) NewGetMarginAccountService() *GetMarginAccountService {
	return &GetMarginAccountService{c: c}
}

// NewGetIsolatedMarginAccountService init get isolated margin asset service
func (c *Client) NewGetIsolatedMarginAccountService() *GetIsolatedMarginAccountService {
	return &GetIsolatedMarginAccountService{c: c}
}

func (c *Client) NewIsolatedMarginTransferService() *IsolatedMarginTransferService {
	return &IsolatedMarginTransferService{c: c}
}

// NewGetMarginAssetService init get margin asset service
func (c *Client) NewGetMarginAssetService() *GetMarginAssetService {
	return &GetMarginAssetService{c: c}
}

// NewGetMarginPairService init get margin pair service
func (c *Client) NewGetMarginPairService() *GetMarginPairService {
	return &GetMarginPairService{c: c}
}

// NewGetMarginAllPairsService init get margin all pairs service
func (c *Client) NewGetMarginAllPairsService() *GetMarginAllPairsService {
	return &GetMarginAllPairsService{c: c}
}

// NewGetMarginPriceIndexService init get margin price index service
func (c *Client) NewGetMarginPriceIndexService() *GetMarginPriceIndexService {
	return &GetMarginPriceIndexService{c: c}
}

// NewListMarginOpenOrdersService init list margin open orders service
func (c *Client) NewListMarginOpenOrdersService() *ListMarginOpenOrdersService {
	return &ListMarginOpenOrdersService{c: c}
}

// NewListMarginOrdersService init list margin all orders service
func (c *Client) NewListMarginOrdersService() *ListMarginOrdersService {
	return &ListMarginOrdersService{c: c}
}

// NewListMarginTradesService init list margin trades service
func (c *Client) NewListMarginTradesService() *ListMarginTradesService {
	return &ListMarginTradesService{c: c}
}

// NewGetMaxBorrowableService init get max borrowable service
func (c *Client) NewGetMaxBorrowableService() *GetMaxBorrowableService {
	return &GetMaxBorrowableService{c: c}
}

// NewGetMaxTransferableService init get max transferable service
func (c *Client) NewGetMaxTransferableService() *GetMaxTransferableService {
	return &GetMaxTransferableService{c: c}
}

// NewStartMarginUserStreamService init starting margin user stream service
func (c *Client) NewStartMarginUserStreamService() *StartMarginUserStreamService {
	return &StartMarginUserStreamService{c: c}
}

// NewKeepaliveMarginUserStreamService init keep alive margin user stream service
func (c *Client) NewKeepaliveMarginUserStreamService() *KeepaliveMarginUserStreamService {
	return &KeepaliveMarginUserStreamService{c: c}
}

// NewCloseMarginUserStreamService init closing margin user stream service
func (c *Client) NewCloseMarginUserStreamService() *CloseMarginUserStreamService {
	return &CloseMarginUserStreamService{c: c}
}

// NewStartIsolatedMarginUserStreamService init starting margin user stream service
func (c *Client) NewStartIsolatedMarginUserStreamService() *StartIsolatedMarginUserStreamService {
	return &StartIsolatedMarginUserStreamService{c: c}
}

// NewKeepaliveIsolatedMarginUserStreamService init keep alive margin user stream service
func (c *Client) NewKeepaliveIsolatedMarginUserStreamService() *KeepaliveIsolatedMarginUserStreamService {
	return &KeepaliveIsolatedMarginUserStreamService{c: c}
}

// NewCloseIsolatedMarginUserStreamService init closing margin user stream service
func (c *Client) NewCloseIsolatedMarginUserStreamService() *CloseIsolatedMarginUserStreamService {
	return &CloseIsolatedMarginUserStreamService{c: c}
}

// NewMarginInterestHistoryService init margin interest history service
func (c *Client) NewMarginInterestHistoryService() *MarginInterestHistoryService {
	return &MarginInterestHistoryService{c: c}
}

// NewMarginInterestRateHistoryService init margin interest rate history service
func (c *Client) NewMarginInterestRateHistoryService() *MarginInterestRateHistoryService {
	return &MarginInterestRateHistoryService{c: c}
}

// NewMarginNextHourlyInterestRateService init margin next hourly interest rate service
func (c *Client) NewMarginNextHourlyInterestRateService() *MarginNextHourlyInterestRateService {
	return &MarginNextHourlyInterestRateService{c: c}
}

// NewFuturesTransferService init futures transfer service
func (c *Client) NewFuturesTransferService() *FuturesTransferService {
	return &FuturesTransferService{c: c}
}

// NewListFuturesTransferService init list futures transfer service
func (c *Client) NewListFuturesTransferService() *ListFuturesTransferService {
	return &ListFuturesTransferService{c: c}
}

// NewListDustLogService init list dust log service
func (c *Client) NewListDustLogService() *ListDustLogService {
	return &ListDustLogService{c: c}
}

// NewDustTransferService init dust transfer service
func (c *Client) NewDustTransferService() *DustTransferService {
	return &DustTransferService{c: c}
}

// NewListDustService init dust list service
func (c *Client) NewListDustService() *ListDustService {
	return &ListDustService{c: c}
}

// NewTransferToSubAccountService transfer to subaccount service
func (c *Client) NewTransferToSubAccountService() *TransferToSubAccountService {
	return &TransferToSubAccountService{c: c}
}

// NewSubaccountAssetsService init list subaccount assets
func (c *Client) NewSubaccountAssetsService() *SubaccountAssetsService {
	return &SubaccountAssetsService{c: c}
}

// NewSubaccountSpotSummaryService init subaccount spot summary
func (c *Client) NewSubaccountSpotSummaryService() *SubaccountSpotSummaryService {
	return &SubaccountSpotSummaryService{c: c}
}

// NewSubaccountDepositAddressService init subaccount deposit address service
func (c *Client) NewSubaccountDepositAddressService() *SubaccountDepositAddressService {
	return &SubaccountDepositAddressService{c: c}
}

// NewAssetDividendService init the asset dividend list service
func (c *Client) NewAssetDividendService() *AssetDividendService {
	return &AssetDividendService{c: c}
}

// NewUserUniversalTransferService
func (c *Client) NewUserUniversalTransferService() *CreateUserUniversalTransferService {
	return &CreateUserUniversalTransferService{c: c}
}

// NewAllCoinsInformation
func (c *Client) NewGetAllCoinsInfoService() *GetAllCoinsInfoService {
	return &GetAllCoinsInfoService{c: c}
}

// NewDustTransferService init Get All Margin Assets service
func (c *Client) NewGetAllMarginAssetsService() *GetAllMarginAssetsService {
	return &GetAllMarginAssetsService{c: c}
}

// NewFiatDepositWithdrawHistoryService init the fiat deposit/withdraw history service
func (c *Client) NewFiatDepositWithdrawHistoryService() *FiatDepositWithdrawHistoryService {
	return &FiatDepositWithdrawHistoryService{c: c}
}

// NewFiatPaymentsHistoryService init the fiat payments history service
func (c *Client) NewFiatPaymentsHistoryService() *FiatPaymentsHistoryService {
	return &FiatPaymentsHistoryService{c: c}
}

// NewPayTransactionService init the pay transaction service
func (c *Client) NewPayTradeHistoryService() *PayTradeHistoryService {
	return &PayTradeHistoryService{c: c}
}

// NewFiatPaymentsHistoryService init the spot rebate history service
func (c *Client) NewSpotRebateHistoryService() *SpotRebateHistoryService {
	return &SpotRebateHistoryService{c: c}
}

// NewConvertTradeHistoryService init the convert trade history service
func (c *Client) NewConvertTradeHistoryService() *ConvertTradeHistoryService {
	return &ConvertTradeHistoryService{c: c}
}

// NewConvertExchangeInfoService init the convert exchange info service
func (c *Client) NewConvertExchangeInfoService() *ConvertExchangeInfoService {
	return &ConvertExchangeInfoService{c: c}
}

// NewConvertAssetInfoService init the convert asset info service
func (c *Client) NewConvertAssetInfoService() *ConvertAssetInfoService {
	return &ConvertAssetInfoService{c: c}
}

// NewConvertQuoteService init the convert quote service
func (c *Client) NewConvertQuoteService() *ConvertGetQuoteService {
	return &ConvertGetQuoteService{c: c}
}

// NewConvertAcceptQuoteService init the convert accept quote service
func (c *Client) NewConvertAcceptQuoteService() *ConvertAcceptQuoteService {
	return &ConvertAcceptQuoteService{c: c}
}

// NewConvertOrderStatusService init the convert order status service
func (c *Client) NewConvertOrderStatusService() *ConvertOrderStatusService {
	return &ConvertOrderStatusService{c: c}
}

// NewGetIsolatedMarginAllPairsService init get isolated margin all pairs service
func (c *Client) NewGetIsolatedMarginAllPairsService() *GetIsolatedMarginAllPairsService {
	return &GetIsolatedMarginAllPairsService{c: c}
}

// NewInterestHistoryService init the interest history service
func (c *Client) NewInterestHistoryService() *InterestHistoryService {
	return &InterestHistoryService{c: c}
}

// NewTradeFeeService init the trade fee service
func (c *Client) NewTradeFeeService() *TradeFeeService {
	return &TradeFeeService{c: c}
}

// NewC2CTradeHistoryService init the c2c trade history service
func (c *Client) NewC2CTradeHistoryService() *C2CTradeHistoryService {
	return &C2CTradeHistoryService{c: c}
}

// NewStakingProductPositionService init the staking product position service
func (c *Client) NewStakingProductPositionService() *StakingProductPositionService {
	return &StakingProductPositionService{c: c}
}

// NewStakingHistoryService init the staking history service
func (c *Client) NewStakingHistoryService() *StakingHistoryService {
	return &StakingHistoryService{c: c}
}

// NewGetAllLiquidityPoolService init the get all swap pool service
func (c *Client) NewGetAllLiquidityPoolService() *GetAllLiquidityPoolService {
	return &GetAllLiquidityPoolService{c: c}
}

// NewGetLiquidityPoolDetailService init the get liquidity pool detail service
func (c *Client) NewGetLiquidityPoolDetailService() *GetLiquidityPoolDetailService {
	return &GetLiquidityPoolDetailService{c: c}
}

// NewAddLiquidityPreviewService init the add liquidity preview service
func (c *Client) NewAddLiquidityPreviewService() *AddLiquidityPreviewService {
	return &AddLiquidityPreviewService{c: c}
}

// NewGetSwapQuoteService init the add liquidity preview service
func (c *Client) NewGetSwapQuoteService() *GetSwapQuoteService {
	return &GetSwapQuoteService{c: c}
}

// NewSwapService init the swap service
func (c *Client) NewSwapService() *SwapService {
	return &SwapService{c: c}
}

// NewAddLiquidityService init the add liquidity service
func (c *Client) NewAddLiquidityService() *AddLiquidityService {
	return &AddLiquidityService{c: c}
}

// NewGetUserSwapRecordsService init the service for listing the swap records
func (c *Client) NewGetUserSwapRecordsService() *GetUserSwapRecordsService {
	return &GetUserSwapRecordsService{c: c}
}

// NewClaimRewardService init the service for liquidity pool rewarding
func (c *Client) NewClaimRewardService() *ClaimRewardService {
	return &ClaimRewardService{c: c}
}

// NewRemoveLiquidityService init the service to remove liquidity
func (c *Client) NewRemoveLiquidityService() *RemoveLiquidityService {
	return &RemoveLiquidityService{c: c, assets: []string{}}
}

// NewQueryClaimedRewardHistoryService init the service to query reward claiming history
func (c *Client) NewQueryClaimedRewardHistoryService() *QueryClaimedRewardHistoryService {
	return &QueryClaimedRewardHistoryService{c: c}
}

// NewGetBNBBurnService init the service to get BNB Burn on spot trade and margin interest
func (c *Client) NewGetBNBBurnService() *GetBNBBurnService {
	return &GetBNBBurnService{c: c}
}

// NewToggleBNBBurnService init the service to toggle BNB Burn on spot trade and margin interest
func (c *Client) NewToggleBNBBurnService() *ToggleBNBBurnService {
	return &ToggleBNBBurnService{c: c}
}

// NewInternalUniversalTransferService Universal Transfer (For Master Account)
func (c *Client) NewInternalUniversalTransferService() *InternalUniversalTransferService {
	return &InternalUniversalTransferService{c: c}
}

// NewInternalUniversalTransferHistoryService Query Universal Transfer History (For Master Account)
func (c *Client) NewInternalUniversalTransferHistoryService() *InternalUniversalTransferHistoryService {
	return &InternalUniversalTransferHistoryService{c: c}
}

// NewSubAccountListService Query Sub-account List (For Master Account)
func (c *Client) NewSubAccountListService() *SubAccountListService {
	return &SubAccountListService{c: c}
}

// NewGetUserAsset Get user assets, just for positive data
func (c *Client) NewGetUserAsset() *GetUserAssetService {
	return &GetUserAssetService{c: c}
}

// NewManagedSubAccountDepositService Deposit Assets Into The Managed Sub-account（For Investor Master Account）
func (c *Client) NewManagedSubAccountDepositService() *ManagedSubAccountDepositService {
	return &ManagedSubAccountDepositService{c: c}
}

// NewManagedSubAccountWithdrawalService Withdrawal Assets From The Managed Sub-account（For Investor Master Account）
func (c *Client) NewManagedSubAccountWithdrawalService() *ManagedSubAccountWithdrawalService {
	return &ManagedSubAccountWithdrawalService{c: c}
}

// NewManagedSubAccountAssetsService Withdrawal Assets From The Managed Sub-account（For Investor Master Account）
func (c *Client) NewManagedSubAccountAssetsService() *ManagedSubAccountAssetsService {
	return &ManagedSubAccountAssetsService{c: c}
}

// NewSubAccountFuturesAccountService Get Detail on Sub-account's Futures Account (For Master Account)
func (c *Client) NewSubAccountFuturesAccountService() *SubAccountFuturesAccountService {
	return &SubAccountFuturesAccountService{c: c}
}

// NewSubAccountFuturesSummaryV1Service Get Summary of Sub-account's Futures Account (For Master Account)
func (c *Client) NewSubAccountFuturesSummaryV1Service() *SubAccountFuturesSummaryV1Service {
	return &SubAccountFuturesSummaryV1Service{c: c}
}

// NewSubAccountFuturesTransferV1Service Futures Transfer for Sub-account (For Master Account)
func (c *Client) NewSubAccountFuturesTransferV1Service() *SubAccountFuturesTransferV1Service {
	return &SubAccountFuturesTransferV1Service{c: c}
}

// NewSubAccountTransferHistoryService Transfer History for Sub-account (For Sub-account)
func (c *Client) NewSubAccountTransferHistoryService() *SubAccountTransferHistoryService {
	return &SubAccountTransferHistoryService{c: c}
}

// NewListUserUniversalTransferService Query User Universal Transfer History
func (c *Client) NewListUserUniversalTransferService() *ListUserUniversalTransferService {
	return &ListUserUniversalTransferService{c: c}
}

// Create virtual sub-account
func (c *Client) NewCreateVirtualSubAccountService() *CreateVirtualSubAccountService {
	return &CreateVirtualSubAccountService{c: c}
}

// Sub-Account spot transfer history
func (c *Client) NewSubAccountSpotTransferHistoryService() *SubAccountSpotTransferHistoryService {
	return &SubAccountSpotTransferHistoryService{c: c}
}

// Sub-Account futures transfer history
func (c *Client) NewSubAccountFuturesTransferHistoryService() *SubAccountFuturesTransferHistoryService {
	return &SubAccountFuturesTransferHistoryService{c: c}
}

// Get sub account deposit record
func (c *Client) NewSubAccountDepositRecordService() *SubAccountDepositRecordService {
	return &SubAccountDepositRecordService{c: c}
}

// Get sub account margin futures status
func (c *Client) NewSubAccountMarginFuturesStatusService() *SubAccountMarginFuturesStatusService {
	return &SubAccountMarginFuturesStatusService{c: c}
}

// sub account margin enable
func (c *Client) NewSubAccountMarginEnableService() *SubAccountMarginEnableService {
	return &SubAccountMarginEnableService{c: c}
}

// get sub-account margin account detail
func (c *Client) NewSubAccountMarginAccountInfoService() *SubAccountMarginAccountInfoService {
	return &SubAccountMarginAccountInfoService{c: c}
}

// get sub-account margin account summary
func (c *Client) NewSubAccountMarginAccountSummaryService() *SubAccountMarginAccountSummaryService {
	return &SubAccountMarginAccountSummaryService{c: c}
}

func (c *Client) NewSubAccountFuturesEnableService() *SubAccountFuturesEnableService {
	return &SubAccountFuturesEnableService{c: c}
}

// get sub-account futures account summary, include U-M and C-M, v2 interface
func (c *Client) NewSubAccountFuturesAccountSummaryService() *SubAccountFuturesAccountSummaryService {
	return &SubAccountFuturesAccountSummaryService{c: c}
}

// get target sub-account futures position information, include U-M and C-M, v2 interface.
func (c *Client) NewSubAccountFuturesPositionsService() *SubAccountFuturesPositionsService {
	return &SubAccountFuturesPositionsService{c: c}
}

// execute sub-account margin account transfer
func (c *Client) NewSubAccountMarginTransferService() *SubAccountMarginTransferService {
	return &SubAccountMarginTransferService{c: c}
}

// sub-account transfer balance to master-account
func (c *Client) NewSubAccountTransferSubToMasterService() *SubAccountTransferSubToMasterService {
	return &SubAccountTransferSubToMasterService{c: c}
}

// Universal transfer of master and sub accounts
func (c *Client) NewSubAccountUniversalTransferService() *SubAccountUniversalTransferService {
	return &SubAccountUniversalTransferService{c: c}
}

// Query the universal transfer history of sub and master accounts
func (c *Client) NewSubAccUniversalTransferHistoryService() *SubAccUniversalTransferHistoryService {
	return &SubAccUniversalTransferHistoryService{c: c}
}

// Binance Leveraged Tokens enable
func (c *Client) NewSubAccountBlvtEnableService() *SubAccountBlvtEnableService {
	return &SubAccountBlvtEnableService{c: c}
}

// query sub-account api ip restriction
func (c *Client) NewSubAccountApiIpRestrictionService() *SubAccountApiIpRestrictionService {
	return &SubAccountApiIpRestrictionService{c: c}
}

// delete sub-account ip restriction
func (c *Client) NewSubAccountApiDeleteIpRestrictionService() *SubAccountApiDeleteIpRestrictionService {
	return &SubAccountApiDeleteIpRestrictionService{c: c}
}

// add sub-account ip restriction
func (c *Client) NewSubAccountApiAddIpRestrictionService() *SubAccountApiAddIpRestrictionService {
	return &SubAccountApiAddIpRestrictionService{c: c}
}

func (c *Client) NewManagedSubAccountWithdrawService() *ManagedSubAccountWithdrawService {
	return &ManagedSubAccountWithdrawService{c: c}
}

// Query asset snapshot of managed-sub account
func (c *Client) NewManagedSubAccountSnapshotService() *ManagedSubAccountSnapshotService {
	return &ManagedSubAccountSnapshotService{c: c}
}

// managed-sub account query transfer log, this interface is for investor
func (c *Client) NewManagedSubAccountQueryTransferLogForInvestorService() *ManagedSubAccountQueryTransferLogForInvestorService {
	return &ManagedSubAccountQueryTransferLogForInvestorService{c: c}
}

func (c *Client) NewManagedSubAccountQueryTransferLogForTradeParentService() *ManagedSubAccountQueryTransferLogForTradeParentService {
	return &ManagedSubAccountQueryTransferLogForTradeParentService{c: c}
}

// Investor account inquiry custody account futures assets
func (c *Client) NewManagedSubAccountQueryFuturesAssetService() *ManagedSubAccountQueryFuturesAssetService {
	return &ManagedSubAccountQueryFuturesAssetService{c: c}
}

// Investor account inquiry for leveraged assets in custodial accounts
func (c *Client) NewManagedSubAccountQueryMarginAssetService() *ManagedSubAccountQueryMarginAssetService {
	return &ManagedSubAccountQueryMarginAssetService{c: c}
}

// Query sub account assets, v4 interface.
func (c *Client) NewSubAccountAssetService() *SubAccountAssetService {
	return &SubAccountAssetService{c: c}
}

// Query the list of managed-accounts
func (c *Client) NewManagedSubAccountInfoService() *ManagedSubAccountInfoService {
	return &ManagedSubAccountInfoService{c: c}
}

// Obtain the recharge address for the custody account
func (c *Client) NewManagedSubAccountDepositAddressService() *ManagedSubAccountDepositAddressService {
	return &ManagedSubAccountDepositAddressService{c: c}
}

func (c *Client) NewSubAccountOptionsEnableService() *SubAccountOptionsEnableService {
	return &SubAccountOptionsEnableService{c: c}
}

// Query transfer records of managed-sub accounts
func (c *Client) NewManagedSubAccountQueryTransferLogService() *ManagedSubAccountQueryTransferLogService {
	return &ManagedSubAccountQueryTransferLogService{c: c}
}

// Execute sub account futures balance transfer
func (c *Client) NewSubAccountFuturesInternalTransferService() *SubAccountFuturesInternalTransferService {
	return &SubAccountFuturesInternalTransferService{c: c}
}

// Query sub account transaction volume statistics list
func (c *Client) NewSubAccountTransactionStatisticsService() *SubAccountTransactionStatisticsService {
	return &SubAccountTransactionStatisticsService{c: c}
}

// get the target sub-account futures account detail, v2 interface.
func (c *Client) NewSubAccountFuturesAccountV2Service() *SubAccountFuturesAccountV2Service {
	return &SubAccountFuturesAccountV2Service{c: c}
}

// Futures order book history service
func (c *Client) NewFuturesOrderBookHistoryService() *FuturesOrderBookHistoryService {
	return &FuturesOrderBookHistoryService{c: c}
}

// NewCreateFuturesAlgoVpOrderService create futures algo vp order
func (c *Client) NewCreateFuturesAlgoVpOrderService() *CreateFuturesAlgoVpOrderService {
	return &CreateFuturesAlgoVpOrderService{c: c}
}

// NewCreateFuturesAlgoTwapOrderService create futures algo twap order
func (c *Client) NewCreateFuturesAlgoTwapOrderService() *CreateFuturesAlgoTwapOrderService {
	return &CreateFuturesAlgoTwapOrderService{c: c}
}

// NewListOpenFuturesAlgoOrdersService list open futures algo orders
func (c *Client) NewListOpenFuturesAlgoOrdersService() *ListOpenFuturesAlgoOrdersService {
	return &ListOpenFuturesAlgoOrdersService{c: c}
}

// NewListHistoryFuturesAlgoOrdersService list history futures algo orders
func (c *Client) NewListHistoryFuturesAlgoOrdersService() *ListHistoryFuturesAlgoOrdersService {
	return &ListHistoryFuturesAlgoOrdersService{c: c}
}

// NewCancelFuturesAlgoOrderService cancel future algo order
func (c *Client) NewCancelFuturesAlgoOrderService() *CancelFuturesAlgoOrderService {
	return &CancelFuturesAlgoOrderService{c: c}
}

// NewGetFuturesAlgoSubOrdersService get futures algo sub orders
func (c *Client) NewGetFuturesAlgoSubOrdersService() *GetFuturesAlgoSubOrdersService {
	return &GetFuturesAlgoSubOrdersService{c: c}
}

// ----- simple earn service -----
func (c *Client) NewSimpleEarnService() *SimpleEarnService {
	return &SimpleEarnService{c: c}
}

// ----- end simple earn service -----

func (c *Client) NewDualInvestmentService() *DualInvestmentService {
	return &DualInvestmentService{c: c}
}
