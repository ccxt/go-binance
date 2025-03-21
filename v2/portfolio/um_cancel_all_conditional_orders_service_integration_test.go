package portfolio

import (
	"context"
	"testing"
)

type umCancelAllConditionalOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMCancelAllConditionalOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umCancelAllConditionalOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelAllUMConditionalOrders", func(t *testing.T) {
		service := suite.client.NewUMCancelAllConditionalOrdersService()
		res, err := service.Symbol("BTCUSDC").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel all UM conditional orders: %v", err)
		}

		// Basic validation of returned data
		if res.Code != "200" {
			t.Errorf("Expected code 200, got %s", res.Code)
		}
		if res.Msg == "" {
			t.Error("Expected non-empty message")
		}
	})

	t.Run("CancelAllUMConditionalOrders_Error_NoSymbol", func(t *testing.T) {
		service := suite.client.NewUMCancelAllConditionalOrdersService()
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
