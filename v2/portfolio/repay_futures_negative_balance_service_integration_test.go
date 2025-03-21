//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type repayFuturesNegativeBalanceServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestRepayFuturesNegativeBalanceServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &repayFuturesNegativeBalanceServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("RepayFuturesNegativeBalance", func(t *testing.T) {
		service := &RepayFuturesNegativeBalanceService{c: suite.client}
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to repay futures negative balance: %v", err)
		}

		if res.Msg != "success" {
			t.Errorf("Expected success message, got %v", res.Msg)
		}
	})
}
