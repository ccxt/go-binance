//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type umTradeHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMTradeHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umTradeHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMTradeHistoryDownloadID", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago

		service := suite.client.NewGetUMTradeHistoryDownloadIDService()
		res, err := service.
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get UM trade history download ID: %v", err)
		}

		if res.DownloadID == "" {
			t.Error("Expected non-empty download ID")
		}

		if res.AvgCostTimestampOfLast30d <= 0 {
			t.Error("Expected positive average cost timestamp")
		}
	})
}
