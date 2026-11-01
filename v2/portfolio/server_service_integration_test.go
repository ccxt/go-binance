//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type serverServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestServerServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &serverServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("Ping", func(t *testing.T) {
		service := &PingService{c: suite.client}
		err := service.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to ping server: %v", err)
		}
	})
}
