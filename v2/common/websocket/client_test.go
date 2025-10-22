package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/suite"
)

type testApiRequest struct {
	Id     string                 `json:"id"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

func (s *clientTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
}

type clientTestSuite struct {
	suite.Suite
	apiKey    string
	secretKey string
}

// findAvailablePort finds and returns an available port on localhost
func findAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, fmt.Errorf("failed to find available port: %w", err)
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}

// createTestWebSocketURL creates a websocket URL for the given port
func createTestWebSocketURL(port int) string {
	return fmt.Sprintf("ws://localhost:%d/ws", port)
}

func TestClient(t *testing.T) {
	suite.Run(t, new(clientTestSuite))
}

func (s *clientTestSuite) TestReadWriteSync() {
	// Find an available port
	port, err := findAvailablePort()
	s.Require().NoError(err)

	stopCh := make(chan struct{})
	go func() {
		startWsTestServer(port, stopCh)
	}()
	defer func() {
		stopCh <- struct{}{}
	}()

	// Give server time to start
	time.Sleep(100 * time.Millisecond)

	conn, err := NewConnection(func() (*websocket.Conn, error) {
		Dialer := websocket.Dialer{
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		}

		wsURL := createTestWebSocketURL(port)
		c, _, err := Dialer.Dial(wsURL, nil)
		if err != nil {
			return nil, err
		}

		return c, nil
	}, true, 10*time.Second)
	s.Require().NoError(err)

	client, err := NewClient(conn)
	s.Require().NoError(err)

	tests := []struct {
		name         string
		testCallback func()
	}{
		{
			name: "WriteSync success",
			testCallback: func() {
				id, err := uuid.NewRandom()
				s.Require().NoError(err)
				requestID := id.String()

				req := testApiRequest{
					Id:     requestID,
					Method: "some-method",
					Params: map[string]interface{}{},
				}
				reqRaw, err := json.Marshal(req)
				s.Require().NoError(err)

				responseRaw, err := client.WriteSync(requestID, reqRaw, 5*time.Second)
				s.Require().NoError(err)
				s.Require().Equal(reqRaw, responseRaw)
			},
		},
		{
			name: "WriteSync success read message with parallel writing",
			testCallback: func() {
				id, err := uuid.NewRandom()
				s.Require().NoError(err)
				requestID := id.String()

				req := testApiRequest{
					Id:     "some-other-request-id",
					Method: "some-method",
					Params: map[string]interface{}{},
				}
				reqRaw, err := json.Marshal(req)
				s.Require().NoError(err)

				err = client.Write(requestID, reqRaw)
				s.Require().NoError(err)

				req = testApiRequest{
					Id:     requestID,
					Method: "some-method",
					Params: map[string]interface{}{},
				}
				reqRaw, err = json.Marshal(req)
				s.Require().NoError(err)

				responseRaw, err := client.WriteSync(requestID, reqRaw, 5*time.Second)
				s.Require().NoError(err)
				s.Require().Equal(reqRaw, responseRaw)
			},
		},
		{
			name: "WriteSync timeout expired",
			testCallback: func() {
				id, err := uuid.NewRandom()
				s.Require().NoError(err)
				requestID := id.String()

				req := testApiRequest{
					Id:     requestID,
					Method: "some-method",
					Params: map[string]interface{}{
						"timeout": "true",
					},
				}
				reqRaw, err := json.Marshal(req)
				s.Require().NoError(err)

				responseRaw, err := client.WriteSync(requestID, reqRaw, 500*time.Millisecond)
				s.Require().Nil(responseRaw)
				s.Require().ErrorIs(err, ErrorWsReadConnectionTimeout)
			},
		},
		{
			name: "WriteAsync success",
			testCallback: func() {
				id, err := uuid.NewRandom()
				s.Require().NoError(err)
				requestID := id.String()

				req := testApiRequest{
					Id:     requestID,
					Method: "some-method",
					Params: map[string]interface{}{},
				}
				reqRaw, err := json.Marshal(req)
				s.Require().NoError(err)

				err = client.Write(requestID, reqRaw)
				s.Require().NoError(err)

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				select {
				case <-ctx.Done():
					s.T().Fatal("timeout waiting for write")
				case responseRaw := <-client.GetReadChannel():
					s.Require().Equal(reqRaw, responseRaw)
				case err := <-client.GetReadErrorChannel():
					s.T().Fatalf("unexpected error: '%v'", err)
				}
			},
		},
		{
			name: "WriteAsync success with waiter channel",
			testCallback: func() {
				id, err := uuid.NewRandom()
				s.Require().NoError(err)
				requestID := id.String()

				req := testApiRequest{
					Id:     requestID,
					Method: "some-method-with-waiter",
					Params: map[string]interface{}{},
				}
				reqRaw, err := json.Marshal(req)
				s.Require().NoError(err)

				waiter := make(chan []byte)

				err = client.Write(requestID, reqRaw, WithWaiter(waiter))
				s.Require().NoError(err)

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				select {
				case <-ctx.Done():
					s.T().Fatal("timeout waiting for write")
				case responseRaw := <-waiter:
					s.Require().Equal(reqRaw, responseRaw)
				case err := <-client.GetReadErrorChannel():
					s.T().Fatalf("unexpected error: '%v'", err)
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		s.T().Run(tt.name, func(t *testing.T) {
			tt.testCallback()
		})
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	conn.SetPingHandler(func(appData string) error {
		log.Println("Received ping:", appData)
		err := conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(10*time.Second))
		if err != nil {
			log.Println("Error sending pong:", err)
		}
		return nil
	})

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Printf("Received message: %s\n", message)

		req := testApiRequest{}
		if err := json.Unmarshal(message, &req); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		if req.Params["timeout"] == "true" {
			continue
		}

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func startWsTestServer(port int, stopCh chan struct{}) {
	addr := fmt.Sprintf("localhost:%d", port)
	server := &http.Server{
		Addr: addr,
	}

	http.HandleFunc("/ws", wsHandler)
	log.Printf("WebSocket server started on :%d", port)

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("WebSocket server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	<-stopCh

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("WebSocket shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}

// Memory benchmark tests for channel references in RequestList
func BenchmarkRequestList_ChannelMemory(b *testing.B) {
	tests := []struct {
		name        string
		mapSize     int
		channelType string
	}{
		{"SameChannel_1000entries", 1000, "same"},
		{"SameChannel_10000entries", 10000, "same"},
		{"DifferentChannels_1000entries", 1000, "different"},
		{"DifferentChannels_10000entries", 10000, "different"},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()

			for n := 0; n < b.N; n++ {
				requestList := NewRequestList()

				if tt.channelType == "same" {
					// Test: Multiple map entries with SAME channel reference
					sharedChannel := make(chan []byte, 1)

					for i := 0; i < tt.mapSize; i++ {
						requestID := fmt.Sprintf("req_%d_%d", n, i)
						requestList.Add(requestID, sharedChannel)
					}

				} else {
					// Test: Multiple map entries with DIFFERENT channels
					for i := 0; i < tt.mapSize; i++ {
						requestID := fmt.Sprintf("req_%d_%d", n, i)
						uniqueChannel := make(chan []byte, 1)
						requestList.Add(requestID, uniqueChannel)
					}
				}

				// Verify the map size
				if requestList.Len() != tt.mapSize {
					b.Fatalf("Expected %d entries, got %d", tt.mapSize, requestList.Len())
				}
			}
		})
	}
}

// Benchmark to demonstrate memory efficiency of channel references
func BenchmarkChannelReference_vs_ChannelCopy(b *testing.B) {
	const numEntries = 1000

	b.Run("ChannelReferences", func(b *testing.B) {
		b.ReportAllocs()

		for n := 0; n < b.N; n++ {
			// Simulate the actual RequestList behavior
			requests := make(map[string]chan []byte, numEntries)
			sharedChannel := make(chan []byte, 1)

			for i := 0; i < numEntries; i++ {
				requestID := fmt.Sprintf("req_%d", i)
				requests[requestID] = sharedChannel // Store reference
			}

			// Verify all entries point to the same channel
			firstChan := requests["req_0"]
			for i := 1; i < numEntries; i++ {
				requestID := fmt.Sprintf("req_%d", i)
				if requests[requestID] != firstChan {
					b.Fatal("Channels should be identical references")
				}
			}
		}
	})

	b.Run("UniqueChannels", func(b *testing.B) {
		b.ReportAllocs()

		for n := 0; n < b.N; n++ {
			requests := make(map[string]chan []byte, numEntries)

			for i := 0; i < numEntries; i++ {
				requestID := fmt.Sprintf("req_%d", i)
				requests[requestID] = make(chan []byte, 1) // Create unique channel
			}

			// Verify all channels are different
			for i := 1; i < numEntries; i++ {
				req0 := fmt.Sprintf("req_%d", 0)
				reqI := fmt.Sprintf("req_%d", i)
				if requests[req0] == requests[reqI] {
					b.Fatal("Channels should be unique")
				}
			}
		}
	})
}
