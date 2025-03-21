//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type marginBorrowServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginBorrowServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginBorrowServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetMaxBorrow", func(t *testing.T) {
		service := &GetMarginMaxBorrowService{c: suite.client}
		maxBorrow, err := service.Asset("USDT").Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get max borrow info: %v", err)
		}

		// Basic validation of returned data
		if maxBorrow.Amount == "" {
			t.Error("Expected non-empty amount")
		}

		if maxBorrow.BorrowLimit == "" {
			t.Error("Expected non-empty borrow limit")
		}
	})
}
