//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umTradeDownloadLinkServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMTradeDownloadLinkServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umTradeDownloadLinkServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMTradeDownloadLink", func(t *testing.T) {
		downloadID := "953365044620652544" // This should be a valid download ID from a previous request
		service := suite.client.NewGetUMTradeDownloadLinkService()
		res, err := service.
			DownloadID(downloadID).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get UM trade download link: %v", err)
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
