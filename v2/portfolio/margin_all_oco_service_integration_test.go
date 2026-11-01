package portfolio

import (
	"context"
	"testing"
	"time"
)

type marginAllOCOServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginAllOCOServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginAllOCOServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("AllOCO_Basic", func(t *testing.T) {
		service := suite.client.NewMarginAllOCOService()
		orders, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all OCO orders: %v", err)
		}

		for _, order := range orders {
			if order.OrderListID == 0 {
				t.Error("Expected non-zero OrderListID")
			}
			if order.ContingencyType != "OCO" {
				t.Errorf("Expected ContingencyType OCO, got %s", order.ContingencyType)
			}
		}
	})

	t.Run("AllOCO_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago
		service := suite.client.NewMarginAllOCOService()
		orders, err := service.StartTime(startTime).
			EndTime(endTime).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get OCO orders with time range: %v", err)
		}

		for _, order := range orders {
			if order.TransactionTime < startTime || order.TransactionTime > endTime {
				t.Errorf("Order time %d outside range [%d, %d]",
					order.TransactionTime, startTime, endTime)
			}
		}
	})

	t.Run("AllOCO_WithLimit", func(t *testing.T) {
		limit := 10
		service := suite.client.NewMarginAllOCOService()
		orders, err := service.Limit(limit).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get OCO orders with limit: %v", err)
		}

		if len(orders) > limit {
			t.Errorf("Expected max %d orders, got %d", limit, len(orders))
		}
	})

	t.Run("AllOCO_WithFromID", func(t *testing.T) {
		fromID := int64(1)
		service := suite.client.NewMarginAllOCOService()
		orders, err := service.FromID(fromID).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get OCO orders from ID: %v", err)
		}

		for _, order := range orders {
			if order.OrderListID < fromID {
				t.Errorf("Expected OrderListID >= %d, got %d", fromID, order.OrderListID)
			}
		}
	})
}
