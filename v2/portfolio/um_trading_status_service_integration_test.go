//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umTradingStatusServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMTradingStatusServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umTradingStatusServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetTradingStatus", func(t *testing.T) {
		service := &GetUMTradingStatusService{c: suite.client}
		status, err := service.Symbol("BTCUSDT").Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get trading status: %v", err)
		}

		if status.UpdateTime == 0 {
			t.Error("Expected non-zero update time")
		}

		for symbol, indicators := range status.Indicators {
			if symbol == "" {
				t.Error("Expected non-empty symbol")
			}
			if len(indicators) == 0 {
				t.Error("Expected non-empty indicators")
			}
		}
	})
}
