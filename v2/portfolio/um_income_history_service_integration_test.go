//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type umIncomeHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMIncomeHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umIncomeHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMIncomeHistory", func(t *testing.T) {
		service := &GetUMIncomeHistoryService{c: suite.client}
		endTime := time.Now().UnixMilli()
		startTime := endTime - 7*24*60*60*1000 // 7 days ago

		incomes, err := service.
			StartTime(startTime).
			EndTime(endTime).
			Limit(100).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get UM income history: %v", err)
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
