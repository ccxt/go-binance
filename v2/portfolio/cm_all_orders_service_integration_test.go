package portfolio

import (
	"context"
	"testing"
	"time"
)

type cmAllOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMAllOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmAllOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetAllCMOrders_BySymbol", func(t *testing.T) {
		service := suite.client.NewCMAllOrdersService()
		orders, err := service.Symbol("BTCUSD_200925").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all CM orders: %v", err)
		}

		// Validate returned data
		for _, order := range orders {
			if order.Symbol != "BTCUSD_200925" {
				t.Errorf("Expected symbol BTCUSD_200925, got %s", order.Symbol)
			}
		}
	})

	t.Run("GetAllCMOrders_ByPair", func(t *testing.T) {
		service := suite.client.NewCMAllOrdersService()
		orders, err := service.Pair("BTCUSD").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all CM orders by pair: %v", err)
		}

		// Validate returned data
		for _, order := range orders {
			if order.Pair != "BTCUSD" {
				t.Errorf("Expected pair BTCUSD, got %s", order.Pair)
			}
		}
	})

	t.Run("GetAllCMOrders_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - 24*60*60*1000 // 24 hours ago

		service := suite.client.NewCMAllOrdersService()
		orders, err := service.Symbol("BTCUSD_200925").
			StartTime(startTime).
			EndTime(endTime).
			Limit(10).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get all CM orders with time range: %v", err)
		}

		// Validate time range
		for _, order := range orders {
			if order.Time < startTime || order.Time > endTime {
				t.Errorf("Order time %d outside requested range [%d, %d]",
					order.Time, startTime, endTime)
			}
		}
	})

	t.Run("GetAllCMOrders_Error_NoSymbolOrPair", func(t *testing.T) {
		service := suite.client.NewCMAllOrdersService()
		_, err := service.Do(context.Background())
		if err == nil {
			t.Fatal("Expected an error when neither symbol nor pair is provided")
		}

		// Verify it's a Portfolio error
		portfolioErr, ok := err.(*Error)
		if !ok {
			t.Fatalf("Expected Error, got %T", err)
		}
		if portfolioErr.Code != ErrMandatoryParamEmptyOrMalformed {
			t.Errorf("Expected error code %d, got %d", ErrMandatoryParamEmptyOrMalformed, portfolioErr.Code)
		}
	})
}
