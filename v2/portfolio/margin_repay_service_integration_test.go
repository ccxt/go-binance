//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type marginRepayServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginRepayServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginRepayServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("RepayLoan", func(t *testing.T) {
		service := &MarginRepayService{c: suite.client}
		response, err := service.
			Asset("USDC").
			Amount("10").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to repay loan: %v", err)
		}

		// Basic validation of returned data
		if response.TranID == 0 {
			t.Error("Expected non-zero transaction ID")
		}
	})
}
