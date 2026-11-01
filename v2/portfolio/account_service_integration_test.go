//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type accountServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestAccountServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &accountServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetAccount", func(t *testing.T) {
		service := &GetAccountService{c: suite.client}
		account, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get account info: %v", err)
		}

		// Basic validation of returned data
		if account.AccountStatus == "" {
			t.Error("Expected non-empty account status")
		}

		if account.AccountEquity == "" {
			t.Error("Expected non-empty account equity")
		}

		if account.UniMMR == "" {
			t.Error("Expected non-empty uniMMR")
		}

		if account.UpdateTime == 0 {
			t.Error("Expected non-zero update time")
		}
	})
}
