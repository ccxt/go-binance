package portfolio

import (
	"context"
	"testing"
	"time"
)

type getMarginForceOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestGetMarginForceOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &getMarginForceOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetForceOrders_Basic", func(t *testing.T) {
		service := suite.client.NewGetMarginForceOrdersService()
		res, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get force orders: %v", err)
		}

		if res.Total < 0 {
			t.Errorf("Expected non-negative total, got %d", res.Total)
		}
	})

	t.Run("GetForceOrders_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago
		service := suite.client.NewGetMarginForceOrdersService()
		res, err := service.
			StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get force orders with time range: %v", err)
		}

		for _, order := range res.Rows {
			if order.UpdatedTime < startTime || order.UpdatedTime > endTime {
				t.Errorf("Order time %d outside range [%d, %d]",
					order.UpdatedTime, startTime, endTime)
			}
		}
	})

	t.Run("GetForceOrders_WithPagination", func(t *testing.T) {
		size := int64(10)
		current := int64(1)
		service := suite.client.NewGetMarginForceOrdersService()
		res, err := service.
			Size(size).
			Current(current).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get force orders with pagination: %v", err)
		}

		if int64(len(res.Rows)) > size {
			t.Errorf("Expected max %d orders, got %d", size, len(res.Rows))
		}
	})
}
