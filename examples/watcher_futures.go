package main

import (
	"fmt"
	"os"
	"os/signal"
	"context"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

func WatchFuturesUserDataStream() {
	futures.UseDemo = true
	apiKey := ""
    secret := ""
	client := binance.NewFuturesClient(apiKey, secret)

	listenKey, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	userDataHandler := func(event *futures.WsUserDataEvent) {
    fmt.Printf("Event: %s, Time: %d\n", event.Event, event.Time)
    
    switch event.Event {
	case futures.UserDataEventTypeAlgoUpdate:
		fmt.Printf("ALGO update: %+v\n", event.AlgoUpdate)
	case futures.UserDataEventTypeTradeLite:
		fmt.Printf("Trade lite: %+v\n", event)
	}
	}

	errHandler := func(err error) {
		panic(err)
	}

	doneC, stopC, err := futures.WsUserDataServe(listenKey, userDataHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case <-c:
		stopC <- struct{}{}
	}
	<-doneC
}
