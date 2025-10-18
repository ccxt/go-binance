package websocket

// WebSocketServiceOption represents a functional option for WebSocket services
type WebSocketServiceOption func(serviceOpt *WebSocketServiceCreateOption)

type WebSocketServiceCreateOption struct {
	Client Client
}

// WithWebSocketClient creates an option to set the websocket Client for any WebSocket service
func WithWebSocketClient(client Client) WebSocketServiceOption {
	return func(opt *WebSocketServiceCreateOption) {
		opt.Client = client
	}
}
