package binance

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestMarketDataServiceSBEIntegration tests SBE for market data services (Phase 1)
func TestMarketDataServiceSBEIntegration(t *testing.T) {
	suite := SetupTest(t)

	t.Run("Depth_SBE", func(t *testing.T) {
		symbol := "BTCUSDT"
		limit := 100

		// JSON version
		startJSON := time.Now()
		depthJSON, err := suite.client.NewDepthService().
			Symbol(symbol).
			Limit(limit).
			Do(context.Background())
		jsonElapsed := time.Since(startJSON)
		if err != nil {
			t.Fatalf("Failed to get depth (JSON): %v", err)
		}

		// SBE version
		startSBE := time.Now()
		depthSBE, err := suite.client.NewDepthService().
			Symbol(symbol).
			Limit(limit).
			DoSBE(context.Background())
		sbeElapsed := time.Since(startSBE)
		if err != nil {
			t.Fatalf("Failed to get depth (SBE): %v", err)
		}

		// Validate SBE response
		if depthSBE.LastUpdateID == 0 {
			t.Error("Expected non-zero LastUpdateID")
		}
		if len(depthSBE.Bids) == 0 {
			t.Error("Expected at least one bid")
		}
		if len(depthSBE.Asks) == 0 {
			t.Error("Expected at least one ask")
		}

		// Validate bid/ask structure
		if depthSBE.Bids[0].Price == "" {
			t.Error("Expected non-empty bid price")
		}
		if depthSBE.Bids[0].Quantity == "" {
			t.Error("Expected non-empty bid quantity")
		}
		if depthSBE.Asks[0].Price == "" {
			t.Error("Expected non-empty ask price")
		}
		if depthSBE.Asks[0].Quantity == "" {
			t.Error("Expected non-empty ask quantity")
		}

		// Compare results
		if len(depthJSON.Bids) != len(depthSBE.Bids) {
			t.Logf("Warning: Different number of bids: JSON=%d, SBE=%d", len(depthJSON.Bids), len(depthSBE.Bids))
		}
		if len(depthJSON.Asks) != len(depthSBE.Asks) {
			t.Logf("Warning: Different number of asks: JSON=%d, SBE=%d", len(depthJSON.Asks), len(depthSBE.Asks))
		}

		fmt.Printf("Depth (%s): LastUpdateID=%d, Bids=%d, Asks=%d [JSON: %v, SBE: %v]\n",
			symbol, depthSBE.LastUpdateID, len(depthSBE.Bids), len(depthSBE.Asks), jsonElapsed, sbeElapsed)
		fmt.Printf("  Best Bid: %s @ %s\n", depthSBE.Bids[0].Quantity, depthSBE.Bids[0].Price)
		fmt.Printf("  Best Ask: %s @ %s\n", depthSBE.Asks[0].Quantity, depthSBE.Asks[0].Price)
	})

	t.Run("RecentTrades_SBE", func(t *testing.T) {
		symbol := "ETHUSDT"
		limit := 100

		// JSON version
		startJSON := time.Now()
		tradesJSON, err := suite.client.NewRecentTradesService().
			Symbol(symbol).
			Limit(limit).
			Do(context.Background())
		jsonElapsed := time.Since(startJSON)
		if err != nil {
			t.Fatalf("Failed to get recent trades (JSON): %v", err)
		}

		// SBE version
		startSBE := time.Now()
		tradesSBE, err := suite.client.NewRecentTradesService().
			Symbol(symbol).
			Limit(limit).
			DoSBE(context.Background())
		sbeElapsed := time.Since(startSBE)
		if err != nil {
			t.Fatalf("Failed to get recent trades (SBE): %v", err)
		}

		// Validate SBE response
		if len(tradesSBE) == 0 {
			t.Error("Expected at least one trade")
		}

		trade := tradesSBE[0]
		if trade.ID == 0 {
			t.Error("Expected non-zero trade ID")
		}
		if trade.Price == "" {
			t.Error("Expected non-empty price")
		}
		if trade.Quantity == "" {
			t.Error("Expected non-empty quantity")
		}
		if trade.Time == 0 {
			t.Error("Expected non-zero time")
		}

		// Compare results
		if len(tradesJSON) != len(tradesSBE) {
			t.Logf("Warning: Different number of trades: JSON=%d, SBE=%d", len(tradesJSON), len(tradesSBE))
		}

		fmt.Printf("Recent Trades (%s): Count=%d [JSON: %v, SBE: %v]\n",
			symbol, len(tradesSBE), jsonElapsed, sbeElapsed)
		fmt.Printf("  Latest Trade: %s @ %s [ID: %d, Time: %v]\n",
			trade.Quantity, trade.Price, trade.ID, time.UnixMilli(trade.Time))
	})

	t.Run("AggTrades_SBE", func(t *testing.T) {
		symbol := "BNBUSDT"
		limit := 100

		// JSON version
		startJSON := time.Now()
		aggTradesJSON, err := suite.client.NewAggTradesService().
			Symbol(symbol).
			Limit(limit).
			Do(context.Background())
		jsonElapsed := time.Since(startJSON)
		if err != nil {
			t.Fatalf("Failed to get agg trades (JSON): %v", err)
		}

		// SBE version
		startSBE := time.Now()
		aggTradesSBE, err := suite.client.NewAggTradesService().
			Symbol(symbol).
			Limit(limit).
			DoSBE(context.Background())
		sbeElapsed := time.Since(startSBE)
		if err != nil {
			t.Fatalf("Failed to get agg trades (SBE): %v", err)
		}

		// Validate SBE response
		if len(aggTradesSBE) == 0 {
			t.Error("Expected at least one aggregated trade")
		}

		aggTrade := aggTradesSBE[0]
		if aggTrade.AggTradeID == 0 {
			t.Error("Expected non-zero agg trade ID")
		}
		if aggTrade.Price == "" {
			t.Error("Expected non-empty price")
		}
		if aggTrade.Quantity == "" {
			t.Error("Expected non-empty quantity")
		}
		if aggTrade.Timestamp == 0 {
			t.Error("Expected non-zero timestamp")
		}

		// Compare results
		if len(aggTradesJSON) != len(aggTradesSBE) {
			t.Logf("Warning: Different number of agg trades: JSON=%d, SBE=%d", len(aggTradesJSON), len(aggTradesSBE))
		}

		fmt.Printf("Agg Trades (%s): Count=%d [JSON: %v, SBE: %v]\n",
			symbol, len(aggTradesSBE), jsonElapsed, sbeElapsed)
		fmt.Printf("  Latest AggTrade: %s @ %s [ID: %d]\n",
			aggTrade.Quantity, aggTrade.Price, aggTrade.AggTradeID)
	})

	t.Run("AggTrades_SBE_WithTimeRange", func(t *testing.T) {
		symbol := "ADAUSDT"
		endTime := time.Now().UnixMilli()
		startTime := endTime - (60 * 60 * 1000) // 1 hour ago
		limit := 500

		aggTrades, err := suite.client.NewAggTradesService().
			Symbol(symbol).
			StartTime(startTime).
			EndTime(endTime).
			Limit(limit).
			DoSBE(context.Background())
		if err != nil {
			t.Fatalf("Failed to get agg trades with time range (SBE): %v", err)
		}

		// Validate response
		if len(aggTrades) == 0 {
			t.Log("No aggregated trades in the specified time range (this is okay)")
		} else {
			fmt.Printf("Agg Trades with time range (%s): Count=%d\n", symbol, len(aggTrades))
		}
	})

	t.Run("Klines_SBE", func(t *testing.T) {
		symbol := "SOLUSDT"
		interval := "1h"
		limit := 100

		// JSON version
		startJSON := time.Now()
		klinesJSON, err := suite.client.NewKlinesService().
			Symbol(symbol).
			Interval(interval).
			Limit(limit).
			Do(context.Background())
		jsonElapsed := time.Since(startJSON)
		if err != nil {
			t.Fatalf("Failed to get klines (JSON): %v", err)
		}

		// SBE version
		startSBE := time.Now()
		klinesSBE, err := suite.client.NewKlinesService().
			Symbol(symbol).
			Interval(interval).
			Limit(limit).
			DoSBE(context.Background())
		sbeElapsed := time.Since(startSBE)
		if err != nil {
			t.Fatalf("Failed to get klines (SBE): %v", err)
		}

		// Validate SBE response
		if len(klinesSBE) == 0 {
			t.Error("Expected at least one kline")
		}

		kline := klinesSBE[0]
		if kline.OpenTime == 0 {
			t.Error("Expected non-zero open time")
		}
		if kline.Open == "" {
			t.Error("Expected non-empty open price")
		}
		if kline.High == "" {
			t.Error("Expected non-empty high price")
		}
		if kline.Low == "" {
			t.Error("Expected non-empty low price")
		}
		if kline.Close == "" {
			t.Error("Expected non-empty close price")
		}
		if kline.Volume == "" {
			t.Error("Expected non-empty volume")
		}

		// Compare results
		if len(klinesJSON) != len(klinesSBE) {
			t.Logf("Warning: Different number of klines: JSON=%d, SBE=%d", len(klinesJSON), len(klinesSBE))
		}

		fmt.Printf("Klines (%s, %s): Count=%d [JSON: %v, SBE: %v]\n",
			symbol, interval, len(klinesSBE), jsonElapsed, sbeElapsed)
		fmt.Printf("  Latest Candle: O:%s H:%s L:%s C:%s V:%s\n",
			kline.Open, kline.High, kline.Low, kline.Close, kline.Volume)
	})

	t.Run("Klines_SBE_WithTimeRange", func(t *testing.T) {
		symbol := "BTCUSDT"
		interval := "15m"
		endTime := time.Now().UnixMilli()
		startTime := endTime - (24 * 60 * 60 * 1000) // 24 hours ago

		klines, err := suite.client.NewKlinesService().
			Symbol(symbol).
			Interval(interval).
			StartTime(startTime).
			EndTime(endTime).
			DoSBE(context.Background())
		if err != nil {
			t.Fatalf("Failed to get klines with time range (SBE): %v", err)
		}

		// Validate response
		if len(klines) == 0 {
			t.Error("Expected at least one kline in 24h period")
		}

		fmt.Printf("Klines with time range (%s, %s): Count=%d\n", symbol, interval, len(klines))
	})

	// Performance comparison test
	t.Run("Performance_Comparison_MarketData", func(t *testing.T) {
		symbol := "BTCUSDT"
		iterations := 10

		fmt.Printf("\n=== Market Data Performance Comparison (%d iterations) ===\n", iterations)

		// Test Depth
		var depthJSONTotal, depthSBETotal time.Duration
		for i := 0; i < iterations; i++ {
			start := time.Now()
			_, err := suite.client.NewDepthService().Symbol(symbol).Limit(100).Do(context.Background())
			if err == nil {
				depthJSONTotal += time.Since(start)
			}

			start = time.Now()
			_, err = suite.client.NewDepthService().Symbol(symbol).Limit(100).DoSBE(context.Background())
			if err == nil {
				depthSBETotal += time.Since(start)
			}
		}

		// Test Recent Trades
		var tradesJSONTotal, tradesSBETotal time.Duration
		for i := 0; i < iterations; i++ {
			start := time.Now()
			_, err := suite.client.NewRecentTradesService().Symbol(symbol).Limit(100).Do(context.Background())
			if err == nil {
				tradesJSONTotal += time.Since(start)
			}

			start = time.Now()
			_, err = suite.client.NewRecentTradesService().Symbol(symbol).Limit(100).DoSBE(context.Background())
			if err == nil {
				tradesSBETotal += time.Since(start)
			}
		}

		// Test Klines
		var klinesJSONTotal, klinesSBETotal time.Duration
		for i := 0; i < iterations; i++ {
			start := time.Now()
			_, err := suite.client.NewKlinesService().Symbol(symbol).Interval("1h").Limit(100).Do(context.Background())
			if err == nil {
				klinesJSONTotal += time.Since(start)
			}

			start = time.Now()
			_, err = suite.client.NewKlinesService().Symbol(symbol).Interval("1h").Limit(100).DoSBE(context.Background())
			if err == nil {
				klinesSBETotal += time.Since(start)
			}
		}

		// Calculate averages and print results
		fmt.Printf("\nDepth Service:\n")
		fmt.Printf("  JSON Avg: %v\n", depthJSONTotal/time.Duration(iterations))
		fmt.Printf("  SBE Avg:  %v\n", depthSBETotal/time.Duration(iterations))
		fmt.Printf("  Speedup:  %.2fx\n", float64(depthJSONTotal)/float64(depthSBETotal))

		fmt.Printf("\nRecent Trades Service:\n")
		fmt.Printf("  JSON Avg: %v\n", tradesJSONTotal/time.Duration(iterations))
		fmt.Printf("  SBE Avg:  %v\n", tradesSBETotal/time.Duration(iterations))
		fmt.Printf("  Speedup:  %.2fx\n", float64(tradesJSONTotal)/float64(tradesSBETotal))

		fmt.Printf("\nKlines Service:\n")
		fmt.Printf("  JSON Avg: %v\n", klinesJSONTotal/time.Duration(iterations))
		fmt.Printf("  SBE Avg:  %v\n", klinesSBETotal/time.Duration(iterations))
		fmt.Printf("  Speedup:  %.2fx\n", float64(klinesJSONTotal)/float64(klinesSBETotal))

		fmt.Printf("\n======================================================\n\n")
	})
}
