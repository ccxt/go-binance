//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umCancelAllOrdersServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMCancelAllOrdersServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umCancelAllOrdersServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelAllOrders", func(t *testing.T) {
		// // First place a few orders to cancel
		// orderService := &UMOrderService{c: suite.client}

		// // Place first order
		// _, err := orderService.
		// 	Symbol("BTCUSDC").
		// 	Side(SideTypeBuy).
		// 	Type(OrderTypeLimit).
		// 	TimeInForce(TimeInForceTypeGTC).
		// 	Quantity("0.001").
		// 	Price("20000").
		// 	Do(context.Background())
		// if err != nil {
		// 	t.Fatalf("Failed to place first order: %v", err)
		// }

		// // Place second order
		// _, err = orderService.
		// 	Symbol("BTCUSDC").
		// 	Side(SideTypeBuy).
		// 	Type(OrderTypeLimit).
		// 	TimeInForce(TimeInForceTypeGTC).
		// 	Quantity("0.001").
		// 	Price("19000").
		// 	Do(context.Background())
		// if err != nil {
		// 	t.Fatalf("Failed to place second order: %v", err)
		// }

		// // Wait a bit to ensure orders are in the system
		// time.Sleep(time.Second)

		// Now cancel all orders
		cancelService := &UMCancelAllOrdersService{c: suite.client}
		response, err := cancelService.
			Symbol("BTCUSDC").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel all orders: %v", err)
		}

		// Validate the response
		if response.Code != 200 {
			t.Errorf("Expected code 200, got %d", response.Code)
		}

		if response.Msg != "The operation of cancel all open order is done." {
			t.Errorf("Unexpected response message: %s", response.Msg)
		}

		// // Verify no orders remain
		// time.Sleep(time.Second)
		// openOrders, err := suite.client.NewUMCancelAllOrdersService().
		// 	Symbol("BTCUSDT").
		// 	Do(context.Background())
		// if err != nil {
		// 	t.Fatalf("Failed to get open orders: %v", err)
		// }

		// if len(openOrders) > 0 {
		// 	t.Errorf("Expected no open orders, found %d", len(openOrders))
		// }
	})
}
