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
