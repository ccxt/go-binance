package portfolio

import (
	"context"
	"testing"
	"time"
)

type umModifyOrderHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMModifyOrderHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umModifyOrderHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetModifyOrderHistory_ByOrderID", func(t *testing.T) {
		service := suite.client.NewUMModifyOrderHistoryService()
		history, err := service.Symbol("BTCUSDT").
			OrderID(20072994037).
			Limit(500).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get modify order history: %v", err)
		}

		// Validate returned data
		for _, record := range history {
			if record.Symbol != "BTCUSDT" {
				t.Errorf("Expected symbol BTCUSDT, got %s", record.Symbol)
			}
			if record.OrderID != 20072994037 {
				t.Errorf("Expected orderID 20072994037, got %d", record.OrderID)
			}
		}
	})

	t.Run("GetModifyOrderHistory_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 1 day ago

		service := suite.client.NewUMModifyOrderHistoryService()
		history, err := service.Symbol("BTCUSDT").
			StartTime(startTime).
			EndTime(endTime).
			Limit(1000).
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
		service := suite.client.NewUMModifyOrderHistoryService()
		_, err := service.Symbol("BTCUSDT").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected error when neither orderId nor origClientOrderId is provided")
		}
	})
}
