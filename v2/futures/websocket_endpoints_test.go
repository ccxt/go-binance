package futures

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

// Integration tests for the /public, /market, /private WS endpoint routing.
//
// Verifies that each stream category connects and receives data on the correct
// endpoint path for all three environments (main, testnet, demo).
//
// Run:
//   BINANCE_APIKEY=... BINANCE_SECRET=... go test -v -run TestWsEndpoints -count=1 -timeout 120s ./futures/

func requireKeys(t *testing.T) {
	t.Helper()
	if os.Getenv("BINANCE_APIKEY") == "" || os.Getenv("BINANCE_SECRET") == "" {
		t.Skip("BINANCE_APIKEY and BINANCE_SECRET not set")
	}
}

// --- Demo environment tests ---

func TestWsEndpoints_Demo_Market_AggTrade(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAggTradeServe("BTCUSDT", func(e *WsAggTradeEvent) {
		fmt.Printf("  [demo/market] aggTrade symbol=%s price=%s qty=%s\n", e.Symbol, e.Price, e.Quantity)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_Kline(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsKlineServe("BTCUSDT", "1m", func(e *WsKlineEvent) {
		fmt.Printf("  [demo/market] kline symbol=%s open=%s close=%s\n", e.Symbol, e.Kline.Open, e.Kline.Close)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_MarkPrice(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsMarkPriceServe("BTCUSDT", func(e *WsMarkPriceEvent) {
		fmt.Printf("  [demo/market] markPrice symbol=%s price=%s funding=%s\n", e.Symbol, e.MarkPrice, e.FundingRate)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_Ticker(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsMarketTickerServe("BTCUSDT", func(e *WsMarketTickerEvent) {
		fmt.Printf("  [demo/market] ticker symbol=%s last=%s\n", e.Symbol, e.ClosePrice)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_MiniTicker(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsMiniMarketTickerServe("BTCUSDT", func(e *WsMiniMarketTickerEvent) {
		fmt.Printf("  [demo/market] miniTicker symbol=%s close=%s\n", e.Symbol, e.ClosePrice)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_CombinedAggTrade(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedAggTradeServe([]string{"BTCUSDT", "ETHUSDT"}, func(e *WsAggTradeEvent) {
		fmt.Printf("  [demo/market] combined aggTrade symbol=%s price=%s\n", e.Symbol, e.Price)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_CombinedKline(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedKlineServe(map[string]string{"BTCUSDT": "1m", "ETHUSDT": "1m"}, func(e *WsKlineEvent) {
		fmt.Printf("  [demo/market] combined kline symbol=%s\n", e.Symbol)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_BookTicker(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsBookTickerServe("BTCUSDT", func(e *WsBookTickerEvent) {
		fmt.Printf("  [demo/public] bookTicker symbol=%s bid=%s ask=%s\n", e.Symbol, e.BestBidPrice, e.BestAskPrice)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_Depth(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsPartialDepthServe("BTCUSDT", 5, func(e *WsDepthEvent) {
		fmt.Printf("  [demo/public] depth bids=%d asks=%d\n", len(e.Bids), len(e.Asks))
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_DiffDepth(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsDiffDepthServe("BTCUSDT", func(e *WsDepthEvent) {
		fmt.Printf("  [demo/public] diffDepth bids=%d asks=%d\n", len(e.Bids), len(e.Asks))
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_CombinedBookTicker(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedBookTickerServe([]string{"BTCUSDT", "ETHUSDT"}, func(e *WsBookTickerEvent) {
		fmt.Printf("  [demo/public] combined bookTicker symbol=%s\n", e.Symbol)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_CombinedDepth(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedDepthServe(map[string]string{"BTCUSDT": "5", "ETHUSDT": "5"}, func(e *WsDepthEvent) {
		fmt.Printf("  [demo/public] combined depth bids=%d\n", len(e.Bids))
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_CombinedDiffDepth(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedDiffDepthServe([]string{"BTCUSDT", "ETHUSDT"}, func(e *WsDepthEvent) {
		fmt.Printf("  [demo/public] combined diffDepth bids=%d\n", len(e.Bids))
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Private_UserData(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	client := NewClient(os.Getenv("BINANCE_APIKEY"), os.Getenv("BINANCE_SECRET"))
	listenKey, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		t.Fatalf("listen key: %v", err)
	}
	t.Logf("listen key: %s...", listenKey[:8])

	doneC, stopC, err := WsUserDataServe(listenKey, func(e *WsUserDataEvent) {
		fmt.Printf("  [demo/private] userData event\n")
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)

	select {
	case <-doneC:
		t.Fatal("connection closed unexpectedly")
	case <-time.After(3 * time.Second):
		t.Log("PASS: connection held for 3s")
	}
}

// --- Parameterized variant tests (WithRate, MultiInterval, ContinuousKline, All*, etc.) ---

func TestWsEndpoints_Demo_Market_MarkPriceWithRate(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsMarkPriceServeWithRate("BTCUSDT", 3*time.Second, func(e *WsMarkPriceEvent) {
		fmt.Printf("  [demo/market] markPrice@1s symbol=%s price=%s\n", e.Symbol, e.MarkPrice)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_AllMarkPrice(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAllMarkPriceServe(func(events WsAllMarkPriceEvent) {
		if len(events) > 0 {
			fmt.Printf("  [demo/market] !markPrice@arr count=%d first=%s\n", len(events), events[0].Symbol)
		}
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_AllMarkPriceWithRate(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAllMarkPriceServeWithRate(3*time.Second, func(events WsAllMarkPriceEvent) {
		if len(events) > 0 {
			fmt.Printf("  [demo/market] !markPrice@arr@1s count=%d\n", len(events))
		}
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_CombinedMarkPriceWithRate(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedMarkPriceServeWithRate(
		map[string]time.Duration{"BTCUSDT": 3 * time.Second, "ETHUSDT": 3 * time.Second},
		func(e *WsMarkPriceEvent) {
			fmt.Printf("  [demo/market] combined markPrice@1s symbol=%s\n", e.Symbol)
			select {
			case got <- struct{}{}:
			default:
			}
		}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_CombinedKlineMultiInterval(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedKlineServeMultiInterval(
		map[string][]string{"BTCUSDT": {"1m", "5m"}},
		func(e *WsKlineEvent) {
			fmt.Printf("  [demo/market] combined kline multi symbol=%s interval=%s\n", e.Symbol, e.Kline.Interval)
			select {
			case got <- struct{}{}:
			default:
			}
		}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_ContinuousKline(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsContinuousKlineServe(
		&WsContinuousKlineSubscribeArgs{Pair: "BTCUSDT", ContractType: "perpetual", Interval: "1m"},
		func(e *WsContinuousKlineEvent) {
			fmt.Printf("  [demo/market] continuousKline pair=%s interval=%s\n", e.PairSymbol, e.Kline.Interval)
			select {
			case got <- struct{}{}:
			default:
			}
		}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_CombinedContinuousKline(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedContinuousKlineServe(
		[]*WsContinuousKlineSubscribeArgs{
			{Pair: "BTCUSDT", ContractType: "perpetual", Interval: "1m"},
			{Pair: "ETHUSDT", ContractType: "perpetual", Interval: "1m"},
		},
		func(e *WsContinuousKlineEvent) {
			fmt.Printf("  [demo/market] combined continuousKline pair=%s\n", e.PairSymbol)
			select {
			case got <- struct{}{}:
			default:
			}
		}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_AllMiniTicker(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAllMiniMarketTickerServe(func(events WsAllMiniMarketTickerEvent) {
		if len(events) > 0 {
			fmt.Printf("  [demo/market] !miniTicker@arr count=%d\n", len(events))
		}
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_AllTicker(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAllMarketTickerServe(func(events WsAllMarketTickerEvent) {
		if len(events) > 0 {
			fmt.Printf("  [demo/market] !ticker@arr count=%d\n", len(events))
		}
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_AllBookTicker(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAllBookTickerServe(func(e *WsBookTickerEvent) {
		fmt.Printf("  [demo/public] !bookTicker symbol=%s\n", e.Symbol)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_PartialDepthWithRate(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	rate := 500 * time.Millisecond
	doneC, stopC, err := WsPartialDepthServeWithRate("BTCUSDT", 10, rate, func(e *WsDepthEvent) {
		fmt.Printf("  [demo/public] depth@10@500ms bids=%d asks=%d\n", len(e.Bids), len(e.Asks))
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_DiffDepthWithRate(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	rate := 500 * time.Millisecond
	doneC, stopC, err := WsDiffDepthServeWithRate("BTCUSDT", rate, func(e *WsDepthEvent) {
		fmt.Printf("  [demo/public] diffDepth@500ms bids=%d asks=%d\n", len(e.Bids), len(e.Asks))
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_LiquidationOrder(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	// Liquidation events are rare; just verify the connection holds without error
	doneC, stopC, err := WsLiquidationOrderServe("BTCUSDT", func(e *WsLiquidationOrderEvent) {
		fmt.Printf("  [demo/market] forceOrder symbol=%s\n", e.Event)
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(3 * time.Second):
		t.Log("PASS: connection held 3s (liquidation events are rare)")
	}
}

func TestWsEndpoints_Demo_Market_AllLiquidationOrder(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	doneC, stopC, err := WsAllLiquidationOrderServe(func(e *WsLiquidationOrderEvent) {
		fmt.Printf("  [demo/market] !forceOrder@arr\n")
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(3 * time.Second):
		t.Log("PASS: connection held 3s")
	}
}

func TestWsEndpoints_Demo_Market_CompositeIndex(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCompositiveIndexServe("DEFIUSDT", func(e *WsCompositeIndexEvent) {
		fmt.Printf("  [demo/market] compositeIndex symbol=%s price=%s baseAssetType=%s compositions=%d\n",
			e.Symbol, e.Price, e.BaseAssetType, len(e.Composition))
		if len(e.Composition) > 0 {
			fmt.Printf("    first: base=%s quote=%s weight=%s indexPrice=%s\n",
				e.Composition[0].BaseAsset, e.Composition[0].QuoteAsset, e.Composition[0].WeighPercent, e.Composition[0].IndexPrice)
		}
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Fatalf("compositeIndex unmarshal error (should not happen): %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(10 * time.Second):
		t.Fatal("timeout")
	}
}

// --- Testnet environment tests ---

func TestWsEndpoints_Testnet_Market_AggTrade(t *testing.T) {
	requireKeys(t)
	UseTestnet = true
	defer func() { UseTestnet = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAggTradeServe("BTCUSDT", func(e *WsAggTradeEvent) {
		fmt.Printf("  [testnet/market] aggTrade symbol=%s price=%s\n", e.Symbol, e.Price)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Testnet_Market_CombinedMarkPrice(t *testing.T) {
	requireKeys(t)
	UseTestnet = true
	defer func() { UseTestnet = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedMarkPriceServe([]string{"BTCUSDT", "ETHUSDT"}, func(e *WsMarkPriceEvent) {
		fmt.Printf("  [testnet/market] combined markPrice symbol=%s\n", e.Symbol)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Testnet_Public_BookTicker(t *testing.T) {
	requireKeys(t)
	UseTestnet = true
	defer func() { UseTestnet = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsBookTickerServe("BTCUSDT", func(e *WsBookTickerEvent) {
		fmt.Printf("  [testnet/public] bookTicker symbol=%s bid=%s ask=%s\n", e.Symbol, e.BestBidPrice, e.BestAskPrice)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Testnet_Public_CombinedDepth(t *testing.T) {
	requireKeys(t)
	UseTestnet = true
	defer func() { UseTestnet = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsCombinedDepthServe(map[string]string{"BTCUSDT": "5"}, func(e *WsDepthEvent) {
		fmt.Printf("  [testnet/public] combined depth bids=%d\n", len(e.Bids))
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Testnet_Private_UserData(t *testing.T) {
	requireKeys(t)
	UseTestnet = true
	defer func() { UseTestnet = false }()

	client := NewClient(os.Getenv("BINANCE_APIKEY"), os.Getenv("BINANCE_SECRET"))
	listenKey, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		t.Fatalf("listen key: %v (testnet may need different keys)", err)
	}
	t.Logf("listen key: %s...", listenKey[:8])

	doneC, stopC, err := WsUserDataServe(listenKey, func(e *WsUserDataEvent) {
		fmt.Printf("  [testnet/private] userData event\n")
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)

	select {
	case <-doneC:
		t.Fatal("connection closed unexpectedly")
	case <-time.After(3 * time.Second):
		t.Log("PASS: connection held for 3s")
	}
}

// --- New stream tests: contractInfo, assetIndex, private query-param ---

func TestWsEndpoints_Demo_Market_ContractInfo(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	// contractInfo only fires when contract info changes — rare event, just verify connection holds
	doneC, stopC, err := WsContractInfoServe(func(e *WsContractInfoEvent) {
		fmt.Printf("  [demo/market] contractInfo event=%s dataLen=%d\n", e.Event, len(e.Data))
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(3 * time.Second):
		t.Log("PASS: connection held 3s (contractInfo is a rare event)")
	}
}

func TestWsEndpoints_Demo_Market_AssetIndex(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAssetIndexServe("BTCUSD", func(e *WsAssetIndexEvent) {
		fmt.Printf("  [demo/market] assetIndex symbol=%s index=%s bidRate=%s askRate=%s\n",
			e.Symbol, e.Index, e.BidRate, e.AskRate)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Market_AllAssetIndex(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAllAssetIndexServe(func(events WsAllAssetIndexEvent) {
		if len(events) > 0 {
			fmt.Printf("  [demo/market] !assetIndex@arr count=%d first=%s index=%s\n",
				len(events), events[0].Symbol, events[0].Index)
		}
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Private_UserDataWithEvents(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	client := NewClient(os.Getenv("BINANCE_APIKEY"), os.Getenv("BINANCE_SECRET"))
	listenKey, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		t.Fatalf("listen key: %v", err)
	}
	t.Logf("listen key: %s...", listenKey[:8])

	events := []string{"ORDER_TRADE_UPDATE", "ACCOUNT_UPDATE"}
	doneC, stopC, err := WsUserDataServeWithEvents(listenKey, events, func(e *WsUserDataEvent) {
		fmt.Printf("  [demo/private] userData with events: %s\n", e.Event)
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)

	select {
	case <-doneC:
		t.Fatal("connection closed unexpectedly")
	case <-time.After(3 * time.Second):
		t.Log("PASS: connection held 3s with events filter")
	}
}

func TestWsEndpoints_Demo_Private_UserDataMultiple(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	client := NewClient(os.Getenv("BINANCE_APIKEY"), os.Getenv("BINANCE_SECRET"))
	listenKey, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		t.Fatalf("listen key: %v", err)
	}
	t.Logf("listen key: %s...", listenKey[:8])

	configs := []WsPrivateStreamConfig{
		{ListenKey: listenKey, Events: []string{"ORDER_TRADE_UPDATE"}},
		{ListenKey: listenKey, Events: []string{"ACCOUNT_UPDATE"}},
	}
	doneC, stopC, err := WsUserDataServeMultiple(configs, func(e *WsUserDataEvent) {
		fmt.Printf("  [demo/private] userData multiple: %s\n", e.Event)
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)

	select {
	case <-doneC:
		t.Fatal("connection closed unexpectedly")
	case <-time.After(3 * time.Second):
		t.Log("PASS: stream mode connection held 3s with multiple configs")
	}
}

func TestWsEndpoints_Testnet_Market_AssetIndex(t *testing.T) {
	requireKeys(t)
	UseTestnet = true
	defer func() { UseTestnet = false }()

	got := make(chan struct{}, 1)
	doneC, stopC, err := WsAssetIndexServe("BTCUSD", func(e *WsAssetIndexEvent) {
		fmt.Printf("  [testnet/market] assetIndex symbol=%s index=%s\n", e.Symbol, e.Index)
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Testnet_Private_UserDataWithEvents(t *testing.T) {
	requireKeys(t)
	UseTestnet = true
	defer func() { UseTestnet = false }()

	client := NewClient(os.Getenv("BINANCE_APIKEY"), os.Getenv("BINANCE_SECRET"))
	listenKey, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		t.Fatalf("listen key: %v", err)
	}
	t.Logf("listen key: %s...", listenKey[:8])

	events := []string{"ORDER_TRADE_UPDATE", "ACCOUNT_UPDATE"}
	doneC, stopC, err := WsUserDataServeWithEvents(listenKey, events, func(e *WsUserDataEvent) {
		fmt.Printf("  [testnet/private] userData with events: %s\n", e.Event)
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)

	select {
	case <-doneC:
		t.Fatal("connection closed unexpectedly")
	case <-time.After(3 * time.Second):
		t.Log("PASS: connection held 3s")
	}
}

// --- BLVT and rate variant tests ---

func TestWsEndpoints_Demo_Market_BLVTInfo(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	// BLVT tokens may not be active on demo; just verify connection holds
	doneC, stopC, err := WsBLVTInfoServe("BTCDOWN", func(e *WsBLVTInfoEvent) {
		fmt.Printf("  [demo/market] BLVT tokenNav symbol=%s nav=%f leverage=%f\n",
			e.Symbol, e.Nav, e.Leverage)
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(3 * time.Second):
		t.Log("PASS: connection held 3s (BLVT may not be active on demo)")
	}
}

func TestWsEndpoints_Demo_Market_BLVTKline(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	doneC, stopC, err := WsBLVTKlineServe("BTCDOWN", "1m", func(e *WsBLVTKlineEvent) {
		fmt.Printf("  [demo/market] BLVT kline symbol=%s open=%s close=%s\n",
			e.Symbol, e.Kline.OpenPrice, e.Kline.ClosePrice)
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(3 * time.Second):
		t.Log("PASS: connection held 3s (BLVT may not be active on demo)")
	}
}

func TestWsEndpoints_Demo_Public_PartialDepthWith100ms(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	rate := 100 * time.Millisecond
	doneC, stopC, err := WsPartialDepthServeWithRate("BTCUSDT", 5, rate, func(e *WsDepthEvent) {
		fmt.Printf("  [demo/public] depth@5@100ms bids=%d asks=%d\n", len(e.Bids), len(e.Asks))
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}

func TestWsEndpoints_Demo_Public_DiffDepthWith100ms(t *testing.T) {
	requireKeys(t)
	UseDemo = true
	defer func() { UseDemo = false }()

	got := make(chan struct{}, 1)
	rate := 100 * time.Millisecond
	doneC, stopC, err := WsDiffDepthServeWithRate("BTCUSDT", rate, func(e *WsDepthEvent) {
		fmt.Printf("  [demo/public] diffDepth@100ms bids=%d asks=%d\n", len(e.Bids), len(e.Asks))
		select {
		case got <- struct{}{}:
		default:
		}
	}, func(err error) { t.Logf("err: %v", err) })
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer close(stopC)
	select {
	case <-got:
		t.Log("PASS")
	case <-doneC:
		t.Fatal("closed")
	case <-time.After(15 * time.Second):
		t.Fatal("timeout")
	}
}
