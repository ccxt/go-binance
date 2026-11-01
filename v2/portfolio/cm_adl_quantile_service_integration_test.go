package portfolio

import (
	"context"
	"testing"
)

type cmADLQuantileServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMADLQuantileServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmADLQuantileServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetADLQuantile_SingleSymbol", func(t *testing.T) {
		service := suite.client.NewCMADLQuantileService()
		res, err := service.Symbol("BTCUSD_200925").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get ADL quantile: %v", err)
		}

		if len(res) != 1 {
			t.Fatalf("Expected 1 result, got %d", len(res))
		}

		// Validate returned data
		quantile := res[0]
		if quantile.Symbol != "BTCUSD_200925" {
			t.Errorf("Expected symbol BTCUSD_200925, got %s", quantile.Symbol)
		}

		// Validate quantile values are in valid range (0-4)
		if quantile.ADLQuantile.LONG < 0 || quantile.ADLQuantile.LONG > 4 {
			t.Errorf("LONG quantile %d outside valid range [0-4]",
				quantile.ADLQuantile.LONG)
		}
		if quantile.ADLQuantile.SHORT < 0 || quantile.ADLQuantile.SHORT > 4 {
			t.Errorf("SHORT quantile %d outside valid range [0-4]",
				quantile.ADLQuantile.SHORT)
		}
	})

	t.Run("GetADLQuantile_AllSymbols", func(t *testing.T) {
		service := suite.client.NewCMADLQuantileService()
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get ADL quantiles: %v", err)
		}

		// Validate all returned quantiles
		for _, quantile := range res {
			if quantile.Symbol == "" {
				t.Error("Empty symbol in response")
			}
			if quantile.ADLQuantile.LONG < 0 || quantile.ADLQuantile.LONG > 4 {
				t.Errorf("LONG quantile %d outside valid range [0-4] for symbol %s",
					quantile.ADLQuantile.LONG, quantile.Symbol)
			}
			if quantile.ADLQuantile.SHORT < 0 || quantile.ADLQuantile.SHORT > 4 {
				t.Errorf("SHORT quantile %d outside valid range [0-4] for symbol %s",
					quantile.ADLQuantile.SHORT, quantile.Symbol)
			}
		}
	})
}
