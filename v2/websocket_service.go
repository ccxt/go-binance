package binance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2/common/websocket"
	"github.com/google/uuid"
	gorilla "github.com/gorilla/websocket"
)

var (
	// Endpoints
	BaseWsMainURL          = "wss://stream.binance.com:9443/ws"
	BaseWsTestnetURL       = "wss://stream.testnet.binance.vision/ws"
	BaseCombinedMainURL    = "wss://stream.binance.com:9443/stream?streams="
	BaseCombinedTestnetURL = "wss://stream.testnet.binance.vision/stream?streams="
	BaseWsApiMainURL       = "wss://ws-api.binance.com:443/ws-api/v3"
	BaseWsApiTestnetURL    = "wss://ws-api.testnet.binance.vision/ws-api/v3"

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
	if UseTestnet {
		return BaseWsTestnetURL
	}
	return BaseWsMainURL
}

// getCombinedEndpoint return the base endpoint of the combined stream according the UseTestnet flag
func getCombinedEndpoint() string {
	if UseTestnet {
		return BaseCombinedTestnetURL
	}
	return BaseCombinedMainURL
}

// WsPartialDepthEvent define websocket partial depth book event
type WsPartialDepthEvent struct {
	Symbol       string
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

// WsPartialDepthHandler handle websocket partial depth event
type WsPartialDepthHandler func(event *WsPartialDepthEvent)

// WsPartialDepthServe serve websocket partial depth handler with a symbol, using 1sec updates
func WsPartialDepthServe(symbol string, levels string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth%s", getWsEndpoint(), strings.ToLower(symbol), levels)
	return wsPartialDepthServe(endpoint, symbol, handler, errHandler)
}

// WsPartialDepthServe100Ms serve websocket partial depth handler with a symbol, using 100msec updates
func WsPartialDepthServe100Ms(symbol string, levels string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth%s@100ms", getWsEndpoint(), strings.ToLower(symbol), levels)
	return wsPartialDepthServe(endpoint, symbol, handler, errHandler)
}

// WsPartialDepthServe serve websocket partial depth handler with a symbol
func wsPartialDepthServe(endpoint string, symbol string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsPartialDepthEvent)
		event.Symbol = symbol
		event.LastUpdateID = j.Get("lastUpdateId").MustInt64()
		bidsLen := len(j.Get("bids").MustArray())
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := j.Get("bids").GetIndex(i)
			event.Bids[i] = Bid{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		asksLen := len(j.Get("asks").MustArray())
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {
			item := j.Get("asks").GetIndex(i)
			event.Asks[i] = Ask{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCombinedPartialDepthServe is similar to WsPartialDepthServe, but it for multiple symbols
func WsCombinedPartialDepthServe(symbolLevels map[string]string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for s, l := range symbolLevels {
		endpoint += fmt.Sprintf("%s@depth%s", strings.ToLower(s), l) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsPartialDepthEvent)
		stream := j.Get("stream").MustString()
		symbol := strings.Split(stream, "@")[0]
		event.Symbol = strings.ToUpper(symbol)
		data := j.Get("data").MustMap()
		event.LastUpdateID, _ = data["lastUpdateId"].(json.Number).Int64()
		bidsLen := len(data["bids"].([]interface{}))
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := data["bids"].([]interface{})[i].([]interface{})
			event.Bids[i] = Bid{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		asksLen := len(data["asks"].([]interface{}))
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {

			item := data["asks"].([]interface{})[i].([]interface{})
			event.Asks[i] = Ask{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsDepthHandler handle websocket depth event
type WsDepthHandler func(event *WsDepthEvent)

// WsDepthServe serve websocket depth handler with a symbol, using 1sec updates
func WsDepthServe(symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth", getWsEndpoint(), strings.ToLower(symbol))
	return wsDepthServe(endpoint, handler, errHandler)
}

// WsDepthServe100Ms serve websocket depth handler with a symbol, using 100msec updates
func WsDepthServe100Ms(symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth@100ms", getWsEndpoint(), strings.ToLower(symbol))
	return wsDepthServe(endpoint, handler, errHandler)
}

// WsDepthServe serve websocket depth handler with an arbitrary endpoint address
func wsDepthServe(endpoint string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsDepthEvent)
		event.Event = j.Get("e").MustString()
		event.Time = j.Get("E").MustInt64()
		event.Symbol = j.Get("s").MustString()
		event.LastUpdateID = j.Get("u").MustInt64()
		event.FirstUpdateID = j.Get("U").MustInt64()
		bidsLen := len(j.Get("b").MustArray())
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := j.Get("b").GetIndex(i)
			event.Bids[i] = Bid{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		asksLen := len(j.Get("a").MustArray())
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {
			item := j.Get("a").GetIndex(i)
			event.Asks[i] = Ask{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsDepthEvent define websocket depth event
type WsDepthEvent struct {
	Event         string `json:"e"`
	Time          int64  `json:"E"`
	Symbol        string `json:"s"`
	LastUpdateID  int64  `json:"u"`
	FirstUpdateID int64  `json:"U"`
	Bids          []Bid  `json:"b"`
	Asks          []Ask  `json:"a"`
}

// WsCombinedDepthServe is similar to WsDepthServe, but it for multiple symbols
func WsCombinedDepthServe(symbols []string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@depth", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return wsCombinedDepthServe(endpoint, handler, errHandler)
}

func WsCombinedDepthServe100Ms(symbols []string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@depth@100ms", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return wsCombinedDepthServe(endpoint, handler, errHandler)
}

func wsCombinedDepthServe(endpoint string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsDepthEvent)
		stream := j.Get("stream").MustString()
		symbol := strings.Split(stream, "@")[0]
		event.Symbol = strings.ToUpper(symbol)
		data := j.Get("data").MustMap()
		event.Time, _ = data["E"].(json.Number).Int64()
		event.LastUpdateID, _ = data["u"].(json.Number).Int64()
		event.FirstUpdateID, _ = data["U"].(json.Number).Int64()
		bidsLen := len(data["b"].([]interface{}))
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := data["b"].([]interface{})[i].([]interface{})
			event.Bids[i] = Bid{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		asksLen := len(data["a"].([]interface{}))
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {

			item := data["a"].([]interface{})[i].([]interface{})
			event.Asks[i] = Ask{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsKlineHandler handle websocket kline event
type WsKlineHandler func(event *WsKlineEvent)

// WsCombinedKlineServe is similar to WsKlineServe, but it handles multiple symbols with it interval
func WsCombinedKlineServe(symbolIntervalPair map[string]string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for symbol, interval := range symbolIntervalPair {
		endpoint += fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		data := j.Get("data").MustMap()

		symbol := strings.Split(stream, "@")[0]

		jsonData, _ := json.Marshal(data)

		event := new(WsKlineEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}
		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCombinedKlineServeMultiInterval is similar to WsCombinedKlineServe, but it supports multiple intervals per symbol
func WsCombinedKlineServeMultiInterval(symbolIntervals map[string][]string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for symbol, intervals := range symbolIntervals {
		for _, interval := range intervals {
			endpoint += fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval) + "/"
		}
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		data := j.Get("data").MustMap()

		symbol := strings.Split(stream, "@")[0]

		jsonData, _ := json.Marshal(data)

		event := new(WsKlineEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}
		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func WsKlineServe(symbol string, interval string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@kline_%s", getWsEndpoint(), strings.ToLower(symbol), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsKlineEvent define websocket kline event
type WsKlineEvent struct {
	Event  string  `json:"e"`
	Time   int64   `json:"E"`
	Symbol string  `json:"s"`
	Kline  WsKline `json:"k"`
}

// WsKline define websocket kline
type WsKline struct {
	StartTime            int64  `json:"t"`
	EndTime              int64  `json:"T"`
	Symbol               string `json:"s"`
	Interval             string `json:"i"`
	FirstTradeID         int64  `json:"f"`
	LastTradeID          int64  `json:"L"`
	Open                 string `json:"o"`
	Close                string `json:"c"`
	High                 string `json:"h"`
	Low                  string `json:"l"`
	Volume               string `json:"v"`
	TradeNum             int64  `json:"n"`
	IsFinal              bool   `json:"x"`
	QuoteVolume          string `json:"q"`
	ActiveBuyVolume      string `json:"V"`
	ActiveBuyQuoteVolume string `json:"Q"`
}

// WsAggTradeHandler handle websocket aggregate trade event
type WsAggTradeHandler func(event *WsAggTradeEvent)

// WsAggTradeServe serve websocket aggregate handler with a symbol
func WsAggTradeServe(symbol string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsAggTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCombinedAggTradeServe is similar to WsAggTradeServe, but it handles multiple symbolx
func WsCombinedAggTradeServe(symbols []string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for s := range symbols {
		endpoint += fmt.Sprintf("%s@aggTrade", strings.ToLower(symbols[s])) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		data := j.Get("data").MustMap()

		symbol := strings.Split(stream, "@")[0]

		jsonData, _ := json.Marshal(data)

		event := new(WsAggTradeEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}

		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAggTradeEvent define websocket aggregate trade event
type WsAggTradeEvent struct {
	Event                 string `json:"e"`
	Time                  int64  `json:"E"`
	Symbol                string `json:"s"`
	AggTradeID            int64  `json:"a"`
	Price                 string `json:"p"`
	Quantity              string `json:"q"`
	FirstBreakdownTradeID int64  `json:"f"`
	LastBreakdownTradeID  int64  `json:"l"`
	TradeTime             int64  `json:"T"`
	IsBuyerMaker          bool   `json:"m"`
	Placeholder           bool   `json:"M"` // add this field to avoid case insensitive unmarshalling
}

// WsTradeHandler handle websocket trade event
type WsTradeHandler func(event *WsTradeEvent)
type WsCombinedTradeHandler func(event *WsCombinedTradeEvent)

// WsTradeServe serve websocket handler with a symbol
func WsTradeServe(symbol string, handler WsTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@trade", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

func WsCombinedTradeServe(symbols []string, handler WsCombinedTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@trade/", strings.ToLower(s))
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsCombinedTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsTradeEvent define websocket trade event
type WsTradeEvent struct {
	Event         string `json:"e"`
	Time          int64  `json:"E"`
	Symbol        string `json:"s"`
	TradeID       int64  `json:"t"`
	Price         string `json:"p"`
	Quantity      string `json:"q"`
	BuyerOrderID  int64  `json:"b"`
	SellerOrderID int64  `json:"a"`
	TradeTime     int64  `json:"T"`
	IsBuyerMaker  bool   `json:"m"`
	Placeholder   bool   `json:"M"` // add this field to avoid case insensitive unmarshalling
}

type WsCombinedTradeEvent struct {
	Stream string       `json:"stream"`
	Data   WsTradeEvent `json:"data"`
}

// WsUserDataEvent define user data event
type WsUserDataEvent struct {
	Event         UserDataEventType `json:"e"`
	Time          int64             `json:"E"`
	AccountUpdate WsAccountUpdateList
	BalanceUpdate WsBalanceUpdate
	OrderUpdate   WsOrderUpdate
	OCOUpdate     WsOCOUpdate
}

type WsAccountUpdateList struct {
	AccountUpdateTime int64             `json:"u"`
	WsAccountUpdates  []WsAccountUpdate `json:"B"`
}

// WsAccountUpdate define account update
type WsAccountUpdate struct {
	Asset  string `json:"a"`
	Free   string `json:"f"`
	Locked string `json:"l"`
}

type WsBalanceUpdate struct {
	Asset           string `json:"a"`
	Change          string `json:"d"`
	TransactionTime int64  `json:"T"`
}

type WsOrderUpdate struct {
	Symbol                  string          `json:"s"`
	ClientOrderId           string          `json:"c"`
	Side                    string          `json:"S"`
	Type                    string          `json:"o"`
	TimeInForce             TimeInForceType `json:"f"`
	Volume                  string          `json:"q"`
	Price                   string          `json:"p"`
	StopPrice               string          `json:"P"`
	IceBergVolume           string          `json:"F"`
	OrderListId             int64           `json:"g"` // for OCO
	OrigCustomOrderId       string          `json:"C"` // customized order ID for the original order
	ExecutionType           string          `json:"x"` // execution type for this event NEW/TRADE...
	Status                  string          `json:"X"` // order status
	RejectReason            string          `json:"r"`
	Id                      int64           `json:"i"` // order id
	LatestVolume            string          `json:"l"` // quantity for the latest trade
	FilledVolume            string          `json:"z"`
	LatestPrice             string          `json:"L"` // price for the latest trade
	FeeAsset                string          `json:"N"`
	FeeCost                 string          `json:"n"`
	TransactionTime         int64           `json:"T"`
	TradeId                 int64           `json:"t"`
	IgnoreI                 int64           `json:"I"` // ignore
	IsInOrderBook           bool            `json:"w"` // is the order in the order book?
	IsMaker                 bool            `json:"m"` // is this order maker?
	IgnoreM                 bool            `json:"M"` // ignore
	CreateTime              int64           `json:"O"`
	FilledQuoteVolume       string          `json:"Z"` // the quote volume that already filled
	LatestQuoteVolume       string          `json:"Y"` // the quote volume for the latest trade
	QuoteVolume             string          `json:"Q"`
	SelfTradePreventionMode string          `json:"V"`

	//These are fields that appear in the payload only if certain conditions are met.
	TrailingDelta              int64  `json:"d"` // Appears only for trailing stop orders.
	TrailingTime               int64  `json:"D"`
	StrategyId                 int64  `json:"j"` // Appears only if the strategyId parameter was provided upon order placement.
	StrategyType               int64  `json:"J"` // Appears only if the strategyType parameter was provided upon order placement.
	PreventedMatchId           int64  `json:"v"` // Appears only for orders that expired due to STP.
	PreventedQuantity          string `json:"A"`
	LastPreventedQuantity      string `json:"B"`
	TradeGroupId               int64  `json:"u"`
	CounterOrderId             int64  `json:"U"`
	CounterSymbol              string `json:"Cs"`
	PreventedExecutionQuantity string `json:"pl"`
	PreventedExecutionPrice    string `json:"pL"`
	PreventedExecutionQuoteQty string `json:"pY"`
	WorkingTime                int64  `json:"W"` // Appears when the order is working on the book
	MatchType                  string `json:"b"`
	AllocationId               int64  `json:"a"`
	WorkingFloor               string `json:"k"`  // Appears for orders that could potentially have allocations
	UsedSor                    bool   `json:"uS"` // Appears for orders that used SOR
}

type WsOCOUpdate struct {
	Symbol          string `json:"s"`
	OrderListId     int64  `json:"g"`
	ContingencyType string `json:"c"`
	ListStatusType  string `json:"l"`
	ListOrderStatus string `json:"L"`
	RejectReason    string `json:"r"`
	ClientOrderId   string `json:"C"` // List Client Order ID
	TransactionTime int64  `json:"T"`
	Orders          WsOCOOrderList
}

type WsOCOOrderList struct {
	WsOCOOrders []WsOCOOrder `json:"O"`
}

type WsOCOOrder struct {
	Symbol        string `json:"s"`
	OrderId       int64  `json:"i"`
	ClientOrderId string `json:"c"`
}

// WsUserDataHandler handle WsUserDataEvent
type WsUserDataHandler func(event *WsUserDataEvent)

// WsUserDataServe serve user data handler with listen key
// Deprecated: Listen key management is deprecated. Use WsUserDataServeSignature instead.
func WsUserDataServe(listenKey string, handler WsUserDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s", getWsEndpoint(), listenKey)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		event := new(WsUserDataEvent)

		err = json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}

		switch UserDataEventType(j.Get("e").MustString()) {
		case UserDataEventTypeOutboundAccountPosition:
			err = json.Unmarshal(message, &event.AccountUpdate)
			if err != nil {
				errHandler(err)
				return
			}
		case UserDataEventTypeBalanceUpdate:
			err = json.Unmarshal(message, &event.BalanceUpdate)
			if err != nil {
				errHandler(err)
				return
			}
		case UserDataEventTypeExecutionReport:
			err = json.Unmarshal(message, &event.OrderUpdate)
			if err != nil {
				errHandler(err)
				return
			}
		case UserDataEventTypeListStatus:
			err = json.Unmarshal(message, &event.OCOUpdate)
			if err != nil {
				errHandler(err)
				return
			}
		}

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsUserDataServeSignature serves user data handler using signature-based subscription via WebSocket API.
// This is the recommended method as listen key management has been deprecated by Binance.
// It connects to the WebSocket API endpoint and subscribes to user data stream using signature authentication.
func WsUserDataServeSignature(apiKey, secretKey string, keyType string, timeOffset int64, handler WsUserDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(getWsApiEndpoint())

	doneC = make(chan struct{})
	stopC = make(chan struct{})

	conn, err := WsGetReadWriteConnection(cfg)
	if err != nil {
		return nil, nil, err
	}

	// Subscribe to user data stream using signature
	reqData := websocket.NewRequestData(
		uuid.New().String(),
		apiKey,
		secretKey,
		timeOffset,
		keyType,
	)
	subscribeRequest, err := websocket.CreateRequest(
		reqData,
		websocket.UserDataStreamSubscribeSignatureSpotWsApiMethod,
		map[string]interface{}{},
	)
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	// Send subscription request
	err = conn.WriteMessage(gorilla.TextMessage, subscribeRequest)
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	go func() {
		defer close(doneC)
		defer conn.Close()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				select {
				case <-stopC:
					return
				default:
					errHandler(err)
					return
				}
			}

			// Check if this is a subscription response
			j, err := newJSON(message)
			if err != nil {
				select {
				case <-stopC:
					continue
				default:
					errHandler(err)
					continue
				}
			}

			// Skip subscription confirmation messages
			if j.Get("id").MustString() != "" && j.Get("status").MustInt() == 200 {
				continue
			}

			// Some WS API pushes wrap the payload inside an envelope like "event", "result" or "data".
			// Determine the actual event payload to unmarshal from.
			payload := message
			// Try unwrap { event: {...} }
			if ev, ok := j.CheckGet("event"); ok {
				if m, mErr := ev.Map(); mErr == nil {
					if raw, encErr := json.Marshal(m); encErr == nil {
						payload = raw
					}
				}
			}
			// Try unwrap { result: {...} }
			if res, ok := j.CheckGet("result"); ok {
				if m, mErr := res.Map(); mErr == nil {
					if raw, encErr := json.Marshal(m); encErr == nil {
						payload = raw
					}
				}
			}
			// Try unwrap { data: {...} } (combined-like envelope)
			if string(payload) == string(message) {
				if data, ok := j.CheckGet("data"); ok {
					if m, mErr := data.Map(); mErr == nil {
						if raw, encErr := json.Marshal(m); encErr == nil {
							payload = raw
						}
					}
				}
			}

			// Parse user data event from the resolved payload
			event := new(WsUserDataEvent)
			if err = json.Unmarshal(payload, event); err != nil {
				select {
				case <-stopC:
					continue
				default:
					errHandler(err)
					continue
				}
			}

			// Determine event type from payload
			jj, _ := newJSON(payload)
			evtType := UserDataEventType(jj.Get("e").MustString())
			// ensure top-level fields are populated even if struct sub-unmarshal doesn't set them
			event.Event = evtType
			event.Time = jj.Get("E").MustInt64()

			switch evtType {
			case UserDataEventTypeOutboundAccountPosition:
				if err = json.Unmarshal(payload, &event.AccountUpdate); err != nil {
					select {
					case <-stopC:
						continue
					default:
						errHandler(err)
						continue
					}
				}
			case UserDataEventTypeBalanceUpdate:
				if err = json.Unmarshal(payload, &event.BalanceUpdate); err != nil {
					select {
					case <-stopC:
						continue
					default:
						errHandler(err)
						continue
					}
				}
			case UserDataEventTypeExecutionReport:
				if err = json.Unmarshal(payload, &event.OrderUpdate); err != nil {
					select {
					case <-stopC:
						continue
					default:
						errHandler(err)
						continue
					}
				}
			case UserDataEventTypeListStatus:
				if err = json.Unmarshal(payload, &event.OCOUpdate); err != nil {
					select {
					case <-stopC:
						continue
					default:
						errHandler(err)
						continue
					}
				}
			default:
				// Unknown event; still forward to handler so callers can log/inspect
			}
			handler(event)
		}
	}()

	return doneC, stopC, nil
}

// WsMarketStatHandler handle websocket that push single market statistics for 24hr
type WsMarketStatHandler func(event *WsMarketStatEvent)

// WsCombinedMarketStatServe is similar to WsMarketStatServe, but it handles multiple symbolx
func WsCombinedMarketStatServe(symbols []string, handler WsMarketStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for s := range symbols {
		endpoint += fmt.Sprintf("%s@ticker", strings.ToLower(symbols[s])) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)

	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		data := j.Get("data").MustMap()

		symbol := strings.Split(stream, "@")[0]

		jsonData, _ := json.Marshal(data)

		event := new(WsMarketStatEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}

		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsMarketStatServe serve websocket that push 24hr statistics for single market every second
func WsMarketStatServe(symbol string, handler WsMarketStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@ticker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsMarketStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(&event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMarketsStatHandler handle websocket that push all markets statistics for 24hr
type WsAllMarketsStatHandler func(event WsAllMarketsStatEvent)

// WsAllMarketsStatServe serve websocket that push 24hr statistics for all market every second
func WsAllMarketsStatServe(handler WsAllMarketsStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!ticker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMarketsStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMarketsStatEvent define array of websocket market statistics events
type WsAllMarketsStatEvent []*WsMarketStatEvent

// WsMarketStatEvent define websocket market statistics event
type WsMarketStatEvent struct {
	Event              string `json:"e"`
	Time               int64  `json:"E"`
	Symbol             string `json:"s"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	PrevClosePrice     string `json:"x"`
	LastPrice          string `json:"c"`
	CloseQty           string `json:"Q"`
	BidPrice           string `json:"b"`
	BidQty             string `json:"B"`
	AskPrice           string `json:"a"`
	AskQty             string `json:"A"`
	OpenPrice          string `json:"o"`
	HighPrice          string `json:"h"`
	LowPrice           string `json:"l"`
	BaseVolume         string `json:"v"`
	QuoteVolume        string `json:"q"`
	OpenTime           int64  `json:"O"`
	CloseTime          int64  `json:"C"`
	FirstID            int64  `json:"F"`
	LastID             int64  `json:"L"`
	Count              int64  `json:"n"`
}

// WsAllMiniMarketsStatServeHandler handle websocket that push all mini-ticker market statistics for 24hr
type WsAllMiniMarketsStatServeHandler func(event WsAllMiniMarketsStatEvent)

// WsAllMiniMarketsStatServe serve websocket that push mini version of 24hr statistics for all market every second
func WsAllMiniMarketsStatServe(handler WsAllMiniMarketsStatServeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!miniTicker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMiniMarketsStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMiniMarketsStatEvent define array of websocket market mini-ticker statistics events
type WsAllMiniMarketsStatEvent []*WsMiniMarketsStatEvent

// WsMiniMarketsStatEvent define websocket market mini-ticker statistics event
type WsMiniMarketsStatEvent struct {
	Event       string `json:"e"`
	Time        int64  `json:"E"`
	Symbol      string `json:"s"`
	LastPrice   string `json:"c"`
	OpenPrice   string `json:"o"`
	HighPrice   string `json:"h"`
	LowPrice    string `json:"l"`
	BaseVolume  string `json:"v"`
	QuoteVolume string `json:"q"`
}

// WsBookTickerEvent define websocket best book ticker event.
type WsBookTickerEvent struct {
	UpdateID     int64  `json:"u"`
	Symbol       string `json:"s"`
	BestBidPrice string `json:"b"`
	BestBidQty   string `json:"B"`
	BestAskPrice string `json:"a"`
	BestAskQty   string `json:"A"`
}

type WsCombinedBookTickerEvent struct {
	Data   *WsBookTickerEvent `json:"data"`
	Stream string             `json:"stream"`
}

// WsBookTickerHandler handle websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
type WsBookTickerHandler func(event *WsBookTickerEvent)

// WsBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
func WsBookTickerServe(symbol string, handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@bookTicker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBookTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsCombinedBookTickerServe is similar to WsBookTickerServe, but it is for multiple symbols
func WsCombinedBookTickerServe(symbols []string, handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@bookTicker", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsCombinedBookTickerEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event.Data)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for all symbols.
func WsAllBookTickerServe(handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!bookTicker", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBookTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsApiInitReadWriteConn create and serve connection
func WsApiInitReadWriteConn() (*gorilla.Conn, error) {
	cfg := newWsConfig(getWsApiEndpoint())
	conn, err := WsGetReadWriteConnection(cfg)
	if err != nil {
		return nil, err
	}

	return conn, err
}

// getWsApiEndpoint return the base endpoint of the API WS according the UseTestnet flag
func getWsApiEndpoint() string {
	if UseTestnet {
		return BaseWsApiTestnetURL
	}
	return BaseWsApiMainURL
}
