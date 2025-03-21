package portfolio

import (
	"context"
	"testing"
)

type umCancelConditionalOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMCancelConditionalOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umCancelConditionalOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelUMConditionalOrder_ByStrategyID", func(t *testing.T) {
		service := suite.client.NewUMCancelConditionalOrderService()
		res, err := service.Symbol("BTCUSDT").
			StrategyID(123445).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel UM conditional order: %v", err)
		}

		// Basic validation of returned data
		if res.StrategyID != 123445 {
			t.Errorf("Expected strategy ID 123445, got %d", res.StrategyID)
		}
		if res.StrategyStatus != "CANCELED" {
			t.Errorf("Expected status CANCELED, got %s", res.StrategyStatus)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
	})

	t.Run("CancelUMConditionalOrder_ByClientStrategyID", func(t *testing.T) {
		service := suite.client.NewUMCancelConditionalOrderService()
		res, err := service.Symbol("BTCUSDT").
			NewClientStrategyID("myOrder1").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel UM conditional order: %v", err)
		}

		// Basic validation of returned data
		if res.NewClientStrategyID != "myOrder1" {
			t.Errorf("Expected client strategy ID myOrder1, got %s", res.NewClientStrategyID)
		}
		if res.StrategyStatus != "CANCELED" {
			t.Errorf("Expected status CANCELED, got %s", res.StrategyStatus)
		}
	})

	t.Run("CancelUMConditionalOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewUMCancelConditionalOrderService()
		_, err := service.Symbol("BTCUSDT").
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
