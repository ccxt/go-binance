package options

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
	Proxy    *string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
		Proxy:    getWsProxyUrl(),
	}
}

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	proxy := http.ProxyFromEnvironment
	if cfg.Proxy != nil {
		u, err := url.Parse(*cfg.Proxy)
		if err != nil {
			return nil, nil, err
		}
		proxy = http.ProxyURL(u)
	}

	Dialer := websocket.Dialer{
		Proxy:             proxy,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: true,
	}

	c, _, err := Dialer.Dial(cfg.Endpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	c.SetReadLimit(655350)
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			c.Close()
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()

	c.SetPingHandler(func(pingData string) error {
		// Respond with Pong using the server's PING payload
		err := c.WriteControl(
			websocket.PongMessage,
			[]byte(pingData),
			time.Now().Add(WebsocketPongTimeout), // Short deadline to ensure timely response
		)
		if err != nil {
			return err
		}

		lastResponse = time.Now()

		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}
