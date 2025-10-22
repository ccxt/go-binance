package options

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Endpoints
const (
	baseWsMainUrl          = "wss://nbstream.binance.com/eoptions/ws"
	baseWsTestnetUrl       = "" // unknown now
	baseCombinedMainURL    = "wss://nbstream.binance.com/eoptions/stream?streams="
	baseCombinedTestnetURL = "" // unknown now
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 600
	// WebsocketPongTimeout is an interval for sending a PONG frame in response to PING frame from server
	WebsocketPongTimeout = time.Second * 10
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
	// UseTestnet switch all the WS streams from production to the testnet
	UseTestnet = false

	ProxyUrl = ""
	Header   = make(http.Header)
)

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint() string {
	if UseTestnet {
		return baseWsTestnetUrl
	}
	return baseWsMainUrl
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

// getCombinedEndpoint return the base endpoint of the combined stream according the UseTestnet flag
func getCombinedEndpoint() string {
	if UseTestnet {
		return baseCombinedTestnetURL
	}
	return baseCombinedMainURL
}

type WsTradeEvent struct {
	Event     string `json:"e"`
	Time      int64  `json:"E"`
	Symbol    string `json:"s"`
	TradeId   string `json:"t"`
	Price     string `json:"p"`
	Quantity  string `json:"q"`
	BuyId     string `json:"b"`
	SellId    string `json:"a"`
	TradeTime int64  `json:"T"`
	Side      string `json:"S"` // -1 active sell, 1 active buy
}
type WsTradeHandler func(event *WsTradeEvent)

type WsIndexEvent struct {
	Event  string `json:"e"`
	Time   int64  `json:"E"`
	Symbol string `json:"s"`
	Price  string `json:"p"`
}
type WsIndexHandler func(event *WsIndexEvent)

type WsMarkPriceEvent struct {
	Event     string `json:"e"`
	Time      int64  `json:"E"`
	Symbol    string `json:"s"`
	MarkPrice string `json:"mp"`
}
type WsMarkPriceHandler func(events []*WsMarkPriceEvent)

type WsKline struct {
	StartTime        int64  `json:"t"`
	EndTime          int64  `json:"T"`
	Symbol           string `json:"s"`
	Interval         string `json:"i"`
	FirstTradeID     string `json:"F"`
	LastTradeID      string `json:"L"`
	Open             string `json:"o"`
	Close            string `json:"c"`
	High             string `json:"h"`
	Low              string `json:"l"`
	Volume           string `json:"v"`
	TradeNum         int64  `json:"n"`
	IsFinal          bool   `json:"x"`
	QuoteVolume      string `json:"q"`
	TakerVolume      string `json:"V"`
	TakerQuoteVolume string `json:"Q"`
}
type WsKlineEvent struct {
	Event  string   `json:"e"`
	Time   int64    `json:"E"`
	Symbol string   `json:"s"`
	Kline  *WsKline `json:"k"`
}
type WsKlineHandler func(events *WsKlineEvent)

type WsTickerEvent struct {
	Event              string `json:"e"`
	Time               int64  `json:"E"`
	TradeTime          int64  `json:"T"`
	Symbol             string `json:"s"`
	Open               string `json:"o"`
	High               string `json:"h"`
	Low                string `json:"l"`
	Close              string `json:"c"`
	Volume             string `json:"V"`
	Amount             string `json:"A"`
	PriceChangePercent string `json:"P"`
	PriceChange        string `json:"p"`
	LastQty            string `json:"Q"`
	FirstID            string `json:"F"`
	LastID             string `json:"L"`
	TradeNum           int64  `json:"n"`
	BidOpenPrice       string `json:"bo"`
	AskOpenPrice       string `json:"ao"`
	BidQty             string `json:"bq"`
	AskQty             string `json:"aq"`
	BidIV              string `json:"b"`
	AskIV              string `json:"a"`
	Delta              string `json:"d"`
	Theta              string `json:"t"`
	Gamma              string `json:"g"`
	Vega               string `json:"v"`
	Volatility         string `json:"vo"`
	MarkPrice          string `json:"mp"`
	HighPriceLimit     string `json:"hl"` // highest limit price
	LowPriceLimit      string `json:"ll"`
	ExercisPrice       string `json:"eep"` // Perhaps there is a better one word to express certain meaning
	RiskFreeInterest   string `json:"r"`
}
type WsTickerHandler func(events []*WsTickerEvent)

type WsOpenInterestEvent struct {
	Event        string `json:"e"`
	Time         int64  `json:"E"`
	Symbol       string `json:"s"`
	OpenInterest string `json:"o"` // open interest (in number of sheets)
	Hold         string `json:"h"` // hold (in usdt)
}
type WsOpenInterestHandler func(events []*WsOpenInterestEvent)

type WsOptionPairEvent struct {
	Event        string `json:"e"`
	Time         int64  `json:"E"`
	Id           int    `json:"id"`  // option id
	CId          int    `json:"cid"` // Contract ID
	Underlying   string `json:"u"`
	QuoteAsset   string `json:"qa"`
	Symbol       string `json:"s"`
	Unit         int    `json:"unit"` // The number of targets represented by a contract
	MinQuantity  string `json:"mq"`   // Minimum transaction quantity
	Type         string `json:"d"`    // datatype is better ?
	StrikePrice  string `json:"sp"`
	ExerciseDate int64  `json:"ed"`
}
type WsOptionPairHandler func(events *WsOptionPairEvent)

type PL struct {
	Price    string `json:"b"`
	Quantity string `json:"a"`
}

type WsDepthEvent struct {
	Event            string `json:"e"`
	Time             int64  `json:"E"`
	TransactionTime  int64  `json:"T"`
	Symbol           string `json:"s"`
	LastUpdateID     int64  `json:"u"`
	PrevLastUpdateID int64  `json:"pu"`
	Bids             []Bid  `json:"b"`
	Asks             []Ask  `json:"a"`
}
type WsDepthHandler func(events *WsDepthEvent)

// maybe there's a better one implement with template, but now that's it

func wsTradeParse(message []byte) (*WsTradeEvent, error) {
	event := new(WsTradeEvent)
	err := json.Unmarshal(message, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func wsTradeServeHandler(message []byte, handler WsTradeHandler, errHandler ErrHandler) {
	event, err := wsTradeParse(message)
	if err != nil {
		errHandler(err)
		return
	}
	handler(event)
}

// WsTradeServe serve websocket that push trade information that is aggregated for a single taker order.
func WsTradeServe(symbol string, handler WsTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@trade", getWsEndpoint(), strings.ToUpper(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		wsTradeServeHandler(message, handler, errHandler)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

func wsIndexParse(message []byte) (*WsIndexEvent, error) {
	event := new(WsIndexEvent)
	err := json.Unmarshal(message, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func wsIndexServeHandler(message []byte, handler WsIndexHandler, errHandler ErrHandler) {
	event, err := wsIndexParse(message)
	if err != nil {
		errHandler(err)
		return
	}
	handler(event)
}

// WsIndexServe serve websocket that push trade information that is aggregated for a single taker order.
func WsIndexServe(symbol string, handler WsIndexHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@index", getWsEndpoint(), strings.ToUpper(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		wsIndexServeHandler(message, handler, errHandler)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

func wsMarkPriceParse(message []byte) ([]*WsMarkPriceEvent, error) {
	event := make([]*WsMarkPriceEvent, 0)
	err := json.Unmarshal(message, &event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func wsMarkPriceServeHandler(message []byte, handler WsMarkPriceHandler, errHandler ErrHandler) {
	event, err := wsMarkPriceParse(message)
	if err != nil {
		errHandler(err)
		return
	}
	handler(event)
}

func WsMarkPriceServe(symbol string, handler WsMarkPriceHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@markPrice", getWsEndpoint(), strings.ToUpper(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		wsMarkPriceServeHandler(message, handler, errHandler)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

func wsKlineParse(message []byte) (*WsKlineEvent, error) {
	event := new(WsKlineEvent)
	err := json.Unmarshal(message, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func wsKlineServeHandler(message []byte, handler WsKlineHandler, errHandler ErrHandler) {
	event, err := wsKlineParse(message)
	if err != nil {
		errHandler(err)
		return
	}
	handler(event)
}

func WsKlineServe(symbol string, interval string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@kline_%s", getWsEndpoint(), strings.ToUpper(symbol), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		wsKlineServeHandler(message, handler, errHandler)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

func wsTickerParse(message []byte) ([]*WsTickerEvent, error) {
	event := make([]*WsTickerEvent, 1)
	err := json.Unmarshal(message, &event[0])
	if err != nil {
		return nil, err
	}
	return event, nil
}

func wsTickerServeHandler(message []byte, handler WsTickerHandler, errHandler ErrHandler) {
	event, err := wsTickerParse(message)
	if err != nil {
		errHandler(err)
		return
	}
	handler(event)
}

func WsTickerServe(symbol string, handler WsTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@ticker", getWsEndpoint(), strings.ToUpper(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		wsTickerServeHandler(message, handler, errHandler)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

func wsTickerExpireParse(message []byte) ([]*WsTickerEvent, error) {
	event := make([]*WsTickerEvent, 0)
	err := json.Unmarshal(message, &event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func wsTickerExpireServeHandler(message []byte, handler WsTickerHandler, errHandler ErrHandler) {
	event, err := wsTickerExpireParse(message)
	if err != nil {
		errHandler(err)
		return
	}
	handler(event)
}

// expireDate: for example 220930
// underlying: for example ETH
func WsTickerWithExpireServe(underlying string, expireDate string, handler WsTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@ticker@%s", getWsEndpoint(), strings.ToUpper(underlying), expireDate)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		wsTickerExpireServeHandler(message, handler, errHandler)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

func wsOpenInterestParse(message []byte) ([]*WsOpenInterestEvent, error) {
	event := make([]*WsOpenInterestEvent, 0)
	err := json.Unmarshal(message, &event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func wsOpenInterestServeHandler(message []byte, handler WsOpenInterestHandler, errHandler ErrHandler) {
	event, err := wsOpenInterestParse(message)
	if err != nil {
		errHandler(err)
		return
	}
	handler(event)
}

// expireDate: for example 220930
// underlying: for example ETH
func WsOpenInterestServe(underlying string, expireDate string, handler WsOpenInterestHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@openInterest@%s", getWsEndpoint(), strings.ToUpper(underlying), expireDate)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		wsOpenInterestServeHandler(message, handler, errHandler)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

func wsOptionPairParse(message []byte) (*WsOptionPairEvent, error) {
	event := new(WsOptionPairEvent)
	err := json.Unmarshal(message, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func wsOptionPairServeHandler(message []byte, handler WsOptionPairHandler, errHandler ErrHandler) {
	event, err := wsOptionPairParse(message)
	if err != nil {
		errHandler(err)
		return
	}
	handler(event)
}

func WsOptionPairServe(handler WsOptionPairHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/option_pair", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		wsOptionPairServeHandler(message, handler, errHandler)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

func wsDepthParse(message []byte) (*WsDepthEvent, error) {
	j, err := newJSON(message)
	if err != nil {
		return nil, err
	}
	event := new(WsDepthEvent)
	event.Event = j.Get("e").MustString()
	event.Time = j.Get("E").MustInt64()
	event.TransactionTime = j.Get("T").MustInt64()
	event.Symbol = j.Get("s").MustString()
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
	return event, nil
}

func wsDepthServeHandler(message []byte, handler WsDepthHandler, errHandler ErrHandler) {
	event, err := wsDepthParse(message)
	if err != nil {
		errHandler(err)
		return
	}
	handler(event)
}

// levels: [10, 20, 50, 100, 1000]
// rate: [100, 500, 100] ms, default 500ms while rate is nil
func WsDepthServe(symbol string, levels string, rate *time.Duration, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	switch levels {
	case "10":
	case "20":
	case "50":
	case "100":
	case "1000":
	default:
		return nil, nil, fmt.Errorf("invalid level %s", levels)
	}
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
			return nil, nil, errors.New("invalid rate")
		}
	}
	endpoint := fmt.Sprintf("%s/%s@depth%s%s", getWsEndpoint(), strings.ToUpper(symbol), levels, rateStr)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		wsDepthServeHandler(message, handler, errHandler)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// reference: https://binance-docs.github.io/apidocs/voptions/en/#websocket-market-streams
//
// streamName: you should collaborate stream names through official documentation or other function above defined,
//
//	the legitimacy of parameters needs to be guaranteed by the caller
//
// handler: a map of handler function, its key needs to correspond to the handler of the incoming streamname,
//
//	handler's key should be in ["trade", "index", "markPrice", "kline", "ticker", "openInterest", "option_pair", "depth"]
//
// for example:
//
//	WsCombinedServe({"ETH-240927-5500-P@depth10"}, map[string]interface{}{"depth": func(*WsDepthEvent) {}}, func(error){})
//	WsCombinedServe({"ETH-240927-5500-P@depth10", "ETH-240927-5500-P@kline_1m"},
//				    map[string]interface{}{"depth": func(*WsDepthEvent) {}, "kline": func(*WsKlineEvent){}}, func(error){})
//
// note: the symbol(underlying) of streamName should be upper.
func WsCombinedServe(streamName []string, handler map[string]interface{}, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	if len(streamName) <= 0 || len(handler) <= 0 {
		return nil, nil, errors.New("streamName is empty or handler is empty")
	}
	endpoint := getCombinedEndpoint()
	for _, s := range streamName {
		endpoint += s + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)

	// TODO: use template after go 1.8
	tradeKey := "trade"
	tradeHandler := func(key string, handler map[string]interface{}, errHandler ErrHandler, jsonData []byte, stream string) {
		h, exist := handler[key]
		if !exist {
			errHandler(fmt.Errorf("stream=%s, not found target handler, stream=%s handler=%v", key, stream, handler))
			return
		}
		fn, ok := h.(WsTradeHandler)
		if !ok {
			fn, ok = h.(func(*WsTradeEvent))
			if !ok {
				errHandler(fmt.Errorf("stream=%s, found target handler, but convert to target function type failed, fn=%T", key, h))
				return
			}
		}
		wsTradeServeHandler(jsonData, fn, errHandler)
	}
	indexKey := "index"
	indexHandler := func(key string, handler map[string]interface{}, errHandler ErrHandler, jsonData []byte, stream string) {
		h, exist := handler[key]
		if !exist {
			errHandler(fmt.Errorf("stream=%s, not found target handler, stream=%s handler=%v", key, stream, handler))
			return
		}
		fn, ok := h.(WsIndexHandler)
		if !ok {
			fn, ok = h.(func(*WsIndexEvent))
			if !ok {
				errHandler(fmt.Errorf("stream=%s, found target handler, but convert to target function type failed, fn=%T", key, h))
				return
			}
		}
		wsIndexServeHandler(jsonData, fn, errHandler)
	}
	markPriceKey := "markPrice"
	markPriceHandler := func(key string, handler map[string]interface{}, errHandler ErrHandler, jsonData []byte, stream string) {
		h, exist := handler[key]
		if !exist {
			errHandler(fmt.Errorf("stream=%s, not found target handler, stream=%s handler=%v", key, stream, handler))
			return
		}
		fn, ok := h.(WsMarkPriceHandler)
		if !ok {
			fn, ok = h.(func([]*WsMarkPriceEvent))
			if !ok {
				errHandler(fmt.Errorf("stream=%s, found target handler, but convert to target function type failed, fn=%T", key, h))
				return
			}
		}
		wsMarkPriceServeHandler(jsonData, fn, errHandler)
	}
	klineKey := "kline"
	klineHandler := func(key string, handler map[string]interface{}, errHandler ErrHandler, jsonData []byte, stream string) {
		h, exist := handler[key]
		if !exist {
			errHandler(fmt.Errorf("stream=%s, not found target handler, stream=%s handler=%v", key, stream, handler))
			return
		}
		fn, ok := h.(WsKlineHandler)
		if !ok {
			fn, ok = h.(func(*WsKlineEvent))
			if !ok {
				errHandler(fmt.Errorf("stream=%s, found target handler, but convert to target function type failed, fn=%T", key, h))
				return
			}
		}
		wsKlineServeHandler(jsonData, fn, errHandler)
	}
	tickerKey := "ticker"
	tickerHandler := func(key string, handler map[string]interface{}, errHandler ErrHandler, jsonData []byte, stream string) {
		h, exist := handler[key]
		if !exist {
			errHandler(fmt.Errorf("stream=%s, not found target handler, stream=%s handler=%v", key, stream, handler))
			return
		}
		fn, ok := h.(WsTickerHandler)
		if !ok {
			fn, ok = h.(func([]*WsTickerEvent))
			if !ok {
				errHandler(fmt.Errorf("stream=%s, found target handler, but convert to target function type failed, fn=%T", key, h))
				return
			}
		}
		splitStream := strings.Split(stream, "@")
		if len(splitStream) <= 2 {
			wsTickerServeHandler(jsonData, fn, errHandler)
			return
		}
		wsTickerExpireServeHandler(jsonData, fn, errHandler)
	}
	openInterestKey := "openInterest"
	openInterestHandler := func(key string, handler map[string]interface{}, errHandler ErrHandler, jsonData []byte, stream string) {
		h, exist := handler[key]
		if !exist {
			errHandler(fmt.Errorf("stream=%s, not found target handler, stream=%s handler=%v", key, stream, handler))
			return
		}
		fn, ok := h.(WsOpenInterestHandler)
		if !ok {
			fn, ok = h.(func([]*WsOpenInterestEvent))
			if !ok {
				errHandler(fmt.Errorf("stream=%s, found target handler, but convert to target function type failed, fn=%T", key, h))
				return
			}
		}
		wsOpenInterestServeHandler(jsonData, fn, errHandler)
	}
	optionPairKey := "option_pair"
	optionPairHandler := func(key string, handler map[string]interface{}, errHandler ErrHandler, jsonData []byte, stream string) {
		h, exist := handler[key]
		if !exist {
			errHandler(fmt.Errorf("stream=%s, not found target handler, stream=%s handler=%v", key, stream, handler))
			return
		}
		fn, ok := h.(WsOptionPairHandler)
		if !ok {
			fn, ok = h.(func(*WsOptionPairEvent))
			if !ok {
				errHandler(fmt.Errorf("stream=%s, found target handler, but convert to target function type failed, fn=%T", key, h))
				return
			}
		}
		wsOptionPairServeHandler(jsonData, fn, errHandler)
	}
	depthKey := "depth"
	depthHandler := func(key string, handler map[string]interface{}, errHandler ErrHandler, jsonData []byte, stream string) {
		h, exist := handler[key]
		if !exist {
			errHandler(fmt.Errorf("stream=%s, not found target handler, stream=%s handler=%v", key, stream, handler))
			return
		}
		fn, ok := h.(WsDepthHandler)
		if !ok {
			fn, ok = h.(func(*WsDepthEvent))
			if !ok {
				errHandler(fmt.Errorf("stream=%s, found target handler, but convert to target function type failed, fn=%T", key, h))
				return
			}
		}
		wsDepthServeHandler(jsonData, fn, errHandler)
	}

	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		var jsonData []byte
		data := j.Get("data").MustMap()
		if data == nil {
			jsonData, _ = json.Marshal(j.Get("data").MustArray())
		} else {
			jsonData, _ = json.Marshal(data)
		}

		if strings.Contains(stream, tradeKey) {
			tradeHandler(tradeKey, handler, errHandler, jsonData, stream)
		} else if strings.Contains(stream, indexKey) {
			indexHandler(indexKey, handler, errHandler, jsonData, stream)
		} else if strings.Contains(stream, markPriceKey) {
			markPriceHandler(markPriceKey, handler, errHandler, jsonData, stream)
		} else if strings.Contains(stream, klineKey) {
			klineHandler(klineKey, handler, errHandler, jsonData, stream)
		} else if strings.Contains(stream, tickerKey) {
			tickerHandler(tickerKey, handler, errHandler, jsonData, stream)
		} else if strings.Contains(stream, openInterestKey) {
			openInterestHandler(openInterestKey, handler, errHandler, jsonData, stream)
		} else if strings.Contains(stream, optionPairKey) {
			optionPairHandler(optionPairKey, handler, errHandler, jsonData, stream)
		} else if strings.Contains(stream, depthKey) {
			depthHandler(depthKey, handler, errHandler, jsonData, stream)
		} else {
			errHandler(fmt.Errorf("wsHandler: streamName=%s not found target key in defined key, data=%v", stream, jsonData))
		}
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsUserDataEvent define user data event
type WsUserDataEvent struct {
	Event UserDataEventType `json:"e"`
	Time  int64             `json:"E"`

	RLCStatus            *string `json:"s"` // RLC = RISK_LEVEL_CHANGE
	RLCMarginBalance     *string `json:"mb"`
	RLCMaintenanceMargin *string `json:"mm"`

	AUBalance  []*WsBalance  `json:"B"` // AU = ACCOUNT_UPDATE
	AUGreek    []*WsGreek    `json:"G"`
	AUPosition []*WsPosition `json:"P"`
	AUUid      *int64        `json:"uid"`

	OTU []*WsOrderTradeUpdate `json:"o"` // OTU = ORDER_TRADE_UPDATE
}

// WsBalance define balance
type WsBalance struct {
	Balance           string `json:"b"`
	Merit             string `json:"m"`
	UnrealizedPnL     string `json:"u"`
	MaintenanceMargin string `json:"M"`
	InitialMargin     string `json:"i"`
	Asset             string `json:"a"`
}

type WsGreek struct {
	UnderlyingId string  `json:"ui"`
	Delta        float64 `json:"d"`
	Theta        float64 `json:"t"`
	Gamma        float64 `json:"g"`
	Vega         float64 `json:"v"`
}

// WsPosition define position
type WsPosition struct {
	Symbol      string `json:"s"`
	CountQty    string `json:"c"`
	ReduleQty   string `json:"r"`
	PositionVal string `json:"p"`
	AvgPrice    string `json:"a"`
}

type WsFilled struct {
	TradeId    string `json:"t"`
	Price      string `json:"p"`
	Quantity   string `json:"q"`
	TradedTime int64  `json:"T"`
	Marker     string `json:"m"`
	Fee        string `json:"f"`
}

// WsOrderTradeUpdate define order trade update
type WsOrderTradeUpdate struct {
	CreateTime    int64      `json:"T"`
	UpdateTime    int64      `json:"t"`
	Symbol        string     `json:"s"`
	ClientOrderID string     `json:"c"`
	OrderId       string     `json:"oid"`
	Price         string     `json:"p"`
	Quantity      string     `json:"q"`
	Stp           int        `json:"stp"` //don't know
	ReduleOnly    bool       `json:"r"`
	PostOnly      bool       `json:"po"`
	Status        string     `json:"S"`
	ExecutedQty   string     `json:"e"`
	ExecutedCost  string     `json:"ec"` // maybe there's a better one word
	Fee           string     `json:"f"`
	TimeInForce   string     `json:"tif"`
	OrderType     string     `json:"oty"`
	Filled        []WsFilled `json:"fi"`
}

// WsUserDataHandler handle WsUserDataEvent
type WsUserDataHandler func(event *WsUserDataEvent)

func WsUserDataServe(listenKey string, handler WsUserDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s", getWsEndpoint(), listenKey)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsUserDataEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(fmt.Errorf("err=%v message=%v", err, string(message)))
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}
