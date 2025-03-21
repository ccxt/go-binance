package portfolio

import (
	"context"
	"testing"
)

type umGetADLQuantileServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMGetADLQuantileServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umGetADLQuantileServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetADLQuantile_SingleSymbol", func(t *testing.T) {
		symbol := "BTCUSDT"
		service := suite.client.NewUMGetADLQuantileService()
		res, err := service.Symbol(symbol).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get ADL quantile: %v", err)
		}

		if len(res) == 0 {
			t.Fatal("Expected at least one result")
		}

		for _, quantile := range res {
			if quantile.Symbol != symbol {
				t.Errorf("Expected symbol %s, got %s", symbol, quantile.Symbol)
			}
			validateADLQuantile(t, &quantile.ADLQuantile)
		}
	})

	t.Run("GetADLQuantile_AllSymbols", func(t *testing.T) {
		service := suite.client.NewUMGetADLQuantileService()
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get ADL quantiles: %v", err)
		}

		for _, quantile := range res {
			validateADLQuantile(t, &quantile.ADLQuantile)
		}
	})
}

func validateADLQuantile(t *testing.T, quantile *ADLQuantile) {
	// Validate LONG and SHORT values are in range [0-4]
	if quantile.LONG < 0 || quantile.LONG > 4 {
		t.Errorf("LONG quantile %d outside valid range [0-4]", quantile.LONG)
	}
	if quantile.SHORT < 0 || quantile.SHORT > 4 {
		t.Errorf("SHORT quantile %d outside valid range [0-4]", quantile.SHORT)
	}

	// Check if either BOTH or HEDGE is present (but not both)
	if quantile.BOTH != nil && quantile.HEDGE != nil {
		t.Error("Both BOTH and HEDGE values present")
	}

	// Validate BOTH or HEDGE values if present
	if quantile.BOTH != nil {
		if *quantile.BOTH < 0 || *quantile.BOTH > 4 {
			t.Errorf("BOTH quantile %d outside valid range [0-4]", *quantile.BOTH)
		}
	}
	if quantile.HEDGE != nil {
		// For HEDGE, we only care that it exists as it's just a sign
		// The value itself should be ignored according to the API docs
	}

	// For cross-margined hedge mode, LONG and SHORT should have the same value
	if quantile.HEDGE != nil && quantile.LONG != quantile.SHORT {
		t.Errorf("In hedge mode, expected LONG (%d) and SHORT (%d) to be equal",
			quantile.LONG, quantile.SHORT)
	}
}
