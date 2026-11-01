//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type fundCollectionByAssetServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestFundCollectionByAssetServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &fundCollectionByAssetServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("FundCollectionByAsset", func(t *testing.T) {
		service := &FundCollectionByAssetService{c: suite.client}
		res, err := service.Asset("USDT").Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to execute fund collection by asset: %v", err)
		}

		if res.Msg != "success" {
			t.Errorf("Expected success message, got %v", res.Msg)
		}
	})
}
