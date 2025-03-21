//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umPositionModeGetServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMPositionModeGetServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umPositionModeGetServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetPositionMode", func(t *testing.T) {
		service := &GetUMPositionModeService{c: suite.client}
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get position mode: %v", err)
		}

		// Basic validation that we got a response
		// The actual value could be either true or false depending on the account settings
		t.Logf("Current position mode: Hedge Mode: %v", res.DualSidePosition)
	})
}
