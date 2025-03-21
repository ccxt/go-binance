//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type bnbTransferServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestBNBTransferServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &bnbTransferServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("BNBTransfer", func(t *testing.T) {
		service := &BNBTransferService{c: suite.client}
		res, err := service.
			Amount("0.1").
			TransferSide(TransferSideToUM).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to execute BNB transfer: %v", err)
		}

		if res.TranID == 0 {
			t.Error("Expected non-zero transaction ID")
		}

		// Transfer back
		_, err = service.
			Amount("0.1").
			TransferSide(TransferSideFromUM).
			Do(context.Background())

		if err != nil {
			t.Fatalf("Failed to transfer BNB back: %v", err)
		}
	})
}
