//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umAccountDetailServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMAccountDetailServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umAccountDetailServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMAccountDetail", func(t *testing.T) {
		service := &GetUMAccountDetailService{c: suite.client}
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get UM account detail: %v", err)
		}

		if len(res.Assets) == 0 {
			t.Error("Expected non-empty assets")
		}
		for _, asset := range res.Assets {
			if asset.Asset == "" {
				t.Error("Expected non-empty asset name")
			}
		}

		if len(res.Positions) == 0 {
			t.Error("Expected non-empty positions")
		}
		for _, position := range res.Positions {
			if position.Symbol == "" {
				t.Error("Expected non-empty symbol")
			}
		}
	})
}
