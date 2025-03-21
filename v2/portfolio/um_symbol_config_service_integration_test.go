//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umSymbolConfigServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMSymbolConfigServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umSymbolConfigServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMSymbolConfig", func(t *testing.T) {
		service := &UMSymbolConfigService{c: suite.client}

		// Test without symbol parameter first
		configs, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get UM symbol configs: %v", err)
		}

		if len(configs) == 0 {
			t.Error("Expected at least one symbol configuration")
		}

		// Test with specific symbol
		symbol := "BTCUSDT"
		configs, err = service.Symbol(symbol).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get UM symbol config for %s: %v", symbol, err)
		}

		// Validate returned data
		for _, config := range configs {
			if config.Symbol == "" {
				t.Error("Expected non-empty symbol")
			}
			if config.MarginType == "" {
				t.Error("Expected non-empty margin type")
			}
			if config.Leverage <= 0 {
				t.Error("Expected positive leverage")
			}
			if config.MaxNotionalValue == "" {
				t.Error("Expected non-empty maxNotionalValue")
			}
		}
	})
}
