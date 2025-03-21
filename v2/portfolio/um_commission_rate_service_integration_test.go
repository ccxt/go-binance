//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umCommissionRateServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMCommissionRateServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umCommissionRateServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetCommissionRate", func(t *testing.T) {
		service := &GetUMCommissionRateService{c: suite.client}
		rates, err := service.Symbol("BTCUSDT").Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get commission rates: %v", err)
		}

		if rates.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %v", rates.Symbol)
		}
		if rates.MakerCommissionRate == "" {
			t.Error("Expected non-empty maker commission rate")
		}
		if rates.TakerCommissionRate == "" {
			t.Error("Expected non-empty taker commission rate")
		}
	})
}
