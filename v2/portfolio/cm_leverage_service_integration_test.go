//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type cmLeverageServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMLeverageServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmLeverageServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("ChangeLeverage", func(t *testing.T) {
		service := &ChangeCMInitialLeverageService{c: suite.client}
		res, err := service.
			Symbol("BTCUSD_PERP").
			Leverage(20).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to change leverage: %v", err)
		}

		// Basic validation of returned data
		if res.Symbol != "BTCUSD_PERP" {
			t.Errorf("Expected symbol BTCUSD_PERP, got %v", res.Symbol)
		}

		if res.Leverage != 20 {
			t.Errorf("Expected leverage 20, got %v", res.Leverage)
		}

		if res.MaxQty == "" {
			t.Error("Expected non-empty max quantity")
		}
	})
}
