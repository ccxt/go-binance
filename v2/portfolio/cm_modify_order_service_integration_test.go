package portfolio

import (
	"context"
	"testing"
)

type cmModifyOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestCMModifyOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &cmModifyOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("ModifyCMOrder_ByOrderID", func(t *testing.T) {
		service := suite.client.NewCMModifyOrderService()
		res, err := service.Symbol("BTCUSD_PERP").
			Side(SideTypeBuy).
			Quantity("1").
			Price("30005").
			OrderID(20072994037).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to modify CM order: %v", err)
		}

		// Basic validation of returned data
		if res.OrderID != 20072994037 {
			t.Errorf("Expected order ID 20072994037, got %d", res.OrderID)
		}
		if res.Symbol != "BTCUSD_PERP" {
			t.Errorf("Expected symbol BTCUSD_PERP, got %s", res.Symbol)
		}
		if res.Pair != "BTCUSD" {
			t.Errorf("Expected pair BTCUSD, got %s", res.Pair)
		}
	})

	t.Run("ModifyCMOrder_ByClientOrderID", func(t *testing.T) {
		service := suite.client.NewCMModifyOrderService()
		res, err := service.Symbol("BTCUSD_PERP").
			Side(SideTypeBuy).
			Quantity("1").
			Price("30005").
			OrigClientOrderID("LJ9R4QZDihCaS8UAOOLpgW").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to modify CM order: %v", err)
		}

		// Basic validation of returned data
		if res.ClientOrderID != "LJ9R4QZDihCaS8UAOOLpgW" {
			t.Errorf("Expected client order ID LJ9R4QZDihCaS8UAOOLpgW, got %s", res.ClientOrderID)
		}
		if res.Status != "NEW" {
			t.Errorf("Expected status NEW, got %s", res.Status)
		}
	})

	t.Run("ModifyCMOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewCMModifyOrderService()
		_, err := service.Symbol("BTCUSD_PERP").
			Side(SideTypeBuy).
			Quantity("1").
			Price("30005").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected an error when neither orderId nor origClientOrderId is provided")
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
