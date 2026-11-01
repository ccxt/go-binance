package portfolio

import (
	"context"
	"testing"
	"time"
)

type umAllConditionalOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMAllConditionalOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umAllConditionalOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetAllConditionalOrders_SingleSymbol", func(t *testing.T) {
		service := suite.client.NewUMAllConditionalOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all conditional orders: %v", err)
		}

		// Validate returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSDT" {
				t.Errorf("Expected symbol BTCUSDT, got %s", order.Symbol)
			}
		}
	})

	t.Run("GetAllConditionalOrders_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago

		service := suite.client.NewUMAllConditionalOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			StartTime(startTime).
			EndTime(endTime).
			Limit(10).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get conditional orders with time range: %v", err)
		}

		// Validate time range
		for _, order := range orders {
			if order.BookTime < startTime || order.BookTime > endTime {
				t.Errorf("Order time %d outside requested range [%d, %d]",
					order.BookTime, startTime, endTime)
			}
		}
	})

	t.Run("GetAllConditionalOrders_WithStrategyID", func(t *testing.T) {
		service := suite.client.NewUMAllConditionalOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			StrategyID(123445).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get conditional orders with strategyId: %v", err)
		}

		// Validate that returned order matches strategyId if any orders returned
		for _, order := range orders {
			if order.StrategyID != 123445 {
				t.Errorf("Expected strategyId 123445, got %d", order.StrategyID)
			}
		}
	})
}
