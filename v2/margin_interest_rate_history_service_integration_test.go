package binance

import (
	"context"
	"testing"
	"time"
)

type marginInterestRateHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginInterestRateHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginInterestRateHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetMarginInterestRateHistory_Basic", func(t *testing.T) {
		service := suite.client.NewMarginInterestRateHistoryService()
		history, err := service.
			Asset("USDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin interest rate history: %v", err)
		}

		for _, item := range *history {
			if item.Asset == "" {
				t.Error("Expected non-empty asset")
			}
			if item.DailyInterestRate == "" {
				t.Error("Expected non-empty dailyInterestRate value")
			}
		}
	})

	t.Run("GetMarginInterestRateHistory_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 30*24*60*60*1000 // 30 days ago (max allowed range)

		service := suite.client.NewMarginInterestRateHistoryService()
		history, err := service.
			Asset("USDT").
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin interest rate history with time range: %v", err)
		}

		// Validate time range
		for _, item := range *history {
			if item.Timestamp < startTime || item.Timestamp > endTime {
				t.Errorf("Interest record time %d outside requested range [%d, %d]",
					item.Timestamp, startTime, endTime)
			}
		}
	})
}
