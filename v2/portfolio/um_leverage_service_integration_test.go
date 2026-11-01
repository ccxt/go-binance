//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umLeverageServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMLeverageServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umLeverageServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("ChangeLeverage", func(t *testing.T) {
		service := &ChangeUMInitialLeverageService{c: suite.client}
		res, err := service.
			Symbol("BTCUSDT").
			Leverage(20).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to change leverage: %v", err)
		}

		// Basic validation of returned data
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %v", res.Symbol)
		}

		if res.Leverage != 20 {
			t.Errorf("Expected leverage 20, got %v", res.Leverage)
		}

		if res.MaxNotionalValue == "" {
			t.Error("Expected non-empty max notional value")
		}
	})
}
