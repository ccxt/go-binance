//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type balanceServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestBalanceServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &balanceServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetBalance", func(t *testing.T) {
		service := &GetBalanceService{c: suite.client}
		balances, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get balance: %v", err)
		}

		if len(balances) == 0 {
			t.Error("Expected non-empty balances")
		}

		// You might want to add more specific assertions here
		for _, balance := range balances {
			if balance.Asset == "" {
				t.Error("Expected non-empty asset")
			}
		}
	})
	t.Run("GetBalance of Asset", func(t *testing.T) {
		service := &GetBalanceService{c: suite.client}
		balances, err := service.Asset("USDT").Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get balance: %v", err)
		}

		if len(balances) == 0 {
			t.Error("Expected non-empty balances")
		}

		for _, balance := range balances {
			if balance.Asset == "" {
				t.Error("Expected non-empty asset")
			}
		}
	})
}
