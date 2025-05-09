package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/adshao/go-binance/v2"
)

func WatchMiniMarketsStat() {
	binance.UseTestnet = true

	doneC, stopC, err := binance.WsAllMiniMarketsStatServe(func(event binance.WsAllMiniMarketsStatEvent) {
		fmt.Printf("got %d events\n", len(event))
	}, func(err error) {
		fmt.Printf("%v", err)
	})
	if err != nil {
		fmt.Printf("%v", err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case <-c:
		stopC <- struct{}{}
	}
	<-doneC

}
