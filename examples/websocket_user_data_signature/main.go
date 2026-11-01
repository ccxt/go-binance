package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/google/uuid"
)

func main() {
	// Get API credentials from environment variables or replace with your actual credentials
	apiKey := os.Getenv("BINANCE_APIKEY")
	secretKey := os.Getenv("BINANCE_SECRET")
	useTestnet := os.Getenv("BINANCE_TESTNET") == "true" || os.Getenv("BINANCE_TESTNET") == "1"

	if apiKey == "" || secretKey == "" {
		log.Fatal("Please set BINANCE_API_KEY and BINANCE_SECRET_KEY environment variables")
	}

	// Enable testnet mode if requested
	if useTestnet {
		binance.UseTestnet = true
		fmt.Println("🧪 Running in TESTNET mode")
		fmt.Println("Testnet WebSocket API: wss://ws-api.testnet.binance.vision/ws-api/v3")
	} else {
		fmt.Println("🚀 Running in PRODUCTION mode")
		fmt.Println("Production WebSocket API: wss://ws-api.binance.com:443/ws-api/v3")
	}
	fmt.Println()

	// Auto-detect key type based on secret key format
	keyType := "HMAC"
	if strings.HasPrefix(secretKey, "-----BEGIN PRIVATE KEY-----") {
		// PEM-encoded private key detected
		// Ed25519 keys are much smaller (~120 bytes) than RSA keys (typically 1600+ bytes for 2048-bit)
		if len(secretKey) < 200 {
			keyType = "ED25519"
			fmt.Println("🔑 Detected Ed25519 private key")
		} else {
			keyType = "RSA"
			fmt.Println("🔑 Detected RSA private key")
		}
	} else {
		// HMAC secret key (default)
		fmt.Println("🔑 Using HMAC signature")
	}

	timeOffset := int64(0) // Time offset in milliseconds

	// Create handlers for user data events
	userDataHandler := func(event *binance.WsUserDataEvent) {
		fmt.Printf("\n=== User Data Event ===\n")
		fmt.Printf("Event Type: %s\n", event.Event)
		fmt.Printf("Time: %d\n", event.Time)

		switch event.Event {
		case binance.UserDataEventTypeOutboundAccountPosition:
			fmt.Printf("Account Update:\n")
			fmt.Printf("  Update Time: %d\n", event.AccountUpdate.AccountUpdateTime)
			for _, balance := range event.AccountUpdate.WsAccountUpdates {
				fmt.Printf("  Asset: %s, Free: %s, Locked: %s\n",
					balance.Asset, balance.Free, balance.Locked)
			}

		case binance.UserDataEventTypeBalanceUpdate:
			fmt.Printf("Balance Update:\n")
			fmt.Printf("  Asset: %s\n", event.BalanceUpdate.Asset)
			fmt.Printf("  Change: %s\n", event.BalanceUpdate.Change)
			fmt.Printf("  Transaction Time: %d\n", event.BalanceUpdate.TransactionTime)

		case binance.UserDataEventTypeExecutionReport:
			fmt.Printf("Order Update:\n")
			fmt.Printf("  Symbol: %s\n", event.OrderUpdate.Symbol)
			fmt.Printf("  Side: %s\n", event.OrderUpdate.Side)
			fmt.Printf("  Order Type: %s\n", event.OrderUpdate.Type)
			fmt.Printf("  Order ID: %d\n", event.OrderUpdate.Id)
			fmt.Printf("  Client Order ID: %s\n", event.OrderUpdate.ClientOrderId)
			fmt.Printf("  Execution Type: %s\n", event.OrderUpdate.ExecutionType)
			fmt.Printf("  Order Status: %s\n", event.OrderUpdate.Status)
			fmt.Printf("  Price: %s\n", event.OrderUpdate.Price)
			fmt.Printf("  Quantity: %s\n", event.OrderUpdate.Volume)
			fmt.Printf("  Filled: %s\n", event.OrderUpdate.FilledVolume)

		case binance.UserDataEventTypeListStatus:
			fmt.Printf("OCO Update:\n")
			fmt.Printf("  Symbol: %s\n", event.OCOUpdate.Symbol)
			fmt.Printf("  Order List ID: %d\n", event.OCOUpdate.OrderListId)
			fmt.Printf("  Contingency Type: %s\n", event.OCOUpdate.ContingencyType)
			fmt.Printf("  List Status: %s\n", event.OCOUpdate.ListStatusType)
			fmt.Printf("  List Order Status: %s\n", event.OCOUpdate.ListOrderStatus)

		default:
			fmt.Printf("Unknown event type: %s\n", event.Event)
		}
		fmt.Println("======================")
	}

	errHandler := func(err error) {
		log.Printf("WebSocket Error: %v\n", err)
	}

	fmt.Println("Starting user data stream with signature-based authentication...")
	fmt.Println("This uses the new WebSocket API method (listen key management is deprecated)")
	fmt.Println()
	if useTestnet {
		fmt.Println("ℹ️  Testnet Tips:")
		fmt.Println("   - Get testnet API keys from: https://testnet.binance.vision/")
		fmt.Println("   - Use testnet website to place orders: https://testnet.binance.vision/")
		fmt.Println("   - Testnet has free test assets (USDT, BTC, etc.)")
		fmt.Println()
	}
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println()

	// Subscribe to user data stream using signature-based authentication
	// This is the recommended method as listen key management has been deprecated
	doneC, stopC, err := binance.WsUserDataServeSignature(
		apiKey,
		secretKey,
		keyType,
		timeOffset,
		userDataHandler,
		errHandler,
	)
	if err != nil {
		log.Fatalf("Failed to start user data stream: %v", err)
	}

	// Set up signal handling for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	// Give the subscription a brief moment to become active
	time.Sleep(1 * time.Second)

	// Place a small MARKET order after subscribing (uses WebSocket API order.place)
	// Adjust symbol/amount as needed, especially if not on testnet.
	orderWs, err := binance.NewOrderCreateWsService(apiKey, secretKey)
	if err != nil {
		log.Printf("Failed to init order WS service: %v", err)
	} else {
		orderWs.KeyType = keyType
		orderWs.TimeOffset = timeOffset

		req := binance.NewOrderCreateWsRequest().
			Symbol("BTCUSDT").
			Side(binance.SideTypeBuy).
			Type(binance.OrderTypeMarket).
			QuoteOrderQty("11").
			NewOrderRespType(binance.NewOrderRespTypeRESULT)

		fmt.Println("Placing test MARKET order via WebSocket API (BTCUSDT, 11 USDT)...")
		err := orderWs.Do(uuid.New().String(), req)
		if err != nil {
			log.Printf("Order placement error: %v", err)
		}
		fmt.Println("Order placed")
	}
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal or connection to close
	select {
	case <-interrupt:
		fmt.Println("\nReceived interrupt signal, closing connection...")
		close(stopC)
		<-doneC
		fmt.Println("Connection closed successfully")
	case <-doneC:
		fmt.Println("Connection closed")
	}
}
