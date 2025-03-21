//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type marginWithdrawServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginWithdrawServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginWithdrawServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetMaxWithdraw", func(t *testing.T) {
		service := &GetMarginMaxWithdrawService{c: suite.client}
		maxWithdraw, err := service.Asset("USDT").Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get max withdraw info: %v", err)
		}

		// Basic validation of returned data
		if maxWithdraw.Amount == "" {
			t.Error("Expected non-empty amount")
		}
	})
}
