//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type getMarginLoanServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestGetMarginLoanServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &getMarginLoanServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetMarginLoan", func(t *testing.T) {
		service := &GetMarginLoanService{c: suite.client}
		endTime := time.Now().UnixMilli()
		startTime := endTime - 7*24*60*60*1000 // 7 days ago

		loans, err := service.
			Asset("BNB").
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get margin loans: %v", err)
		}

		if loans.Total < 0 {
			t.Error("Expected non-negative total")
		}

		for _, loan := range loans.Rows {
			if loan.Asset == "" {
				t.Error("Expected non-empty asset")
			}
			if loan.Status == "" {
				t.Error("Expected non-empty status")
			}
		}
	})
}
