package portfolio

import (
	"context"
	"testing"
)

type marginLoanServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginLoanServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginLoanServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("MarginLoan", func(t *testing.T) {
		service := &MarginLoanService{c: suite.client}
		res, err := service.Asset("USDC").
			Amount("10").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to borrow margin loan: %v", err)
		}

		// Basic validation of returned data
		if res.TranID == 0 {
			t.Error("Expected non-zero transaction ID")
		}
	})

	t.Run("MarginLoan_USDT_Error", func(t *testing.T) {
		service := &MarginLoanService{c: suite.client}
		_, err := service.Asset("USDT").
			Amount("10").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected an error for USDT margin loan")
		}

		// Verify it's a Portfolio error with the expected code
		portfolioErr, ok := err.(*Error)
		if !ok {
			t.Fatalf("Expected Error, got %T", err)
		}
		if portfolioErr.Code != 51138 {
			t.Errorf("Expected error code 51138, got %d", portfolioErr.Code)
		}
	})
}
