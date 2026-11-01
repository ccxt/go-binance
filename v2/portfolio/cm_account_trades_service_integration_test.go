package portfolio

import (
	"context"
	"testing"
	"time"
)

type cmAccountTradesServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMAccountTradesServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmAccountTradesServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetTrades_BySymbol", func(t *testing.T) {
		symbol := "BTCUSD_200626"
		service := suite.client.NewCMAccountTradesService()
		trades, err := service.Symbol(symbol).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get trades by symbol: %v", err)
		}

		for _, trade := range trades {
			if trade.Symbol != symbol {
				t.Errorf("Expected symbol %s, got %s", symbol, trade.Symbol)
			}
		}
	})

	t.Run("GetTrades_ByPair", func(t *testing.T) {
		pair := "BTCUSD"
		service := suite.client.NewCMAccountTradesService()
		trades, err := service.Pair(pair).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get trades by pair: %v", err)
		}

		for _, trade := range trades {
			if trade.Pair != pair {
				t.Errorf("Expected pair %s, got %s", pair, trade.Pair)
			}
		}
	})

	t.Run("GetTrades_WithTimeRange", func(t *testing.T) {
		symbol := "BTCUSD_200626"
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago
		service := suite.client.NewCMAccountTradesService()
		trades, err := service.Symbol(symbol).
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get trades with time range: %v", err)
		}

		for _, trade := range trades {
			if trade.Time < startTime || trade.Time > endTime {
				t.Errorf("Trade time %d outside range [%d, %d]",
					trade.Time, startTime, endTime)
			}
		}
	})

	t.Run("GetTrades_WithLimit", func(t *testing.T) {
		symbol := "BTCUSD_200626"
		limit := 10
		service := suite.client.NewCMAccountTradesService()
		trades, err := service.Symbol(symbol).
			Limit(limit).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get trades with limit: %v", err)
		}

		if len(trades) > limit {
			t.Errorf("Expected max %d trades, got %d", limit, len(trades))
		}
	})
}
