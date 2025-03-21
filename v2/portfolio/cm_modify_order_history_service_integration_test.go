package portfolio

import (
	"context"
	"testing"
	"time"
)

type cmModifyOrderHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMModifyOrderHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmModifyOrderHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetModifyOrderHistory_ByOrderID", func(t *testing.T) {
		service := suite.client.NewCMModifyOrderHistoryService()
		history, err := service.Symbol("BTCUSD_PERP").
			OrderID(20072994037).
			Limit(50).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get modify order history: %v", err)
		}

		// Validate returned data
		for _, record := range history {
			if record.Symbol != "BTCUSD_PERP" {
				t.Errorf("Expected symbol BTCUSD_PERP, got %s", record.Symbol)
			}
			if record.Pair != "BTCUSD" {
				t.Errorf("Expected pair BTCUSD, got %s", record.Pair)
			}
		}
	})

	t.Run("GetModifyOrderHistory_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 1 day ago

		service := suite.client.NewCMModifyOrderHistoryService()
		history, err := service.Symbol("BTCUSD_PERP").
			StartTime(startTime).
			EndTime(endTime).
			Limit(100).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get modify order history with time range: %v", err)
		}

		// Validate time range
		for _, record := range history {
			if record.Time < startTime || record.Time > endTime {
				t.Errorf("Modification time %d outside range [%d, %d]",
					record.Time, startTime, endTime)
			}
		}
	})

	t.Run("GetModifyOrderHistory_NoIDProvided", func(t *testing.T) {
		service := suite.client.NewCMModifyOrderHistoryService()
		_, err := service.Symbol("BTCUSD_PERP").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected error when neither orderId nor origClientOrderId is provided")
		}
	})
}
