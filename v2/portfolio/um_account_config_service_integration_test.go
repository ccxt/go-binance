//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umAccountConfigServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMAccountConfigServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umAccountConfigServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMAccountConfig", func(t *testing.T) {
		service := &UMAccountConfigService{c: suite.client}
		config, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get UM account config: %v", err)
		}

		// Basic validation of returned data
		if config.UpdateTime == 0 {
			t.Error("Expected non-zero update time")
		}

		// Validate boolean fields are set
		if !config.CanTrade && !config.CanDeposit && !config.CanWithdraw {
			t.Error("Expected at least one permission to be true")
		}

		// Validate fee tier
		if config.FeeTier < 0 {
			t.Error("Expected non-negative fee tier")
		}

		// Validate trade group ID
		if config.TradeGroupId == 0 {
			t.Error("Expected non-zero trade group ID")
		}
	})
}
