//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umPositionModeServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMPositionModeServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umPositionModeServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("ChangePositionMode", func(t *testing.T) {
		service := &ChangeUMPositionModeService{c: suite.client}
		res, err := service.
			DualSidePosition(true). // Enable Hedge Mode
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to change position mode: %v", err)
		}

		// Basic validation of returned data
		if res.Code != 200 {
			t.Errorf("Expected code 200, got %v", res.Code)
		}

		if res.Msg != "success" {
			t.Errorf("Expected msg 'success', got %v", res.Msg)
		}

		// Test changing back to One-way Mode
		res, err = service.
			DualSidePosition(false).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to change position mode back: %v", err)
		}
	})
}
