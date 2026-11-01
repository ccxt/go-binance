package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func Ticker() {
	apiKey := ""
	secret := ""
	client := binance.NewClient(apiKey, secret)

	// spot ticker
	ticker, err := client.NewTradingDayTickerService().Symbol("BTCUSDT").Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ticker := range ticker {
		fmt.Println(ticker)
	}

	// futures ticker
	futuresClient := binance.NewFuturesClient(apiKey, secret)
	futuresTicker, err2 := futuresClient.NewListBookTickersService().Symbol("BTCUSDT").Do(context.Background())
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	for _, ticker := range futuresTicker {
		fmt.Println(ticker)
	}

}
