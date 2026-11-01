//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type cmOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("PlaceOrder", func(t *testing.T) {
		service := &CMOrderService{c: suite.client}
		order, err := service.Symbol("BTCUSD_PERP").
			Side(SideTypeBuy).
			Type(OrderTypeLimit).
			TimeInForce(TimeInForceTypeGTC).
			Quantity("1").
			Price("20000").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to place order: %v", err)
		}

		// Basic validation of returned data
		if order.Symbol != "BTCUSD_PERP" {
			t.Error("Expected symbol to be BTCUSD_PERP")
		}

		if order.Side != SideTypeBuy {
			t.Error("Expected side to be BUY")
		}

		if order.Type != OrderTypeLimit {
			t.Error("Expected type to be LIMIT")
		}

		if order.TimeInForce != TimeInForceTypeGTC {
			t.Error("Expected timeInForce to be GTC")
		}

		if order.Status == "" {
			t.Error("Expected non-empty status")
		}

		if order.OrderID == 0 {
			t.Error("Expected non-zero order ID")
		}
	})
}
