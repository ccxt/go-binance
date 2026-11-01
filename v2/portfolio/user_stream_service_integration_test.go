//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
	"time"
)

type userStreamServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUserStreamServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &userStreamServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("TestUserStreamLifecycle", func(t *testing.T) {
		// Start user stream
		startService := &StartUserStreamService{c: suite.client}
		listenKey, err := startService.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to start user stream: %v", err)
		}
		if listenKey == "" {
			t.Error("Expected non-empty listen key")
		}

		// Keep alive user stream
		keepaliveService := &KeepaliveUserStreamService{c: suite.client}
		err = keepaliveService.ListenKey(listenKey).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to keepalive user stream: %v", err)
		}

		// Wait a bit before closing
		time.Sleep(2 * time.Second)

		// Close user stream
		closeService := &CloseUserStreamService{c: suite.client}
		err = closeService.ListenKey(listenKey).Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to close user stream: %v", err)
		}
	})

	t.Run("TestInvalidListenKey", func(t *testing.T) {
		// Try to keepalive with invalid listen key
		keepaliveService := &KeepaliveUserStreamService{c: suite.client}
		err := keepaliveService.ListenKey("invalid_listen_key").Do(context.Background())
		if err == nil {
			t.Error("Expected error with invalid listen key")
		}

		// Try to close with invalid listen key
		closeService := &CloseUserStreamService{c: suite.client}
		err = closeService.ListenKey("invalid_listen_key").Do(context.Background())
		if err == nil {
			t.Error("Expected error with invalid listen key")
		}
	})
}
