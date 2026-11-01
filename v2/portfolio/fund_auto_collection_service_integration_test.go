//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type fundAutoCollectionServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestFundAutoCollectionServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &fundAutoCollectionServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("FundAutoCollection", func(t *testing.T) {
		service := &FundAutoCollectionService{c: suite.client}
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to execute fund auto-collection: %v", err)
		}

		if res.Msg != "success" {
			t.Errorf("Expected success message, got %v", res.Msg)
		}
	})
}
