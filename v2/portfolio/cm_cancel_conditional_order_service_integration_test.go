package portfolio

import (
	"context"
	"testing"
)

type cmCancelConditionalOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMCancelConditionalOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmCancelConditionalOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelCMConditionalOrder_ByStrategyID", func(t *testing.T) {
		service := suite.client.NewCMCancelConditionalOrderService()
		res, err := service.Symbol("BTCUSD").
			StrategyID(123445).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel CM conditional order: %v", err)
		}

		// Basic validation of returned data
		if res.StrategyID != 123445 {
			t.Errorf("Expected strategy ID 123445, got %d", res.StrategyID)
		}
		if res.StrategyStatus != "CANCELED" {
			t.Errorf("Expected status CANCELED, got %s", res.StrategyStatus)
		}
		if res.Symbol != "BTCUSD" {
			t.Errorf("Expected symbol BTCUSD, got %s", res.Symbol)
		}
	})

	t.Run("CancelCMConditionalOrder_ByClientStrategyID", func(t *testing.T) {
		service := suite.client.NewCMCancelConditionalOrderService()
		res, err := service.Symbol("BTCUSD").
			NewClientStrategyID("myOrder1").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel CM conditional order: %v", err)
		}

		// Basic validation of returned data
		if res.NewClientStrategyID != "myOrder1" {
			t.Errorf("Expected client strategy ID myOrder1, got %s", res.NewClientStrategyID)
		}
		if res.StrategyStatus != "CANCELED" {
			t.Errorf("Expected status CANCELED, got %s", res.StrategyStatus)
		}
	})

	t.Run("CancelCMConditionalOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewCMCancelConditionalOrderService()
		_, err := service.Symbol("BTCUSD").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected an error when neither strategyId nor newClientStrategyId is provided")
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
