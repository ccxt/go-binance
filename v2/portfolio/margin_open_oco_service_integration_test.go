package portfolio

import (
	"context"
	"testing"
)

type marginOpenOCOServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginOpenOCOServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginOpenOCOServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetOpenOCO_Basic", func(t *testing.T) {
		service := suite.client.NewMarginOpenOCOService()
		orders, err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open OCO orders: %v", err)
		}

		for _, order := range orders {
			if order.ContingencyType != "OCO" {
				t.Errorf("Expected contingencyType OCO, got %s", order.ContingencyType)
			}
			if len(order.Orders) != 2 {
				t.Errorf("Expected 2 orders in OCO, got %d", len(order.Orders))
			}
			for _, o := range order.Orders {
				if o.Symbol != order.Symbol {
					t.Errorf("Expected symbol %s, got %s", order.Symbol, o.Symbol)
				}
			}
		}
	})

	t.Run("GetOpenOCO_WithRecvWindow", func(t *testing.T) {
		service := suite.client.NewMarginOpenOCOService()
		orders, err := service.
			RecvWindow(5000).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get open OCO orders with recvWindow: %v", err)
		}

		for _, order := range orders {
			if order.ContingencyType != "OCO" {
				t.Errorf("Expected contingencyType OCO, got %s", order.ContingencyType)
			}
		}
	})
}
