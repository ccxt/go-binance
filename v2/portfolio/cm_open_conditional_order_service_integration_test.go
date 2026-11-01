package portfolio

import (
	"context"
	"testing"
)

type cmOpenConditionalOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMOpenConditionalOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmOpenConditionalOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetOpenConditionalOrder_ByStrategyID", func(t *testing.T) {
		service := suite.client.NewCMOpenConditionalOrderService()
		order, err := service.Symbol("BTCUSD").
			StrategyID(123445).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open conditional order: %v", err)
		}

		// Validate returned data
		if order.Symbol != "BTCUSD" {
			t.Errorf("Expected symbol BTCUSD, got %s", order.Symbol)
		}
		if order.StrategyStatus != "NEW" {
			t.Errorf("Expected status NEW, got %s", order.StrategyStatus)
		}
	})

	t.Run("GetOpenConditionalOrder_ByClientStrategyID", func(t *testing.T) {
		service := suite.client.NewCMOpenConditionalOrderService()
		order, err := service.Symbol("BTCUSD").
			NewClientStrategyID("abc").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open conditional order: %v", err)
		}

		if order.NewClientStrategyID != "abc" {
			t.Errorf("Expected newClientStrategyId abc, got %s", order.NewClientStrategyID)
		}
	})

	t.Run("GetOpenConditionalOrder_NoIDs", func(t *testing.T) {
		service := suite.client.NewCMOpenConditionalOrderService()
		_, err := service.Symbol("BTCUSD").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected error when neither strategyId nor newClientStrategyId is provided")
		}
	})
}
