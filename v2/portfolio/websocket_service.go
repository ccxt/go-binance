package portfolio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bitly/go-simplejson"
)

// Endpoints
var (
	BaseWsMainUrl = "wss://fstream.binance.com/pm"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 600
	// WebsocketPongTimeout is an interval for sending a PONG frame in response to PING frame from server
	WebsocketPongTimeout = time.Second * 10
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
	// WebsocketTimeoutReadWriteConnection is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	// using for websocket API (read/write)
	WebsocketTimeoutReadWriteConnection = time.Second * 10
	ProxyUrl                            = ""
	Header                              = make(http.Header)
)

func getWsProxyUrl() *string {
	if ProxyUrl == "" {
		return nil
	}
	return &ProxyUrl
}

func SetWsProxyUrl(url string) {
	ProxyUrl = url
}

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint() string {
	return BaseWsMainUrl
}

// WsUserDataEvent define user data event
type WsUserDataEvent struct {
	Event           UserDataEventType `json:"e"`
	Time            int64             `json:"E"`
	TransactionTime int64             `json:"T"`

	// listenKeyExpired only have Event and Time
	//

	// MARGIN_CALL
	WsUserDataMarginCall

	// ACCOUNT_UPDATE
	WsUserDataAccountUpdate

	// ORDER_TRADE_UPDATE
	WsUserDataOrderTradeUpdate

	// ACCOUNT_CONFIG_UPDATE
	WsUserDataAccountConfigUpdate

	// TRADE_LITE
	WsUserDataTradeLite
}

type WsUserDataAccountConfigUpdate struct {
	AccountConfigUpdate WsAccountConfigUpdate `json:"ac"`
}

type WsUserDataAccountUpdate struct {
	AccountUpdate WsAccountUpdate `json:"a"`
}

type WsUserDataMarginCall struct {
	CrossWalletBalance  string       `json:"cw"`
	MarginCallPositions []WsPosition `json:"p"`
}

type WsUserDataOrderTradeUpdate struct {
	OrderTradeUpdate WsOrderTradeUpdate `json:"o"`
}

type WsUserDataTradeLite struct {
	Symbol          string   `json:"s"`
	OriginalQty     string   `json:"q"`
	OriginalPrice   string   //`json:"p"`
	IsMaker         bool     `json:"m"`
	ClientOrderID   string   `json:"c"`
	Side            SideType `json:"S"`
	LastFilledPrice string   `json:"L"`
	LastFilledQty   string   `json:"l"`
	TradeID         int64    `json:"t"`
	OrderID         int64    `json:"i"`
}

func (w *WsUserDataTradeLite) fromSimpleJson(j *simplejson.Json) (err error) {
	w.Symbol = j.Get("s").MustString()
	w.OriginalQty = j.Get("q").MustString()
	w.OriginalPrice = j.Get("p").MustString()
	w.IsMaker = j.Get("m").MustBool()
	w.ClientOrderID = j.Get("c").MustString()
	w.Side = SideType(j.Get("S").MustString())
	w.LastFilledPrice = j.Get("L").MustString()
	w.LastFilledQty = j.Get("l").MustString()
	w.TradeID = j.Get("t").MustInt64()
	w.OrderID = j.Get("i").MustInt64()
	return nil
}

func (e *WsUserDataEvent) UnmarshalJSON(data []byte) error {
	j, err := newJSON(data)
	if err != nil {
		return err
	}
	e.Event = UserDataEventType(j.Get("e").MustString())
	e.Time = j.Get("E").MustInt64()
	if v, ok := j.CheckGet("T"); ok {
		e.TransactionTime = v.MustInt64()
	}

	eventMaps := map[UserDataEventType]any{
		UserDataEventTypeMarginCall:          &e.WsUserDataMarginCall,
		UserDataEventTypeAccountUpdate:       &e.WsUserDataAccountUpdate,
		UserDataEventTypeOrderTradeUpdate:    &e.WsUserDataOrderTradeUpdate,
		UserDataEventTypeAccountConfigUpdate: &e.WsUserDataAccountConfigUpdate,
	}

	switch e.Event {
	case UserDataEventTypeTradeLite:
		return e.WsUserDataTradeLite.fromSimpleJson(j)
	case UserDataEventTypeListenKeyExpired:
		// noting
	default:
		if v, ok := eventMaps[e.Event]; ok {
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unexpected event type: %v", e.Event)
		}
	}
	return nil
}

// WsAccountUpdate define account update
type WsAccountUpdate struct {
	Reason    UserDataEventReasonType `json:"m"`
	Balances  []WsBalance             `json:"B"`
	Positions []WsPosition            `json:"P"`
}

// WsBalance define balance
type WsBalance struct {
	Asset              string `json:"a"`
	Balance            string `json:"wb"`
	CrossWalletBalance string `json:"cw"`
	ChangeBalance      string `json:"bc"`
}

// WsPosition define position
type WsPosition struct {
	Symbol              string           `json:"s"`
	Side                PositionSideType `json:"ps"`
	Amount              string           `json:"pa"`
	EntryPrice          string           `json:"ep"`
	UnrealizedPnL       string           `json:"up"`
	AccumulatedRealized string           `json:"cr"`
	BreakEvenPrice      float64          `json:"bep"`
}

// WsOrderTradeUpdate define order trade update
type WsOrderTradeUpdate struct {
	Symbol               string             `json:"s"`   // Symbol
	ClientOrderID        string             `json:"c"`   // Client order ID
	Side                 SideType           `json:"S"`   // Side
	Type                 OrderType          `json:"o"`   // Order type
	TimeInForce          TimeInForceType    `json:"f"`   // Time in force
	OriginalQty          string             `json:"q"`   // Original quantity
	OriginalPrice        string             `json:"p"`   // Original price
	AveragePrice         string             `json:"ap"`  // Average price
	StopPrice            string             `json:"sp"`  // Stop price. Please ignore with TRAILING_STOP_MARKET order
	ExecutionType        OrderExecutionType `json:"x"`   // Execution type
	Status               OrderStatusType    `json:"X"`   // Order status
	ID                   int64              `json:"i"`   // Order ID
	LastFilledQty        string             `json:"l"`   // Order Last Filled Quantity
	AccumulatedFilledQty string             `json:"z"`   // Order Filled Accumulated Quantity
	LastFilledPrice      string             `json:"L"`   // Last Filled Price
	CommissionAsset      string             `json:"N"`   // Commission Asset, will not push if no commission
	Commission           string             `json:"n"`   // Commission, will not push if no commission
	TradeTime            int64              `json:"T"`   // Order Trade Time
	TradeID              int64              `json:"t"`   // Trade ID
	BidsNotional         string             `json:"b"`   // Bids Notional
	AsksNotional         string             `json:"a"`   // Asks Notional
	IsMaker              bool               `json:"m"`   // Is this trade the maker side?
	IsReduceOnly         bool               `json:"R"`   // Is this reduce only
	PositionSide         PositionSideType   `json:"ps"`  // Position Side
	PriceProtect         bool               `json:"pP"`  // If price protection is turned on
	RealizedPnL          string             `json:"rp"`  // Realized Profit of the trade
	STP                  string             `json:"V"`   // STP mode
	PriceMode            string             `json:"pm"`  // Price match mode
	GTD                  int64              `json:"gtd"` // TIF GTD order auto cancel time
}

// WsAccountConfigUpdate define account config update
type WsAccountConfigUpdate struct {
	Symbol   string `json:"s"`
	Leverage int64  `json:"l"`
}

// WsConditionalOrderTradeUpdate represents a conditional order trade update event
type WsConditionalOrderTradeUpdate struct {
	EventType string `json:"e"`
	TransTime int64  `json:"T"`
	EventTime int64  `json:"E"`
	Business  string `json:"fs"`
	Order     struct {
		Symbol          string `json:"s"`
		ClientOrderID   string `json:"c"`
		StrategyID      int64  `json:"si"`
		Side            string `json:"S"`
		StrategyType    string `json:"st"`
		TimeInForce     string `json:"f"`
		Quantity        string `json:"q"`
		Price           string `json:"p"`
		StopPrice       string `json:"sp"`
		OrderStatus     string `json:"os"`
		OrderTime       int64  `json:"T"`
		UpdateTime      int64  `json:"ut"`
		ReduceOnly      bool   `json:"R"`
		WorkingType     string `json:"wt"`
		PositionSide    string `json:"ps"`
		ClosePosition   bool   `json:"cp"`
		ActivationPrice string `json:"AP"`
		CallbackRate    string `json:"cr"`
		OrderID         int64  `json:"i"`
		STPMode         string `json:"V"`
		GTD             int64  `json:"gtd"`
	} `json:"so"`
}

// WsOpenOrderLossUpdate represents an open order loss update event
type WsOpenOrderLossUpdate struct {
	EventType string `json:"e"` // "openOrderLoss"
	EventTime int64  `json:"E"`
	Orders    []struct {
		Asset  string `json:"a"`
		Amount string `json:"o"`
	} `json:"O"`
}

// WsMarginAccountUpdate represents a margin account update event
type WsMarginAccountUpdate struct {
	EventType      string            `json:"e"` // "outboundAccountPosition"
	EventTime      int64             `json:"E"`
	LastUpdateTime int64             `json:"u"`
	UpdateID       int64             `json:"U"`
	Balances       []WsMarginBalance `json:"B"`
}

type WsMarginBalance struct {
	Asset  string `json:"a"`
	Free   string `json:"f"`
	Locked string `json:"l"`
}

// WsLiabilityUpdate represents a liability update event
type WsLiabilityUpdate struct {
	EventType      string `json:"e"` // "liabilityChange"
	EventTime      int64  `json:"E"`
	Asset          string `json:"a"`
	Type           string `json:"t"` // e.g., "BORROW"
	TransactionID  int64  `json:"T"`
	Principal      string `json:"p"`
	Interest       string `json:"i"`
	TotalLiability string `json:"l"`
}

// WsMarginOrderUpdate represents a margin order update event
type WsMarginOrderUpdate struct {
	EventType               string `json:"e"` // "executionReport"
	EventTime               int64  `json:"E"`
	Symbol                  string `json:"s"`
	ClientOrderID           string `json:"c"`
	Side                    string `json:"S"`
	OrderType               string `json:"o"`
	TimeInForce             string `json:"f"`
	Quantity                string `json:"q"`
	Price                   string `json:"p"`
	StopPrice               string `json:"P"`
	TrailingDelta           int    `json:"d"`
	IcebergQuantity         string `json:"F"`
	OrderListID             int64  `json:"g"`
	OrigClientOrderID       string `json:"C"`
	ExecutionType           string `json:"x"`
	OrderStatus             string `json:"X"`
	RejectReason            string `json:"r"`
	OrderID                 int64  `json:"i"`
	LastExecutedQty         string `json:"l"`
	CumulativeFilledQty     string `json:"z"`
	LastExecutedPrice       string `json:"L"`
	CommissionAmount        string `json:"n"`
	CommissionAsset         string `json:"N"`
	TransactionTime         int64  `json:"T"`
	TradeID                 int64  `json:"t"`
	PreventedMatchID        int64  `json:"v"`
	UpdateID                int64  `json:"I"`
	IsOnBook                bool   `json:"w"`
	IsMaker                 bool   `json:"m"`
	OrderCreationTime       int64  `json:"O"`
	QuoteQtyFilled          string `json:"Z"`
	LastQuoteQtyFilled      string `json:"Y"`
	QuoteOrderQty           string `json:"Q"`
	TrailingTime            int64  `json:"D"`
	StrategyID              int64  `json:"j"`
	StrategyType            int64  `json:"J"`
	WorkingTime             int64  `json:"W"`
	SelfTradePreventionMode string `json:"V"`
	TradeGroupID            int64  `json:"u"`
	CounterOrderID          int64  `json:"U"`
	PreventedQuantity       string `json:"A"`
	LastPreventedQuantity   string `json:"B"`
}

// WsFuturesOrderUpdate represents a futures order update event
type WsFuturesOrderUpdate struct {
	EventType       string             `json:"e"`  // "ORDER_TRADE_UPDATE"
	BusinessUnit    string             `json:"fs"` // "UM" or "CM"
	EventTime       int64              `json:"E"`
	TransactionTime int64              `json:"T"`
	AccountAlias    string             `json:"i"`
	Order           WsFuturesOrderData `json:"o"`
}

type WsFuturesOrderData struct {
	Symbol          string             `json:"s"`
	ClientOrderID   string             `json:"c"`
	Side            SideType           `json:"S"`
	OrderType       OrderType          `json:"o"`
	TimeInForce     TimeInForceType    `json:"f"`
	OriginalQty     string             `json:"q"`
	OriginalPrice   string             `json:"p"`
	AveragePrice    string             `json:"ap"`
	StopPrice       string             `json:"sp"`
	ExecutionType   OrderExecutionType `json:"x"`
	OrderStatus     OrderStatusType    `json:"X"`
	OrderID         int64              `json:"i"`
	LastFilledQty   string             `json:"l"`
	FilledAccumQty  string             `json:"z"`
	LastFilledPrice string             `json:"L"`
	CommissionAsset string             `json:"N"`
	Commission      string             `json:"n"`
	TradeTime       int64              `json:"T"`
	TradeID         int64              `json:"t"`
	BidsNotional    string             `json:"b"`
	AskNotional     string             `json:"a"`
	IsMaker         bool               `json:"m"`
	IsReduceOnly    bool               `json:"R"`
	PositionSide    PositionSideType   `json:"ps"`
	RealizedProfit  string             `json:"rp"`
	StrategyType    string             `json:"st"`
	StrategyID      int64              `json:"si"`
	STPMode         string             `json:"V"`
	GTD             int64              `json:"gtd"`
}

// WsFuturesAccountUpdate represents a futures account update event
type WsFuturesAccountUpdate struct {
	EventType       string `json:"e"`  // "ACCOUNT_UPDATE"
	BusinessUnit    string `json:"fs"` // "UM" or "CM"
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	AccountAlias    string `json:"i"`
	AccountData     struct {
		ReasonType string              `json:"m"`
		Balances   []WsFuturesBalance  `json:"B"`
		Positions  []WsFuturesPosition `json:"P"`
	} `json:"a"`
}

type WsFuturesBalance struct {
	Asset              string `json:"a"`
	WalletBalance      string `json:"wb"`
	CrossWalletBalance string `json:"cw"`
	BalanceChange      string `json:"bc"`
}

type WsFuturesPosition struct {
	Symbol              string           `json:"s"`
	PositionAmount      string           `json:"pa"`
	EntryPrice          string           `json:"ep"`
	AccumulatedRealized string           `json:"cr"`
	UnrealizedPnL       string           `json:"up"`
	PositionSide        PositionSideType `json:"ps"`
	BreakEvenPrice      string           `json:"bep"`
}

// WsFuturesAccountConfigUpdate represents a futures account configuration update event
type WsFuturesAccountConfigUpdate struct {
	EventType       string `json:"e"`  // "ACCOUNT_CONFIG_UPDATE"
	BusinessUnit    string `json:"fs"` // "UM" or "CM"
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	AccountConfig   struct {
		Symbol   string `json:"s"`
		Leverage int64  `json:"l"`
	} `json:"ac"`
}

// Constants for account update reason types
const (
	ReasonDeposit             = "DEPOSIT"
	ReasonWithdraw            = "WITHDRAW"
	ReasonOrder               = "ORDER"
	ReasonFundingFee          = "FUNDING_FEE"
	ReasonWithdrawReject      = "WITHDRAW_REJECT"
	ReasonAdjustment          = "ADJUSTMENT"
	ReasonInsuranceClear      = "INSURANCE_CLEAR"
	ReasonAdminDeposit        = "ADMIN_DEPOSIT"
	ReasonAdminWithdraw       = "ADMIN_WITHDRAW"
	ReasonMarginTransfer      = "MARGIN_TRANSFER"
	ReasonMarginTypeChange    = "MARGIN_TYPE_CHANGE"
	ReasonAssetTransfer       = "ASSET_TRANSFER"
	ReasonOptionsPremiumFee   = "OPTIONS_PREMIUM_FEE"
	ReasonOptionsSettleProfit = "OPTIONS_SETTLE_PROFIT"
	ReasonAutoExchange        = "AUTO_EXCHANGE"
	ReasonCoinSwapDeposit     = "COIN_SWAP_DEPOSIT"
	ReasonCoinSwapWithdraw    = "COIN_SWAP_WITHDRAW"
)

// WsRiskLevelChange represents a risk level change event
type WsRiskLevelChange struct {
	EventType            string `json:"e"` // "riskLevelChange"
	EventTime            int64  `json:"E"`
	UniMMRLevel          string `json:"u"`
	Status               string `json:"s"` // MARGIN_CALL, REDUCE_ONLY, FORCE_LIQUIDATION
	EquityUSD            string `json:"eq"`
	ActualEquityUSD      string `json:"ae"`
	MaintenanceMarginUSD string `json:"m"`
}

// Risk level status constants
const (
	RiskStatusMarginCall       = "MARGIN_CALL"
	RiskStatusReduceOnly       = "REDUCE_ONLY"
	RiskStatusForceLiquidation = "FORCE_LIQUIDATION"
)

// WsMarginBalanceUpdate represents a margin balance update event
type WsMarginBalanceUpdate struct {
	EventType    string `json:"e"` // "balanceUpdate"
	EventTime    int64  `json:"E"`
	Asset        string `json:"a"`
	BalanceDelta string `json:"d"`
	UpdateID     int64  `json:"U"`
	ClearTime    int64  `json:"T"`
}

// WsListenKeyExpired represents a listen key expired event
type WsListenKeyExpired struct {
	EventType string `json:"e"` // "listenKeyExpired"
	EventTime int64  `json:"E"`
}

// WsUserDataHandler represents a handler for user data events
type WsUserDataHandler interface {
	HandleListenKeyExpired(*WsListenKeyExpired)
	HandleMarginBalanceUpdate(*WsMarginBalanceUpdate)
	HandleRiskLevelChange(*WsRiskLevelChange)
	HandleFuturesAccountConfigUpdate(*WsFuturesAccountConfigUpdate)
	HandleFuturesAccountUpdate(*WsFuturesAccountUpdate)
	HandleFuturesOrderUpdate(*WsFuturesOrderUpdate)
	HandleMarginOrderUpdate(*WsMarginOrderUpdate)
	HandleLiabilityUpdate(*WsLiabilityUpdate)
	HandleMarginAccountUpdate(*WsMarginAccountUpdate)
	HandleOpenOrderLossUpdate(*WsOpenOrderLossUpdate)
	HandleConditionalOrderTradeUpdate(*WsConditionalOrderTradeUpdate)
}

func wsUserDataHandler(handler WsUserDataHandler) func(message []byte) {
	return func(message []byte) {
		var event struct {
			EventType string `json:"e"`
			// fix golang bug: https://github.com/golang/go/issues/14750
			EventTime int64 `json:"E"`
		}
		if err := json.Unmarshal(message, &event); err != nil {
			return
		}

		switch event.EventType {
		case "listenKeyExpired":
			var expiredEvent WsListenKeyExpired
			if err := json.Unmarshal(message, &expiredEvent); err != nil {
				return
			}
			handler.HandleListenKeyExpired(&expiredEvent)
		case "balanceUpdate":
			var balanceUpdate WsMarginBalanceUpdate
			if err := json.Unmarshal(message, &balanceUpdate); err != nil {
				return
			}
			handler.HandleMarginBalanceUpdate(&balanceUpdate)
		case "riskLevelChange":
			var riskUpdate WsRiskLevelChange
			if err := json.Unmarshal(message, &riskUpdate); err != nil {
				return
			}
			handler.HandleRiskLevelChange(&riskUpdate)
		case "ACCOUNT_CONFIG_UPDATE":
			var configUpdate WsFuturesAccountConfigUpdate
			if err := json.Unmarshal(message, &configUpdate); err != nil {
				return
			}
			handler.HandleFuturesAccountConfigUpdate(&configUpdate)
		case "ACCOUNT_UPDATE":
			var accountUpdate WsFuturesAccountUpdate
			if err := json.Unmarshal(message, &accountUpdate); err != nil {
				return
			}
			handler.HandleFuturesAccountUpdate(&accountUpdate)
		case "ORDER_TRADE_UPDATE":
			var orderUpdate WsFuturesOrderUpdate
			if err := json.Unmarshal(message, &orderUpdate); err != nil {
				return
			}
			handler.HandleFuturesOrderUpdate(&orderUpdate)
		case "executionReport":
			var orderUpdate WsMarginOrderUpdate
			if err := json.Unmarshal(message, &orderUpdate); err != nil {
				return
			}
			handler.HandleMarginOrderUpdate(&orderUpdate)
		case "liabilityChange":
			var liabilityUpdate WsLiabilityUpdate
			if err := json.Unmarshal(message, &liabilityUpdate); err != nil {
				return
			}
			handler.HandleLiabilityUpdate(&liabilityUpdate)
		case "outboundAccountPosition":
			var accountUpdate WsMarginAccountUpdate
			if err := json.Unmarshal(message, &accountUpdate); err != nil {
				return
			}
			handler.HandleMarginAccountUpdate(&accountUpdate)
		case "openOrderLoss":
			var openOrderLoss WsOpenOrderLossUpdate
			if err := json.Unmarshal(message, &openOrderLoss); err != nil {
				return
			}
			handler.HandleOpenOrderLossUpdate(&openOrderLoss)
		case "CONDITIONAL_ORDER_TRADE_UPDATE":
			var conditionalOrderUpdate WsConditionalOrderTradeUpdate
			if err := json.Unmarshal(message, &conditionalOrderUpdate); err != nil {
				return
			}
			handler.HandleConditionalOrderTradeUpdate(&conditionalOrderUpdate)
		}
	}
}

// WsUserDataServe enhanced with automatic listen key renewal
func WsUserDataServe(listenKey string, handler WsUserDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/ws/%s", getWsEndpoint(), listenKey)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event struct {
			EventType string `json:"e"`
			// fix golang bug: https://github.com/golang/go/issues/14750
			EventTime int64 `json:"E"`
		}
		if err := json.Unmarshal(message, &event); err != nil {
			errHandler(err)
			return
		}

		wsUserDataHandler(handler)(message)
	}

	return wsServe(cfg, wsHandler, errHandler)
}
