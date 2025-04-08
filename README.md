### go-binance

A Golang SDK for [binance](https://accounts.binance.com/register?ref=PGDFCE46) API.

[![Telegram Chat](https://patrolavia.github.io/telegram-badge/chat.png)](https://t.me/gobinancegroup)
[![Build Status](https://travis-ci.org/adshao/go-binance.svg?branch=master)](https://travis-ci.org/adshao/go-binance)
[![GoDoc](https://godoc.org/github.com/adshao/go-binance?status.svg)](https://godoc.org/github.com/adshao/go-binance)
[![Go Report Card](https://goreportcard.com/badge/github.com/adshao/go-binance)](https://goreportcard.com/report/github.com/adshao/go-binance)
[![codecov](https://codecov.io/gh/adshao/go-binance/branch/master/graph/badge.svg)](https://codecov.io/gh/adshao/go-binance)

All the REST APIs listed in [binance API document](https://github.com/binance-exchange/binance-official-api-docs) are implemented, as well as the websocket APIs.

For best compatibility, please use Go >= 1.8.

Make sure you have read binance API document before continuing.


## Community
Join our growing community on Telegram to get help, or just chat!

https://t.me/gobinancegroup

### API List

Name | Description | Status
------------ | ------------ | ------------
[rest-api.md](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md) | Details on the Rest API (/api) | <input type="checkbox" checked> Implemented
[web-socket-streams.md](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md) | Details on available streams and payloads | <input type="checkbox" checked>  Implemented
[user-data-stream.md](https://github.com/binance/binance-spot-api-docs/blob/master/user-data-stream.md) | Details on the dedicated account stream | <input type="checkbox" checked>  Implemented
[margin-api.md](https://binance-docs.github.io/apidocs/spot/en) | Details on the Margin API (/sapi) | <input type="checkbox" checked>  Implemented
[futures-api.md](https://binance-docs.github.io/apidocs/futures/en/#general-info) | Details on the Futures API (/fapi) | <input type="checkbox" checked>  Implemented
[delivery-api.md](https://binance-docs.github.io/apidocs/delivery/en/#general-info) | Details on the Coin-M Futures API (/dapi) | <input type="checkbox" checked>  Implemented
[options-api.md](https://binance-docs.github.io/apidocs/voptions/en/#general-info) | Details on the Options API(/eapi) | <input type="checkbox" checked>  Implemented
[portfolio-margin-api.md](https://developers.binance.com/docs/derivatives/portfolio-margin/general-info) | Details on the Portfolio Margin API(/papi) | <input type="checkbox" checked>  Implemented


If you find an unimplemented interface, please submit an issue. It's great if you can open a PR to fix it.

### Installation

```shell
go get github.com/adshao/go-binance/v2
```

For v1 API, it has been moved to `v1` branch, please use:

```shell
go get github.com/adshao/go-binance/v1
```

### Importing

```golang
import (
    // for spot and other interfaces contained in https://binance-docs.github.io/apidocs/spot/en/#change-log
    "github.com/adshao/go-binance/v2"

    "github.com/adshao/go-binance/v2/futures" // optional package
    "github.com/adshao/go-binance/v2/delivery" // optional package
    "github.com/adshao/go-binance/v2/options" // optional package
)
```

### Documentation

[![GoDoc](https://godoc.org/github.com/adshao/go-binance?status.svg)](https://godoc.org/github.com/adshao/go-binance)

### REST API

#### Setup

Init client for API services. Get APIKey/SecretKey from your binance account.

```golang
var (
    apiKey = "your api key"
    secretKey = "your secret key"
)
client := binance.NewClient(apiKey, secretKey)
futuresClient := binance.NewFuturesClient(apiKey, secretKey)    // USDT-M Futures
deliveryClient := binance.NewDeliveryClient(apiKey, secretKey)  // Coin-M Futures
```

A service instance stands for a REST API endpoint and is initialized by client.NewXXXService function.

Simply call API in chain style. Call Do() in the end to send HTTP request.

Following are some simple examples, please refer to [godoc](https://godoc.org/github.com/adshao/go-binance) for full references.

If you have any questions, please refer to the specific version of the code for specific reference definitions or usage methods

##### Proxy Client

```
proxyUrl := "http://127.0.0.1:7890" // Please replace it with your exact proxy URL.
client := binance.NewProxiedClient(apiKey, apiSecret, proxyUrl)
```


#### Create Order

```golang
order, err := client.NewCreateOrderService().Symbol("BNBETH").
        Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
        TimeInForce(binance.TimeInForceTypeGTC).Quantity("5").
        Price("0.0030000").Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(order)

// Use Test() instead of Do() for testing.
```

#### Get Order

```golang
order, err := client.NewGetOrderService().Symbol("BNBETH").
    OrderID(4432844).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(order)
```

#### Cancel Order

```golang
_, err := client.NewCancelOrderService().Symbol("BNBETH").
    OrderID(4432844).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
```

#### List Open Orders

```golang
openOrders, err := client.NewListOpenOrdersService().Symbol("BNBETH").
    Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, o := range openOrders {
    fmt.Println(o)
}
```

#### List Orders

```golang
orders, err := client.NewListOrdersService().Symbol("BNBETH").
    Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, o := range orders {
    fmt.Println(o)
}
```

#### List Ticker Prices

```golang
prices, err := client.NewListPricesService().Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, p := range prices {
    fmt.Println(p)
}
```

#### Show Depth

```golang
res, err := client.NewDepthService().Symbol("LTCBTC").
    Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
```

#### List Klines

```golang
klines, err := client.NewKlinesService().Symbol("LTCBTC").
    Interval("15m").Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, k := range klines {
    fmt.Println(k)
}
```

#### List Aggregate Trades

```golang
trades, err := client.NewAggTradesService().
    Symbol("LTCBTC").StartTime(1508673256594).EndTime(1508673256595).
    Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, t := range trades {
    fmt.Println(t)
}
```

#### Get Account

```golang
res, err := client.NewGetAccountService().Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
```

#### Start User Stream

```golang
res, err := client.NewStartUserStreamService().Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
```

### Websocket

You don't need Client in websocket API. Just call binance.WsXxxServe(args, handler, errHandler).

> For delivery API you can use `delivery.WsXxxServe(args, handler, errHandler)`.

If you want to use a proxy, you can set `HTTPS_PROXY` or `HTTP_PROXY` in the environment variable, or you can call `SetWsProxyUrl` in the target packages within your code. Then you can call other websocket functions. For example:
```golang
binance.SetWsProxyUrl("http://127.0.0.1:7890")
binance.WsDepthServe("LTCBTC", wsDepthHandler, errHandler)
```

#### Depth

```golang
wsDepthHandler := func(event *binance.WsDepthEvent) {
    fmt.Println(event)
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, stopC, err := binance.WsDepthServe("LTCBTC", wsDepthHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
// use stopC to exit
go func() {
    time.Sleep(5 * time.Second)
    stopC <- struct{}{}
}()
// remove this if you do not want to be blocked here
<-doneC
```

#### Kline

```golang
wsKlineHandler := func(event *binance.WsKlineEvent) {
    fmt.Println(event)
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, _, err := binance.WsKlineServe("LTCBTC", "1m", wsKlineHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
<-doneC
```

#### Aggregate

```golang
wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
    fmt.Println(event)
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, _, err := binance.WsAggTradeServe("LTCBTC", wsAggTradeHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
<-doneC
```

#### User Data

```golang
wsHandler := func(message []byte) {
    fmt.Println(string(message))
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, _, err := binance.WsUserDataServe(listenKey, wsHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
<-doneC
```

#### Setting Server Time

Your system time may be incorrect and you may use following function to set the time offset based off Binance Server Time:

```golang
// use the client future for Futures
client.NewSetServerTimeService().Do(context.Background())
```

Or you can also overwrite the `TimeOffset` yourself:

```golang
client.TimeOffset = 123
```

### Testnet

You can use the testnet by enabling the corresponding flag.

> Note that you can't use your regular API and Secret keys for the testnet. You have to create an account on
> the testnet websites : [https://testnet.binancefuture.com/](https://testnet.binancefuture.com/) for futures and delivery
> or [https://testnet.binance.vision/](https://testnet.binance.vision/) for the Spot Test Network.

#### Spot

Use the `binance.UseTestnet` flag before calling the client creation and the websockets methods.

```go
import (
    "github.com/adshao/go-binance/v2"
)

binance.UseTestnet = true
client := binance.NewClient(apiKey, secretKey)
```

Use the `binance.UseUSDomain` flag to switch domain to binance.us, including both client creation and the wesockets methods.
```go
import (
    "github.com/adshao/go-binance/v2"
)

binance.UseUSDomain = true
client := binance.NewClient(apiKey, secretKey)
```


#### Futures (usd(s)-m futures)

Use the `futures.UseTestnet` flag before calling the client creation and the websockets methods

```go
import (
    "github.com/adshao/go-binance/v2/futures"
)

futures.UseTestnet = true
BinanceClient = futures.NewClient(ApiKey, SecretKey)
```

#### Delivery (coin-m futures)

Use the `delivery.UseTestnet` flag before calling the client creation and the websockets methods

```go
import (
    "github.com/adshao/go-binance/v2/delivery"
)

delivery.UseTestnet = true
BinanceClient = delivery.NewClient(ApiKey, SecretKey)
```

#### Websocket client
##### Order place
##### Async write/read
```go
func main() {
    orderPlaceService, _ := futures.NewOrderPlaceWsService(apiKey, secretKey)
    
    ctx, cancel := context.WithCancel(context.Background())
    
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        select {
            case <-c:
            cancel()
        }
    }()

    request := futures.NewOrderPlaceWsRequest()
    request.
        Symbol("BTCUSDT").
        Side(futures.SideTypeSell).
        Type(futures.OrderTypeLimit).
        Price("68198.00").
        Quantity("0.002").
        TimeInForce(futures.TimeInForceTypeGTC)

    // sender
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                err := orderPlaceService.Do("id", request)
                if err != nil {
                    return
                }
            }
        }
    }()

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go listenOrderPlaceResponse(ctx, wg, orderPlaceService)
    wg.Wait()

    log.Println("exit")
}

func listenOrderPlaceResponse(ctx context.Context, wg *sync.WaitGroup, orderPlaceService *futures.OrderPlaceWsService) {
    defer wg.Done()

    go func() {
        for msg := range orderPlaceService.GetReadChannel() {
            log.Println("order place response", string(msg))
        }
    }()

    go func() {
        for err := range orderPlaceService.GetReadErrorChannel() {
            log.Println("order place error", err)
        }
    }()

    select {
    case <-ctx.Done():
        orderPlaceService.ReceiveAllDataBeforeStop(10 * time.Second)
    }
}
```
##### Sync write/read
```go
func main() {
    orderPlaceService, _ := futures.NewOrderPlaceWsService(apiKey, secretKey)

    id := "some-id"
    request := futures.NewOrderPlaceWsRequest()
    request.
        Symbol("BTCUSDT").
        Side(futures.SideTypeSell).
        Type(futures.OrderTypeLimit).
        Price("68198.00").
        Quantity("0.002").
        TimeInForce(futures.TimeInForceTypeGTC)

    response, err := orderPlaceService.SyncDo(id, request)
    if err != nil {
        log.Fatal(err)
    }

    // handle response
}
```

## Star history

[![Star History Chart](https://api.star-history.com/svg?repos=ccxt/go-binance&type=Date)](https://star-history.com/#ccxt/go-binance&Date)

## Check out some of the other packages

- Check out [CCXT](https://github.com/ccxt/ccxt)  for more than 100 crypto exchanges with a unified trading API in 5 different languages.
- Check out [Python-Binance](https://github.com/sammchardy/python-binance) for a complete Python Wrapper.
- Check out [Node-Binance-API](https://github.com/carlosmiei/node-binance-api) for a node.js sdk.
- Check out [Binance-Trade-Bot](https://github.com/ccxt/binance-trade-bot) for a binance bot in python





