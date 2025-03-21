//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type autoRepayFuturesStatusServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestAutoRepayFuturesStatusServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &autoRepayFuturesStatusServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetAutoRepayFuturesStatus", func(t *testing.T) {
		service := &GetAutoRepayFuturesStatusService{c: suite.client}
		status, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get auto repay futures status: %v", err)
		}

		t.Logf("Auto repay futures is %v", status.AutoRepay)
	})
}
