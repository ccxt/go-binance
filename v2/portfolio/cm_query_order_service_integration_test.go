package portfolio

import (
	"context"
	"testing"
)

type cmQueryOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMQueryOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmQueryOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("QueryCMOrder_ByOrderID", func(t *testing.T) {
		service := suite.client.NewCMQueryOrderService()
		res, err := service.Symbol("BTCUSD_200925").
			OrderID(1917641).
			Do(context.Background())
		if err != nil {
			// Check if error is "Order does not exist" which is expected for old canceled/expired orders
			if err.Error() != "Order does not exist" {
				t.Fatalf("Failed to query CM order: %v", err)
			}
			return
		}

		// Basic validation of returned data
		if res.OrderID != 1917641 {
			t.Errorf("Expected order ID 1917641, got %d", res.OrderID)
		}
		if res.Symbol != "BTCUSD_200925" {
			t.Errorf("Expected symbol BTCUSD_200925, got %s", res.Symbol)
		}
	})

	t.Run("QueryCMOrder_ByClientOrderID", func(t *testing.T) {
		service := suite.client.NewCMQueryOrderService()
		res, err := service.Symbol("BTCUSD_200925").
			OrigClientOrderID("abc").
			Do(context.Background())
		if err != nil {
			// Check if error is "Order does not exist" which is expected for old canceled/expired orders
			if err.Error() != "Order does not exist" {
				t.Fatalf("Failed to query CM order: %v", err)
			}
			return
		}

		// Basic validation of returned data
		if res.ClientOrderID != "abc" {
			t.Errorf("Expected client order ID abc, got %s", res.ClientOrderID)
		}
		if res.Symbol != "BTCUSD_200925" {
			t.Errorf("Expected symbol BTCUSD_200925, got %s", res.Symbol)
		}
	})

	t.Run("QueryCMOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewCMQueryOrderService()
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
