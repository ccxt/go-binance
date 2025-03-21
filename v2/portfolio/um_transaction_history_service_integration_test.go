//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type umTransactionHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMTransactionHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umTransactionHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMTransactionHistoryDownloadID", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago

		service := suite.client.NewGetUMTransactionHistoryDownloadIDService()
		res, err := service.
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get UM transaction history download ID: %v", err)
		}

		if res.DownloadID == "" {
			t.Error("Expected non-empty download ID")
		}

		if res.AvgCostTimestampOfLast30d <= 0 {
			t.Error("Expected positive average cost timestamp")
		}
	})
}
