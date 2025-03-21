package portfolio

import (
	"context"
	"testing"
)

type getMarginOpenOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestGetMarginOpenOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &getMarginOpenOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetOpenOrders_SingleSymbol", func(t *testing.T) {
		symbol := "BNBBTC"
		service := suite.client.NewGetMarginOpenOrdersService()
		orders, err := service.Symbol(symbol).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open orders for symbol: %v", err)
		}

		for _, order := range orders {
			if order.Symbol != symbol {
				t.Errorf("Expected symbol %s, got %s", symbol, order.Symbol)
			}
			if order.Status != "NEW" {
				t.Errorf("Expected status NEW, got %s", order.Status)
			}
		}
	})

	t.Run("GetOpenOrders_AllSymbols", func(t *testing.T) {
		service := suite.client.NewGetMarginOpenOrdersService()
		orders, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all open orders: %v", err)
		}

		for _, order := range orders {
			if order.Status != "NEW" {
				t.Errorf("Expected status NEW, got %s", order.Status)
			}
		}
	})

	t.Run("GetOpenOrders_WithRecvWindow", func(t *testing.T) {
		symbol := "BNBBTC"
		service := suite.client.NewGetMarginOpenOrdersService()
		orders, err := service.Symbol(symbol).
			RecvWindow(5000).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open orders with recvWindow: %v", err)
		}

		for _, order := range orders {
			if order.Symbol != symbol {
				t.Errorf("Expected symbol %s, got %s", symbol, order.Symbol)
			}
		}
	})
}
