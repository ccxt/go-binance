package binance

import (
	"encoding/json"
	"fmt"
	"strings"
)

func WsCombinedDepthAndTradeServe(symbol string, depthHandler WsDepthHandler, tradeHandler WsTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth/%s@trade", getWsEndpoint(), strings.ToLower(symbol), strings.ToLower(symbol))
	return wsCombinedDepthAndTradeServe(endpoint, depthHandler, tradeHandler, errHandler)
}

func wsCombinedDepthAndTradeServe(endpoint string, depthHandler WsDepthHandler, tradeHandler WsTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		eventType := j.Get("e").MustString()
		if eventType == "depthUpdate" {
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
			depthHandler(event)
		} else if eventType == "trade" {
			event := new(WsTradeEvent)
			err := json.Unmarshal(message, event)
			if err != nil {
				errHandler(err)
				return
			}
			tradeHandler(event)
		}
	}
	return wsServe(cfg, wsHandler, errHandler)
}
