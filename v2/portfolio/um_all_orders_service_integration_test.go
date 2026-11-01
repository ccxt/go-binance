package portfolio

import (
	"context"
	"testing"
	"time"
)

type umAllOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMAllOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umAllOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetAllUMOrders_Basic", func(t *testing.T) {
		service := suite.client.NewUMAllOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all UM orders: %v", err)
		}

		if len(orders) == 0 {
			t.Log("No orders found, which is possible")
			return
		}

		// Basic validation of returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSDT" {
				t.Errorf("Expected symbol BTCUSDT, got %s", order.Symbol)
			}
		}
	})

	t.Run("GetAllUMOrders_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago

		service := suite.client.NewUMAllOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			StartTime(startTime).
			EndTime(endTime).
			Limit(10).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all UM orders with time range: %v", err)
		}

		// Validate time range
		for _, order := range orders {
			if order.Time < startTime || order.Time > endTime {
				t.Errorf("Order time %d outside requested range [%d, %d]",
					order.Time, startTime, endTime)
			}
		}
	})

	t.Run("GetAllUMOrders_WithOrderID", func(t *testing.T) {
		service := suite.client.NewUMAllOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			OrderID(1917641).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all UM orders with orderID: %v", err)
		}

		// Validate that all returned orders have orderID >= specified orderID
		for _, order := range orders {
			if order.OrderID < 1917641 {
				t.Errorf("Got order with ID %d, which is less than requested ID 1917641",
					order.OrderID)
			}
		}
	})
}
