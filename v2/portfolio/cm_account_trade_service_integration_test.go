package portfolio

import (
	"context"
	"testing"
	"time"
)

type cmAccountTradeServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMAccountTradeServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmAccountTradeServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetAccountTrades_BySymbol", func(t *testing.T) {
		service := suite.client.NewCMAccountTradeService()
		trades, err := service.Symbol("BTCUSD_200626").
			Limit(50).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get account trades by symbol: %v", err)
		}

		// Validate returned data
		for _, trade := range trades {
			if trade.Symbol != "BTCUSD_200626" {
				t.Errorf("Expected symbol BTCUSD_200626, got %s", trade.Symbol)
			}
		}
	})

	t.Run("GetAccountTrades_ByPair", func(t *testing.T) {
		service := suite.client.NewCMAccountTradeService()
		trades, err := service.Pair("BTCUSD").
			Limit(50).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get account trades by pair: %v", err)
		}

		// Validate returned data
		for _, trade := range trades {
			if trade.Pair != "BTCUSD" {
				t.Errorf("Expected pair BTCUSD, got %s", trade.Pair)
			}
		}
	})

	t.Run("GetAccountTrades_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago

		service := suite.client.NewCMAccountTradeService()
		trades, err := service.Symbol("BTCUSD_200626").
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
		service := suite.client.NewCMAccountTradeService()
		trades, err := service.Symbol("BTCUSD_200626").
			FromID(6).
			Limit(100).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get account trades from ID: %v", err)
		}

		if len(trades) > 0 {
			if trades[0].ID < 6 {
				t.Errorf("Expected trades with ID >= 6, got %d", trades[0].ID)
			}
		}
	})
}
