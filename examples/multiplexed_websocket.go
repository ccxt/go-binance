package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/adshao/go-binance/v2/common/websocket"
	"github.com/adshao/go-binance/v2/futures"
)

const (
	requestsCount     = 100
	responseTimeout   = 30 * time.Second
	requestRateLimit  = 100 * time.Millisecond
	maxConcurrency    = 10
	channelBufferSize = 10
)

// MultiplexedFuturesWebSocketExample demonstrates concurrent WebSocket requests
// with waiter channels for 1:1 request/response matching
func MultiplexedFuturesWebSocketExample() {
	manager, err := NewWebSocketManager()
	if err != nil {
		slog.Error("Failed to create WebSocket manager", "error", err)
		return
	}

	manager.logger.Info("=== Enhanced Multiplexed Futures WebSocket Example ===")
	manager.logger.Info("Testing request/response tracking and waiter channels")

	// Create context for this operation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2) // 2 services will run

	// Error channel to collect errors from goroutines
	errorChan := make(chan error, 3) // Buffer for all potential errors

	// Start monitoring and services
	go manager.startFallbackMonitoring(ctx)
	go manager.startOrderPlaceService(ctx, &wg, errorChan)
	go manager.startOrderCancelService(ctx, &wg, errorChan)

	// Monitor for errors in a separate goroutine
	go func() {
		for err := range errorChan {
			if err != nil {
				manager.logger.Error("Service error", "error", err)
			}
		}
	}()

	wg.Wait()

	manager.logger.Info("All services completed")
	manager.printFinalStats()
}

// createWebSocketClient creates a new WebSocket client
func createWebSocketClient(logger *slog.Logger) (websocket.Client, error) {
	conn, err := websocket.NewConnection(
		futures.WsApiInitReadWriteConn,
		futures.WebsocketKeepalive,
		futures.WebsocketTimeoutReadWriteConnection,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create WebSocket connection: %w", err)
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to create WebSocket client: %w", err)
	}

	logger.Info("Shared futures WebSocket client created successfully")
	return client, nil
}

// startFallbackMonitoring monitors the shared channel for fallback messages
func (wm *WebSocketManager) startFallbackMonitoring(ctx context.Context) {
	wm.logger.Info("Starting fallback message monitoring")

	for {
		select {
		case data := <-wm.client.GetReadChannel():
			wm.logger.Error("FALLBACK: Message received on shared channel (should be routed to waiter)",
				"data", string(data))
		case err := <-wm.client.GetReadErrorChannel():
			if err != nil {
				wm.logger.Error("Shared channel error", "error", err)
				return
			}
		case <-ctx.Done():
			wm.logger.Info("Fallback monitoring stopped")
			return
		}
	}
}

// startOrderPlaceService handles order placement requests
func (wm *WebSocketManager) startOrderPlaceService(ctx context.Context, wg *sync.WaitGroup, errorChan chan<- error) {
	defer wg.Done()

	// Create order place service
	orderPlaceService, err := futures.NewOrderPlaceWsService(
		AppConfig.APIKey,
		AppConfig.SecretKey,
		websocket.WithWebSocketClient(wm.client),
	)
	if err != nil {
		errorChan <- fmt.Errorf("failed to create order place service: %w", err)
		return
	}

	// Create dedicated waiter channel with proper buffer
	waiterChannel := make(chan []byte, channelBufferSize)
	defer close(waiterChannel)

	// Initialize stats
	wm.placeStats.SetStartTime()

	// Start request sender and response handler concurrently
	var serviceWg sync.WaitGroup
	serviceWg.Add(2)

	requestErrors := make(chan error, requestsCount) // Buffer for all possible request errors

	go func() {
		defer serviceWg.Done()
		wm.sendOrderPlaceRequests(ctx, orderPlaceService, waiterChannel, requestErrors)
		close(requestErrors)
	}()

	go func() {
		defer serviceWg.Done()
		wm.handleOrderPlaceResponses(ctx, waiterChannel, requestErrors, errorChan)
	}()

	serviceWg.Wait()
}

// handleOrderPlaceResponses handles responses from order place service
func (wm *WebSocketManager) handleOrderPlaceResponses(ctx context.Context, waiterChannel <-chan []byte, requestErrors <-chan error, errorChan chan<- error) {
	serviceLogger := wm.logger.With("service", "order_place")
	timeout := time.NewTimer(responseTimeout)
	defer timeout.Stop()

	// Monitor request errors
	go func() {
		for err := range requestErrors {
			if err != nil {
				serviceLogger.Error("Request error", "error", err)
			}
		}
	}()

	for {
		select {
		case data, ok := <-waiterChannel:
			if !ok {
				serviceLogger.Info("WaiterChannel closed")
				wm.placeStats.SetEndTime()
				return
			}
			response := string(data)
			serviceLogger.Debug("Response received", "response", response)

			if !strings.Contains(response, "place_") {
				err := fmt.Errorf("response does not contain expected prefix 'place_': %s", response)
				errorChan <- err
				wm.placeStats.SetEndTime()
				return
			}

			wm.placeStats.IncrementResponses()
			sent, received := wm.placeStats.GetCounts()
			serviceLogger.Info("Progress", "responses_received", received, "requests_sent", sent)

			if received >= requestsCount {
				serviceLogger.Info("All responses received successfully", "total", received)
				wm.placeStats.SetEndTime()
				return
			}

		case <-timeout.C:
			sent, received := wm.placeStats.GetCounts()
			err := fmt.Errorf("timeout reached: sent=%d, received=%d, missing=%d", sent, received, sent-received)
			errorChan <- err
			wm.placeStats.SetEndTime()
			return

		case <-ctx.Done():
			serviceLogger.Info("Response handler stopped")
			wm.placeStats.SetEndTime()
			return
		}
	}
}

// sendOrderPlaceRequests sends multiple order placement requests concurrently
func (wm *WebSocketManager) sendOrderPlaceRequests(ctx context.Context, orderPlaceService *futures.OrderPlaceWsService, waiterChannel chan []byte, errorChan chan<- error) {
	serviceLogger := wm.logger.With("service", "order_place")

	// Build request
	request := futures.NewOrderPlaceWsRequest().
		Symbol("BTCUSDT").
		Side(futures.SideTypeBuy).
		Type(futures.OrderTypeMarket).
		TimeInForce(futures.TimeInForceTypeFOK).
		Quantity("0.00001")

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrency)

	for i := 0; i < requestsCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire semaphore
			defer func() { <-semaphore }() // Release semaphore

			requestID := fmt.Sprintf("place_%d_%d", time.Now().UnixNano(), index)
			serviceLogger.Debug("Sending request", "request_id", requestID, "index", index)

			if err := orderPlaceService.Do(requestID, request, websocket.WithWaiter(waiterChannel)); err != nil {
				errorChan <- fmt.Errorf("failed to send request %s: %w", requestID, err)
				return
			}
			wm.placeStats.IncrementRequests()
		}(i)

		// Rate limiting
		time.Sleep(requestRateLimit)
	}

	wg.Wait()
	serviceLogger.Info("All order place requests sent")
}

// startOrderCancelService handles order cancellation requests
func (wm *WebSocketManager) startOrderCancelService(ctx context.Context, wg *sync.WaitGroup, errorChan chan<- error) {
	defer wg.Done()

	// Create order cancel service
	orderCancelService, err := futures.NewOrderCancelWsService(
		AppConfig.APIKey,
		AppConfig.SecretKey,
		websocket.WithWebSocketClient(wm.client),
	)
	if err != nil {
		errorChan <- fmt.Errorf("failed to create order cancel service: %w", err)
		return
	}

	// Create dedicated waiter channel with proper buffer
	waiterChannel := make(chan []byte, channelBufferSize)
	defer close(waiterChannel)

	// Initialize stats
	wm.cancelStats.SetStartTime()

	// Start request sender and response handler concurrently
	var serviceWg sync.WaitGroup
	serviceWg.Add(2)

	requestErrors := make(chan error, requestsCount) // Buffer for all possible request errors

	go func() {
		defer serviceWg.Done()
		wm.sendOrderCancelRequests(ctx, orderCancelService, waiterChannel, requestErrors)
		close(requestErrors)
	}()

	go func() {
		defer serviceWg.Done()
		wm.handleOrderCancelResponses(ctx, waiterChannel, requestErrors, errorChan)
	}()

	serviceWg.Wait()
}

// handleOrderCancelResponses handles responses from order cancel service
func (wm *WebSocketManager) handleOrderCancelResponses(ctx context.Context, waiterChannel <-chan []byte, requestErrors <-chan error, errorChan chan<- error) {
	serviceLogger := wm.logger.With("service", "order_cancel")
	timeout := time.NewTimer(responseTimeout)
	defer timeout.Stop()

	// Monitor request errors
	go func() {
		for err := range requestErrors {
			if err != nil {
				serviceLogger.Error("Request error", "error", err)
			}
		}
	}()

	for {
		select {
		case data, ok := <-waiterChannel:
			if !ok {
				serviceLogger.Info("WaiterChannel closed")
				wm.cancelStats.SetEndTime()
				return
			}
			response := string(data)
			serviceLogger.Debug("Response received", "response", response)

			if !strings.Contains(response, "cancel_") {
				err := fmt.Errorf("response does not contain expected prefix 'cancel_': %s", response)
				errorChan <- err
				wm.cancelStats.SetEndTime()
				return
			}

			wm.cancelStats.IncrementResponses()
			sent, received := wm.cancelStats.GetCounts()
			serviceLogger.Info("Progress", "responses_received", received, "requests_sent", sent)

			if received >= requestsCount {
				serviceLogger.Info("All responses received successfully", "total", received)
				wm.cancelStats.SetEndTime()
				return
			}

		case <-timeout.C:
			sent, received := wm.cancelStats.GetCounts()
			err := fmt.Errorf("timeout reached: sent=%d, received=%d, missing=%d", sent, received, sent-received)
			errorChan <- err
			wm.cancelStats.SetEndTime()
			return

		case <-ctx.Done():
			serviceLogger.Info("Response handler stopped")
			wm.cancelStats.SetEndTime()
			return
		}
	}
}

// sendOrderCancelRequests sends multiple order cancellation requests concurrently
func (wm *WebSocketManager) sendOrderCancelRequests(ctx context.Context, orderCancelService *futures.OrderCancelWsService, waiterChannel chan []byte, errorChan chan<- error) {
	serviceLogger := wm.logger.With("service", "order_cancel")

	// Build request
	request := futures.NewOrderCancelRequest().
		Symbol("BTCUSDT").
		OrderID(123123) // Non-existing order for testing

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrency)

	for i := 0; i < requestsCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire semaphore
			defer func() { <-semaphore }() // Release semaphore

			requestID := fmt.Sprintf("cancel_%d_%d", time.Now().UnixNano(), index)
			serviceLogger.Debug("Sending request", "request_id", requestID, "index", index)

			if err := orderCancelService.Do(requestID, request, websocket.WithWaiter(waiterChannel)); err != nil {
				errorChan <- fmt.Errorf("failed to send request %s: %w", requestID, err)
				return
			}
			wm.cancelStats.IncrementRequests()
		}(i)

		// Rate limiting
		time.Sleep(requestRateLimit)
	}

	wg.Wait()
	serviceLogger.Info("All order cancel requests sent")
}

// printFinalStats prints comprehensive statistics for both services
func (wm *WebSocketManager) printFinalStats() {
	wm.logger.Info("=== FINAL STATISTICS ===")

	// Get counts atomically to avoid race conditions
	placeSent, placeReceived := wm.placeStats.GetCounts()
	placeDuration := wm.placeStats.GetDuration()
	placeSuccess := placeSent == placeReceived && placeSent == requestsCount

	// Calculate average response time safely
	var placeAvgTime time.Duration
	if placeReceived > 0 {
		placeAvgTime = placeDuration / time.Duration(placeReceived)
	}

	wm.logger.Info("Order Place Service Statistics",
		"requests_sent", placeSent,
		"responses_received", placeReceived,
		"expected", requestsCount,
		"missing", placeSent-placeReceived,
		"success", placeSuccess,
		"duration", placeDuration,
		"avg_response_time", placeAvgTime,
	)

	// Get counts atomically to avoid race conditions
	cancelSent, cancelReceived := wm.cancelStats.GetCounts()
	cancelDuration := wm.cancelStats.GetDuration()
	cancelSuccess := cancelSent == cancelReceived && cancelSent == requestsCount

	// Calculate average response time safely
	var cancelAvgTime time.Duration
	if cancelReceived > 0 {
		cancelAvgTime = cancelDuration / time.Duration(cancelReceived)
	}

	wm.logger.Info("Order Cancel Service Statistics",
		"requests_sent", cancelSent,
		"responses_received", cancelReceived,
		"expected", requestsCount,
		"missing", cancelSent-cancelReceived,
		"success", cancelSuccess,
		"duration", cancelDuration,
		"avg_response_time", cancelAvgTime,
	)

	// Overall stats
	totalSent := placeSent + cancelSent
	totalReceived := placeReceived + cancelReceived
	totalExpected := int64(requestsCount * 2)
	overallSuccess := totalSent == totalReceived && totalSent == totalExpected

	// Calculate success rate safely
	var successRate float64
	if totalSent > 0 {
		successRate = float64(totalReceived) / float64(totalSent) * 100
	}

	wm.logger.Info("Overall Statistics",
		"total_requests_sent", totalSent,
		"total_responses_received", totalReceived,
		"total_expected", totalExpected,
		"total_missing", totalSent-totalReceived,
		"overall_success", overallSuccess,
		"success_rate", fmt.Sprintf("%.2f%%", successRate),
	)

	if !overallSuccess {
		wm.logger.Error("VERIFICATION FAILED: Request/Response counts do not match!",
			"sent", totalSent,
			"received", totalReceived,
			"expected", totalExpected,
		)
	} else {
		wm.logger.Info("✅ VERIFICATION PASSED: All requests received responses!")
	}
}

// WebSocketManager encapsulates the WebSocket client and statistics
type WebSocketManager struct {
	logger      *slog.Logger
	placeStats  *Stats
	cancelStats *Stats
	client      websocket.Client
}

// NewWebSocketManager creates a new manager with proper initialization
func NewWebSocketManager() (*WebSocketManager, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	client, err := createWebSocketClient(logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create WebSocket client: %w", err)
	}

	return &WebSocketManager{
		logger:      logger,
		placeStats:  &Stats{},
		cancelStats: &Stats{},
		client:      client,
	}, nil
}

// Stats tracks request/response statistics using atomic operations
type Stats struct {
	requestsSent      int64
	responsesReceived int64
	startTime         time.Time
	endTime           time.Time
	mu                sync.RWMutex // Only for time operations
}

func (s *Stats) IncrementRequests() {
	atomic.AddInt64(&s.requestsSent, 1)
}

func (s *Stats) IncrementResponses() {
	atomic.AddInt64(&s.responsesReceived, 1)
}

func (s *Stats) GetCounts() (int64, int64) {
	return atomic.LoadInt64(&s.requestsSent), atomic.LoadInt64(&s.responsesReceived)
}

func (s *Stats) SetStartTime() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.startTime = time.Now()
}

func (s *Stats) SetEndTime() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.endTime = time.Now()
}

func (s *Stats) GetDuration() time.Duration {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.endTime.IsZero() {
		return time.Since(s.startTime)
	}
	return s.endTime.Sub(s.startTime)
}
