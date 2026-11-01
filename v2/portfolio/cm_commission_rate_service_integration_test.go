//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type cmCommissionRateServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMCommissionRateServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmCommissionRateServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetCommissionRate", func(t *testing.T) {
		service := &GetCMCommissionRateService{c: suite.client}
		rates, err := service.Symbol("BTCUSD_PERP").Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get commission rates: %v", err)
		}

		if rates.Symbol != "BTCUSD_PERP" {
			t.Errorf("Expected symbol BTCUSD_PERP, got %v", rates.Symbol)
		}
		if rates.MakerCommissionRate == "" {
			t.Error("Expected non-empty maker commission rate")
		}
		if rates.TakerCommissionRate == "" {
			t.Error("Expected non-empty taker commission rate")
		}
	})
}
