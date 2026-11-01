package portfolio

import (
	"context"
	"testing"
)

type umFeeBurnServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMFeeBurnServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umFeeBurnServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("ToggleFeeBurn_Enable", func(t *testing.T) {
		service := suite.client.NewUMFeeBurnService()
		res, err := service.FeeBurn(true).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to enable fee burn: %v", err)
		}

		if res.Code != 200 {
			t.Errorf("Expected code 200, got %d", res.Code)
		}
		if res.Msg != "success" {
			t.Errorf("Expected msg 'success', got %s", res.Msg)
		}
	})

	t.Run("ToggleFeeBurn_Disable", func(t *testing.T) {
		service := suite.client.NewUMFeeBurnService()
		res, err := service.FeeBurn(false).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to disable fee burn: %v", err)
		}

		if res.Code != 200 {
			t.Errorf("Expected code 200, got %d", res.Code)
		}
		if res.Msg != "success" {
			t.Errorf("Expected msg 'success', got %s", res.Msg)
		}
	})

	t.Run("ToggleFeeBurn_WithRecvWindow", func(t *testing.T) {
		service := suite.client.NewUMFeeBurnService()
		res, err := service.
			FeeBurn(true).
			RecvWindow(5000).
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to toggle fee burn with recvWindow: %v", err)
		}

		if res.Code != 200 {
			t.Errorf("Expected code 200, got %d", res.Code)
		}
		if res.Msg != "success" {
			t.Errorf("Expected msg 'success', got %s", res.Msg)
		}
	})
}
