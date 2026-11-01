package portfolio

import (
	"context"
	"testing"
)

type cmCancelOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMCancelOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmCancelOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelCMOrder_ByOrderID", func(t *testing.T) {
		service := suite.client.NewCMCancelOrderService()
		res, err := service.Symbol("BTCUSD_200925").
			OrderID(283194212).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel CM order: %v", err)
		}

		// Basic validation of returned data
		if res.OrderID != 283194212 {
			t.Errorf("Expected order ID 283194212, got %d", res.OrderID)
		}
		if res.Status != "CANCELED" {
			t.Errorf("Expected status CANCELED, got %s", res.Status)
		}
		if res.Symbol != "BTCUSD_200925" {
			t.Errorf("Expected symbol BTCUSD_200925, got %s", res.Symbol)
		}
	})

	t.Run("CancelCMOrder_ByClientOrderID", func(t *testing.T) {
		service := suite.client.NewCMCancelOrderService()
		res, err := service.Symbol("BTCUSD_200925").
			OrigClientOrderID("myOrder1").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel CM order: %v", err)
		}

		// Basic validation of returned data
		if res.ClientOrderID != "myOrder1" {
			t.Errorf("Expected client order ID myOrder1, got %s", res.ClientOrderID)
		}
		if res.Status != "CANCELED" {
			t.Errorf("Expected status CANCELED, got %s", res.Status)
		}
	})

	t.Run("CancelCMOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewCMCancelOrderService()
		_, err := service.Symbol("BTCUSD_200925").
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
