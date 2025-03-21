package portfolio

import (
	"context"
	"testing"
	"time"
)

type umAccountTradeServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMAccountTradeServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umAccountTradeServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetAccountTrades_DefaultParams", func(t *testing.T) {
		service := suite.client.NewUMAccountTradeService()
		trades, err := service.Symbol("BTCUSDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get account trades: %v", err)
		}

		// Validate default limit
		if len(trades) > 500 {
			t.Errorf("Expected max 500 trades with default limit, got %d", len(trades))
		}
	})

	t.Run("GetAccountTrades_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago

		service := suite.client.NewUMAccountTradeService()
		trades, err := service.Symbol("BTCUSDT").
			StartTime(startTime).
			EndTime(endTime).
			Limit(1000).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get account trades with time range: %v", err)
		}

		// Validate time range
		for _, trade := range trades {
			if trade.Time < startTime || trade.Time > endTime {
				t.Errorf("Trade time %d outside range [%d, %d]",
					trade.Time, startTime, endTime)
			}
		}
	})

	t.Run("GetAccountTrades_FromID", func(t *testing.T) {
		service := suite.client.NewUMAccountTradeService()
		trades, err := service.Symbol("BTCUSDT").
			FromID(67880589).
			Limit(100).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get account trades from ID: %v", err)
		}

		if len(trades) > 0 {
			if trades[0].ID < 67880589 {
				t.Errorf("Expected trades with ID >= 67880589, got %d", trades[0].ID)
			}
		}
	})
}
