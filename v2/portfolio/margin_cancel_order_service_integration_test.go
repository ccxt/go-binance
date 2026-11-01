package portfolio

import (
	"context"
	"testing"
)

type marginCancelOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginCancelOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginCancelOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelMarginOrder_ByOrderID", func(t *testing.T) {
		service := suite.client.NewMarginCancelOrderService()
		res, err := service.Symbol("LTCBTC").
			OrderID(28).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel margin order: %v", err)
		}

		// Basic validation of returned data
		if res.OrderID != 28 {
			t.Errorf("Expected order ID 28, got %d", res.OrderID)
		}
		if res.Status != "CANCELED" {
			t.Errorf("Expected status CANCELED, got %s", res.Status)
		}
		if res.Symbol != "LTCBTC" {
			t.Errorf("Expected symbol LTCBTC, got %s", res.Symbol)
		}
	})

	t.Run("CancelMarginOrder_ByClientOrderID", func(t *testing.T) {
		service := suite.client.NewMarginCancelOrderService()
		res, err := service.Symbol("LTCBTC").
			OrigClientOrderID("myOrder1").
			NewClientOrderID("cancelMyOrder1").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel margin order: %v", err)
		}

		// Basic validation of returned data
		if res.OrigClientOrderID != "myOrder1" {
			t.Errorf("Expected orig client order ID myOrder1, got %s", res.OrigClientOrderID)
		}
		if res.ClientOrderID != "cancelMyOrder1" {
			t.Errorf("Expected client order ID cancelMyOrder1, got %s", res.ClientOrderID)
		}
		if res.Status != "CANCELED" {
			t.Errorf("Expected status CANCELED, got %s", res.Status)
		}
	})

	t.Run("CancelMarginOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewMarginCancelOrderService()
		_, err := service.Symbol("LTCBTC").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected an error when neither orderId nor origClientOrderId is provided")
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
