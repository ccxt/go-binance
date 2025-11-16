package futures

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

// Endpoints
var (
	BaseAPIMainURL    = "https://fapi.binance.com"
	BaseAPITestnetURL = "https://testnet.binancefuture.com"
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
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	OrderTypeTrailingStopMarket OrderType = "TRAILING_STOP_MARKET"
	OrderTypeLiquidation        OrderType = "LIQUIDATION"

	TimeInForceTypeGTC    TimeInForceType = "GTC"     // Good Till Cancel
	TimeInForceTypeGTD    TimeInForceType = "GTD"     // Good Till Date
	TimeInForceTypeGTEGTC TimeInForceType = "GTE_GTC" // https://github.com/ccxt/go-binance/issues/681
	TimeInForceTypeIOC    TimeInForceType = "IOC"     // Immediate or Cancel
	TimeInForceTypeFOK    TimeInForceType = "FOK"     // Fill or Kill
	TimeInForceTypeGTX    TimeInForceType = "GTX"     // Good Till Crossing (Post Only)

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

	UserDataEventTypeListenKeyExpired              UserDataEventType = "listenKeyExpired"
	UserDataEventTypeMarginCall                    UserDataEventType = "MARGIN_CALL"
	UserDataEventTypeAccountUpdate                 UserDataEventType = "ACCOUNT_UPDATE"
	UserDataEventTypeOrderTradeUpdate              UserDataEventType = "ORDER_TRADE_UPDATE"
	UserDataEventTypeAccountConfigUpdate           UserDataEventType = "ACCOUNT_CONFIG_UPDATE"
	UserDataEventTypeTradeLite                     UserDataEventType = "TRADE_LITE"
	UserDataEventTypeConditionalOrderTriggerReject UserDataEventType = "CONDITIONAL_ORDER_TRIGGER_REJECT"

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
	if UseTestnet {
		return BaseAPITestnetURL
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
		apiErr := new(common.APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s\n", e)
		}
		if !apiErr.IsValid() {
			apiErr.Response = data
		}
		return nil, &res.Header, apiErr
	}
	return data, &res.Header, nil
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

// NewContinuousKlinesService init continuous klines service
func (c *Client) NewContinuousKlinesService() *ContinuousKlinesService {
	return &ContinuousKlinesService{c: c}
}

// NewIndexPriceKlinesService init index price klines service
func (c *Client) NewIndexPriceKlinesService() *IndexPriceKlinesService {
	return &IndexPriceKlinesService{c: c}
}

// NewMarkPriceKlinesService init markPriceKlines service
func (c *Client) NewMarkPriceKlinesService() *MarkPriceKlinesService {
	return &MarkPriceKlinesService{c: c}
}

// NewListPriceChangeStatsService init list prices change stats service
func (c *Client) NewListPriceChangeStatsService() *ListPriceChangeStatsService {
	return &ListPriceChangeStatsService{c: c}
}

// NewListPricesService init listing prices service
func (c *Client) NewListPricesService() *ListPricesService {
	return &ListPricesService{c: c}
}

// NewListBookTickersService init listing booking tickers service
func (c *Client) NewListBookTickersService() *ListBookTickersService {
	return &ListBookTickersService{c: c}
}

// NewCreateOrderService init creating order service
func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{c: c}
}

// NewModifyOrderService init creating order service
func (c *Client) NewModifyOrderService() *ModifyOrderService {
	return &ModifyOrderService{c: c}
}

// NewCreateBatchOrdersService init creating batch order service
func (c *Client) NewCreateBatchOrdersService() *CreateBatchOrdersService {
	return &CreateBatchOrdersService{c: c}
}

// NewModifyBatchOrdersService init modifying batch order service
func (c *Client) NewModifyBatchOrdersService() *ModifyBatchOrdersService {
	return &ModifyBatchOrdersService{c: c}
}

// NewGetOrderService init get order service
func (c *Client) NewGetOrderService() *GetOrderService {
	return &GetOrderService{c: c}
}

// NewCancelOrderService init cancel order service
func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{c: c}
}

// NewCancelAllOpenOrdersService init cancel all open orders service
func (c *Client) NewCancelAllOpenOrdersService() *CancelAllOpenOrdersService {
	return &CancelAllOpenOrdersService{c: c}
}

// NewCancelMultipleOrdersService init cancel multiple orders service
func (c *Client) NewCancelMultipleOrdersService() *CancelMultiplesOrdersService {
	return &CancelMultiplesOrdersService{c: c}
}

// NewGetOpenOrderService init get open order service
func (c *Client) NewGetOpenOrderService() *GetOpenOrderService {
	return &GetOpenOrderService{c: c}
}

// NewListOpenOrdersService init list open orders service
func (c *Client) NewListOpenOrdersService() *ListOpenOrdersService {
	return &ListOpenOrdersService{c: c}
}

// NewListOrdersService init listing orders service
func (c *Client) NewListOrdersService() *ListOrdersService {
	return &ListOrdersService{c: c}
}

// NewGetAccountService init getting account service
func (c *Client) NewGetAccountService() *GetAccountService {
	return &GetAccountService{c: c}
}

// NewGetAccountV3Service init getting account service
func (c *Client) NewGetAccountV3Service() *GetAccountV3Service {
	return &GetAccountV3Service{c: c}
}

// NewGetBalanceService init getting balance service
func (c *Client) NewGetBalanceService() *GetBalanceService {
	return &GetBalanceService{c: c}
}

// NewGetAccountConfigService init get futures account configuration service
func (c *Client) NewGetAccountConfigService() *AccountConfigService {
	return &AccountConfigService{c: c}
}

// NewGetSymbolConfigService init get futures symbol configuration service
func (c *Client) NewGetSymbolConfigService() *SymbolConfigService {
	return &SymbolConfigService{c: c}
}

func (c *Client) NewGetPositionRiskService() *GetPositionRiskService {
	return &GetPositionRiskService{c: c}
}

func (c *Client) NewGetPositionRiskV3Service() *GetPositionRiskV3Service {
	return &GetPositionRiskV3Service{c: c}
}

// NewGetPositionMarginHistoryService init getting position margin history service
func (c *Client) NewGetPositionMarginHistoryService() *GetPositionMarginHistoryService {
	return &GetPositionMarginHistoryService{c: c}
}

// NewGetIncomeHistoryService init getting income history service
func (c *Client) NewGetIncomeHistoryService() *GetIncomeHistoryService {
	return &GetIncomeHistoryService{c: c}
}

// NewHistoricalTradesService init listing trades service
func (c *Client) NewHistoricalTradesService() *HistoricalTradesService {
	return &HistoricalTradesService{c: c}
}

// NewListAccountTradeService init account trade list service
func (c *Client) NewListAccountTradeService() *ListAccountTradeService {
	return &ListAccountTradeService{c: c}
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

// NewPremiumIndexService init premium index service
func (c *Client) NewPremiumIndexService() *PremiumIndexService {
	return &PremiumIndexService{c: c}
}

// NewPremiumIndexKlinesService init premium index klines service
func (c *Client) NewPremiumIndexKlinesService() *PremiumIndexKlinesService {
	return &PremiumIndexKlinesService{c: c}
}

// NewFundingRateService init funding rate service
func (c *Client) NewFundingRateService() *FundingRateService {
	return &FundingRateService{c: c}
}

// NewFundingRateInfoService init funding rate info service
func (c *Client) NewFundingRateInfoService() *FundingRateInfoService {
	return &FundingRateInfoService{c: c}
}

// NewListUserLiquidationOrdersService init list user's liquidation orders service
func (c *Client) NewListUserLiquidationOrdersService() *ListUserLiquidationOrdersService {
	return &ListUserLiquidationOrdersService{c: c}
}

// NewListLiquidationOrdersService init funding rate service
func (c *Client) NewListLiquidationOrdersService() *ListLiquidationOrdersService {
	return &ListLiquidationOrdersService{c: c}
}

// NewChangeLeverageService init change leverage service
func (c *Client) NewChangeLeverageService() *ChangeLeverageService {
	return &ChangeLeverageService{c: c}
}

// NewGetLeverageBracketService init change leverage service
func (c *Client) NewGetLeverageBracketService() *GetLeverageBracketService {
	return &GetLeverageBracketService{c: c}
}

// NewChangeMarginTypeService init change margin type service
func (c *Client) NewChangeMarginTypeService() *ChangeMarginTypeService {
	return &ChangeMarginTypeService{c: c}
}

// NewUpdatePositionMarginService init update position margin
func (c *Client) NewUpdatePositionMarginService() *UpdatePositionMarginService {
	return &UpdatePositionMarginService{c: c}
}

// NewChangePositionModeService init change position mode service
func (c *Client) NewChangePositionModeService() *ChangePositionModeService {
	return &ChangePositionModeService{c: c}
}

// NewGetPositionModeService init get position mode service
func (c *Client) NewGetPositionModeService() *GetPositionModeService {
	return &GetPositionModeService{c: c}
}

// NewChangeMultiAssetModeService init change multi-asset mode service
func (c *Client) NewChangeMultiAssetModeService() *ChangeMultiAssetModeService {
	return &ChangeMultiAssetModeService{c: c}
}

// NewGetMultiAssetModeService init get multi-asset mode service
func (c *Client) NewGetMultiAssetModeService() *GetMultiAssetModeService {
	return &GetMultiAssetModeService{c: c}
}

// NewGetRebateNewUserService init get rebate_newuser service
func (c *Client) NewGetRebateNewUserService() *GetRebateNewUserService {
	return &GetRebateNewUserService{c: c}
}

// NewCommissionRateService returns commission rate
func (c *Client) NewCommissionRateService() *CommissionRateService {
	return &CommissionRateService{c: c}
}

// NewGetOpenInterestService init open interest service
func (c *Client) NewGetOpenInterestService() *GetOpenInterestService {
	return &GetOpenInterestService{c: c}
}

// NewOpenInterestStatisticsService init open interest statistics service
func (c *Client) NewOpenInterestStatisticsService() *OpenInterestStatisticsService {
	return &OpenInterestStatisticsService{c: c}
}

// NewLongShortRatioService init open interest statistics service
func (c *Client) NewLongShortRatioService() *LongShortRatioService {
	return &LongShortRatioService{c: c}
}

func (c *Client) NewDeliveryPriceService() *DeliveryPriceService {
	return &DeliveryPriceService{c: c}
}

func (c *Client) NewTopLongShortAccountRatioService() *TopLongShortAccountRatioService {
	return &TopLongShortAccountRatioService{c: c}
}

func (c *Client) NewTopLongShortPositionRatioService() *TopLongShortPositionRatioService {
	return &TopLongShortPositionRatioService{c: c}
}

func (c *Client) NewTakerLongShortRatioService() *TakerLongShortRatioService {
	return &TakerLongShortRatioService{c: c}
}

func (c *Client) NewBasisService() *BasisService {
	return &BasisService{c: c}
}

func (c *Client) NewIndexInfoService() *IndexInfoService {
	return &IndexInfoService{c: c}
}

func (c *Client) NewAssetIndexService() *AssetIndexService {
	return &AssetIndexService{c: c}
}

func (c *Client) NewConstituentsService() *ConstituentsService {
	return &ConstituentsService{c: c}
}

func (c *Client) NewLvtKlinesService() *LvtKlinesService {
	return &LvtKlinesService{c: c}
}

func (c *Client) NewGetFeeBurnService() *GetFeeBurnService {
	return &GetFeeBurnService{c: c}
}

func (c *Client) NewFeeBurnService() *FeeBurnService {
	return &FeeBurnService{c: c}
}

// NewListConvertAssetsService init list convert assets service
func (c *Client) NewListConvertExchangeInfoService() *ListConvertExchangeInfoService {
	return &ListConvertExchangeInfoService{c: c}
}

// NewCreateConvertQuoteService init create convert quote service
func (c *Client) NewCreateConvertQuoteService() *CreateConvertQuoteService {
	return &CreateConvertQuoteService{c: c}
}

// NewCreateConvertService init accept convert quote service
func (c *Client) NewConvertAcceptService() *ConvertAcceptService {
	return &ConvertAcceptService{c: c}
}

// NewGetConvertStatusService init get convert status service
func (c *Client) NewGetConvertStatusService() *ConvertStatusService {
	return &ConvertStatusService{c: c}
}

// NewApiTradingStatusService init get api trading status service
func (c *Client) NewApiTradingStatusService() *ApiTradingStatusService {
	return &ApiTradingStatusService{c: c}
}
