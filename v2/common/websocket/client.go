package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
)

//go:generate mockgen -source client.go -destination mock/client.go -package mock

const (
	// reconnectMinInterval define reconnect min interval
	reconnectMinInterval = 100 * time.Millisecond

	// reconnectMaxInterval define reconnect max interval
	reconnectMaxInterval = 10 * time.Second
)

var (
	// ErrorWsReadConnectionTimeout defines that connection read timeout expired
	ErrorWsReadConnectionTimeout = errors.New("ws error: read connection timeout")

	// ErrorWsIdAlreadySent defines that request with the same id was already sent
	ErrorWsIdAlreadySent = errors.New("ws error: request with same id already sent")

	// KeepAlivePingDeadline defines deadline to send ping frame
	KeepAlivePingDeadline = 10 * time.Second

	// WaitCheckInternal defines interval for ticker when it checks pending requests while stop application
	WaitCheckInternal = 300 * time.Millisecond
)

// messageId define id field of request/response
type messageId struct {
	Id string `json:"id"`
}

// client define API websocket client
type client struct {
	Debug                       bool
	logger                      *log.Logger
	conn                        Connection
	connMu                      sync.Mutex
	reconnectSignal             chan struct{}
	connectionEstablishedSignal chan struct{}
	requestsList                RequestList
	readC                       chan []byte
	readErrChan                 chan error
	reconnectCount              int64
}

func (c *client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.logger.Println(fmt.Sprintf(format, v...))
	}
}

// NewClient init client
func NewClient(conn Connection) (Client, error) {
	client := &client{
		logger:                      log.New(os.Stderr, "Binance-golang ", log.LstdFlags),
		conn:                        conn,
		connMu:                      sync.Mutex{},
		reconnectSignal:             make(chan struct{}, 1),
		connectionEstablishedSignal: make(chan struct{}, 1),
		requestsList:                NewRequestList(),
		readErrChan:                 make(chan error, 1),
		readC:                       make(chan []byte),
	}

	go client.handleReconnect()
	go client.read()

	return client, nil
}

type Client interface {
	Write(id string, data []byte) error
	WriteSync(id string, data []byte, timeout time.Duration) ([]byte, error)
	GetReadChannel() <-chan []byte
	GetReadErrorChannel() <-chan error
	GetReconnectCount() int64
	Wait(timeout time.Duration)
}

// Write sends data into websocket connection
func (c *client) Write(id string, data []byte) error {
	c.connMu.Lock()
	defer c.connMu.Unlock()

	if c.requestsList.IsAlreadyInList(id) {
		return ErrorWsIdAlreadySent
	}

	if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		c.debug("write: unable to write message into websocket conn '%v'", err)
		return err
	}

	c.requestsList.Add(id)

	return nil
}

// WriteSync sends data to the websocket connection and waits for a response synchronously
// Should be used separately from the asynchronous Write method (do not send anything in parallel)
func (c *client) WriteSync(id string, data []byte, timeout time.Duration) ([]byte, error) {
	c.connMu.Lock()
	defer c.connMu.Unlock()

	if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		c.debug("write sync: unable to write message into websocket conn '%v'", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			c.debug("write sync: timeout expired")
			return nil, ErrorWsReadConnectionTimeout
		case rawData := <-c.readC:
			// check that the correct response from websocket has been read
			msg := messageId{}
			err := json.Unmarshal(rawData, &msg)
			if err != nil {
				return nil, err
			}
			if msg.Id != id {
				c.debug("write sync: wrong response with id '%v' has been read", msg.Id)
				continue
			}

			return rawData, nil
		case err := <-c.readErrChan:
			c.debug("write sync: error read '%v'", err)
			return nil, err
		}
	}
}

func (c *client) GetReadChannel() <-chan []byte {
	return c.readC
}

func (c *client) GetReadErrorChannel() <-chan error {
	return c.readErrChan
}

func (c *client) Wait(timeout time.Duration) {
	c.wait(timeout)
}

// read data from connection
func (c *client) read() {
	defer func() {
		// reading from closed connection 1000 times caused panic
		// prevent panic for any case
		if r := recover(); r != nil {
		}
	}()

	for {
		c.debug("read: waiting for message")
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.debug("read: error reading message '%v'", err)
			c.reconnectSignal <- struct{}{}
			c.readErrChan <- err

			c.debug("read: wait to get connected")
			<-c.connectionEstablishedSignal

			// refresh map after reconnect to avoid useless waiting after stop application
			c.requestsList.RecreateList()

			c.debug("read: connection established")
			continue
		}
		c.debug("read: got new message")

		msg := messageId{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			c.debug("read: error unmarshalling message '%v'", err)
			c.readErrChan <- err
			continue
		}

		c.debug("read: sending message into read channel '%v'", msg)
		c.readC <- message

		c.debug("read: remove message from request list '%v'", msg)
		c.requestsList.Remove(msg.Id)
	}
}

// wait until all responses received
// make sure that you are not sending requests
func (c *client) wait(timeout time.Duration) {
	timeoutTimer := time.NewTimer(timeout)
	defer timeoutTimer.Stop()

	checkTicker := time.NewTicker(WaitCheckInternal)
	defer checkTicker.Stop()

	for {
		select {
		case <-checkTicker.C:
			if c.requestsList.Len() == 0 {
				return
			}
		case <-timeoutTimer.C:
			return
		}
	}
}

// handleReconnect waits for reconnect signal and starts reconnect
func (c *client) handleReconnect() {
	for range c.reconnectSignal {
		c.debug("reconnect: received signal")

		b := &backoff.Backoff{
			Min:    reconnectMinInterval,
			Max:    reconnectMaxInterval,
			Factor: 1.8,
			Jitter: false,
		}

		conn := c.startReconnect(b)

		b.Reset()

		c.connMu.Lock()
		c.conn = conn
		c.connMu.Unlock()

		c.debug("reconnect: connected")
		c.connectionEstablishedSignal <- struct{}{}
	}
}

// startReconnect starts reconnect loop with increasing delay
func (c *client) startReconnect(b *backoff.Backoff) Connection {
	for {
		atomic.AddInt64(&c.reconnectCount, 1)
		conn, err := c.conn.RestoreConnection()
		if err != nil {
			delay := b.Duration()
			c.debug("reconnect: error while reconnecting. try in %s", delay.Round(time.Millisecond))
			time.Sleep(delay)
			continue
		}

		return conn
	}
}

// GetReconnectCount returns reconnect counter value
func (c *client) GetReconnectCount() int64 {
	return atomic.LoadInt64(&c.reconnectCount)
}

// NewRequestList creates request list
func NewRequestList() RequestList {
	return RequestList{
		mu:       sync.Mutex{},
		requests: make(map[string]struct{}), // TODO preallocate buckets
	}
}

// RequestList state of requests that was sent/received
type RequestList struct {
	mu       sync.Mutex
	requests map[string]struct{}
}

// Add adds request into list
func (l *RequestList) Add(id string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.requests[id] = struct{}{}
}

// RecreateList creates new request list
func (l *RequestList) RecreateList() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.requests = make(map[string]struct{})
}

// Remove adds request from list
func (l *RequestList) Remove(id string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.requests, id)
}

// Len get list length
func (l *RequestList) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.requests)
}

// IsAlreadyInList checks if id is presented in requests list
func (l *RequestList) IsAlreadyInList(id string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.requests[id]; ok {
		return true
	}

	return false
}

// NewConnection constructor for connection
func NewConnection(
	initUnderlyingWsConnFn func() (*websocket.Conn, error),
	isKeepAliveNeeded bool,
	keepaliveTimeout time.Duration,
) (Connection, error) {
	underlyingWsConn, err := initUnderlyingWsConnFn()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	wsConn := &connection{
		conn:                   underlyingWsConn,
		connectionMu:           sync.Mutex{},
		lastResponseMu:         sync.Mutex{},
		initUnderlyingWsConnFn: initUnderlyingWsConnFn,
		keepaliveTimeout:       keepaliveTimeout,
		isKeepAliveNeeded:      isKeepAliveNeeded,
		ctx:                    ctx,
		cancel:                 cancel,
	}

	if isKeepAliveNeeded {
		go wsConn.keepAlive(keepaliveTimeout)
	}

	return wsConn, nil
}

// connection is an instance of single ws connection with keepalive handler
type connection struct {
	conn                   *websocket.Conn
	connectionMu           sync.Mutex
	lastResponse           time.Time
	lastResponseMu         sync.Mutex
	initUnderlyingWsConnFn func() (*websocket.Conn, error)
	keepaliveTimeout       time.Duration
	isKeepAliveNeeded      bool
	ctx                    context.Context
	cancel                 context.CancelFunc
}

type Connection interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
	RestoreConnection() (Connection, error)
}

// WriteMessage is a thread-safe method for conn.WriteMessage
func (c *connection) WriteMessage(messageType int, data []byte) error {
	c.connectionMu.Lock()
	defer c.connectionMu.Unlock()
	return c.conn.WriteMessage(messageType, data)
}

// ReadMessage wrapper for conn.ReadMessage
func (c *connection) ReadMessage() (int, []byte, error) {
	msgType, msg, err := c.conn.ReadMessage()
	if err != nil {
		c.cancel()
	}
	return msgType, msg, err
}

// RestoreConnection recreates ws connection with the same underlying connection callback and keepalive timeout
func (c *connection) RestoreConnection() (Connection, error) {
	return NewConnection(c.initUnderlyingWsConnFn, c.isKeepAliveNeeded, c.keepaliveTimeout)
}

// keepAlive handles ping-pong for connection
func (c *connection) keepAlive(timeout time.Duration) {
	c.updateLastResponse()

	c.conn.SetPongHandler(func(msg string) error {
		c.updateLastResponse()
		return nil
	})

	go func() {
		ticker := time.NewTicker(timeout)
		defer ticker.Stop()

		for {
			select {
			case <-c.ctx.Done():
				return
			case <-ticker.C:
				err := c.ping()
				if err != nil {
					return
				}

				if c.isLastResponseOutdated(timeout) {
					c.close()
					return
				}
			}
		}
	}()
}

// updateLastResponse sets lastResponse now
func (c *connection) updateLastResponse() {
	c.lastResponseMu.Lock()
	defer c.lastResponseMu.Unlock()
	c.lastResponse = time.Now()
}

// isLastResponseOutdated checks that time since last pong message exceeded timeout
func (c *connection) isLastResponseOutdated(timeout time.Duration) bool {
	c.lastResponseMu.Lock()
	defer c.lastResponseMu.Unlock()
	return time.Since(c.lastResponse) > timeout
}

// close thread-safe method for closing connection
func (c *connection) close() error {
	c.connectionMu.Lock()
	defer c.connectionMu.Unlock()
	return c.conn.Close()
}

// ping thread-safe method sending ping message
func (c *connection) ping() error {
	c.connectionMu.Lock()
	defer c.connectionMu.Unlock()

	deadline := time.Now().Add(KeepAlivePingDeadline)
	return c.conn.WriteControl(websocket.PingMessage, []byte{}, deadline)
}
