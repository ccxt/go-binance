package binance

import (
	"context"
	"testing"
	"time"
)

type marginInterestHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginInterestHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginInterestHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetMarginInterestHistory_Basic", func(t *testing.T) {
		service := suite.client.NewMarginInterestHistoryService()
		history, err := service.
			Asset("USDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin interest history: %v", err)
		}

		// Validate returned data
		if history.Total == 0 {
			t.Log("No interest history records found")
			return
		}

		for _, item := range history.Rows {
			if item.Asset == "" {
				t.Error("Expected non-empty asset")
			}
			if item.Interest == "" {
				t.Error("Expected non-empty interest value")
			}
			if item.InterestRate == "" {
				t.Error("Expected non-empty interest rate")
			}
			if item.Principal == "" {
				t.Error("Expected non-empty principal")
			}
			if item.Type == "" {
				t.Error("Expected non-empty type")
			}
		}
	})

	t.Run("GetMarginInterestHistory_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 30*24*60*60*1000 // 30 days ago (max allowed range)

		service := suite.client.NewMarginInterestHistoryService()
		history, err := service.
			Asset("USDT").
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin interest history with time range: %v", err)
		}

		// Validate time range
		for _, item := range history.Rows {
			if item.InterestAccuredTime < startTime || item.InterestAccuredTime > endTime {
				t.Errorf("Interest record time %d outside requested range [%d, %d]",
					item.InterestAccuredTime, startTime, endTime)
			}
		}
	})

	t.Run("GetMarginInterestHistory_WithPagination", func(t *testing.T) {
		size := int64(10)
		current := int64(1)

		service := suite.client.NewMarginInterestHistoryService()
		history, err := service.
			Asset("USDT").
			Size(size).
			Current(current).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin interest history with pagination: %v", err)
		}

		// Validate pagination
		if int64(len(history.Rows)) > size {
			t.Errorf("Expected maximum %d records, got %d", size, len(history.Rows))
		}
	})

	t.Run("GetMarginInterestHistory_WithIsolatedSymbol", func(t *testing.T) {
		service := suite.client.NewMarginInterestHistoryService()
		history, err := service.
			Asset("USDT").
			IsolatedSymbol("BTCUSDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin interest history for isolated symbol: %v", err)
		}

		// Validate isolated symbol
		for _, item := range history.Rows {
			if item.IsolatedSymbol != "BTCUSDT" {
				t.Errorf("Expected isolated symbol BTCUSDT, got %s", item.IsolatedSymbol)
			}
		}
	})
}
