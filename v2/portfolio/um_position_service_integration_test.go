//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umPositionServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMPositionServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umPositionServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetPositionRisk", func(t *testing.T) {
		service := &GetUMPositionRiskService{c: suite.client}
		positions, err := service.Symbol("BTCUSDT").Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get position risk info: %v", err)
		}

		for _, position := range positions {
			if position.Symbol == "" {
				t.Error("Expected non-empty symbol")
			}

			if position.PositionSide == "" {
				t.Error("Expected non-empty position side")
			}

			if position.UpdateTime == 0 {
				t.Error("Expected non-zero update time")
			}
		}
	})

	t.Run("GetAllPositionsRisk", func(t *testing.T) {
		service := &GetUMPositionRiskService{c: suite.client}
		positions, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all positions risk info: %v", err)
		}

		for _, position := range positions {
			if position.Symbol == "" {
				t.Error("Expected non-empty symbol")
			}
		}
	})
}
