package binance

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseIntegrationTestSuite struct {
	suite.Suite
	client *Client
}

func SetupTest(t *testing.T) *baseIntegrationTestSuite {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	proxyURL := os.Getenv("BINANCE_PROXY_URL")
	useTestnet := true
	if os.Getenv("BINANCE_USE_TESTNET") == "false" {
		useTestnet = false
	}

	if apiKey == "" || secretKey == "" {
		t.Skip("API key and secret are required for integration tests")
	}

	var client *Client
	if proxyURL != "" {
		client = NewProxiedClient(apiKey, secretKey, proxyURL)
	} else {
		client = NewClient(apiKey, secretKey)
	}

	client.Debug = true
	UseTestnet = useTestnet // Set the global testnet flag

	return &baseIntegrationTestSuite{
		client: client,
	}
}
