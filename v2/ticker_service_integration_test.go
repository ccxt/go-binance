package binance

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type tickerServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestTickerServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &tickerServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("ListBookTickers_All", func(t *testing.T) {
		service := suite.client.NewListBookTickersService()
		tickers, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get book tickers: %v", err)
		}

		// Validate returned data
		if len(tickers) == 0 {
			t.Error("Expected at least one book ticker")
		}

		for _, ticker := range tickers {
			if ticker.Symbol == "" {
				t.Error("Expected non-empty symbol")
			}
			if ticker.BidPrice == "" {
				t.Error("Expected non-empty bid price")
			}
			if ticker.BidQuantity == "" {
				t.Error("Expected non-empty bid quantity")
			}
			if ticker.AskPrice == "" {
				t.Error("Expected non-empty ask price")
			}
			if ticker.AskQuantity == "" {
				t.Error("Expected non-empty ask quantity")
			}
		}
	})

	t.Run("ListBookTickers_SingleSymbol", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewListBookTickersService()
		tickers, err := service.Symbol(symbol).Do(context.Background())
		fmt.Printf("tickers: %+v\n", tickers)
		if err != nil {
			t.Fatalf("Failed to get book ticker for %s: %v", symbol, err)
		}

		// Validate returned data
		if len(tickers) != 1 {
			t.Errorf("Expected exactly 1 ticker, got %d", len(tickers))
		}

		if tickers[0].Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, tickers[0].Symbol)
		}
	})

	t.Run("ListPrices_All", func(t *testing.T) {
		service := suite.client.NewListPricesService()
		prices, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get prices: %v", err)
		}

		// Validate returned data
		if len(prices) == 0 {
			t.Error("Expected at least one price")
		}

		for _, price := range prices {
			if price.Symbol == "" {
				t.Error("Expected non-empty symbol")
			}
			if price.Price == "" {
				t.Error("Expected non-empty price")
			}
		}
	})

	t.Run("ListPrices_SingleSymbol", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewListPricesService()
		prices, err := service.Symbol(symbol).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get price for %s: %v", symbol, err)
		}

		// Validate returned data
		if len(prices) != 1 {
			t.Errorf("Expected exactly 1 price, got %d", len(prices))
		}

		if prices[0].Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, prices[0].Symbol)
		}
	})

	t.Run("ListPrices_MultipleSymbols", func(t *testing.T) {
		symbols := []string{"BTCUSDT", "ETHUSDT"}
		service := suite.client.NewListPricesService()
		prices, err := service.Symbols(symbols).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get prices for multiple symbols: %v", err)
		}

		// Validate returned data
		if len(prices) != len(symbols) {
			t.Errorf("Expected %d prices, got %d", len(symbols), len(prices))
		}

		// Check that all requested symbols are present
		symbolMap := make(map[string]bool)
		for _, price := range prices {
			symbolMap[price.Symbol] = true
		}

		for _, symbol := range symbols {
			if !symbolMap[symbol] {
				t.Errorf("Expected symbol %s in response", symbol)
			}
		}
	})

	t.Run("ListPriceChangeStats_All", func(t *testing.T) {
		service := suite.client.NewListPriceChangeStatsService()
		stats, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get price change stats: %v", err)
		}

		// Validate returned data
		if len(stats) == 0 {
			t.Error("Expected at least one price change stat")
		}

		for _, stat := range stats {
			if stat.Symbol == "" {
				t.Error("Expected non-empty symbol")
			}
			// Other fields can be empty (like zero price change), so just check symbol
		}
	})

	t.Run("ListPriceChangeStats_SingleSymbol", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewListPriceChangeStatsService()
		stats, err := service.Symbol(symbol).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get price change stats for %s: %v", symbol, err)
		}

		// Validate returned data
		if len(stats) != 1 {
			t.Errorf("Expected exactly 1 stat, got %d", len(stats))
		}

		if stats[0].Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, stats[0].Symbol)
		}

		// Validate required fields are present
		if stats[0].OpenTime == 0 {
			t.Error("Expected non-zero open time")
		}
		if stats[0].CloseTime == 0 {
			t.Error("Expected non-zero close time")
		}
	})

	t.Run("ListPriceChangeStats_MultipleSymbols", func(t *testing.T) {
		symbols := []string{"BTCUSDT", "ETHUSDT"}
		service := suite.client.NewListPriceChangeStatsService()
		stats, err := service.Symbols(symbols).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get price change stats for multiple symbols: %v", err)
		}

		// Validate returned data
		if len(stats) != len(symbols) {
			t.Errorf("Expected %d stats, got %d", len(symbols), len(stats))
		}

		// Check that all requested symbols are present
		symbolMap := make(map[string]bool)
		for _, stat := range stats {
			symbolMap[stat.Symbol] = true
		}

		for _, symbol := range symbols {
			if !symbolMap[symbol] {
				t.Errorf("Expected symbol %s in response", symbol)
			}
		}
	})

	t.Run("AveragePrice", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewAveragePriceService()
		avgPrice, err := service.Symbol(symbol).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get average price for %s: %v", symbol, err)
		}

		// Validate returned data
		if avgPrice.Mins == 0 {
			t.Error("Expected non-zero mins")
		}
		if avgPrice.Price == "" {
			t.Error("Expected non-empty price")
		}
	})

	t.Run("ListSymbolTicker_SingleSymbol", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewListSymbolTickerService()
		tickers, err := service.Symbol(symbol).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get symbol ticker for %s: %v", symbol, err)
		}

		// Validate returned data
		if len(tickers) != 1 {
			t.Errorf("Expected exactly 1 ticker, got %d", len(tickers))
		}

		ticker := tickers[0]
		if ticker.Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, ticker.Symbol)
		}
		if ticker.OpenTime == 0 {
			t.Error("Expected non-zero open time")
		}
		if ticker.CloseTime == 0 {
			t.Error("Expected non-zero close time")
		}
	})

	t.Run("ListSymbolTicker_WithWindowSize", func(t *testing.T) {
		symbol := "BTCUSDT"
		windowSize := "1h"
		service := suite.client.NewListSymbolTickerService()
		tickers, err := service.Symbol(symbol).WindowSize(windowSize).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get symbol ticker for %s with window size %s: %v", symbol, windowSize, err)
		}

		// Validate returned data
		if len(tickers) != 1 {
			t.Errorf("Expected exactly 1 ticker, got %d", len(tickers))
		}

		ticker := tickers[0]
		if ticker.Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, ticker.Symbol)
		}
	})

	t.Run("ListSymbolTicker_MultipleSymbols", func(t *testing.T) {
		symbols := []string{"BTCUSDT", "ETHUSDT"}
		service := suite.client.NewListSymbolTickerService()
		tickers, err := service.Symbols(symbols).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get symbol tickers for multiple symbols: %v", err)
		}

		// Validate returned data
		if len(tickers) != len(symbols) {
			t.Errorf("Expected %d tickers, got %d", len(symbols), len(tickers))
		}

		// Check that all requested symbols are present
		symbolMap := make(map[string]bool)
		for _, ticker := range tickers {
			symbolMap[ticker.Symbol] = true
		}

		for _, symbol := range symbols {
			if !symbolMap[symbol] {
				t.Errorf("Expected symbol %s in response", symbol)
			}
		}
	})
}

// TestTickerServiceSBEIntegration tests SBE (Simple Binary Encoding) functionality
func TestTickerServiceSBEIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &tickerServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("ListBookTickers_SBE_SingleSymbol", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewListBookTickersService()

		start := time.Now()
		tickers, err := service.Symbol(symbol).DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get book ticker (SBE) for %s: %v", symbol, err)
		}

		// Validate returned data
		if len(tickers) != 1 {
			t.Errorf("Expected exactly 1 ticker, got %d", len(tickers))
		}

		if tickers[0].Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, tickers[0].Symbol)
		}

		// Validate fields are populated
		if tickers[0].BidPrice == "" {
			t.Error("Expected non-empty bid price")
		}
		if tickers[0].AskPrice == "" {
			t.Error("Expected non-empty ask price")
		}

		fmt.Printf("SBE Book Ticker: %+v [Duration: %v]\n", tickers[0], elapsed)
	})

	t.Run("ListBookTickers_SBE_All", func(t *testing.T) {
		service := suite.client.NewListBookTickersService()

		start := time.Now()
		tickers, err := service.DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get all book tickers (SBE): %v", err)
		}

		// Validate returned data
		if len(tickers) == 0 {
			t.Error("Expected at least one book ticker")
		}

		fmt.Printf("SBE Book Tickers count: %d [Duration: %v]\n", len(tickers), elapsed)
	})

	t.Run("ListPrices_SBE_SingleSymbol", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewListPricesService()

		start := time.Now()
		prices, err := service.Symbol(symbol).DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get price (SBE) for %s: %v", symbol, err)
		}

		// Validate returned data
		if len(prices) != 1 {
			t.Errorf("Expected exactly 1 price, got %d", len(prices))
		}

		if prices[0].Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, prices[0].Symbol)
		}

		if prices[0].Price == "" {
			t.Error("Expected non-empty price")
		}

		fmt.Printf("SBE Price: %+v [Duration: %v]\n", prices[0], elapsed)
	})

	t.Run("ListPrices_SBE_MultipleSymbols", func(t *testing.T) {
		symbols := []string{"BTCUSDT", "ETHUSDT"}
		service := suite.client.NewListPricesService()

		start := time.Now()
		prices, err := service.Symbols(symbols).DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get prices (SBE) for multiple symbols: %v", err)
		}

		// Validate returned data
		if len(prices) != len(symbols) {
			t.Errorf("Expected %d prices, got %d", len(symbols), len(prices))
		}

		// Check that all requested symbols are present
		symbolMap := make(map[string]bool)
		for _, price := range prices {
			symbolMap[price.Symbol] = true
		}

		for _, symbol := range symbols {
			if !symbolMap[symbol] {
				t.Errorf("Expected symbol %s in response", symbol)
			}
		}

		fmt.Printf("SBE Prices count: %d [Duration: %v]\n", len(prices), elapsed)
	})

	t.Run("ListPriceChangeStats_SBE_SingleSymbol", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewListPriceChangeStatsService()

		start := time.Now()
		stats, err := service.Symbol(symbol).DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get price change stats (SBE) for %s: %v", symbol, err)
		}

		// Validate returned data
		if len(stats) != 1 {
			t.Errorf("Expected exactly 1 stat, got %d", len(stats))
		}

		if stats[0].Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, stats[0].Symbol)
		}

		// Validate required fields are present
		if stats[0].OpenTime == 0 {
			t.Error("Expected non-zero open time")
		}
		if stats[0].CloseTime == 0 {
			t.Error("Expected non-zero close time")
		}

		fmt.Printf("SBE 24h Stats: Symbol=%s, Volume=%s, PriceChange=%s%% [Duration: %v]\n",
			stats[0].Symbol, stats[0].Volume, stats[0].PriceChangePercent, elapsed)
	})

	t.Run("ListPriceChangeStats_SBE_MultipleSymbols", func(t *testing.T) {
		symbols := []string{"BTCUSDT", "ETHUSDT"}
		service := suite.client.NewListPriceChangeStatsService()

		start := time.Now()
		stats, err := service.Symbols(symbols).DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get price change stats (SBE) for multiple symbols: %v", err)
		}

		// Validate returned data
		if len(stats) != len(symbols) {
			t.Errorf("Expected %d stats, got %d", len(symbols), len(stats))
		}

		// Check that all requested symbols are present
		symbolMap := make(map[string]bool)
		for _, stat := range stats {
			symbolMap[stat.Symbol] = true
		}

		for _, symbol := range symbols {
			if !symbolMap[symbol] {
				t.Errorf("Expected symbol %s in response", symbol)
			}
		}

		fmt.Printf("SBE 24h Stats count: %d [Duration: %v]\n", len(stats), elapsed)
	})

	t.Run("AveragePrice_SBE", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewAveragePriceService()

		start := time.Now()
		avgPrice, err := service.Symbol(symbol).DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get average price (SBE) for %s: %v", symbol, err)
		}

		// Validate returned data
		if avgPrice.Mins == 0 {
			t.Error("Expected non-zero mins")
		}
		if avgPrice.Price == "" {
			t.Error("Expected non-empty price")
		}

		fmt.Printf("SBE Avg Price: Price=%s, Mins=%d [Duration: %v]\n", avgPrice.Price, avgPrice.Mins, elapsed)
	})

	t.Run("ListSymbolTicker_SBE_SingleSymbol", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewListSymbolTickerService()

		start := time.Now()
		tickers, err := service.Symbol(symbol).DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get symbol ticker (SBE) for %s: %v", symbol, err)
		}

		// Validate returned data
		if len(tickers) != 1 {
			t.Errorf("Expected exactly 1 ticker, got %d", len(tickers))
		}

		ticker := tickers[0]
		if ticker.Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, ticker.Symbol)
		}
		if ticker.OpenTime == 0 {
			t.Error("Expected non-zero open time")
		}
		if ticker.CloseTime == 0 {
			t.Error("Expected non-zero close time")
		}

		fmt.Printf("SBE Ticker: Symbol=%s, LastPrice=%s, Volume=%s [Duration: %v]\n",
			ticker.Symbol, ticker.LastPrice, ticker.Volume, elapsed)
	})

	t.Run("ListSymbolTicker_SBE_WithWindowSize", func(t *testing.T) {
		symbol := "BTCUSDT"
		windowSize := "1h"
		service := suite.client.NewListSymbolTickerService()

		start := time.Now()
		tickers, err := service.Symbol(symbol).WindowSize(windowSize).DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get symbol ticker (SBE) for %s with window size %s: %v", symbol, windowSize, err)
		}

		// Validate returned data
		if len(tickers) != 1 {
			t.Errorf("Expected exactly 1 ticker, got %d", len(tickers))
		}

		ticker := tickers[0]
		if ticker.Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, ticker.Symbol)
		}

		fmt.Printf("SBE Ticker (1h window): Symbol=%s, LastPrice=%s [Duration: %v]\n",
			ticker.Symbol, ticker.LastPrice, elapsed)
	})

	t.Run("ListSymbolTicker_SBE_MultipleSymbols", func(t *testing.T) {
		symbols := []string{"BTCUSDT", "ETHUSDT"}
		service := suite.client.NewListSymbolTickerService()

		start := time.Now()
		tickers, err := service.Symbols(symbols).DoSBE(context.Background())
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to get symbol tickers (SBE) for multiple symbols: %v", err)
		}

		// Validate returned data
		if len(tickers) != len(symbols) {
			t.Errorf("Expected %d tickers, got %d", len(symbols), len(tickers))
		}

		// Check that all requested symbols are present
		symbolMap := make(map[string]bool)
		for _, ticker := range tickers {
			symbolMap[ticker.Symbol] = true
		}

		for _, symbol := range symbols {
			if !symbolMap[symbol] {
				t.Errorf("Expected symbol %s in response", symbol)
			}
		}

		fmt.Printf("SBE Tickers count: %d [Duration: %v]\n", len(tickers), elapsed)
	})

	// Comparison test: JSON vs SBE
	t.Run("Comparison_JSON_vs_SBE", func(t *testing.T) {
		symbol := "BTCUSDT"

		// Get using JSON
		startJSON := time.Now()
		jsonTickers, err := suite.client.NewListBookTickersService().
			Symbol(symbol).
			Do(context.Background())
		jsonElapsed := time.Since(startJSON)

		if err != nil {
			t.Fatalf("Failed to get book ticker (JSON): %v", err)
		}

		// Get using SBE
		startSBE := time.Now()
		sbeTickers, err := suite.client.NewListBookTickersService().
			Symbol(symbol).
			DoSBE(context.Background())
		sbeElapsed := time.Since(startSBE)

		if err != nil {
			t.Fatalf("Failed to get book ticker (SBE): %v", err)
		}

		// Compare that both return data for the same symbol
		if len(jsonTickers) != len(sbeTickers) {
			t.Errorf("Different number of results: JSON=%d, SBE=%d",
				len(jsonTickers), len(sbeTickers))
		}

		if jsonTickers[0].Symbol != sbeTickers[0].Symbol {
			t.Errorf("Different symbols: JSON=%s, SBE=%s",
				jsonTickers[0].Symbol, sbeTickers[0].Symbol)
		}

		// Calculate performance difference
		speedup := float64(jsonElapsed) / float64(sbeElapsed)
		var fasterMsg string
		if speedup > 1 {
			fasterMsg = fmt.Sprintf("SBE is %.2fx faster", speedup)
		} else if speedup < 1 {
			fasterMsg = fmt.Sprintf("JSON is %.2fx faster", 1/speedup)
		} else {
			fasterMsg = "Same speed"
		}

		fmt.Printf("\n=== Performance Comparison ===\n")
		fmt.Printf("JSON Result: Symbol=%s, BidPrice=%s [Duration: %v]\n",
			jsonTickers[0].Symbol, jsonTickers[0].BidPrice, jsonElapsed)
		fmt.Printf("SBE Result:  Symbol=%s, BidPrice=%s [Duration: %v]\n",
			sbeTickers[0].Symbol, sbeTickers[0].BidPrice, sbeElapsed)
		fmt.Printf("Performance: %s (diff: %v)\n", fasterMsg, jsonElapsed-sbeElapsed)
		fmt.Printf("==============================\n\n")
	})

	// Benchmark test: Run 10 times each to compare performance
	t.Run("Benchmark_JSON_vs_SBE_10_Iterations", func(t *testing.T) {
		symbol := "BTCUSDT"
		iterations := 50

		fmt.Printf("\n=== Running Benchmark: %d iterations each ===\n", iterations)

		// Benchmark SBE
		var sbeTotalTime time.Duration
		var sbeMinTime time.Duration = time.Duration(1<<63 - 1) // Max duration
		var sbeMaxTime time.Duration
		var sbeResults []*BookTicker
		for i := 0; i < iterations; i++ {
			start := time.Now()
			results, err := suite.client.NewListBookTickersService().
				Symbol(symbol).
				DoSBE(context.Background())
			elapsed := time.Since(start)

			if err != nil {
				t.Fatalf("Failed to get book ticker (SBE) iteration %d: %v", i+1, err)
			}

			sbeTotalTime += elapsed
			if elapsed < sbeMinTime {
				sbeMinTime = elapsed
			}
			if elapsed > sbeMaxTime {
				sbeMaxTime = elapsed
			}
			if i == 0 {
				sbeResults = results
			}
			fmt.Printf("SBE iteration %d: %v\n", i+1, elapsed)
		}
		sbeAvgTime := sbeTotalTime / time.Duration(iterations)

		// Benchmark JSON
		var jsonTotalTime time.Duration
		var jsonMinTime time.Duration = time.Duration(1<<63 - 1) // Max duration
		var jsonMaxTime time.Duration
		var jsonResults []*BookTicker
		for i := 0; i < iterations; i++ {
			start := time.Now()
			results, err := suite.client.NewListBookTickersService().
				Symbol(symbol).
				Do(context.Background())
			elapsed := time.Since(start)

			if err != nil {
				t.Fatalf("Failed to get book ticker (JSON) iteration %d: %v", i+1, err)
			}

			jsonTotalTime += elapsed
			if elapsed < jsonMinTime {
				jsonMinTime = elapsed
			}
			if elapsed > jsonMaxTime {
				jsonMaxTime = elapsed
			}
			if i == 0 {
				jsonResults = results
			}
			fmt.Printf("JSON iteration %d: %v\n", i+1, elapsed)
		}
		jsonAvgTime := jsonTotalTime / time.Duration(iterations)

		// Validate results match
		if len(jsonResults) != len(sbeResults) {
			t.Errorf("Different number of results: JSON=%d, SBE=%d",
				len(jsonResults), len(sbeResults))
		}

		if jsonResults[0].Symbol != sbeResults[0].Symbol {
			t.Errorf("Different symbols: JSON=%s, SBE=%s",
				jsonResults[0].Symbol, sbeResults[0].Symbol)
		}

		// Calculate performance metrics
		speedup := float64(jsonAvgTime) / float64(sbeAvgTime)
		var fasterMsg string
		if speedup > 1 {
			fasterMsg = fmt.Sprintf("SBE is %.2fx faster", speedup)
		} else if speedup < 1 {
			fasterMsg = fmt.Sprintf("JSON is %.2fx faster", 1/speedup)
		} else {
			fasterMsg = "Same speed"
		}

		// Print summary
		fmt.Printf("\n=== Benchmark Results (%d iterations) ===\n", iterations)
		fmt.Printf("JSON:\n")
		fmt.Printf("  Total time: %v\n", jsonTotalTime)
		fmt.Printf("  Average time: %v\n", jsonAvgTime)
		fmt.Printf("  Min time: %v\n", jsonMinTime)
		fmt.Printf("  Max time: %v\n", jsonMaxTime)
		fmt.Printf("  Range: %v\n", jsonMaxTime-jsonMinTime)
		fmt.Printf("  Result: Symbol=%s, BidPrice=%s\n", jsonResults[0].Symbol, jsonResults[0].BidPrice)
		fmt.Printf("\nSBE:\n")
		fmt.Printf("  Total time: %v\n", sbeTotalTime)
		fmt.Printf("  Average time: %v\n", sbeAvgTime)
		fmt.Printf("  Min time: %v\n", sbeMinTime)
		fmt.Printf("  Max time: %v\n", sbeMaxTime)
		fmt.Printf("  Range: %v\n", sbeMaxTime-sbeMinTime)
		fmt.Printf("  Result: Symbol=%s, BidPrice=%s\n", sbeResults[0].Symbol, sbeResults[0].BidPrice)
		fmt.Printf("\nPerformance: %s\n", fasterMsg)
		fmt.Printf("Average difference: %v\n", jsonAvgTime-sbeAvgTime)
		fmt.Printf("Total time saved: %v\n", jsonTotalTime-sbeTotalTime)
		fmt.Printf("========================================\n\n")
	})
}
