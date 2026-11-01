package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/delivery"
	"github.com/adshao/go-binance/v2/futures"
)

func SpotOrder() {
	binance.UseTestnet = true
	apiKey := ""
	secret := ""
	client := binance.NewClient(apiKey, secret)

	symbol := "BTCUSDT"
	side := binance.SideTypeSell
	orderType := binance.OrderTypeMarket
	quantity := "0.0001"

	res, err := client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func FuturesOrder() {
	futures.UseTestnet = true
	apiKey := ""
	secret := ""
	client := binance.NewFuturesClient(apiKey, secret)

	symbol := "LTCUSDT"
	side := futures.SideTypeSell
	orderType := futures.OrderTypeMarket
	quantity := "0.1"

	res, err := client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).PositionSide(futures.PositionSideTypeLong).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

}

func DeliveryOrder() {
	delivery.UseTestnet = true
	apiKey := ""
	secret := ""
	client := binance.NewDeliveryClient(apiKey, secret)

	symbol := "BTCUSD_PERP"
	side := delivery.SideTypeSell
	orderType := delivery.OrderTypeMarket
	quantity := "1"

	res, err := client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).PositionSide(delivery.PositionSideTypeBoth).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

}
