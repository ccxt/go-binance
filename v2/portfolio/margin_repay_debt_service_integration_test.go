package portfolio

import (
	"context"
	"testing"
)

type marginRepayDebtServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestMarginRepayDebtServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &marginRepayDebtServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("RepayDebt_Basic", func(t *testing.T) {
		asset := "BNB"
		amount := "0.1"
		service := suite.client.NewMarginRepayDebtService()
		res, err := service.Asset(asset).
			Amount(amount).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to repay debt: %v", err)
		}

		if !res.Success {
			t.Error("Expected successful repayment")
		}
		if res.Asset != asset {
			t.Errorf("Expected asset %s, got %s", asset, res.Asset)
		}
	})

	t.Run("RepayDebt_WithSpecificAssets", func(t *testing.T) {
		asset := "BNB"
		specifyRepayAssets := "USDT,BTC"
		service := suite.client.NewMarginRepayDebtService()
		res, err := service.Asset(asset).
			SpecifyRepayAssets(specifyRepayAssets).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to repay debt with specific assets: %v", err)
		}

		if !res.Success {
			t.Error("Expected successful repayment")
		}
		if len(res.SpecifyRepayAssets) == 0 {
			t.Error("Expected specific repay assets in response")
		}
	})

	t.Run("RepayDebt_WithRecvWindow", func(t *testing.T) {
		asset := "BNB"
		amount := "0.1"
		recvWindow := int64(5000)
		service := suite.client.NewMarginRepayDebtService()
		res, err := service.Asset(asset).
			Amount(amount).
			RecvWindow(recvWindow).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to repay debt with recvWindow: %v", err)
		}

		if !res.Success {
			t.Error("Expected successful repayment")
		}
	})
}
