package websocket

// WebSocketServiceOption represents a functional option for WebSocket services
type WebSocketServiceOption func(serviceOpt *WebSocketServiceCreateOption)

type WebSocketServiceCreateOption struct {
	Client     Client
	RecvWindow int64
}

// WithWebSocketClient creates an option to set the websocket Client for any WebSocket service
func WithWebSocketClient(client Client) WebSocketServiceOption {
	return func(opt *WebSocketServiceCreateOption) {
		opt.Client = client
	}
}

// WithRecvWindow creates an option to set the receive window for WebSocket services
func WithRecvWindow(recvWindow int64) WebSocketServiceOption {
	return func(opt *WebSocketServiceCreateOption) {
		opt.RecvWindow = recvWindow
	}
}
