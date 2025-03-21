//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umOrderDownloadLinkServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMOrderDownloadLinkServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umOrderDownloadLinkServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetUMOrderDownloadLink", func(t *testing.T) {
		downloadID := "953366156082814976"
		service := suite.client.NewGetUMOrderDownloadLinkService()
		res, err := service.
			DownloadID(downloadID).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to get UM order download link: %v", err)
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
