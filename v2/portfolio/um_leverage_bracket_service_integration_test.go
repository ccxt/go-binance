//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umLeverageBracketServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMLeverageBracketServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umLeverageBracketServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetLeverageBracket", func(t *testing.T) {
		service := &GetUMLeverageBracketService{c: suite.client}
		brackets, err := service.Symbol("BTCUSDT").Do(context.Background())
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
