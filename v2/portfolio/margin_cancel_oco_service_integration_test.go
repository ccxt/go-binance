package portfolio

import (
	"context"
	"testing"
)

type marginCancelOCOServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginCancelOCOServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginCancelOCOServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("CancelMarginOCO_ByOrderListID", func(t *testing.T) {
		service := suite.client.NewMarginCancelOCOService()
		res, err := service.Symbol("LTCBTC").
			OrderListID(0).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel margin OCO: %v", err)
		}

		// Basic validation of returned data
		if res.OrderListID != 0 {
			t.Errorf("Expected order list ID 0, got %d", res.OrderListID)
		}
		if res.ListStatusType != "ALL_DONE" {
			t.Errorf("Expected list status type ALL_DONE, got %s", res.ListStatusType)
		}
		if res.Symbol != "LTCBTC" {
			t.Errorf("Expected symbol LTCBTC, got %s", res.Symbol)
		}
	})

	t.Run("CancelMarginOCO_ByListClientOrderID", func(t *testing.T) {
		service := suite.client.NewMarginCancelOCOService()
		res, err := service.Symbol("LTCBTC").
			ListClientOrderID("C3wyj4WVEktd7u9aVBRXcN").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to cancel margin OCO: %v", err)
		}

		// Basic validation of returned data
		if res.ListClientOrderID != "C3wyj4WVEktd7u9aVBRXcN" {
			t.Errorf("Expected list client order ID C3wyj4WVEktd7u9aVBRXcN, got %s", res.ListClientOrderID)
		}
		if res.ListStatusType != "ALL_DONE" {
			t.Errorf("Expected list status type ALL_DONE, got %s", res.ListStatusType)
		}
	})

	t.Run("CancelMarginOCO_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewMarginCancelOCOService()
		_, err := service.Symbol("LTCBTC").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected an error when neither orderListId nor listClientOrderId is provided")
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
