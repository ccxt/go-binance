//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type negativeBalanceServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestNegativeBalanceServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &negativeBalanceServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetNegativeBalanceExchangeRecord", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago

		service := suite.client.NewGetNegativeBalanceExchangeRecordService()
		res, err := service.
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get negative balance exchange record: %v", err)
		}

		if res.Total < 0 {
			t.Error("Expected non-negative total")
		}

		for _, row := range res.Rows {
			if row.StartTime <= 0 {
				t.Error("Expected positive start time")
			}
			if row.EndTime <= 0 {
				t.Error("Expected positive end time")
			}
			if row.EndTime < row.StartTime {
				t.Error("Expected end time to be after start time")
			}

			for _, detail := range row.Details {
				if detail.Asset == "" {
					t.Error("Expected non-empty asset")
				}
			}
		}
	})
}
