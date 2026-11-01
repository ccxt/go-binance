//go:build integration
// +build integration

package portfolio

import (
	"context"
	"sync"
	"testing"
	"time"
)

type websocketServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestWebsocketServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &websocketServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("TestUserDataWebsocket", func(t *testing.T) {
		// First get a listen key
		startService := &StartUserStreamService{c: suite.client}
		listenKey, err := startService.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get listen key: %v", err)
		}

		// Create channels to track received events
		receivedEvent := make(chan struct{})
		var wg sync.WaitGroup
		wg.Add(1)

		// Define handlers
		wsHandler := func(event *WsUserDataEvent) {
			defer wg.Done()
			// Verify basic event properties
			if event.Time == 0 {
				t.Error("Expected non-zero event time")
			}
			if event.Event == "" {
				t.Error("Expected non-empty event type")
			}
			close(receivedEvent)
		}

		errHandler := func(err error) {
			t.Errorf("Websocket error: %v", err)
		}

		// Start websocket connection
		doneC, stopC, err := WsUserDataServe(listenKey, wsHandler, errHandler)
		if err != nil {
			t.Fatalf("Failed to serve websocket: %v", err)
		}

		// Cleanup function
		defer func() {
			close(stopC)
			<-doneC
		}()

		// Keep the connection alive
		go func() {
			keepaliveService := &KeepaliveUserStreamService{c: suite.client}
			for {
				select {
				case <-stopC:
					return
				case <-time.After(30 * time.Second):
					err := keepaliveService.ListenKey(listenKey).Do(context.Background())
					if err != nil {
						t.Logf("Keepalive error: %v", err)
					}
				}
			}
		}()

		// Wait for event or timeout
		select {
		case <-receivedEvent:
			// Success - event received
		case <-time.After(2 * time.Minute):
			t.Fatal("Timeout waiting for websocket event")
		}

		// Wait for handler to complete
		wg.Wait()
	})

	t.Run("TestWebsocketReconnection", func(t *testing.T) {
		// Get listen key
		startService := &StartUserStreamService{c: suite.client}
		listenKey, err := startService.Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to get listen key: %v", err)
		}

		wsHandler := func(event *WsUserDataEvent) {
			// Just log events
			t.Logf("Received event: %v", event.Event)
		}

		errHandler := func(err error) {
			t.Logf("Websocket error: %v", err)
		}

		// Test multiple connections
		for i := 0; i < 3; i++ {
			doneC, stopC, err := WsUserDataServe(listenKey, wsHandler, errHandler)
			if err != nil {
				t.Fatalf("Failed to serve websocket on attempt %d: %v", i+1, err)
			}

			// Keep connection for a short period
			time.Sleep(5 * time.Second)

			// Clean up
			close(stopC)
			<-doneC
		}
	})

	t.Run("TestInvalidListenKey", func(t *testing.T) {
		invalidListenKey := "invalid_listen_key"

		wsHandler := func(event *WsUserDataEvent) {
			t.Error("Should not receive any events with invalid listen key")
		}

		errHandler := func(err error) {
			// Expected to get an error
			t.Logf("Got expected error: %v", err)
		}

		_, stopC, err := WsUserDataServe(invalidListenKey, wsHandler, errHandler)
		if err == nil {
			close(stopC)
			t.Error("Expected error with invalid listen key")
		}
	})
}
