//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type cmLeverageBracketServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMLeverageBracketServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmLeverageBracketServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetLeverageBracket", func(t *testing.T) {
		service := &GetCMLeverageBracketService{c: suite.client}
		brackets, err := service.Symbol("BTCUSD_PERP").Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get leverage bracket: %v", err)
		}

		for _, bracket := range brackets {
			if bracket.Symbol == "" {
				t.Error("Expected non-empty symbol")
			}
			if len(bracket.Brackets) == 0 {
				t.Error("Expected non-empty brackets")
			}
		}
	})
}
