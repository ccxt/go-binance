package portfolio

import (
	"context"
	"testing"
	"time"
)

type cmConditionalOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMConditionalOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmConditionalOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetConditionalOrders_SingleSymbol", func(t *testing.T) {
		service := suite.client.NewCMConditionalOrdersService()
		orders, err := service.Symbol("BTCUSD").
			Limit(10).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get conditional orders: %v", err)
		}

		// Validate returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSD" {
				t.Errorf("Expected symbol BTCUSD, got %s", order.Symbol)
			}
		}
	})

	t.Run("GetConditionalOrders_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 7*24*60*60*1000 // 7 days ago

		service := suite.client.NewCMConditionalOrdersService()
		orders, err := service.Symbol("BTCUSD").
			StartTime(startTime).
			EndTime(endTime).
			Limit(500).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get conditional orders with time range: %v", err)
		}

		// Validate time range
		for _, order := range orders {
			if order.BookTime < startTime || order.BookTime > endTime {
				t.Errorf("Order time %d outside range [%d, %d]",
					order.BookTime, startTime, endTime)
			}
		}
	})

	t.Run("GetConditionalOrders_AllSymbols", func(t *testing.T) {
		service := suite.client.NewCMConditionalOrdersService()
		orders, err := service.Limit(500).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all conditional orders: %v", err)
		}

		if len(orders) > 500 {
			t.Errorf("Expected max 500 orders, got %d", len(orders))
		}
	})
}
