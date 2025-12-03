package websocket

import (
	"context"
	"encoding/json"
	"errors"
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

func TestClient(t *testing.T) {
	suite.Run(t, new(clientTestSuite))
}

func (s *clientTestSuite) TestReadWriteSync() {
	stopCh := make(chan struct{})
	readyCh := make(chan struct{})
	go func() {
		startWsTestServer(stopCh, readyCh)
	}()
	defer func() {
		stopCh <- struct{}{}
	}()

	<-readyCh

	conn, err := NewConnection(func() (*websocket.Conn, error) {
		Dialer := websocket.Dialer{
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		}

		c, _, err := Dialer.Dial("ws://localhost:8080/ws", nil)
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

func startWsTestServer(stopCh chan struct{}, readyCh chan struct{}) {
	server := &http.Server{}

	http.HandleFunc("/ws", wsHandler)
	log.Println("WebSocket server started on :8080")

	go func() {
		lis, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("WebSocket server error: %v", err)
		}
		close(readyCh)

		if err := server.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
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
