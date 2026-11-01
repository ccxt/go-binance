//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type umOrderHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMOrderHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umOrderHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMOrderHistoryDownloadID", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago

		service := suite.client.NewGetUMOrderHistoryDownloadIDService()
		res, err := service.
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get UM order history download ID: %v", err)
		}

		if res.DownloadID == "" {
			t.Error("Expected non-empty download ID")
		}

		if res.AvgCostTimestampOfLast30d <= 0 {
			t.Error("Expected positive average cost timestamp")
		}
	})
}
