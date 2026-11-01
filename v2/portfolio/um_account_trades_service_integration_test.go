package portfolio

import (
	"context"
	"testing"
	"time"
)

type umAccountTradesServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMAccountTradesServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umAccountTradesServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetTrades_Basic", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewUMAccountTradesService()
		trades, err := service.Symbol(symbol).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get trades: %v", err)
		}

		for _, trade := range trades {
			if trade.Symbol != symbol {
				t.Errorf("Expected symbol %s, got %s", symbol, trade.Symbol)
			}
		}
	})

	t.Run("GetTrades_WithTimeRange", func(t *testing.T) {
		symbol := "BTCUSDT"
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago
		service := suite.client.NewUMAccountTradesService()
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

	t.Run("GetTrades_WithFromID", func(t *testing.T) {
		symbol := "BTCUSDT"
		fromID := int64(67880589)
		service := suite.client.NewUMAccountTradesService()
		trades, err := service.Symbol(symbol).
			FromID(fromID).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get trades from ID: %v", err)
		}

		for _, trade := range trades {
			if trade.ID < fromID {
				t.Errorf("Expected trade ID >= %d, got %d", fromID, trade.ID)
			}
		}
	})

	t.Run("GetTrades_WithLimit", func(t *testing.T) {
		symbol := "BTCUSDT"
		limit := 10
		service := suite.client.NewUMAccountTradesService()
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
