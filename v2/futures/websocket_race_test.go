package futures

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// TestKeepAliveNoRace verifies that keepAlive's lastResponse variable
// doesn't trigger the race detector when pings arrive concurrently
// with the ticker reads.
func TestKeepAliveNoRace(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close()

		// Send pings rapidly to trigger the ping handler on the client side
		for i := 0; i < 50; i++ {
			if err := c.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(time.Second)); err != nil {
				return
			}
			time.Sleep(5 * time.Millisecond)
		}
	}))
	defer srv.Close()

	wsURL := "ws" + strings.TrimPrefix(srv.URL, "http")
	c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()

	// keepAlive with a short timeout so the ticker fires frequently
	keepAlive(c, 100*time.Millisecond)

	// Read messages to drive the ping handler (ReadMessage dispatches control frames)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				return
			}
		}
	}()

	// Let it run for enough time that ticker and pings overlap
	time.Sleep(300 * time.Millisecond)
	c.Close()
	<-done
}

// TestWsSilentNoRace verifies that the silent variable in wsServe
// doesn't trigger the race detector when stopC is closed during ReadMessage.
func TestWsSilentNoRace(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close()

		// Send a few messages then hang so the client blocks on ReadMessage
		for i := 0; i < 5; i++ {
			c.WriteMessage(websocket.TextMessage, []byte("hello"))
			time.Sleep(10 * time.Millisecond)
		}
		// Hold connection open
		time.Sleep(500 * time.Millisecond)
	}))
	defer srv.Close()

	wsURL := "ws" + strings.TrimPrefix(srv.URL, "http")

	origKeepalive := WebsocketKeepalive
	WebsocketKeepalive = false
	defer func() { WebsocketKeepalive = origKeepalive }()

	cfg := &WsConfig{Endpoint: wsURL}

	var received int
	var mu sync.Mutex
	handler := func(msg []byte) {
		mu.Lock()
		received++
		mu.Unlock()
	}
	errHandler := func(err error) {}

	doneC, stopC, err := wsServe(cfg, handler, errHandler)
	if err != nil {
		t.Fatalf("wsServe: %v", err)
	}

	// Let some messages arrive
	time.Sleep(80 * time.Millisecond)

	// Close stopC which sets silent=true in one goroutine
	// while ReadMessage loop checks it in another — this is the race we fixed
	close(stopC)
	<-doneC

	mu.Lock()
	defer mu.Unlock()
	if received == 0 {
		t.Error("expected to receive at least one message")
	}
}
