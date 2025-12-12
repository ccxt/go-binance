package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/adshao/go-binance/v2"
)

// Config holds the application configuration
type Config struct {
	APIKey     string
	SecretKey  string
	UseTestnet bool
	UseDemo bool
}

// Global configuration instance
var AppConfig = &Config{
	APIKey:     getEnvOrDefault("BINANCE_API_KEY", ""),
	SecretKey:  getEnvOrDefault("BINANCE_SECRET_KEY", ""),
	UseTestnet: getEnvOrDefault("BINANCE_USE_TESTNET", "true") == "true",
	UseDemo: getEnvOrDefault("BINANCE_USE_DEMO", "true") == "true",
}

// getEnvOrDefault returns the environment variable value or a default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetClient creates a new Binance client with the configured credentials
func (c *Config) GetClient() *binance.Client {
	return binance.NewClient(c.APIKey, c.SecretKey)
}

// SetupTestnet configures testnet usage based on the config
func (c *Config) SetupTestnet() {
	binance.UseTestnet = c.UseTestnet
}

// SetupDemo configures demo usage based on the config
func (c *Config) SetupDemo() {
	binance.UseDemo = c.UseDemo
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if strings.TrimSpace(c.APIKey) == "" {
		return fmt.Errorf("API key is required. Set BINANCE_API_KEY environment variable or update config.go")
	}
	if strings.TrimSpace(c.SecretKey) == "" {
		return fmt.Errorf("Secret key is required. Set BINANCE_SECRET_KEY environment variable or update config.go")
	}
	return nil
}
