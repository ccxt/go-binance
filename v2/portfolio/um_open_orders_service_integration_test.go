package portfolio

import (
	"context"
	"testing"
)

type umOpenOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMOpenOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umOpenOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetOpenUMOrders_SingleSymbol", func(t *testing.T) {
		service := suite.client.NewUMOpenOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open UM orders: %v", err)
		}

		// Validate returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSDT" {
				t.Errorf("Expected symbol BTCUSDT, got %s", order.Symbol)
			}
			if order.Status != "NEW" {
				t.Errorf("Expected status NEW, got %s", order.Status)
			}
		}
	})

	t.Run("GetOpenUMOrders_AllSymbols", func(t *testing.T) {
		service := suite.client.NewUMOpenOrdersService()
		orders, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all open UM orders: %v", err)
		}

		// Validate that all returned orders are open
		for _, order := range orders {
			if order.Status != "NEW" {
				t.Errorf("Expected status NEW, got %s for order %d",
					order.Status, order.OrderID)
			}
		}
	})

	t.Run("GetOpenUMOrders_WithRecvWindow", func(t *testing.T) {
		service := suite.client.NewUMOpenOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			RecvWindow(5000).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open UM orders with recvWindow: %v", err)
		}

		// Basic validation of returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSDT" {
				t.Errorf("Expected symbol BTCUSDT, got %s", order.Symbol)
			}
		}
	})
}
