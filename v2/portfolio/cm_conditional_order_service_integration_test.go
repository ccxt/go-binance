//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type cmConditionalOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMConditionalOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmConditionalOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("PlaceTrailingStopOrder", func(t *testing.T) {
		service := &CMConditionalOrderService{c: suite.client}
		order, err := service.Symbol("BTCUSD_PERP").
			Side(SideTypeBuy).
			StrategyType("TRAILING_STOP_MARKET").
			CallbackRate("1").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to place conditional order: %v", err)
		}

		// Basic validation of returned data
		if order.Symbol != "BTCUSD_PERP" {
			t.Error("Expected symbol to be BTCUSD_PERP")
		}

		if order.Side != SideTypeBuy {
			t.Error("Expected side to be BUY")
		}

		if order.StrategyType != "TRAILING_STOP_MARKET" {
			t.Error("Expected strategy type to be TRAILING_STOP_MARKET")
		}

		if order.StrategyStatus == "" {
			t.Error("Expected non-empty strategy status")
		}

		if order.StrategyId == 0 {
			t.Error("Expected non-zero strategy ID")
		}
	})
}
