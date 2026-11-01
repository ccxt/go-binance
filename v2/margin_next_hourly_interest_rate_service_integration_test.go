package binance

import (
	"context"
	"testing"
)

type marginNextHourlyInterestRateServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginNextHourlyInterestRateServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginNextHourlyInterestRateServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetMarginNextHourlyInterestRate_Basic", func(t *testing.T) {
		service := suite.client.NewMarginNextHourlyInterestRateService()
		history, err := service.
			Assets("USDT").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin next hourly interest rate: %v", err)
		}

		for _, item := range *history {
			if item.Asset == "" {
				t.Error("Expected non-empty asset")
			}
			if item.NextHourlyInterestRate == "" {
				t.Error("Expected non-empty interest rate")
			}
		}
	})

	t.Run("GetMarginNextHourlyInterestRate_WithIsolated", func(t *testing.T) {
		service := suite.client.NewMarginNextHourlyInterestRateService()
		history, err := service.
			Assets("USDT").
			Isolated(true).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get margin next hourly interest rate for isolated symbol: %v", err)
		}

		for _, item := range *history {
			if item.Asset == "" {
				t.Error("Expected non-empty asset")
			}
			if item.NextHourlyInterestRate == "" {
				t.Error("Expected non-empty interest rate")
			}
		}
	})
}
