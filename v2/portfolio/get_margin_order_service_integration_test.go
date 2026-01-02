package portfolio

import (
	"context"
	"testing"
)

type getMarginOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestGetMarginOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &getMarginOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetMarginOrder_WithOrderID", func(t *testing.T) {
		symbol := "BNBBTC"
		orderID := int64(213205622) // Replace with a valid order ID
		service := suite.client.NewGetMarginOrderService()
		order, err := service.Symbol(symbol).OrderID(orderID).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin orders for orderID: %v", err)
		}

		if order.Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, order.Symbol)
		}
		if order.Status != "NEW" {
			t.Errorf("Expected status NEW, got %s", order.Status)
		}
	})

	t.Run("GetMarginOrder_WithClientOrderID", func(t *testing.T) {
		symbol := "BNBBTC"
		origClientOrderID := "ZwfQzuDIGpceVhKW5DvCmO" // Replace with a valid origClientOrderID
		service := suite.client.NewGetMarginOrderService()
		order, err := service.Symbol(symbol).OrigClientOrderID(origClientOrderID).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin orders for origClientOrderID: %v", err)
		}

		if order.Symbol != symbol {
			t.Errorf("Expected symbol %s, got %s", symbol, order.Symbol)
		}
		if order.Status != "NEW" {
			t.Errorf("Expected status NEW, got %s", order.Status)
		}
	})
}
