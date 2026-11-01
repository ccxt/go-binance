package portfolio

import (
	"context"
	"testing"
	"time"
)

type marginForceOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginForceOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginForceOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetForceOrders_DefaultParams", func(t *testing.T) {
		service := suite.client.NewMarginForceOrdersService()
		response, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get force orders: %v", err)
		}

		// Validate default pagination
		if len(response.Rows) > 10 {
			t.Errorf("Expected max 10 orders with default size, got %d", len(response.Rows))
		}
	})

	t.Run("GetForceOrders_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 1 day ago

		service := suite.client.NewMarginForceOrdersService()
		response, err := service.StartTime(startTime).
			EndTime(endTime).
			Size(100).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get force orders with time range: %v", err)
		}

		// Validate time range
		for _, order := range response.Rows {
			if order.UpdatedTime < startTime || order.UpdatedTime > endTime {
				t.Errorf("Order time %d outside range [%d, %d]",
					order.UpdatedTime, startTime, endTime)
			}
		}
	})

	t.Run("GetForceOrders_Pagination", func(t *testing.T) {
		service := suite.client.NewMarginForceOrdersService()
		response, err := service.Size(5).
			Current(2).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get force orders with pagination: %v", err)
		}

		if len(response.Rows) > 5 {
			t.Errorf("Expected max 5 orders per page, got %d", len(response.Rows))
		}
	})
}
