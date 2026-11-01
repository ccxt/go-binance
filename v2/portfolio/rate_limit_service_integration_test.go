//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type rateLimitServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestRateLimitServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &rateLimitServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetRateLimit", func(t *testing.T) {
		service := suite.client.NewGetRateLimitService()
		res, err := service.Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get rate limit: %v", err)
		}

		if len(res) == 0 {
			t.Error("Expected at least one rate limit")
		}

		for _, limit := range res {
			if limit.RateLimitType == "" {
				t.Error("Expected non-empty rate limit type")
			}
			if limit.Interval == "" {
				t.Error("Expected non-empty interval")
			}
			if limit.IntervalNum <= 0 {
				t.Error("Expected positive interval number")
			}
			if limit.Limit <= 0 {
				t.Error("Expected positive limit")
			}
		}
	})
}
