/*
Package main provides comprehensive examples of Binance WebSocket API order list operations.

This file demonstrates how to use all the new WebSocket services for order lists:

1. OrderListPlaceOCO() - Creates OCO (One-Cancels-Other) orders using the current API
2. OrderListPlaceOTO() - Creates OTO (One-Triggers-the-Other) orders
3. OrderListPlaceOTOCO() - Creates OTOCO (One-Triggers-One-Cancels-the-Other) orders
4. OrderListCancel() - Cancels existing order lists
5. SorOrderPlace() - Places orders using Smart Order Routing (SOR)
6. SorOrderTest() - Tests SOR orders without execution
7. OrderListPlaceDeprecated() - Uses the deprecated OCO endpoint

Usage:
1. Set your API credentials in each function (replace "your_api_key" and "your_secret_key")
2. Enable testnet for testing (binance.UseTestnet = true)
3. Run individual functions or call RunOrderListExamples() to run all examples

Important Notes:
- All examples use testnet by default for safety
- WebSocket services support both synchronous (SyncDo) and asynchronous (Do) operations
- Request IDs are automatically generated using common.GenerateSpotId()
- SOR orders provide information about whether Smart Order Routing was used
- Commission rate computation can be enabled for SOR test orders

Example Usage:

	// Run all examples
	RunOrderListExamples()

	// Or run individual examples
	OrderListPlaceOTO()
	SorOrderPlace()
*/
package main

import (
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/common"
)

// OrderListPlaceOCO demonstrates creating an OCO order using WebSocket API
func OrderListPlaceOCO() {
	// Setup configuration
	AppConfig.SetupTestnet()

	// Validate configuration
	if err := AppConfig.Validate(); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		return
	}

	client := AppConfig.GetClient()

	// Create OCO WebSocket service
	service, err := client.NewOrderListCreateWsService()
	if err != nil {
		fmt.Printf("Error creating OCO service: %v\n", err)
		return
	}

	// Create OCO order request
	request := binance.NewOrderListCreateWsRequest().
		Symbol("BTCUSDT").
		Side(binance.SideTypeSell).
		Quantity("0.001").
		AboveType(binance.OrderTypeTakeProfitLimit).
		AbovePrice("115000").
		AboveStopPrice("115000").
		AboveTimeInForce(binance.TimeInForceTypeGTC).
		BelowType(binance.OrderTypeStopLossLimit).
		BelowPrice("110000").
		BelowStopPrice("110000").
		BelowTimeInForce(binance.TimeInForceTypeGTC).
		NewOrderRespType(binance.NewOrderRespTypeFULL)

	requestID := common.GenerateSpotId()

	// Send async request
	err = service.Do(requestID, request)
	if err != nil {
		fmt.Printf("Error placing OCO order: %v\n", err)
		return
	}

	fmt.Printf("OCO order sent with request ID: %s\n", requestID)

	// Listen for response
	go func() {
		for {
			select {
			case response := <-service.GetReadChannel():
				fmt.Printf("OCO Response: %s\n", string(response))
				return
			case err := <-service.GetReadErrorChannel():
				fmt.Printf("OCO Error: %v\n", err)
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	service.ReceiveAllDataBeforeStop(2 * time.Second)
}

// OrderListPlaceOTO demonstrates creating an OTO order using WebSocket API
func OrderListPlaceOTO() {
	// Setup configuration
	AppConfig.SetupTestnet()

	// Validate configuration
	if err := AppConfig.Validate(); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		return
	}

	client := AppConfig.GetClient()

	// Create OTO WebSocket service
	service, err := client.NewOrderListPlaceOtoWsService()
	if err != nil {
		fmt.Printf("Error creating OTO service: %v\n", err)
		return
	}

	// Create OTO order request
	request := binance.NewOrderListPlaceOtoWsRequest().
		Symbol("BTCUSDT").
		WorkingType(binance.OrderTypeLimit).
		WorkingSide(binance.SideTypeBuy).
		WorkingPrice("30000").
		WorkingQuantity("0.001").
		PendingType(binance.OrderTypeLimit).
		PendingSide(binance.SideTypeSell).
		WorkingTimeInForce(binance.TimeInForceTypeGTC).
		PendingTimeInForce(binance.TimeInForceTypeGTC).
		PendingPrice("32000").
		PendingQuantity("0.001")

	requestID := common.GenerateSpotId()

	// Send synchronous request
	response, err := service.SyncDo(requestID, request)
	if err != nil {
		fmt.Printf("Error placing OTO order: %v\n", err)
		return
	}

	fmt.Printf("OTO Order Response: %+v\n", response)
}

// OrderListPlaceOTOCO demonstrates creating an OTOCO order using WebSocket API
func OrderListPlaceOTOCO() {
	// Setup configuration
	AppConfig.SetupTestnet()

	// Validate configuration
	if err := AppConfig.Validate(); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		return
	}

	client := AppConfig.GetClient()

	// Create OTOCO WebSocket service
	service, err := client.NewOrderListPlaceOtocoWsService()
	if err != nil {
		fmt.Printf("Error creating OTOCO service: %v\n", err)
		return
	}

	// Create OTOCO order request
	request := binance.NewOrderListPlaceOtocoWsRequest().
		Symbol("BTCUSDT").
		WorkingType(binance.OrderTypeLimit).
		WorkingSide(binance.SideTypeBuy).
		WorkingPrice("30000").
		WorkingQuantity("0.001").
		WorkingTimeInForce(binance.TimeInForceTypeGTC).
		PendingSide(binance.SideTypeSell).
		PendingQuantity("0.001").
		PendingAboveType(binance.OrderTypeLimitMaker).
		PendingAbovePrice("32000").
		PendingBelowType(binance.OrderTypeStopLoss).
		PendingBelowStopPrice("28000").
		ListClientOrderID("testOTOCOList")

	requestID := common.GenerateSpotId()

	// Send synchronous request
	response, err := service.SyncDo(requestID, request)
	if err != nil {
		fmt.Printf("Error placing OTOCO order: %v\n", err)
		return
	}

	fmt.Printf("OTOCO Order Response: %+v\n", response)
}

// OrderListCancel demonstrates canceling an order list using WebSocket API
func OrderListCancel() {
	// Setup configuration
	AppConfig.SetupTestnet()

	// Validate configuration
	if err := AppConfig.Validate(); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		return
	}

	client := AppConfig.GetClient()

	// Create order list cancel WebSocket service
	service, err := client.NewOrderListCancelWsService()
	if err != nil {
		fmt.Printf("Error creating cancel service: %v\n", err)
		return
	}

	// Create cancel request
	request := binance.NewOrderListCancelWsRequest().
		Symbol("BTCUSDT").
		OrderListID(123456789) // Replace with actual order list ID

	requestID := common.GenerateSpotId()

	// Send synchronous request
	response, err := service.SyncDo(requestID, request)
	if err != nil {
		fmt.Printf("Error canceling order list: %v\n", err)
		return
	}

	fmt.Printf("Cancel Order List Response: %+v\n", response)
}

// SorOrderPlace demonstrates placing a SOR order using WebSocket API
func SorOrderPlace() {
	// Setup configuration
	AppConfig.SetupTestnet()

	// Validate configuration
	if err := AppConfig.Validate(); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		return
	}

	client := AppConfig.GetClient()

	// Create SOR order placement WebSocket service
	service, err := client.NewSorOrderPlaceWsService()
	if err != nil {
		fmt.Printf("Error creating SOR service: %v\n", err)
		return
	}

	// Create SOR order request - using ETHUSDT as it has SOR support
	request := binance.NewSorOrderPlaceWsRequest().
		Symbol("ETHUSDT").
		Side(binance.SideTypeBuy).
		Type(binance.OrderTypeLimit).
		Quantity("0.1").
		Price("2000").
		TimeInForce(binance.TimeInForceTypeGTC).
		NewClientOrderID("sBI1KM6nNtOfj5tccZSKly").
		NewOrderRespType(binance.NewOrderRespTypeFULL)

	requestID := common.GenerateSpotId()

	// Send synchronous request
	response, err := service.SyncDo(requestID, request)
	if err != nil {
		fmt.Printf("Error placing SOR order: %v\n", err)
		return
	}

	fmt.Printf("SOR Order Response: %+v\n", response)

	// Check if result array is not empty before accessing
	if len(response.Result) > 0 {
		fmt.Printf("Used SOR: %v\n", response.Result[0].UsedSor)
	} else {
		fmt.Printf("No order results returned\n")
	}
}

// SorOrderTest demonstrates testing a SOR order using WebSocket API
func SorOrderTest() {
	// Setup configuration
	AppConfig.SetupTestnet()

	// Validate configuration
	if err := AppConfig.Validate(); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		return
	}

	client := AppConfig.GetClient()

	// Create SOR order test WebSocket service
	service, err := client.NewSorOrderTestWsService()
	if err != nil {
		fmt.Printf("Error creating SOR test service: %v\n", err)
		return
	}

	// Create SOR test request with commission rates - using ETHUSDT for SOR support
	request := binance.NewSorOrderTestWsRequest().
		Symbol("ETHUSDT").
		Side(binance.SideTypeBuy).
		Type(binance.OrderTypeLimit).
		Quantity("0.1").
		Price("2000").
		TimeInForce(binance.TimeInForceTypeGTC).
		ComputeCommissionRates(true)

	requestID := common.GenerateSpotId()

	// Send synchronous request
	response, err := service.SyncDo(requestID, request)
	if err != nil {
		fmt.Printf("Error testing SOR order: %v\n", err)
		return
	}

	fmt.Printf("SOR Test Response: %+v\n", response)
	if response.Result.StandardCommissionForOrder != nil {
		fmt.Printf("Standard Commission - Maker: %s, Taker: %s\n",
			response.Result.StandardCommissionForOrder.Maker,
			response.Result.StandardCommissionForOrder.Taker)
	}
}

func RunOrderListExamples() {
	fmt.Println("=== Binance Order List WebSocket API Examples ===")

	fmt.Println("1. OCO Order (Current)")
	OrderListPlaceOCO()
	time.Sleep(2 * time.Second)

	fmt.Println("\n2. OTO Order")
	OrderListPlaceOTO()
	time.Sleep(2 * time.Second)

	fmt.Println("\n3. OTOCO Order")
	OrderListPlaceOTOCO()
	time.Sleep(2 * time.Second)

	fmt.Println("\n4. Cancel Order List")
	// OrderListCancel() // Uncomment and provide valid order list ID

	fmt.Println("\n=== Examples completed ===")
}
