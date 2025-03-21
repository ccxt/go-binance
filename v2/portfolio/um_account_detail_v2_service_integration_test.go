//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umAccountDetailV2ServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMAccountDetailV2ServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umAccountDetailV2ServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMAccountDetailV2", func(t *testing.T) {
		service := &UMAccountDetailV2Service{c: suite.client}
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get UM account detail v2: %v", err)
		}

		// Validate assets
		if len(res.Assets) == 0 {
			t.Error("Expected at least one asset")
		}

		for _, asset := range res.Assets {
			if asset.Asset == "" {
				t.Error("Expected non-empty asset name")
			}
		}

		// Validate positions
		for _, position := range res.Positions {
			if position.Symbol == "" {
				t.Error("Expected non-empty symbol")
			}
			if position.PositionSide == "" {
				t.Error("Expected non-empty position side")
			}
			// Position amount and notional can be "0"
		}
	})
}
