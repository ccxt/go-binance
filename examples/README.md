# Binance API Examples

This directory contains examples demonstrating how to use the Binance API with the go-binance library.

## Configuration

All examples now use a centralized configuration system. You can set your API credentials in one of two ways:

### Option 1: Environment Variables (Recommended)

Set the following environment variables:

```bash
export BINANCE_API_KEY="your_api_key_here"
export BINANCE_SECRET_KEY="your_secret_key_here"
export BINANCE_USE_TESTNET="true"  # Set to "false" for production
```

### Option 2: Direct Configuration

Edit the `config.go` file and update the default values:

```go
var AppConfig = &Config{
    APIKey:     "your_api_key_here",
    SecretKey:  "your_secret_key_here", 
    UseTestnet: true,  // Set to false for production
}
```

## Security Notes

- **Never commit API credentials to version control**
- Use environment variables for production deployments
- The examples use testnet by default for safety
- Always validate your configuration before making API calls

## Running Examples

```bash
# Run all examples
go run .

## Available Examples

- `orderList.go` - WebSocket order list operations (OCO, OTO, OTOCO, SOR)
- `order.go` - Spot, Futures, and Delivery order placement
- `ticker.go` - Market data retrieval
- `wallet.go` - Wallet balance queries
- `ohlcv.go` - OHLCV (candlestick) data
- `watcher.go` - WebSocket market data streaming

## Configuration Validation

All examples automatically validate the configuration before making API calls. If credentials are missing or invalid, the examples will display an error message and exit gracefully.

## Testnet vs Production

- **Testnet**: Use for testing and development (default)
- **Production**: Set `BINANCE_USE_TESTNET=false` or `UseTestnet: false` in config

## Error Handling

The centralized configuration includes proper error handling:

```go
if err := AppConfig.Validate(); err != nil {
    fmt.Printf("Configuration error: %v\n", err)
    return
}
```

This ensures that all examples fail gracefully if credentials are not properly configured. 
