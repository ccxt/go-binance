//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type cmIncomeHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMIncomeHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmIncomeHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetCMIncomeHistory", func(t *testing.T) {
		service := &GetCMIncomeHistoryService{c: suite.client}
		endTime := time.Now().UnixMilli()
		startTime := endTime - 7*24*60*60*1000 // 7 days ago

		incomes, err := service.
			StartTime(startTime).
			EndTime(endTime).
			Limit(100).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get CM income history: %v", err)
		}

		for _, income := range incomes {
			if income.Asset == "" {
				t.Error("Expected non-empty asset")
			}
			if income.IncomeType == "" {
				t.Error("Expected non-empty income type")
			}
			if income.TranID == 0 {
				t.Error("Expected non-zero transaction ID")
			}
		}
	})
}
