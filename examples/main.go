package main

import "fmt"

func main() {
	// Validate configuration before running examples
	if err := AppConfig.Validate(); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		fmt.Println("Please set your API credentials using environment variables or update config.go")
		fmt.Println("Example:")
		fmt.Println("  export BINANCE_API_KEY=\"your_api_key\"")
		fmt.Println("  export BINANCE_SECRET_KEY=\"your_secret_key\"")
		return
	}

	// Setup testnet
	AppConfig.SetupTestnet()

	fmt.Println("=== Binance API Examples ===")
	fmt.Printf("Using testnet: %v\n", AppConfig.UseTestnet)
	fmt.Println()

	// Run examples
	// Ticker()
	// Ohlcv()
	// SpotOrder()
	// FuturesOrder()
	// DeliveryOrder()
	// WalletBalance()
	WatchMiniMarketsStat()
	// RunOrderListExamples()
}
