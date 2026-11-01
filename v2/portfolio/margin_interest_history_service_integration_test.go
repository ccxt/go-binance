//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type marginInterestHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginInterestHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginInterestHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetMarginInterestHistory", func(t *testing.T) {
		service := &GetMarginInterestHistoryService{c: suite.client}
		endTime := time.Now().UnixMilli()
		startTime := endTime - 7*24*60*60*1000 // 7 days ago

		history, err := service.
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get margin interest history: %v", err)
		}

		if history.Total < 0 {
			t.Error("Expected non-negative total")
		}

		for _, interest := range history.Rows {
			if interest.Asset == "" {
				t.Error("Expected non-empty asset")
			}
			if interest.Type == "" {
				t.Error("Expected non-empty type")
			}
		}
	})
}
