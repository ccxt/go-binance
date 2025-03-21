package portfolio

import (
	"context"
	"testing"
	"time"
)

type getMarginAllOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestGetMarginAllOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &getMarginAllOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetAllOrders_Basic", func(t *testing.T) {
		symbol := "BNBBTC"
		service := suite.client.NewGetMarginAllOrdersService()
		orders, err := service.Symbol(symbol).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all orders: %v", err)
		}

		for _, order := range orders {
			if order.Symbol != symbol {
				t.Errorf("Expected symbol %s, got %s", symbol, order.Symbol)
			}
		}
	})

	t.Run("GetAllOrders_WithTimeRange", func(t *testing.T) {
		symbol := "BNBBTC"
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago
		service := suite.client.NewGetMarginAllOrdersService()
		orders, err := service.Symbol(symbol).
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get orders with time range: %v", err)
		}

		for _, order := range orders {
			if order.TransactTime < startTime || order.TransactTime > endTime {
				t.Errorf("Order time %d outside range [%d, %d]",
					order.TransactTime, startTime, endTime)
			}
		}
	})

	t.Run("GetAllOrders_WithLimit", func(t *testing.T) {
		symbol := "BNBBTC"
		limit := 10
		service := suite.client.NewGetMarginAllOrdersService()
		orders, err := service.Symbol(symbol).
			Limit(limit).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get orders with limit: %v", err)
		}

		if len(orders) > limit {
			t.Errorf("Expected max %d orders, got %d", limit, len(orders))
		}
	})

	t.Run("GetAllOrders_WithOrderID", func(t *testing.T) {
		symbol := "BNBBTC"
		orderID := int64(41295) // Replace with a valid order ID
		service := suite.client.NewGetMarginAllOrdersService()
		orders, err := service.Symbol(symbol).
			OrderID(orderID).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get orders from orderID: %v", err)
		}

		for _, order := range orders {
			if order.OrderID < orderID {
				t.Errorf("Expected orderID >= %d, got %d",
					orderID, order.OrderID)
			}
		}
	})
}
