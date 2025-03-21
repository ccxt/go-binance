package portfolio

import (
	"context"
	"testing"
)

type cmConditionalOrderHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMConditionalOrderHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmConditionalOrderHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetConditionalOrderHistory_ByStrategyID", func(t *testing.T) {
		service := suite.client.NewCMConditionalOrderHistoryService()
		order, err := service.Symbol("BTCUSD").
			StrategyID(123445).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get conditional order history: %v", err)
		}

		// Validate returned data
		if order.Symbol != "BTCUSD" {
			t.Errorf("Expected symbol BTCUSD, got %s", order.Symbol)
		}
		if order.StrategyStatus == "NEW" {
			t.Error("NEW orders should not be returned in order history")
		}
	})

	t.Run("GetConditionalOrderHistory_ByClientStrategyID", func(t *testing.T) {
		service := suite.client.NewCMConditionalOrderHistoryService()
		order, err := service.Symbol("BTCUSD").
			NewClientStrategyID("abc").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get conditional order history: %v", err)
		}

		if order.NewClientStrategyID != "abc" {
			t.Errorf("Expected newClientStrategyId abc, got %s", order.NewClientStrategyID)
		}
	})

	t.Run("GetConditionalOrderHistory_NoIDs", func(t *testing.T) {
		service := suite.client.NewCMConditionalOrderHistoryService()
		_, err := service.Symbol("BTCUSD").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected error when neither strategyId nor newClientStrategyId is provided")
		}
	})
}
