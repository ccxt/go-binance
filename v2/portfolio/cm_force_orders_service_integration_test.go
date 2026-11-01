package portfolio

import (
	"context"
	"testing"
	"time"
)

type cmForceOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMForceOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmForceOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetForceOrders_SingleSymbol", func(t *testing.T) {
		service := suite.client.NewCMForceOrdersService()
		orders, err := service.Symbol("BTCUSD_200925").
			AutoCloseType("LIQUIDATION").
			Limit(50).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get force orders: %v", err)
		}

		// Validate returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSD_200925" {
				t.Errorf("Expected symbol BTCUSD_200925, got %s", order.Symbol)
			}
			if order.Pair != "BTCUSD" {
				t.Errorf("Expected pair BTCUSD, got %s", order.Pair)
			}
		}
	})

	t.Run("GetForceOrders_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 7*24*60*60*1000 // 7 days ago

		service := suite.client.NewCMForceOrdersService()
		orders, err := service.Symbol("BTCUSD_200925").
			AutoCloseType("ADL").
			StartTime(startTime).
			EndTime(endTime).
			Limit(100).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get force orders with time range: %v", err)
		}

		// Validate time range
		for _, order := range orders {
			if order.Time < startTime || order.Time > endTime {
				t.Errorf("Order time %d outside range [%d, %d]",
					order.Time, startTime, endTime)
			}
		}
	})

	t.Run("GetForceOrders_AllSymbols", func(t *testing.T) {
		service := suite.client.NewCMForceOrdersService()
		orders, err := service.AutoCloseType("LIQUIDATION").
			Limit(50).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all force orders: %v", err)
		}

		if len(orders) > 50 {
			t.Errorf("Expected max 50 orders, got %d", len(orders))
		}
	})
}
