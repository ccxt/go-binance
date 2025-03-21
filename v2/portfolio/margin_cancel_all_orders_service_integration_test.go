package portfolio

import (
	"context"
	"testing"
)

type marginCancelAllOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginCancelAllOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginCancelAllOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelAllMarginOrders", func(t *testing.T) {
		service := suite.client.NewMarginCancelAllOrdersService()
		res, err := service.Symbol("BTCUSDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel all margin orders: %v", err)
		}

		// Basic validation of returned data
		if len(res) == 0 {
			return // No orders to cancel is a valid response
		}

		for _, order := range res {
			if order.Symbol != "BTCUSDT" {
				t.Errorf("Expected symbol BTCUSDT, got %s", order.Symbol)
			}
			if order.Status != "CANCELED" && order.ListOrderStatus != "ALL_DONE" {
				t.Errorf("Expected status CANCELED or ALL_DONE, got %s/%s", order.Status, order.ListOrderStatus)
			}
		}
	})

	t.Run("CancelAllMarginOrders_Error_NoSymbol", func(t *testing.T) {
		service := suite.client.NewMarginCancelAllOrdersService()
		_, err := service.Do(context.Background())
		if err == nil {
			t.Fatal("Expected an error when symbol is not provided")
		}

		// Verify it's a Portfolio error
		portfolioErr, ok := err.(*Error)
		if !ok {
			t.Fatalf("Expected Error, got %T", err)
		}
		if portfolioErr.Code != ErrMandatoryParamEmptyOrMalformed {
			t.Errorf("Expected error code %d, got %d", ErrMandatoryParamEmptyOrMalformed, portfolioErr.Code)
		}
	})
}
