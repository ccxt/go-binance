package portfolio

import (
	"context"
	"testing"
)

type cmCancelAllConditionalOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMCancelAllConditionalOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmCancelAllConditionalOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelAllCMConditionalOrders", func(t *testing.T) {
		service := suite.client.NewCMCancelAllConditionalOrdersService()
		res, err := service.Symbol("LTCBTC").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel all CM conditional orders: %v", err)
		}

		// Basic validation of returned data
		if res.Code != "200" {
			t.Errorf("Expected code 200, got %s", res.Code)
		}
		if res.Msg == "" {
			t.Error("Expected non-empty message")
		}
	})

	t.Run("CancelAllCMConditionalOrders_Error_NoSymbol", func(t *testing.T) {
		service := suite.client.NewCMCancelAllConditionalOrdersService()
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
