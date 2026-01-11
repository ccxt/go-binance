package portfolio

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
	"time"

	"github.com/bitly/go-simplejson"

	"github.com/adshao/go-binance/v2/common"
)

// SideType define side type of order
type SideType string

// PositionSideType define position side type of order
type PositionSideType string

// OrderType define order type
type OrderType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// NewOrderRespType define response JSON verbosity
type NewOrderRespType string

// OrderExecutionType define order execution type
type OrderExecutionType string

// OrderStatusType define order status type
type OrderStatusType string

// PriceMatchType define priceMatch type
// Can't be passed together with price
type PriceMatchType string

// SymbolType define symbol type
type SymbolType string

// SymbolStatusType define symbol status type
type SymbolStatusType string

// SymbolFilterType define symbol filter type
type SymbolFilterType string

// SideEffectType define side effect type for orders
type SideEffectType string

// WorkingType define working type
type WorkingType string

// MarginType define margin type
type MarginType string

// ContractType define contract type
type ContractType string

// UserDataEventType define user data event type
type UserDataEventType string

// UserDataEventReasonType define reason type for user data event
type UserDataEventReasonType string

// ForceOrderCloseType define reason type for force order
type ForceOrderCloseType string

// SelfTradePreventionMode define self trade prevention strategy
type SelfTradePreventionMode string

// StrategyType define strategy type for conditional orders
type StrategyType string

// UMPosition define UM position information
type UMPosition struct {
	Symbol                 string `json:"symbol"`                     // symbol name
	PositionAmt            string `json:"positionAmt"`                // position amount
	EntryPrice             string `json:"entryPrice"`                 // average entry price
	MarkPrice              string `json:"markPrice,omitempty"`        // mark price (only in position risk endpoint)
	UnrealizedProfit       string `json:"unrealizedProfit"`           // unrealized profit
	LiquidationPrice       string `json:"liquidationPrice,omitempty"` // liquidation price (only in position risk endpoint)
	Leverage               string `json:"leverage"`                   // current initial leverage
	MaxNotional            string `json:"maxNotional,omitempty"`      // maximum available notional with current leverage (account detail)
	MaxNotionalValue       string `json:"maxNotionalValue,omitempty"` // maximum notional value (position risk)
	PositionSide           string `json:"positionSide"`               // position side
	InitialMargin          string `json:"initialMargin"`              // initial margin required with current mark price
	MaintMargin            string `json:"maintMargin"`                // maintenance margin required
	PositionInitialMargin  string `json:"positionInitialMargin"`      // initial margin required for positions with current mark price
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`     // initial margin required for open orders with current mark price
	Notional               string `json:"notional,omitempty"`         // notional value (only in position risk endpoint)
	BidNotional            string `json:"bidNotional"`                // bids notional
	AskNotional            string `json:"askNotional"`                // ask notional
	UpdateTime             int64  `json:"updateTime"`                 // last update time
}

// ... existing code ...

// CMPosition define CM position information
type CMPosition struct {
	Symbol                 string `json:"symbol"`                     // Symbol name
	PositionAmt            string `json:"positionAmt"`                // Position amount
	EntryPrice             string `json:"entryPrice"`                 // Average entry price
	UnrealizedProfit       string `json:"unRealizedProfit"`           // Unrealized profit or loss
	PositionSide           string `json:"positionSide"`               // Position side (BOTH, LONG, SHORT)
	InitialMargin          string `json:"initialMargin"`              // Initial margin required
	MaintMargin            string `json:"maintMargin"`                // Maintenance margin required
	PositionInitialMargin  string `json:"positionInitialMargin"`      // Position initial margin
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`     // Open orders initial margin
	Leverage               string `json:"leverage"`                   // Current leverage
	MaxQty                 string `json:"maxQty"`                     // Maximum quantity of base asset
	MarkPrice              string `json:"markPrice,omitempty"`        // Mark price
	LiquidationPrice       string `json:"liquidationPrice,omitempty"` // Liquidation price
	NotionalValue          string `json:"notionalValue,omitempty"`    // Notional value
	UpdateTime             int64  `json:"updateTime"`                 // Last update time
}

// ... rest of the code ...

// Endpoints
var (
	BaseAPIMainURL = "https://papi.binance.com"
	//BaseAPITestnetURL = "https://testnet.binancefuture.com"
)

// Global enums
const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	PositionSideTypeBoth  PositionSideType = "BOTH"
	PositionSideTypeLong  PositionSideType = "LONG"
	PositionSideTypeShort PositionSideType = "SHORT"

	OrderTypeLimit              OrderType = "LIMIT"
	OrderTypeMarket             OrderType = "MARKET"
	OrderTypeStop               OrderType = "STOP"
	OrderTypeStopMarket         OrderType = "STOP_MARKET"
	OrderTypeTakeProfit         OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitLimit    OrderType = "TAKE_PROFIT_LIMIT"
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	OrderTypeTrailingStopMarket OrderType = "TRAILING_STOP_MARKET"
	OrderTypeLiquidation        OrderType = "LIQUIDATION"

	TimeInForceTypeGTC TimeInForceType = "GTC" // Good Till Cancel
	TimeInForceTypeIOC TimeInForceType = "IOC" // Immediate or Cancel
	TimeInForceTypeFOK TimeInForceType = "FOK" // Fill or Kill
	TimeInForceTypeGTX TimeInForceType = "GTX" // Good Till Crossing (Post Only)
	TimeInForceTypeGTD TimeInForceType = "GTD" // Good Till Date

	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeRESULT NewOrderRespType = "RESULT"

	OrderExecutionTypeNew         OrderExecutionType = "NEW"
	OrderExecutionTypePartialFill OrderExecutionType = "PARTIAL_FILL"
	OrderExecutionTypeFill        OrderExecutionType = "FILL"
	OrderExecutionTypeCanceled    OrderExecutionType = "CANCELED"
	OrderExecutionTypeCalculated  OrderExecutionType = "CALCULATED"
	OrderExecutionTypeExpired     OrderExecutionType = "EXPIRED"
	OrderExecutionTypeTrade       OrderExecutionType = "TRADE"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"
	OrderStatusTypeRejected        OrderStatusType = "REJECTED"
	OrderStatusTypeExpired         OrderStatusType = "EXPIRED"
	OrderStatusTypeNewInsurance    OrderStatusType = "NEW_INSURANCE"
	OrderStatusTypeNewADL          OrderStatusType = "NEW_ADL"

	PriceMatchTypeOpponent   PriceMatchType = "OPPONENT"
	PriceMatchTypeOpponent5  PriceMatchType = "OPPONENT_5"
	PriceMatchTypeOpponent10 PriceMatchType = "OPPONENT_10"
	PriceMatchTypeOpponent20 PriceMatchType = "OPPONENT_20"
	PriceMatchTypeQueue      PriceMatchType = "QUEUE"
	PriceMatchTypeQueue5     PriceMatchType = "QUEUE_5"
	PriceMatchTypeQueue10    PriceMatchType = "QUEUE_10"
	PriceMatchTypeQueue20    PriceMatchType = "QUEUE_20"
	PriceMatchTypeNone       PriceMatchType = "NONE"

	SymbolTypeFuture SymbolType = "FUTURE"

	WorkingTypeMarkPrice     WorkingType = "MARK_PRICE"
	WorkingTypeContractPrice WorkingType = "CONTRACT_PRICE"

	SymbolStatusTypePreTrading   SymbolStatusType = "PRE_TRADING"
	SymbolStatusTypeTrading      SymbolStatusType = "TRADING"
	SymbolStatusTypePostTrading  SymbolStatusType = "POST_TRADING"
	SymbolStatusTypeEndOfDay     SymbolStatusType = "END_OF_DAY"
	SymbolStatusTypeHalt         SymbolStatusType = "HALT"
	SymbolStatusTypeAuctionMatch SymbolStatusType = "AUCTION_MATCH"
	SymbolStatusTypeBreak        SymbolStatusType = "BREAK"

	SymbolFilterTypeLotSize          SymbolFilterType = "LOT_SIZE"
	SymbolFilterTypePrice            SymbolFilterType = "PRICE_FILTER"
	SymbolFilterTypePercentPrice     SymbolFilterType = "PERCENT_PRICE"
	SymbolFilterTypeMarketLotSize    SymbolFilterType = "MARKET_LOT_SIZE"
	SymbolFilterTypeMaxNumOrders     SymbolFilterType = "MAX_NUM_ORDERS"
	SymbolFilterTypeMaxNumAlgoOrders SymbolFilterType = "MAX_NUM_ALGO_ORDERS"
	SymbolFilterTypeMinNotional      SymbolFilterType = "MIN_NOTIONAL"

	SideEffectTypeNoSideEffect SideEffectType = "NO_SIDE_EFFECT"
	SideEffectTypeMarginBuy    SideEffectType = "MARGIN_BUY"
	SideEffectTypeAutoRepay    SideEffectType = "AUTO_REPAY"

	MarginTypeIsolated MarginType = "ISOLATED"
	MarginTypeCrossed  MarginType = "CROSSED"

	ContractTypePerpetual      ContractType = "PERPETUAL"
	ContractTypeCurrentQuarter ContractType = "CURRENT_QUARTER"
	ContractTypeNextQuarter    ContractType = "NEXT_QUARTER"

	UserDataEventTypeListenKeyExpired    UserDataEventType = "listenKeyExpired"
	UserDataEventTypeMarginCall          UserDataEventType = "MARGIN_CALL"
	UserDataEventTypeAccountUpdate       UserDataEventType = "ACCOUNT_UPDATE"
	UserDataEventTypeOrderTradeUpdate    UserDataEventType = "ORDER_TRADE_UPDATE"
	UserDataEventTypeAccountConfigUpdate UserDataEventType = "ACCOUNT_CONFIG_UPDATE"
	UserDataEventTypeTradeLite           UserDataEventType = "TRADE_LITE"

	UserDataEventReasonTypeDeposit             UserDataEventReasonType = "DEPOSIT"
	UserDataEventReasonTypeWithdraw            UserDataEventReasonType = "WITHDRAW"
	UserDataEventReasonTypeOrder               UserDataEventReasonType = "ORDER"
	UserDataEventReasonTypeFundingFee          UserDataEventReasonType = "FUNDING_FEE"
	UserDataEventReasonTypeWithdrawReject      UserDataEventReasonType = "WITHDRAW_REJECT"
	UserDataEventReasonTypeAdjustment          UserDataEventReasonType = "ADJUSTMENT"
	UserDataEventReasonTypeInsuranceClear      UserDataEventReasonType = "INSURANCE_CLEAR"
	UserDataEventReasonTypeAdminDeposit        UserDataEventReasonType = "ADMIN_DEPOSIT"
	UserDataEventReasonTypeAdminWithdraw       UserDataEventReasonType = "ADMIN_WITHDRAW"
	UserDataEventReasonTypeMarginTransfer      UserDataEventReasonType = "MARGIN_TRANSFER"
	UserDataEventReasonTypeMarginTypeChange    UserDataEventReasonType = "MARGIN_TYPE_CHANGE"
	UserDataEventReasonTypeAssetTransfer       UserDataEventReasonType = "ASSET_TRANSFER"
	UserDataEventReasonTypeOptionsPremiumFee   UserDataEventReasonType = "OPTIONS_PREMIUM_FEE"
	UserDataEventReasonTypeOptionsSettleProfit UserDataEventReasonType = "OPTIONS_SETTLE_PROFIT"

	ForceOrderCloseTypeLiquidation ForceOrderCloseType = "LIQUIDATION"
	ForceOrderCloseTypeADL         ForceOrderCloseType = "ADL"

	SelfTradePreventionModeNone        SelfTradePreventionMode = "NONE"
	SelfTradePreventionModeExpireTaker SelfTradePreventionMode = "EXPIRE_TAKER"
	SelfTradePreventionModeExpireBoth  SelfTradePreventionMode = "EXPIRE_BOTH"
	SelfTradePreventionModeExpireMaker SelfTradePreventionMode = "EXPIRE_MAKER"

	StrategyTypeStop               StrategyType = "STOP"
	StrategyTypeStopMarket         StrategyType = "STOP_MARKET"
	StrategyTypeTakeProfit         StrategyType = "TAKE_PROFIT"
	StrategyTypeTakeProfitMarket   StrategyType = "TAKE_PROFIT_MARKET"
	StrategyTypeTrailingStopMarket StrategyType = "TRAILING_STOP_MARKET"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

func currentTimestamp() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// getApiEndpoint return the base endpoint of the WS according the UseTestnet flag
func getApiEndpoint() string {
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
		BaseURL:    getApiEndpoint(),
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
		BaseURL:   getApiEndpoint(),
		UserAgent: "Binance/golang",
		HTTPClient: &http.Client{
			Transport: tr,
		},
		Logger: log.New(os.Stderr, "Binance-golang ", log.LstdFlags),
	}
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

	UsedWeight common.UsedWeight
	OrderCount common.OrderCount
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
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
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

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, &http.Header{}, err
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
		return []byte{}, &http.Header{}, err
	}
	c.UsedWeight.UpdateByHeader(res.Header)
	c.OrderCount.UpdateByHeader(res.Header)
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, &http.Header{}, err
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
		// Try to parse the error response
		var apiErr Error
		e := json.Unmarshal(data, &apiErr)
		if e != nil {
			c.debug("failed to unmarshal error response: %s\n", e)
			// If we can't parse the JSON response, return a generic error with the raw response
			return nil, &res.Header, NewErrorFromResponse(int64(res.StatusCode), res.Status, data)
		}
		// Return the parsed error with the raw response included
		return nil, &res.Header, NewErrorFromResponse(apiErr.Code, apiErr.Message, data)
	}
	return data, &res.Header, nil
}

// SetApiEndpoint set api Endpoint
func (c *Client) SetApiEndpoint(url string) *Client {
	c.BaseURL = url
	return c
}

// --------- Market Data ---------

// NewPingService init ping service
func (c *Client) NewPingService() *PingService {
	return &PingService{c: c}
}

// --------- Account ---------

// NewGetBalanceService init server balance service
func (c *Client) NewGetBalanceService() *GetBalanceService {
	return &GetBalanceService{c: c}
}

// NewGetAccountService init get account service
func (c *Client) NewGetAccountService() *GetAccountService {
	return &GetAccountService{c: c}
}

// Add account status constants
type AccountStatus string

const (
	AccountStatusNormal            AccountStatus = "NORMAL"
	AccountStatusMarginCall        AccountStatus = "MARGIN_CALL"
	AccountStatusSupplyMargin      AccountStatus = "SUPPLY_MARGIN"
	AccountStatusReduceOnly        AccountStatus = "REDUCE_ONLY"
	AccountStatusActiveLiquidation AccountStatus = "ACTIVE_LIQUIDATION"
	AccountStatusForceLiquidation  AccountStatus = "FORCE_LIQUIDATION"
	AccountStatusBankrupted        AccountStatus = "BANKRUPTED"
)

// NewGetMarginMaxBorrowService init margin max borrow service
func (c *Client) NewGetMarginMaxBorrowService() *GetMarginMaxBorrowService {
	return &GetMarginMaxBorrowService{c: c}
}

// NewGetMarginMaxWithdrawService init margin max withdraw service
func (c *Client) NewGetMarginMaxWithdrawService() *GetMarginMaxWithdrawService {
	return &GetMarginMaxWithdrawService{c: c}
}

// NewGetUMPositionRiskService init UM position risk service
func (c *Client) NewGetUMPositionRiskService() *GetUMPositionRiskService {
	return &GetUMPositionRiskService{c: c}
}

// NewGetCMPositionRiskService init CM position risk service
func (c *Client) NewGetCMPositionRiskService() *GetCMPositionRiskService {
	return &GetCMPositionRiskService{c: c}
}

// NewChangeUMInitialLeverageService init change UM initial leverage service
func (c *Client) NewChangeUMInitialLeverageService() *ChangeUMInitialLeverageService {
	return &ChangeUMInitialLeverageService{c: c}
}

// NewChangeCMInitialLeverageService init change CM initial leverage service
func (c *Client) NewChangeCMInitialLeverageService() *ChangeCMInitialLeverageService {
	return &ChangeCMInitialLeverageService{c: c}
}

// NewChangeUMPositionModeService init change UM position mode service
func (c *Client) NewChangeUMPositionModeService() *ChangeUMPositionModeService {
	return &ChangeUMPositionModeService{c: c}
}

// NewChangeCMPositionModeService init change CM position mode service
func (c *Client) NewChangeCMPositionModeService() *ChangeCMPositionModeService {
	return &ChangeCMPositionModeService{c: c}
}

// NewGetUMPositionModeService init get UM position mode service
func (c *Client) NewGetUMPositionModeService() *GetUMPositionModeService {
	return &GetUMPositionModeService{c: c}
}

// NewGetCMPositionModeService init get CM position mode service
func (c *Client) NewGetCMPositionModeService() *GetCMPositionModeService {
	return &GetCMPositionModeService{c: c}
}

// NewGetUMLeverageBracketService init get UM leverage bracket service
func (c *Client) NewGetUMLeverageBracketService() *GetUMLeverageBracketService {
	return &GetUMLeverageBracketService{c: c}
}

// NewGetCMLeverageBracketService init get CM leverage bracket service
func (c *Client) NewGetCMLeverageBracketService() *GetCMLeverageBracketService {
	return &GetCMLeverageBracketService{c: c}
}

// NewGetUMTradingStatusService init get UM trading status service
func (c *Client) NewGetUMTradingStatusService() *GetUMTradingStatusService {
	return &GetUMTradingStatusService{c: c}
}

// NewGetUMCommissionRateService init get UM commission rate service
func (c *Client) NewGetUMCommissionRateService() *GetUMCommissionRateService {
	return &GetUMCommissionRateService{c: c}
}

// NewGetCMCommissionRateService init get CM commission rate service
func (c *Client) NewGetCMCommissionRateService() *GetCMCommissionRateService {
	return &GetCMCommissionRateService{c: c}
}

// NewGetMarginLoanService init get margin loan service
func (c *Client) NewGetMarginLoanService() *GetMarginLoanService {
	return &GetMarginLoanService{c: c}
}

// NewGetMarginRepayService init get margin repay service
func (c *Client) NewGetMarginRepayService() *GetMarginRepayService {
	return &GetMarginRepayService{c: c}
}

// NewGetAutoRepayFuturesStatusService init get auto repay futures status service
func (c *Client) NewGetAutoRepayFuturesStatusService() *GetAutoRepayFuturesStatusService {
	return &GetAutoRepayFuturesStatusService{c: c}
}

// NewChangeAutoRepayFuturesStatusService init change auto repay futures status service
func (c *Client) NewChangeAutoRepayFuturesStatusService() *ChangeAutoRepayFuturesStatusService {
	return &ChangeAutoRepayFuturesStatusService{c: c}
}

// NewGetMarginInterestHistoryService init get margin interest history service
func (c *Client) NewGetMarginInterestHistoryService() *GetMarginInterestHistoryService {
	return &GetMarginInterestHistoryService{c: c}
}

// NewRepayFuturesNegativeBalanceService init repay futures negative balance service
func (c *Client) NewRepayFuturesNegativeBalanceService() *RepayFuturesNegativeBalanceService {
	return &RepayFuturesNegativeBalanceService{c: c}
}

// NewGetNegativeBalanceInterestHistoryService init get negative balance interest history service
func (c *Client) NewGetNegativeBalanceInterestHistoryService() *GetNegativeBalanceInterestHistoryService {
	return &GetNegativeBalanceInterestHistoryService{c: c}
}

// NewFundAutoCollectionService init fund auto-collection service
func (c *Client) NewFundAutoCollectionService() *FundAutoCollectionService {
	return &FundAutoCollectionService{c: c}
}

// NewFundCollectionByAssetService init fund collection by asset service
func (c *Client) NewFundCollectionByAssetService() *FundCollectionByAssetService {
	return &FundCollectionByAssetService{c: c}
}

// NewBNBTransferService init BNB transfer service
func (c *Client) NewBNBTransferService() *BNBTransferService {
	return &BNBTransferService{c: c}
}

// NewGetUMIncomeHistoryService init get UM income history service
func (c *Client) NewGetUMIncomeHistoryService() *GetUMIncomeHistoryService {
	return &GetUMIncomeHistoryService{c: c}
}

// NewGetCMIncomeHistoryService init get CM income history service
func (c *Client) NewGetCMIncomeHistoryService() *GetCMIncomeHistoryService {
	return &GetCMIncomeHistoryService{c: c}
}

// NewGetUMAccountDetailService init get UM account detail service
func (c *Client) NewGetUMAccountDetailService() *GetUMAccountDetailService {
	return &GetUMAccountDetailService{c: c}
}

// NewGetCMAccountDetailService init get CM account detail service
func (c *Client) NewGetCMAccountDetailService() *GetCMAccountDetailService {
	return &GetCMAccountDetailService{c: c}
}

// NewGetUMAccountConfigService init get UM futures account configuration service
func (c *Client) NewGetUMAccountConfigService() *UMAccountConfigService {
	return &UMAccountConfigService{c: c}
}

// NewGetUMSymbolConfigService init get UM futures symbol configuration service
func (c *Client) NewGetUMSymbolConfigService() *UMSymbolConfigService {
	return &UMSymbolConfigService{c: c}
}

// NewGetUMAccountDetailV2Service init get UM account detail v2 service
func (c *Client) NewGetUMAccountDetailV2Service() *UMAccountDetailV2Service {
	return &UMAccountDetailV2Service{c: c}
}

// NewGetUMTradeHistoryDownloadIDService init getting um trade history download id service
func (c *Client) NewGetUMTradeHistoryDownloadIDService() *GetUMTradeHistoryDownloadIDService {
	return &GetUMTradeHistoryDownloadIDService{c: c}
}

// NewGetUMTradeDownloadLinkService init getting um trade download link service
func (c *Client) NewGetUMTradeDownloadLinkService() *GetUMTradeDownloadLinkService {
	return &GetUMTradeDownloadLinkService{c: c}
}

// NewGetUMOrderHistoryDownloadIDService init getting um order history download id service
func (c *Client) NewGetUMOrderHistoryDownloadIDService() *GetUMOrderHistoryDownloadIDService {
	return &GetUMOrderHistoryDownloadIDService{c: c}
}

// NewGetUMOrderDownloadLinkService init getting um order download link service
func (c *Client) NewGetUMOrderDownloadLinkService() *GetUMOrderDownloadLinkService {
	return &GetUMOrderDownloadLinkService{c: c}
}

// NewGetUMTransactionHistoryDownloadIDService init getting um transaction history download id service
func (c *Client) NewGetUMTransactionHistoryDownloadIDService() *GetUMTransactionHistoryDownloadIDService {
	return &GetUMTransactionHistoryDownloadIDService{c: c}
}

// NewGetUMTransactionDownloadLinkService init getting um transaction download link service
func (c *Client) NewGetUMTransactionDownloadLinkService() *GetUMTransactionDownloadLinkService {
	return &GetUMTransactionDownloadLinkService{c: c}
}

// NewGetRateLimitService init getting rate limit service
func (c *Client) NewGetRateLimitService() *GetRateLimitService {
	return &GetRateLimitService{c: c}
}

// NewGetNegativeBalanceExchangeRecordService init getting negative balance exchange record service
func (c *Client) NewGetNegativeBalanceExchangeRecordService() *GetNegativeBalanceExchangeRecordService {
	return &GetNegativeBalanceExchangeRecordService{c: c}
}

// --------- Trade ---------

// NewUMOrderService init UM order service
func (c *Client) NewUMOrderService() *UMOrderService {
	return &UMOrderService{c: c}
}

// NewUMConditionalOrderService init UM conditional order service
func (c *Client) NewUMConditionalOrderService() *UMConditionalOrderService {
	return &UMConditionalOrderService{c: c}
}

// NewCMOrderService creates a new CMOrderService
func (c *Client) NewCMOrderService() *CMOrderService {
	return &CMOrderService{c: c}
}

// NewCMConditionalOrderService creates a new CMConditionalOrderService
func (c *Client) NewCMConditionalOrderService() *CMConditionalOrderService {
	return &CMConditionalOrderService{c: c}
}

// NewMarginOrderService creates a new MarginOrderService
func (c *Client) NewMarginOrderService() *MarginOrderService {
	return &MarginOrderService{c: c}
}

// NewMarginLoanService creates a new MarginLoanService
func (c *Client) NewMarginLoanService() *MarginLoanService {
	return &MarginLoanService{c: c}
}

// NewMarginRepayService creates a new MarginRepayService
func (c *Client) NewMarginRepayService() *MarginRepayService {
	return &MarginRepayService{c: c}
}

// NewMarginOCOService creates a new MarginOCOService
func (c *Client) NewMarginOCOService() *MarginOCOService {
	return &MarginOCOService{c: c}
}

// NewUMCancelOrderService creates a new UMCancelOrderService
func (c *Client) NewUMCancelOrderService() *UMCancelOrderService {
	return &UMCancelOrderService{c: c}
}

// NewUMCancelAllOrdersService creates a new UMCancelAllOrdersService
func (c *Client) NewUMCancelAllOrdersService() *UMCancelAllOrdersService {
	return &UMCancelAllOrdersService{c: c}
}

// NewUMCancelConditionalOrderService creates a new UMCancelConditionalOrderService
func (c *Client) NewUMCancelConditionalOrderService() *UMCancelConditionalOrderService {
	return &UMCancelConditionalOrderService{c: c}
}

// NewUMCancelAllConditionalOrdersService creates a new UMCancelAllConditionalOrdersService
func (c *Client) NewUMCancelAllConditionalOrdersService() *UMCancelAllConditionalOrdersService {
	return &UMCancelAllConditionalOrdersService{c: c}
}

// NewCMCancelOrderService creates a new CMCancelOrderService
func (c *Client) NewCMCancelOrderService() *CMCancelOrderService {
	return &CMCancelOrderService{c: c}
}

// NewCMCancelAllOrdersService creates a new CMCancelAllOrdersService
func (c *Client) NewCMCancelAllOrdersService() *CMCancelAllOrdersService {
	return &CMCancelAllOrdersService{c: c}
}

// NewCMCancelConditionalOrderService creates a new CMCancelConditionalOrderService
func (c *Client) NewCMCancelConditionalOrderService() *CMCancelConditionalOrderService {
	return &CMCancelConditionalOrderService{c: c}
}

// NewCMCancelAllConditionalOrdersService creates a new CMCancelAllConditionalOrdersService
func (c *Client) NewCMCancelAllConditionalOrdersService() *CMCancelAllConditionalOrdersService {
	return &CMCancelAllConditionalOrdersService{c: c}
}

// NewMarginCancelOrderService creates a new MarginCancelOrderService
func (c *Client) NewMarginCancelOrderService() *MarginCancelOrderService {
	return &MarginCancelOrderService{c: c}
}

// NewMarginCancelOCOService creates a new MarginCancelOCOService
func (c *Client) NewMarginCancelOCOService() *MarginCancelOCOService {
	return &MarginCancelOCOService{c: c}
}

// NewUMModifyOrderService creates a new UMModifyOrderService
func (c *Client) NewUMModifyOrderService() *UMModifyOrderService {
	return &UMModifyOrderService{c: c}
}

// NewMarginCancelAllOrdersService creates a new MarginCancelAllOrdersService
func (c *Client) NewMarginCancelAllOrdersService() *MarginCancelAllOrdersService {
	return &MarginCancelAllOrdersService{c: c}
}

// NewCMModifyOrderService creates a new CMModifyOrderService
func (c *Client) NewCMModifyOrderService() *CMModifyOrderService {
	return &CMModifyOrderService{c: c}
}

// NewUMQueryOrderService creates a new UMQueryOrderService
func (c *Client) NewUMQueryOrderService() *UMQueryOrderService {
	return &UMQueryOrderService{c: c}
}

// NewUMAllOrdersService creates a new UMAllOrdersService
func (c *Client) NewUMAllOrdersService() *UMAllOrdersService {
	return &UMAllOrdersService{c: c}
}

// NewUMOpenOrderService creates a new UMOpenOrderService
func (c *Client) NewUMOpenOrderService() *UMOpenOrderService {
	return &UMOpenOrderService{c: c}
}

// NewUMOpenOrdersService creates a new UMOpenOrdersService
func (c *Client) NewUMOpenOrdersService() *UMOpenOrdersService {
	return &UMOpenOrdersService{c: c}
}

// NewUMAllConditionalOrdersService creates a new UMAllConditionalOrdersService
func (c *Client) NewUMAllConditionalOrdersService() *UMAllConditionalOrdersService {
	return &UMAllConditionalOrdersService{c: c}
}

// NewUMOpenConditionalOrdersService creates a new UMOpenConditionalOrdersService
func (c *Client) NewUMOpenConditionalOrdersService() *UMOpenConditionalOrdersService {
	return &UMOpenConditionalOrdersService{c: c}
}

// NewUMOpenConditionalOrderService creates a new UMOpenConditionalOrderService
func (c *Client) NewUMOpenConditionalOrderService() *UMOpenConditionalOrderService {
	return &UMOpenConditionalOrderService{c: c}
}

// NewUMConditionalOrderHistoryService creates a new UMConditionalOrderHistoryService
func (c *Client) NewUMConditionalOrderHistoryService() *UMConditionalOrderHistoryService {
	return &UMConditionalOrderHistoryService{c: c}
}

// NewCMQueryOrderService creates a new CMQueryOrderService
func (c *Client) NewCMQueryOrderService() *CMQueryOrderService {
	return &CMQueryOrderService{c: c}
}

// NewCMAllOrdersService creates a new CMAllOrdersService
func (c *Client) NewCMAllOrdersService() *CMAllOrdersService {
	return &CMAllOrdersService{c: c}
}

// NewCMOpenOrderService creates a new CMOpenOrderService
func (c *Client) NewCMOpenOrderService() *CMOpenOrderService {
	return &CMOpenOrderService{c: c}
}

// NewCMOpenOrdersService creates a new CMOpenOrdersService
func (c *Client) NewCMOpenOrdersService() *CMOpenOrdersService {
	return &CMOpenOrdersService{c: c}
}

// NewCMOpenConditionalOrdersService creates a new CMOpenConditionalOrdersService
func (c *Client) NewCMOpenConditionalOrdersService() *CMOpenConditionalOrdersService {
	return &CMOpenConditionalOrdersService{c: c}
}

// NewCMOpenConditionalOrderService creates a new CMOpenConditionalOrderService
func (c *Client) NewCMOpenConditionalOrderService() *CMOpenConditionalOrderService {
	return &CMOpenConditionalOrderService{c: c}
}

// NewCMConditionalOrder creates a new CMConditionalOrder
func (c *Client) NewCMConditionalOrder() *CMConditionalOrderService {
	return &CMConditionalOrderService{c: c}
}

// NewCMConditionalOrdersService creates a new CMConditionalOrdersService
func (c *Client) NewCMConditionalOrdersService() *CMConditionalOrdersService {
	return &CMConditionalOrdersService{c: c}
}

// NewCMConditionalOrderHistoryService creates a new CMConditionalOrderHistoryService
func (c *Client) NewCMConditionalOrderHistoryService() *CMConditionalOrderHistoryService {
	return &CMConditionalOrderHistoryService{c: c}
}

// NewUMForceOrdersService creates a new UMForceOrdersService
func (c *Client) NewUMForceOrdersService() *UMForceOrdersService {
	return &UMForceOrdersService{c: c}
}

// NewCMForceOrdersService creates a new CMForceOrdersService
func (c *Client) NewCMForceOrdersService() *CMForceOrdersService {
	return &CMForceOrdersService{c: c}
}

// NewUMModifyOrderHistoryService creates a new UMModifyOrderHistoryService
func (c *Client) NewUMModifyOrderHistoryService() *UMModifyOrderHistoryService {
	return &UMModifyOrderHistoryService{c: c}
}

// NewCMModifyOrderHistoryService creates a new CMModifyOrderHistoryService
func (c *Client) NewCMModifyOrderHistoryService() *CMModifyOrderHistoryService {
	return &CMModifyOrderHistoryService{c: c}
}

// NewMarginForceOrdersService creates a new MarginForceOrdersService
func (c *Client) NewMarginForceOrdersService() *MarginForceOrdersService {
	return &MarginForceOrdersService{c: c}
}

// NewUMAccountTradeService creates a new UMAccountTradeService
func (c *Client) NewUMAccountTradeService() *UMAccountTradeService {
	return &UMAccountTradeService{c: c}
}

// NewCMAccountTradeService creates a new CMAccountTradeService
func (c *Client) NewCMAccountTradeService() *CMAccountTradeService {
	return &CMAccountTradeService{c: c}
}

// NewUMADLQuantileService creates a new UMADLQuantileService
func (c *Client) NewUMADLQuantileService() *UMADLQuantileService {
	return &UMADLQuantileService{c: c}
}

// NewCMADLQuantileService creates a new CMADLQuantileService
func (c *Client) NewCMADLQuantileService() *CMADLQuantileService {
	return &CMADLQuantileService{c: c}
}

// NewUMFeeBurnService creates a new UMFeeBurnService
func (c *Client) NewUMFeeBurnService() *UMFeeBurnService {
	return &UMFeeBurnService{c: c}
}

// NewGetMarginOrderService creates a new GetMarginOrderService
func (c *Client) NewGetMarginOrderService() *GetMarginOrderService {
	return &GetMarginOrderService{c: c}
}

// NewGetMarginOpenOrdersService creates a new GetMarginOpenOrdersService
func (c *Client) NewGetMarginOpenOrdersService() *GetMarginOpenOrdersService {
	return &GetMarginOpenOrdersService{c: c}
}

// NewGetMarginAllOrdersService creates a new GetMarginAllOrdersService
func (c *Client) NewGetMarginAllOrdersService() *GetMarginAllOrdersService {
	return &GetMarginAllOrdersService{c: c}
}

// NewUMFeeBurnStatusService creates a new UMFeeBurnStatusService
func (c *Client) NewUMFeeBurnStatusService() *UMFeeBurnStatusService {
	return &UMFeeBurnStatusService{c: c}
}

// NewGetMarginForceOrdersService creates a new GetMarginForceOrdersService
func (c *Client) NewGetMarginForceOrdersService() *GetMarginForceOrdersService {
	return &GetMarginForceOrdersService{c: c}
}

// NewUMAccountTradesService creates a new UMAccountTradesService
func (c *Client) NewUMAccountTradesService() *UMAccountTradesService {
	return &UMAccountTradesService{c: c}
}

// NewCMAccountTradesService creates a new CMAccountTradesService
func (c *Client) NewCMAccountTradesService() *CMAccountTradesService {
	return &CMAccountTradesService{c: c}
}

// NewUMGetADLQuantileService creates a new UMGetADLQuantileService
func (c *Client) NewUMGetADLQuantileService() *UMGetADLQuantileService {
	return &UMGetADLQuantileService{c: c}
}

// NewMarginOCOQueryService creates a new MarginOCOQueryService
func (c *Client) NewMarginOCOQueryService() *MarginOCOQueryService {
	return &MarginOCOQueryService{c: c}
}

// NewMarginOpenOCOService creates a new MarginOpenOCOService
func (c *Client) NewMarginOpenOCOService() *MarginOpenOCOService {
	return &MarginOpenOCOService{c: c}
}

// NewMarginAccountTradesService creates a new MarginAccountTradesService
func (c *Client) NewMarginAccountTradesService() *MarginAccountTradesService {
	return &MarginAccountTradesService{c: c}
}

// NewMarginRepayDebtService creates a new MarginRepayDebtService
func (c *Client) NewMarginRepayDebtService() *MarginRepayDebtService {
	return &MarginRepayDebtService{c: c}
}

// NewMarginAllOCOService creates a new MarginAllOCOService
func (c *Client) NewMarginAllOCOService() *MarginAllOCOService {
	return &MarginAllOCOService{c: c}
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
