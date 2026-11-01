package portfolio

import (
	"context"
	"testing"
)

type umFeeBurnStatusServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMFeeBurnStatusServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umFeeBurnStatusServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetFeeBurnStatus", func(t *testing.T) {
		service := suite.client.NewUMFeeBurnStatusService()
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get fee burn status: %v", err)
		}

		// Validate that we received a boolean response
		// The actual value could be either true or false depending on the account settings
		if res == nil {
			t.Fatal("Expected non-nil response")
		}
	})

	t.Run("GetFeeBurnStatus_WithRecvWindow", func(t *testing.T) {
		service := suite.client.NewUMFeeBurnStatusService()
		res, err := service.
			RecvWindow(5000).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get fee burn status with recvWindow: %v", err)
		}

		// Validate that we received a boolean response
		if res == nil {
			t.Fatal("Expected non-nil response")
		}
	})
}
