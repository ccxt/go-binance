package portfolio

import (
	"context"
	"testing"
)

type umOpenConditionalOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMOpenConditionalOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umOpenConditionalOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetOpenConditionalOrders_SingleSymbol", func(t *testing.T) {
		service := suite.client.NewUMOpenConditionalOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open conditional orders: %v", err)
		}

		// Validate returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSDT" {
				t.Errorf("Expected symbol BTCUSDT, got %s", order.Symbol)
			}
			if order.StrategyStatus != "NEW" {
				t.Errorf("Expected status NEW, got %s", order.StrategyStatus)
			}
		}
	})

	t.Run("GetOpenConditionalOrders_AllSymbols", func(t *testing.T) {
		service := suite.client.NewUMOpenConditionalOrdersService()
		orders, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all open conditional orders: %v", err)
		}

		// Validate that all returned orders are open
		for _, order := range orders {
			if order.StrategyStatus != "NEW" {
				t.Errorf("Expected status NEW, got %s for strategy %d",
					order.StrategyStatus, order.StrategyID)
			}
		}
	})

	t.Run("GetOpenConditionalOrders_WithRecvWindow", func(t *testing.T) {
		service := suite.client.NewUMOpenConditionalOrdersService()
		orders, err := service.Symbol("BTCUSDT").
			RecvWindow(5000).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open conditional orders with recvWindow: %v", err)
		}

		// Basic validation of returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSDT" {
				t.Errorf("Expected symbol BTCUSDT, got %s", order.Symbol)
			}
		}
	})
}
