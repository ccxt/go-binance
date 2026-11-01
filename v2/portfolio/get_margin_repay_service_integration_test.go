//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type getMarginRepayServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestGetMarginRepayServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &getMarginRepayServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetMarginRepay", func(t *testing.T) {
		service := &GetMarginRepayService{c: suite.client}
		endTime := time.Now().UnixMilli()
		startTime := endTime - 7*24*60*60*1000 // 7 days ago

		repays, err := service.
			Asset("BNB").
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get margin repays: %v", err)
		}

		if repays.Total < 0 {
			t.Error("Expected non-negative total")
		}

		for _, repay := range repays.Rows {
			if repay.Asset == "" {
				t.Error("Expected non-empty asset")
			}
			if repay.Status == "" {
				t.Error("Expected non-empty status")
			}
		}
	})
}
