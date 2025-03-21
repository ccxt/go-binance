package portfolio

import (
	"context"
	"testing"
)

type umOpenOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMOpenOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umOpenOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetOpenUMOrder_ByOrderID", func(t *testing.T) {
		service := suite.client.NewUMOpenOrderService()
		res, err := service.Symbol("BTCUSDT").
			OrderID(1917641).
			Do(context.Background())
		if err != nil {
			// Check if error is "Order does not exist" which is expected for filled/cancelled orders
			if err.Error() != "Order does not exist" {
				t.Fatalf("Failed to get open UM order: %v", err)
			}
			return
		}

		// Basic validation of returned data
		if res.OrderID != 1917641 {
			t.Errorf("Expected order ID 1917641, got %d", res.OrderID)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
		if res.Status != "NEW" {
			t.Errorf("Expected status NEW, got %s", res.Status)
		}
	})

	t.Run("GetOpenUMOrder_ByClientOrderID", func(t *testing.T) {
		service := suite.client.NewUMOpenOrderService()
		res, err := service.Symbol("BTCUSDT").
			OrigClientOrderID("abc").
			Do(context.Background())
		if err != nil {
			// Check if error is "Order does not exist" which is expected for filled/cancelled orders
			if err.Error() != "Order does not exist" {
				t.Fatalf("Failed to get open UM order: %v", err)
			}
			return
		}

		// Basic validation of returned data
		if res.ClientOrderID != "abc" {
			t.Errorf("Expected client order ID abc, got %s", res.ClientOrderID)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
	})

	t.Run("GetOpenUMOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewUMOpenOrderService()
		_, err := service.Symbol("BTCUSDT").
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
