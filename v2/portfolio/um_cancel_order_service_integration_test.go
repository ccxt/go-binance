//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type umCancelOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMCancelOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umCancelOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelOrder", func(t *testing.T) {
		// First place an order to cancel
		orderService := &UMOrderService{c: suite.client}
		order, err := orderService.
			Symbol("BTCUSDT").
			Side(SideTypeBuy).
			Type(OrderTypeLimit).
			TimeInForce(TimeInForceTypeGTC).
			Quantity("0.001").
			Price("20000").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to place order: %v", err)
		}

		// Wait a bit to ensure order is in the system
		time.Sleep(time.Second)

		// Now cancel the order
		cancelService := &UMCancelOrderService{c: suite.client}
		response, err := cancelService.
			Symbol("BTCUSDT").
			OrderID(order.OrderID).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel order: %v", err)
		}

		// Validate the cancelled order response
		if response.Symbol != "BTCUSDT" {
			t.Error("Expected symbol to be BTCUSDT")
		}

		if response.OrderID != order.OrderID {
			t.Error("Expected matching order IDs")
		}

		if response.Status != "CANCELED" {
			t.Errorf("Expected status CANCELED, got %s", response.Status)
		}

		if response.Type != OrderTypeLimit {
			t.Error("Expected type to be LIMIT")
		}

		if response.Side != SideTypeBuy {
			t.Error("Expected side to be BUY")
		}
	})
}
