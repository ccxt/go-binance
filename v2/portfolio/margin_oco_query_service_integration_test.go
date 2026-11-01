//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type marginOCOQueryServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginOCOQueryServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginOCOQueryServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("QueryOCO", func(t *testing.T) {
		service := &MarginOCOQueryService{c: suite.client}
		res, err := service.OrderListID(27).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to query OCO order: %v", err)
		}

		// Basic validation of returned data
		if res.OrderListID == 0 {
			t.Error("Expected non-zero order list ID")
		}

		if res.ContingencyType != "OCO" {
			t.Error("Expected contingency type to be OCO")
		}

		if len(res.Orders) == 0 {
			t.Error("Expected at least one order in the OCO")
		}
	})
}
