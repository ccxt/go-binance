package portfolio

import (
	"context"
	"testing"
)

type umModifyOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMModifyOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umModifyOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("ModifyUMOrder_ByOrderID", func(t *testing.T) {
		service := suite.client.NewUMModifyOrderService()
		res, err := service.Symbol("BTCUSDT").
			Side(SideTypeBuy).
			Quantity("1").
			Price("30005").
			OrderID(20072994037).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to modify UM order: %v", err)
		}

		// Basic validation of returned data
		if res.OrderID != 20072994037 {
			t.Errorf("Expected order ID 20072994037, got %d", res.OrderID)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
		if res.Side != "BUY" {
			t.Errorf("Expected side BUY, got %s", res.Side)
		}
	})

	t.Run("ModifyUMOrder_ByClientOrderID", func(t *testing.T) {
		service := suite.client.NewUMModifyOrderService()
		res, err := service.Symbol("BTCUSDT").
			Side(SideTypeBuy).
			Quantity("1").
			PriceMatch("OPPONENT").
			OrigClientOrderID("LJ9R4QZDihCaS8UAOOLpgW").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to modify UM order: %v", err)
		}

		// Basic validation of returned data
		if res.ClientOrderID != "LJ9R4QZDihCaS8UAOOLpgW" {
			t.Errorf("Expected client order ID LJ9R4QZDihCaS8UAOOLpgW, got %s", res.ClientOrderID)
		}
		if res.PriceMatch != "OPPONENT" {
			t.Errorf("Expected price match OPPONENT, got %s", res.PriceMatch)
		}
	})

	t.Run("ModifyUMOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewUMModifyOrderService()
		_, err := service.Symbol("BTCUSDT").
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
