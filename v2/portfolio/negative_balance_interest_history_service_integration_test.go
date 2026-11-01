//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type negativeBalanceInterestHistoryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestNegativeBalanceInterestHistoryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &negativeBalanceInterestHistoryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetNegativeBalanceInterestHistory", func(t *testing.T) {
		service := &GetNegativeBalanceInterestHistoryService{c: suite.client}
		endTime := time.Now().UnixMilli()
		startTime := endTime - 7*24*60*60*1000 // 7 days ago

		interests, err := service.
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get negative balance interest history: %v", err)
		}

		for _, interest := range interests {
			if interest.Asset == "" {
				t.Error("Expected non-empty asset")
			}
			if interest.Interest == "" {
				t.Error("Expected non-empty interest")
			}
		}
	})
}
