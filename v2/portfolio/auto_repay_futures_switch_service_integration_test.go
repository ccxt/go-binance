//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type autoRepayFuturesSwitchServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestAutoRepayFuturesSwitchServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &autoRepayFuturesSwitchServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("ChangeAutoRepayFuturesStatus", func(t *testing.T) {
		service := &ChangeAutoRepayFuturesStatusService{c: suite.client}

		// Get current status
		currentStatus, err := suite.client.NewGetAutoRepayFuturesStatusService().Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get current auto repay futures status: %v", err)
		}

		// Change status
		newStatus := !currentStatus.AutoRepay
		res, err := service.AutoRepay(newStatus).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to change auto repay futures status: %v", err)
		}

		if res.Msg != "success" {
			t.Errorf("Expected success message, got %v", res.Msg)
		}

		// Revert to original status
		_, err = service.AutoRepay(currentStatus.AutoRepay).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to revert auto repay futures status: %v", err)
		}
	})
}
