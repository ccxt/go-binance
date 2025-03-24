package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Endpoints
var (
	BaseWsMainUrl    = "wss://dstream.binance.com/ws"
	BaseWsTestnetUrl = "wss://dstream.binancefuture.com/ws"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketPongTimeout is an interval for sending a PONG frame in response to PING frame from server
	WebsocketPongTimeout = time.Second * 10
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
	// UseTestnet switch all the WS streams from production to the testnet
	UseTestnet = false
	ProxyUrl   = ""
)

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint() string {
	if UseTestnet {
		return BaseWsTestnetUrl
	}
	return BaseWsMainUrl
}

func getWsProxyUrl() *string {
	if ProxyUrl == "" {
		return nil
	}
	return &ProxyUrl
}

func SetWsProxyUrl(url string) {
	ProxyUrl = url
}

// WsAggTradeEvent define websocket aggTrde event.
type WsAggTradeEvent struct {
	Event            string `json:"e"`
	Time             int64  `json:"E"`
	AggregateTradeID int64  `json:"a"`
	Symbol           string `json:"s"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int64  `json:"f"`
	LastTradeID      int64  `json:"l"`
	TradeTime        int64  `json:"T"`
	Maker            bool   `json:"m"`
}

// WsAggTradeHandler handle websocket that push trade information that is aggregated for a single taker order.
type WsAggTradeHandler func(event *WsAggTradeEvent)

// WsAggTradeServe serve websocket that push trade information that is aggregated for a single taker order.
func WsAggTradeServe(symbol string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsAggTradeEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsIndexPriceEvent define websocket indexPriceUpdate event.
type WsIndexPriceEvent struct {
	Event      string `json:"e"`
	Time       int64  `json:"E"`
	Pair       string `json:"i"`
	IndexPrice string `json:"p"`
}

// WsIndexPriceHandler handle websocket that push index price for a pair.
type WsIndexPriceHandler func(event *WsIndexPriceEvent)

// WsIndexPriceServe serve websocket that pushes index price for a pair.
func WsIndexPriceServe(symbol string, handler WsIndexPriceHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@indexPrice", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsIndexPriceEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsMarkPriceEvent define websocket markPriceUpdate event.
type WsMarkPriceEvent struct {
	Event                string `json:"e"`
	Time                 int64  `json:"E"`
	Symbol               string `json:"s"`
	MarkPrice            string `json:"p"`
	EstimatedSettlePrice string `json:"P"`
	FundingRate          string `json:"r"`
	NextFundingTime      int64  `json:"T"`
}

// WsMarkPriceHandler handle websocket that pushes price and funding rate for a single symbol.
type WsMarkPriceHandler func(event *WsMarkPriceEvent)

// WsMarkPriceServe serve websocket that pushes price and funding rate for a single symbol.
func WsMarkPriceServe(symbol string, handler WsMarkPriceHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@markPrice", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMarkPriceEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsPairMarkPriceEvent defines an array of websocket markPriceUpdate events.
type WsPairMarkPriceEvent []*WsMarkPriceEvent

// WsPairMarkPriceHandler handle websocket that pushes price and funding rate for all symbol.
type WsPairMarkPriceHandler func(event WsPairMarkPriceEvent)

// WsPairMarkPriceServe serve websocket that pushes price and funding rate for all symbol.
func WsPairMarkPriceServe(handler WsPairMarkPriceHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/markPrice@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsPairMarkPriceEvent
		err := json.Unmarshal(message, &event)
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

// WsKlineHandler handle websocket kline event
type WsKlineHandler func(event *WsKlineEvent)

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

// WsContinuousKlineEvent define websocket continuous kline event
type WsContinuousKlineEvent struct {
	Event        string            `json:"e"`
	Time         int64             `json:"E"`
	Pair         string            `json:"ps"`
	ContractType string            `json:"ct"`
	Kline        WsContinuousKline `json:"k"`
}

// WsContinuousKline define websocket continuous kline
type WsContinuousKline struct {
	StartTime            int64  `json:"t"`
	EndTime              int64  `json:"T"`
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

// WsContinuousKlineHandler handle websocket continuous kline event
type WsContinuousKlineHandler func(event *WsContinuousKlineEvent)

// WsContinuousKlineServe serve websocket kline handler with a pair, a contract type and interval like 15m, 30s
func WsContinuousKlineServe(pair string, contractType string, interval string, handler WsContinuousKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s_%s@continuousKline_%s", getWsEndpoint(), strings.ToLower(pair), strings.ToLower(contractType), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsContinuousKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsIndexPriceKlineEvent define websocket index price kline event
type WsIndexPriceKlineEvent struct {
	Event string            `json:"e"`
	Time  int64             `json:"E"`
	Pair  string            `json:"ps"`
	Kline WsIndexPriceKline `json:"k"`
}

// WsIndexPriceKline define websocket index price kline
type WsIndexPriceKline struct {
	StartTime   int64  `json:"t"`
	EndTime     int64  `json:"T"`
	Interval    string `json:"i"`
	LastTradeId int64  `json:"L"` // LastTradeId Ignore
	Open        string `json:"o"`
	Close       string `json:"c"`
	High        string `json:"h"`
	Low         string `json:"l"`
	TradeNum    int64  `json:"n"`
	IsFinal     bool   `json:"x"`
}

// WsIndexPriceKlineHandler handle websocket index kline event
type WsIndexPriceKlineHandler func(event *WsIndexPriceKlineEvent)

// WsIndexPriceKlineServe serve websocket kline handler with a pair and interval like 15m, 30s
func WsIndexPriceKlineServe(pair string, interval string, handler WsIndexPriceKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@indexPriceKline_%s", getWsEndpoint(), strings.ToLower(pair), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsIndexPriceKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsMarkPriceKlineEvent define websocket market price kline event
type WsMarkPriceKlineEvent struct {
	Event string           `json:"e"`
	Time  int64            `json:"E"`
	Pair  string           `json:"ps"`
	Kline WsMarkPriceKline `json:"k"`
}

// WsMarkPriceKline define websocket market price kline
type WsMarkPriceKline struct {
	StartTime   int64  `json:"t"`
	EndTime     int64  `json:"T"`
	Symbol      string `json:"s"`
	Interval    string `json:"i"`
	LastTradeID int64  `json:"L"` // LastTradeID Ignore
	Open        string `json:"o"`
	Close       string `json:"c"`
	High        string `json:"h"`
	Low         string `json:"l"`
	TradeNum    int64  `json:"n"`
	IsFinal     bool   `json:"x"`
}

// WsMarkPriceKlineHandler handle websocket market price kline event
type WsMarkPriceKlineHandler func(event *WsMarkPriceKlineEvent)

// WsMarkPriceKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func WsMarkPriceKlineServe(symbol string, interval string, handler WsMarkPriceKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@markPriceKline_%s", getWsEndpoint(), strings.ToLower(symbol), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMarkPriceKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsMiniMarketTickerEvent define websocket mini market ticker event.
type WsMiniMarketTickerEvent struct {
	Event       string `json:"e"`
	Time        int64  `json:"E"`
	Symbol      string `json:"s"`
	Pair        string `json:"ps"`
	ClosePrice  string `json:"c"`
	OpenPrice   string `json:"o"`
	HighPrice   string `json:"h"`
	LowPrice    string `json:"l"`
	Volume      string `json:"v"`
	QuoteVolume string `json:"q"`
}

// WsMiniMarketTickerHandler handle websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
type WsMiniMarketTickerHandler func(event *WsMiniMarketTickerEvent)

// WsMiniMarketTickerServe serve websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
func WsMiniMarketTickerServe(symbol string, handler WsMiniMarketTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@miniTicker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMiniMarketTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMiniMarketTickerEvent define an array of websocket mini market ticker events.
type WsAllMiniMarketTickerEvent []*WsMiniMarketTickerEvent

// WsAllMiniMarketTickerHandler handle websocket that pushes price and funding rate for all markets.
type WsAllMiniMarketTickerHandler func(event WsAllMiniMarketTickerEvent)

// WsAllMiniMarketTickerServe serve websocket that pushes price and funding rate for all markets.
func WsAllMiniMarketTickerServe(handler WsAllMiniMarketTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!miniTicker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMiniMarketTickerEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsMarketTickerEvent define websocket market ticker event.
type WsMarketTickerEvent struct {
	Event              string `json:"e"`
	Time               int64  `json:"E"`
	Symbol             string `json:"s"`
	Pair               string `json:"ps"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	ClosePrice         string `json:"c"`
	CloseQty           string `json:"Q"`
	OpenPrice          string `json:"o"`
	HighPrice          string `json:"h"`
	LowPrice           string `json:"l"`
	BaseVolume         string `json:"v"`
	QuoteVolume        string `json:"q"`
	OpenTime           int64  `json:"O"`
	CloseTime          int64  `json:"C"`
	FirstID            int64  `json:"F"`
	LastID             int64  `json:"L"`
	TradeCount         int64  `json:"n"`
}

// WsMarketTickerHandler handle websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
type WsMarketTickerHandler func(event *WsMarketTickerEvent)

// WsMarketTickerServe serve websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
func WsMarketTickerServe(symbol string, handler WsMarketTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@ticker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMarketTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllMarketTickerEvent define an array of websocket mini ticker events.
type WsAllMarketTickerEvent []*WsMarketTickerEvent

// WsAllMarketTickerHandler handle websocket that pushes price and funding rate for all markets.
type WsAllMarketTickerHandler func(event WsAllMarketTickerEvent)

// WsAllMarketTickerServe serve websocket that pushes price and funding rate for all markets.
func WsAllMarketTickerServe(handler WsAllMarketTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!ticker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMarketTickerEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsBookTickerEvent define websocket best book ticker event.
type WsBookTickerEvent struct {
	Event           string `json:"e"`
	UpdateID        int64  `json:"u"`
	Symbol          string `json:"s"`
	Pair            string `json:"ps"`
	BestBidPrice    string `json:"b"`
	BestBidQty      string `json:"B"`
	BestAskPrice    string `json:"a"`
	BestAskQty      string `json:"A"`
	TransactionTime int64  `json:"T"`
	Time            int64  `json:"E"`
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

// WsLiquidationOrderEvent define websocket liquidation order event.
type WsLiquidationOrderEvent struct {
	Event            string             `json:"e"`
	Time             int64              `json:"E"`
	LiquidationOrder WsLiquidationOrder `json:"o"`
}

// WsLiquidationOrder define websocket liquidation order.
type WsLiquidationOrder struct {
	Symbol               string          `json:"s"`
	Pair                 string          `json:"ps"`
	Side                 SideType        `json:"S"`
	OrderType            OrderType       `json:"o"`
	TimeInForce          TimeInForceType `json:"f"`
	OrigQuantity         string          `json:"q"`
	Price                string          `json:"p"`
	AvgPrice             string          `json:"ap"`
	OrderStatus          OrderStatusType `json:"X"`
	LastFilledQty        string          `json:"l"`
	AccumulatedFilledQty string          `json:"z"`
	TradeTime            int64           `json:"T"`
}

// WsLiquidationOrderHandler handle websocket that pushes force liquidation order information for specific symbol.
type WsLiquidationOrderHandler func(event *WsLiquidationOrderEvent)

// WsLiquidationOrderServe serve websocket that pushes force liquidation order information for specific symbol.
func WsLiquidationOrderServe(symbol string, handler WsLiquidationOrderHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@forceOrder", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsLiquidationOrderEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsAllLiquidationOrderServe serve websocket that pushes force liquidation order information for all symbols.
func WsAllLiquidationOrderServe(handler WsLiquidationOrderHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!forceOrder@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsLiquidationOrderEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsDepthEvent define websocket depth book event
type WsDepthEvent struct {
	Event            string `json:"e"`
	Time             int64  `json:"E"`
	TransactionTime  int64  `json:"T"`
	Symbol           string `json:"s"`
	Pair             string `json:"ps"`
	FirstUpdateID    int64  `json:"U"`
	LastUpdateID     int64  `json:"u"`
	PrevLastUpdateID int64  `json:"pu"`
	Bids             []Bid  `json:"b"`
	Asks             []Ask  `json:"a"`
}

// WsDepthHandler handle websocket depth event
type WsDepthHandler func(event *WsDepthEvent)

func wsPartialDepthServe(symbol string, levels int, rate *time.Duration, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	if levels != 5 && levels != 10 && levels != 20 {
		return nil, nil, errors.New("Invalid levels")
	}
	levelsStr := fmt.Sprintf("%d", levels)
	return wsDepthServe(symbol, levelsStr, rate, handler, errHandler)
}

// WsPartialDepthServe serve websocket partial depth handler.
func WsPartialDepthServe(symbol string, levels int, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	return wsPartialDepthServe(symbol, levels, nil, handler, errHandler)
}

// WsPartialDepthServeWithRate serve websocket partial depth handler with rate.
func WsPartialDepthServeWithRate(symbol string, levels int, rate *time.Duration, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	return wsPartialDepthServe(symbol, levels, rate, handler, errHandler)
}

// WsDiffDepthServe serve websocket diff. depth handler.
func WsDiffDepthServe(symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	return wsDepthServe(symbol, "", nil, handler, errHandler)
}

// WsDiffDepthServe serve websocket diff. depth handler with rate.
func WsDiffDepthServeWithRate(symbol string, rate *time.Duration, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	return wsDepthServe(symbol, "", rate, handler, errHandler)
}

func wsDepthServe(symbol string, levels string, rate *time.Duration, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	var rateStr string
	if rate != nil {
		switch *rate {
		case 250 * time.Millisecond:
			rateStr = ""
		case 500 * time.Millisecond:
			rateStr = "@500ms"
		case 100 * time.Millisecond:
			rateStr = "@100ms"
		default:
			return nil, nil, errors.New("Invalid rate")
		}
	}

	endpoint := fmt.Sprintf("%s/%s@depth%s%s", getWsEndpoint(), strings.ToLower(symbol), levels, rateStr)
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
		event.TransactionTime = j.Get("T").MustInt64()
		event.Symbol = j.Get("s").MustString()
		event.Pair = j.Get("ps").MustString()
		event.FirstUpdateID = j.Get("U").MustInt64()
		event.LastUpdateID = j.Get("u").MustInt64()
		event.PrevLastUpdateID = j.Get("pu").MustInt64()
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

// WsUserDataEvent define user data event
type WsUserDataEvent struct {
	Event               UserDataEventType  `json:"e"`
	Time                int64              `json:"E"`
	Alias               string             `json:"i"`
	CrossWalletBalance  string             `json:"cw"`
	MarginCallPositions []WsPosition       `json:"p"`
	TransactionTime     int64              `json:"T"`
	AccountUpdate       WsAccountUpdate    `json:"a"`
	OrderTradeUpdate    WsOrderTradeUpdate `json:"o"`
}

func (e *WsUserDataEvent) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Event               UserDataEventType  `json:"e"`
		Time                interface{}        `json:"E"`
		Alias               string             `json:"i"`
		CrossWalletBalance  string             `json:"cw"`
		MarginCallPositions []WsPosition       `json:"p"`
		TransactionTime     int64              `json:"T"`
		AccountUpdate       WsAccountUpdate    `json:"a"`
		OrderTradeUpdate    WsOrderTradeUpdate `json:"o"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	e.Event = tmp.Event
	switch v := tmp.Time.(type) {
	case float64:
		e.Time = int64(v)
	case string:
		parsedTime, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		e.Time = parsedTime
	default:
		return fmt.Errorf("unexpected type for E: %T", tmp.Time)
	}
	e.Alias = tmp.Alias
	e.CrossWalletBalance = tmp.CrossWalletBalance
	e.MarginCallPositions = tmp.MarginCallPositions
	e.TransactionTime = tmp.TransactionTime
	e.AccountUpdate = tmp.AccountUpdate
	e.OrderTradeUpdate = tmp.OrderTradeUpdate
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
	BalanceChange      string `json:"bc"`
}

// WsPosition define position
type WsPosition struct {
	Symbol                    string           `json:"s"`
	Side                      PositionSideType `json:"ps"`
	Amount                    string           `json:"pa"`
	MarginType                MarginType       `json:"mt"`
	IsolatedWallet            string           `json:"iw"`
	EntryPrice                string           `json:"ep"`
	MarkPrice                 string           `json:"mp"`
	UnrealizedPnL             string           `json:"up"`
	AccumulatedRealized       string           `json:"cr"`
	MaintenanceMarginRequired string           `json:"mm"`
}

// WsOrderTradeUpdate define order trade update
type WsOrderTradeUpdate struct {
	Symbol               string             `json:"s"`
	ClientOrderID        string             `json:"c"`
	Side                 SideType           `json:"S"`
	Type                 OrderType          `json:"o"`
	TimeInForce          TimeInForceType    `json:"f"`
	OriginalQty          string             `json:"q"`
	OriginalPrice        string             `json:"p"`
	AveragePrice         string             `json:"ap"`
	StopPrice            string             `json:"sp"`
	ExecutionType        OrderExecutionType `json:"x"`
	Status               OrderStatusType    `json:"X"`
	ID                   int64              `json:"i"`
	LastFilledQty        string             `json:"l"`
	AccumulatedFilledQty string             `json:"z"`
	LastFilledPrice      string             `json:"L"`
	MarginAsset          string             `json:"ma"`
	CommissionAsset      string             `json:"N"`
	Commission           string             `json:"n"`
	TradeTime            int64              `json:"T"`
	TradeID              int64              `json:"t"`
	RealizedPnL          string             `json:"rp"`
	BidsNotional         string             `json:"b"`
	AsksNotional         string             `json:"a"`
	IsMaker              bool               `json:"m"`
	IsReduceOnly         bool               `json:"R"`
	WorkingType          WorkingType        `json:"wt"`
	OriginalType         OrderType          `json:"ot"`
	PositionSide         PositionSideType   `json:"ps"`
	IsClosingPosition    bool               `json:"cp"`
	ActivationPrice      string             `json:"AP"`
	CallbackRate         string             `json:"cr"`
	IsProtected          bool               `json:"pP"`
}

// WsAccountConfigUpdate define account config update
type WsAccountConfigUpdate struct {
	Symbol   string `json:"s"`
	Leverage int64  `json:"l"`
}

// WsUserDataHandler handle WsUserDataEvent
type WsUserDataHandler func(event *WsUserDataEvent)

// WsUserDataServe serve user data handler with listen key
func WsUserDataServe(listenKey string, handler WsUserDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s", getWsEndpoint(), listenKey)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsUserDataEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}
