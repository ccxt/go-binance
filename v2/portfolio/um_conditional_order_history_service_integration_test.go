package portfolio

import (
	"context"
	"testing"
)

type umConditionalOrderHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMConditionalOrderHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umConditionalOrderHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetConditionalOrderHistory_ByStrategyID", func(t *testing.T) {
		service := suite.client.NewUMConditionalOrderHistoryService()
		res, err := service.Symbol("BTCUSDT").
			StrategyID(123445).
			Do(context.Background())
		if err != nil {
			// Check if error is "Order does not exist" which is expected for NEW orders
			if err.Error() != "Order does not exist" {
				t.Fatalf("Failed to get conditional order history: %v", err)
			}
			return
		}

		// Basic validation of returned data
		if res.StrategyID != 123445 {
			t.Errorf("Expected strategy ID 123445, got %d", res.StrategyID)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
		if res.StrategyStatus == "NEW" {
			t.Error("Expected non-NEW status for history query")
		}
	})

	t.Run("GetConditionalOrderHistory_ByClientStrategyID", func(t *testing.T) {
		service := suite.client.NewUMConditionalOrderHistoryService()
		res, err := service.Symbol("BTCUSDT").
			NewClientStrategyID("abc").
			Do(context.Background())
		if err != nil {
			// Check if error is "Order does not exist" which is expected for NEW orders
			if err.Error() != "Order does not exist" {
				t.Fatalf("Failed to get conditional order history: %v", err)
			}
			return
		}

		// Basic validation of returned data
		if res.NewClientStrategyID != "abc" {
			t.Errorf("Expected client strategy ID abc, got %s", res.NewClientStrategyID)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
	})

	t.Run("GetConditionalOrderHistory_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewUMConditionalOrderHistoryService()
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
