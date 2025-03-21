package portfolio

import (
	"context"
	"testing"
)

type cmOpenOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMOpenOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmOpenOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetOpenCMOrders_SingleSymbol", func(t *testing.T) {
		service := suite.client.NewCMOpenOrdersService()
		orders, err := service.Symbol("BTCUSD_200925").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open CM orders: %v", err)
		}

		// Validate returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSD_200925" {
				t.Errorf("Expected symbol BTCUSD_200925, got %s", order.Symbol)
			}
			if order.Status != "NEW" {
				t.Errorf("Expected status NEW, got %s", order.Status)
			}
		}
	})

	t.Run("GetOpenCMOrders_ByPair", func(t *testing.T) {
		service := suite.client.NewCMOpenOrdersService()
		orders, err := service.Pair("BTCUSD").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open CM orders by pair: %v", err)
		}

		// Validate returned data
		for _, order := range orders {
			if order.Pair != "BTCUSD" {
				t.Errorf("Expected pair BTCUSD, got %s", order.Pair)
			}
			if order.Status != "NEW" {
				t.Errorf("Expected status NEW, got %s", order.Status)
			}
		}
	})

	t.Run("GetOpenCMOrders_AllSymbols", func(t *testing.T) {
		service := suite.client.NewCMOpenOrdersService()
		orders, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all open CM orders: %v", err)
		}

		// Validate that all returned orders are open
		for _, order := range orders {
			if order.Status != "NEW" {
				t.Errorf("Expected status NEW, got %s for order %d",
					order.Status, order.OrderID)
			}
		}
	})

	t.Run("GetOpenCMOrders_WithRecvWindow", func(t *testing.T) {
		service := suite.client.NewCMOpenOrdersService()
		orders, err := service.Symbol("BTCUSD_200925").
			RecvWindow(5000).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open CM orders with recvWindow: %v", err)
		}

		// Basic validation of returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSD_200925" {
				t.Errorf("Expected symbol BTCUSD_200925, got %s", order.Symbol)
			}
		}
	})
}
