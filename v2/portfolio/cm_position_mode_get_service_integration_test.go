//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type cmPositionModeGetServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMPositionModeGetServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmPositionModeGetServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetPositionMode", func(t *testing.T) {
		service := &GetCMPositionModeService{c: suite.client}
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get position mode: %v", err)
		}

		t.Logf("Current position mode: Hedge Mode: %v", res.DualSidePosition)
	})
}
