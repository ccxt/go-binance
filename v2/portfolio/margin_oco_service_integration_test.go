//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type marginOCOServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginOCOServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginOCOServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("PlaceOCOOrder", func(t *testing.T) {
		service := &MarginOCOService{c: suite.client}
		response, err := service.
			Symbol("BTCUSDT").
			Side(SideTypeBuy).
			Quantity("0.001").
			Price("20000").
			StopPrice("90005").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to place OCO order: %v", err)
		}

		// Basic validation of returned data
		if response.Symbol != "BTCUSDT" {
			t.Error("Expected symbol to be BTCUSDT")
		}

		if response.ContingencyType != "OCO" {
			t.Error("Expected contingencyType to be OCO")
		}

		if response.ListStatusType == "" {
			t.Error("Expected non-empty listStatusType")
		}

		if response.ListOrderStatus == "" {
			t.Error("Expected non-empty listOrderStatus")
		}

		if len(response.Orders) != 2 {
			t.Error("Expected exactly 2 orders in OCO")
		}

		if len(response.OrderReports) != 2 {
			t.Error("Expected exactly 2 order reports")
		}

		// Validate individual orders
		for _, order := range response.OrderReports {
			if order.Symbol != "BTCUSDT" {
				t.Error("Expected order symbol to be BTCUSDT")
			}

			if order.OrderID == 0 {
				t.Error("Expected non-zero order ID")
			}

			if order.Side != SideTypeBuy {
				t.Error("Expected order side to be BUY")
			}
		}
	})
}
