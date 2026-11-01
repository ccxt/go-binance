package portfolio

import (
	"context"
	"testing"
)

type umQueryOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMQueryOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umQueryOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("QueryUMOrder_ByOrderID", func(t *testing.T) {
		service := suite.client.NewUMQueryOrderService()
		res, err := service.Symbol("BTCUSDT").
			OrderID(1917641).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to query UM order: %v", err)
		}

		// Basic validation of returned data
		if res.OrderID != 1917641 {
			t.Errorf("Expected order ID 1917641, got %d", res.OrderID)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
	})

	t.Run("QueryUMOrder_ByClientOrderID", func(t *testing.T) {
		service := suite.client.NewUMQueryOrderService()
		res, err := service.Symbol("BTCUSDT").
			OrigClientOrderID("abc").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to query UM order: %v", err)
		}

		// Basic validation of returned data
		if res.ClientOrderID != "abc" {
			t.Errorf("Expected client order ID abc, got %s", res.ClientOrderID)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
	})

	t.Run("QueryUMOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewUMQueryOrderService()
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
