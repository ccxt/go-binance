//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umTransactionDownloadLinkServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMTransactionDownloadLinkServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umTransactionDownloadLinkServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMTransactionDownloadLink", func(t *testing.T) {
		downloadID := "953367130350170112"
		service := suite.client.NewGetUMTransactionDownloadLinkService()
		res, err := service.
			DownloadID(downloadID).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get UM transaction download link: %v", err)
		}

		if res.DownloadID == "" {
			t.Error("Expected non-empty download ID")
		}

		if res.Status != "completed" && res.Status != "processing" {
			t.Error("Expected status to be either 'completed' or 'processing'")
		}

		if res.Status == "completed" && res.URL == "" {
			t.Error("Expected non-empty URL for completed status")
		}

		if res.ExpirationTimestamp <= 0 && res.Status == "completed" {
			t.Error("Expected positive expiration timestamp for completed status")
		}
	})
}
