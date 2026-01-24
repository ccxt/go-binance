package binance

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestTradingAccountServiceSBEIntegration tests SBE for trading and account services (Phase 2)
// NOTE: These tests require API keys with trading permissions and will place real orders on testnet
func TestTradingAccountServiceSBEIntegration(t *testing.T) {
	suite := SetupTest(t)

	// Test symbol with low price for testing (adjust as needed for testnet)
	testSymbol := "BTCUSDT"

	t.Run("GetAccount_SBE", func(t *testing.T) {
		// JSON version
		startJSON := time.Now()
		accountJSON, err := suite.client.NewGetAccountService().
			Do(context.Background())
		jsonElapsed := time.Since(startJSON)
		if err != nil {
			t.Fatalf("Failed to get account (JSON): %v", err)
		}

		// SBE version
		startSBE := time.Now()
		accountSBE, err := suite.client.NewGetAccountService().
			DoSBE(context.Background())
		sbeElapsed := time.Since(startSBE)
		if err != nil {
			t.Fatalf("Failed to get account (SBE): %v", err)
		}

		// Validate SBE response
		if accountSBE.AccountType == "" {
			t.Error("Expected non-empty account type")
		}
		if len(accountSBE.Balances) == 0 {
			t.Error("Expected at least one balance")
		}
		if len(accountSBE.Permissions) == 0 {
			t.Error("Expected at least one permission")
		}

		// Compare results
		if accountJSON.AccountType != accountSBE.AccountType {
			t.Errorf("Different account types: JSON=%s, SBE=%s", accountJSON.AccountType, accountSBE.AccountType)
		}
		if len(accountJSON.Balances) != len(accountSBE.Balances) {
			t.Logf("Warning: Different number of balances: JSON=%d, SBE=%d", len(accountJSON.Balances), len(accountSBE.Balances))
		}

		fmt.Printf("Account Info [JSON: %v, SBE: %v]\n", jsonElapsed, sbeElapsed)
		fmt.Printf("  Account Type: %s\n", accountSBE.AccountType)
		fmt.Printf("  Can Trade: %v\n", accountSBE.CanTrade)
		fmt.Printf("  Can Withdraw: %v\n", accountSBE.CanWithdraw)
		fmt.Printf("  Can Deposit: %v\n", accountSBE.CanDeposit)
		fmt.Printf("  Balances: %d\n", len(accountSBE.Balances))
		fmt.Printf("  Permissions: %v\n", accountSBE.Permissions)

		// Show first few non-zero balances
		count := 0
		for _, balance := range accountSBE.Balances {
			if balance.Free != "0.00000000" || balance.Locked != "0.00000000" {
				fmt.Printf("    %s: Free=%s, Locked=%s\n", balance.Asset, balance.Free, balance.Locked)
				count++
				if count >= 5 {
					break
				}
			}
		}
	})

	t.Run("GetAccount_SBE_OmitZeroBalances", func(t *testing.T) {
		account, err := suite.client.NewGetAccountService().
			OmitZeroBalances(true).
			DoSBE(context.Background())
		if err != nil {
			t.Fatalf("Failed to get account with omitZeroBalances (SBE): %v", err)
		}

		// Validate that no balances are zero
		for _, balance := range account.Balances {
			if balance.Free == "0.00000000" && balance.Locked == "0.00000000" {
				t.Errorf("Found zero balance for %s when omitZeroBalances=true", balance.Asset)
			}
		}

		fmt.Printf("Account (omitZeroBalances=true): %d non-zero balances\n", len(account.Balances))
	})

	t.Run("ListOrders_SBE", func(t *testing.T) {
		// JSON version
		startJSON := time.Now()
		ordersJSON, err := suite.client.NewListOrdersService().
			Symbol(testSymbol).
			Limit(10).
			Do(context.Background())
		jsonElapsed := time.Since(startJSON)
		if err != nil {
			t.Fatalf("Failed to list orders (JSON): %v", err)
		}

		// SBE version
		startSBE := time.Now()
		ordersSBE, err := suite.client.NewListOrdersService().
			Symbol(testSymbol).
			Limit(10).
			DoSBE(context.Background())
		sbeElapsed := time.Since(startSBE)
		if err != nil {
			t.Fatalf("Failed to list orders (SBE): %v", err)
		}

		// Validate SBE response
		if len(ordersSBE) > 0 {
			order := ordersSBE[0]
			if order.Symbol != testSymbol {
				t.Errorf("Expected symbol %s, got %s", testSymbol, order.Symbol)
			}
			if order.OrderID == 0 {
				t.Error("Expected non-zero order ID")
			}
			if order.Status == "" {
				t.Error("Expected non-empty order status")
			}
			if order.Type == "" {
				t.Error("Expected non-empty order type")
			}
			if order.Side == "" {
				t.Error("Expected non-empty order side")
			}

			fmt.Printf("Latest Order: ID=%d, %s %s %s @ %s [%s]\n",
				order.OrderID, order.Side, order.Type, order.OrigQuantity, order.Price, order.Status)
		}

		// Compare results
		if len(ordersJSON) != len(ordersSBE) {
			t.Logf("Warning: Different number of orders: JSON=%d, SBE=%d", len(ordersJSON), len(ordersSBE))
		}

		fmt.Printf("List Orders (%s): Count=%d [JSON: %v, SBE: %v]\n",
			testSymbol, len(ordersSBE), jsonElapsed, sbeElapsed)
	})

	t.Run("ListTrades_SBE", func(t *testing.T) {
		// JSON version
		startJSON := time.Now()
		tradesJSON, err := suite.client.NewListTradesService().
			Symbol(testSymbol).
			Limit(10).
			Do(context.Background())
		jsonElapsed := time.Since(startJSON)
		if err != nil {
			t.Fatalf("Failed to list trades (JSON): %v", err)
		}

		// SBE version
		startSBE := time.Now()
		tradesSBE, err := suite.client.NewListTradesService().
			Symbol(testSymbol).
			Limit(10).
			DoSBE(context.Background())
		sbeElapsed := time.Since(startSBE)
		if err != nil {
			t.Fatalf("Failed to list trades (SBE): %v", err)
		}

		// Validate SBE response
		if len(tradesSBE) > 0 {
			trade := tradesSBE[0]
			if trade.Symbol != testSymbol {
				t.Errorf("Expected symbol %s, got %s", testSymbol, trade.Symbol)
			}
			if trade.ID == 0 {
				t.Error("Expected non-zero trade ID")
			}
			if trade.OrderID == 0 {
				t.Error("Expected non-zero order ID")
			}
			if trade.Price == "" {
				t.Error("Expected non-empty price")
			}
			if trade.Quantity == "" {
				t.Error("Expected non-empty quantity")
			}
			if trade.CommissionAsset == "" {
				t.Error("Expected non-empty commission asset")
			}

			fmt.Printf("Latest Trade: ID=%d, OrderID=%d, %s @ %s (Commission: %s %s) [Buyer: %v, Maker: %v]\n",
				trade.ID, trade.OrderID, trade.Quantity, trade.Price,
				trade.Commission, trade.CommissionAsset, trade.IsBuyer, trade.IsMaker)
		}

		// Compare results
		if len(tradesJSON) != len(tradesSBE) {
			t.Logf("Warning: Different number of trades: JSON=%d, SBE=%d", len(tradesJSON), len(tradesSBE))
		}

		fmt.Printf("List My Trades (%s): Count=%d [JSON: %v, SBE: %v]\n",
			testSymbol, len(tradesSBE), jsonElapsed, sbeElapsed)
	})

	t.Run("ListTrades_SBE_WithTimeRange", func(t *testing.T) {
		endTime := time.Now().UnixMilli()
		startTime := endTime - (7 * 24 * 60 * 60 * 1000) // 7 days ago

		trades, err := suite.client.NewListTradesService().
			Symbol(testSymbol).
			StartTime(startTime).
			EndTime(endTime).
			Limit(100).
			DoSBE(context.Background())
		if err != nil {
			t.Fatalf("Failed to list trades with time range (SBE): %v", err)
		}

		fmt.Printf("My Trades (7 days, %s): Count=%d\n", testSymbol, len(trades))
	})

	// The following tests require placing real orders - only run on testnet
	// They are commented out by default to avoid accidental execution on mainnet
	t.Run("CreateOrder_GetOrder_CancelOrder_SBE_Flow", func(t *testing.T) {
		t.Skip("Skipping order creation tests - requires testnet and manual activation")

		/*
			// Step 1: Create a limit order (SBE)
			fmt.Println("\n=== Creating Limit Order (SBE) ===")
			createOrder, err := suite.client.NewCreateOrderService().
				Symbol(testSymbol).
				Side(SideTypeBuy).
				Type(OrderTypeLimit).
				TimeInForce(TimeInForceTypeGTC).
				Quantity("0.001"). // Very small quantity for testing
				Price("10000.00"). // Well below market to avoid fill
				DoSBE(context.Background())
			if err != nil {
				t.Fatalf("Failed to create order (SBE): %v", err)
			}

			fmt.Printf("Order Created: ID=%d, Symbol=%s, Status=%s\n",
				createOrder.OrderID, createOrder.Symbol, createOrder.Status)

			// Step 2: Query the order (SBE)
			fmt.Println("\n=== Querying Order (SBE) ===")
			queryOrder, err := suite.client.NewGetOrderService().
				Symbol(testSymbol).
				OrderID(createOrder.OrderID).
				DoSBE(context.Background())
			if err != nil {
				t.Fatalf("Failed to query order (SBE): %v", err)
			}

			// Validate query response
			if queryOrder.OrderID != createOrder.OrderID {
				t.Errorf("Expected order ID %d, got %d", createOrder.OrderID, queryOrder.OrderID)
			}
			if queryOrder.Symbol != testSymbol {
				t.Errorf("Expected symbol %s, got %s", testSymbol, queryOrder.Symbol)
			}

			fmt.Printf("Order Queried: ID=%d, Status=%s, Side=%s, Type=%s, Price=%s, Qty=%s\n",
				queryOrder.OrderID, queryOrder.Status, queryOrder.Side, queryOrder.Type,
				queryOrder.Price, queryOrder.OrigQuantity)

			// Step 3: Cancel the order (SBE)
			fmt.Println("\n=== Canceling Order (SBE) ===")
			cancelOrder, err := suite.client.NewCancelOrderService().
				Symbol(testSymbol).
				OrderID(createOrder.OrderID).
				DoSBE(context.Background())
			if err != nil {
				t.Fatalf("Failed to cancel order (SBE): %v", err)
			}

			// Validate cancel response
			if cancelOrder.OrderID != createOrder.OrderID {
				t.Errorf("Expected order ID %d, got %d", createOrder.OrderID, cancelOrder.OrderID)
			}

			fmt.Printf("Order Canceled: ID=%d, Symbol=%s, Status=%s\n",
				cancelOrder.OrderID, cancelOrder.Symbol, cancelOrder.Status)

			// Step 4: Verify order is canceled
			fmt.Println("\n=== Verifying Cancellation (SBE) ===")
			verifyOrder, err := suite.client.NewGetOrderService().
				Symbol(testSymbol).
				OrderID(createOrder.OrderID).
				DoSBE(context.Background())
			if err != nil {
				t.Fatalf("Failed to verify order (SBE): %v", err)
			}

			if verifyOrder.Status != string(OrderStatusTypeCanceled) {
				t.Errorf("Expected order status CANCELED, got %s", verifyOrder.Status)
			}

			fmt.Printf("Order Verified: ID=%d, Status=%s (successfully canceled)\n",
				verifyOrder.OrderID, verifyOrder.Status)
		*/
	})

	t.Run("CreateOrder_Types_SBE", func(t *testing.T) {
		t.Skip("Skipping order creation tests - requires testnet and manual activation")

		/*
			// Test different order response types (ACK, RESULT, FULL)

			// ACK response (template 300)
			fmt.Println("\n=== Creating Order with ACK Response (SBE) ===")
			ackOrder, err := suite.client.NewCreateOrderService().
				Symbol(testSymbol).
				Side(SideTypeBuy).
				Type(OrderTypeLimit).
				TimeInForce(TimeInForceTypeGTC).
				Quantity("0.001").
				Price("10000.00").
				NewOrderRespType(NewOrderRespTypeACK).
				DoSBE(context.Background())
			if err != nil {
				t.Fatalf("Failed to create order with ACK (SBE): %v", err)
			}
			fmt.Printf("ACK Order: ID=%d, Status=%s\n", ackOrder.OrderID, ackOrder.Status)

			// Cancel ACK order
			_, err = suite.client.NewCancelOrderService().
				Symbol(testSymbol).
				OrderID(ackOrder.OrderID).
				DoSBE(context.Background())
			if err != nil {
				t.Logf("Failed to cancel ACK order: %v", err)
			}

			// RESULT response (template 301)
			fmt.Println("\n=== Creating Order with RESULT Response (SBE) ===")
			resultOrder, err := suite.client.NewCreateOrderService().
				Symbol(testSymbol).
				Side(SideTypeBuy).
				Type(OrderTypeLimit).
				TimeInForce(TimeInForceTypeGTC).
				Quantity("0.001").
				Price("10000.00").
				NewOrderRespType(NewOrderRespTypeRESULT).
				DoSBE(context.Background())
			if err != nil {
				t.Fatalf("Failed to create order with RESULT (SBE): %v", err)
			}
			fmt.Printf("RESULT Order: ID=%d, Status=%s, ExecutedQty=%s\n",
				resultOrder.OrderID, resultOrder.Status, resultOrder.ExecutedQuantity)

			// Cancel RESULT order
			_, err = suite.client.NewCancelOrderService().
				Symbol(testSymbol).
				OrderID(resultOrder.OrderID).
				DoSBE(context.Background())
			if err != nil {
				t.Logf("Failed to cancel RESULT order: %v", err)
			}

			// FULL response (template 302) - includes fills
			fmt.Println("\n=== Creating Order with FULL Response (SBE) ===")
			fullOrder, err := suite.client.NewCreateOrderService().
				Symbol(testSymbol).
				Side(SideTypeBuy).
				Type(OrderTypeLimit).
				TimeInForce(TimeInForceTypeGTC).
				Quantity("0.001").
				Price("10000.00").
				NewOrderRespType(NewOrderRespTypeFULL).
				DoSBE(context.Background())
			if err != nil {
				t.Fatalf("Failed to create order with FULL (SBE): %v", err)
			}
			fmt.Printf("FULL Order: ID=%d, Status=%s, ExecutedQty=%s, Fills=%d\n",
				fullOrder.OrderID, fullOrder.Status, fullOrder.ExecutedQuantity, len(fullOrder.Fills))

			// Cancel FULL order
			_, err = suite.client.NewCancelOrderService().
				Symbol(testSymbol).
				OrderID(fullOrder.OrderID).
				DoSBE(context.Background())
			if err != nil {
				t.Logf("Failed to cancel FULL order: %v", err)
			}
		*/
	})

	t.Run("CancelOpenOrders_SBE", func(t *testing.T) {
		t.Skip("Skipping cancel open orders test - requires testnet and manual activation")

		/*
			// First, create a few orders
			fmt.Println("\n=== Creating Multiple Orders ===")
			for i := 0; i < 3; i++ {
				_, err := suite.client.NewCreateOrderService().
					Symbol(testSymbol).
					Side(SideTypeBuy).
					Type(OrderTypeLimit).
					TimeInForce(TimeInForceTypeGTC).
					Quantity("0.001").
					Price(fmt.Sprintf("1000%d.00", i)). // Different prices
					DoSBE(context.Background())
				if err != nil {
					t.Logf("Failed to create order %d: %v", i+1, err)
				} else {
					fmt.Printf("Created order %d\n", i+1)
				}
			}

			// Cancel all open orders
			fmt.Println("\n=== Canceling All Open Orders (SBE) ===")
			canceledOrders, err := suite.client.NewCancelOpenOrdersService().
				Symbol(testSymbol).
				DoSBE(context.Background())
			if err != nil {
				t.Fatalf("Failed to cancel open orders (SBE): %v", err)
			}

			fmt.Printf("Canceled %d open orders\n", len(canceledOrders))
			for i, order := range canceledOrders {
				fmt.Printf("  Order %d: ID=%d, Status=%s\n", i+1, order.OrderID, order.Status)
			}
		*/
	})

	// Performance comparison
	t.Run("Performance_Comparison_TradingAccount", func(t *testing.T) {
		iterations := 5

		fmt.Printf("\n=== Trading/Account Performance Comparison (%d iterations) ===\n", iterations)

		// Test GetAccount
		var accountJSONTotal, accountSBETotal time.Duration
		for i := 0; i < iterations; i++ {
			start := time.Now()
			_, err := suite.client.NewGetAccountService().Do(context.Background())
			if err == nil {
				accountJSONTotal += time.Since(start)
			}

			start = time.Now()
			_, err = suite.client.NewGetAccountService().DoSBE(context.Background())
			if err == nil {
				accountSBETotal += time.Since(start)
			}
		}

		// Test ListOrders
		var ordersJSONTotal, ordersSBETotal time.Duration
		for i := 0; i < iterations; i++ {
			start := time.Now()
			_, err := suite.client.NewListOrdersService().Symbol(testSymbol).Limit(10).Do(context.Background())
			if err == nil {
				ordersJSONTotal += time.Since(start)
			}

			start = time.Now()
			_, err = suite.client.NewListOrdersService().Symbol(testSymbol).Limit(10).DoSBE(context.Background())
			if err == nil {
				ordersSBETotal += time.Since(start)
			}
		}

		// Test ListTrades
		var tradesJSONTotal, tradesSBETotal time.Duration
		for i := 0; i < iterations; i++ {
			start := time.Now()
			_, err := suite.client.NewListTradesService().Symbol(testSymbol).Limit(10).Do(context.Background())
			if err == nil {
				tradesJSONTotal += time.Since(start)
			}

			start = time.Now()
			_, err = suite.client.NewListTradesService().Symbol(testSymbol).Limit(10).DoSBE(context.Background())
			if err == nil {
				tradesSBETotal += time.Since(start)
			}
		}

		// Calculate averages and print results
		fmt.Printf("\nGetAccount Service:\n")
		fmt.Printf("  JSON Avg: %v\n", accountJSONTotal/time.Duration(iterations))
		fmt.Printf("  SBE Avg:  %v\n", accountSBETotal/time.Duration(iterations))
		if accountSBETotal > 0 {
			fmt.Printf("  Speedup:  %.2fx\n", float64(accountJSONTotal)/float64(accountSBETotal))
		}

		fmt.Printf("\nListOrders Service:\n")
		fmt.Printf("  JSON Avg: %v\n", ordersJSONTotal/time.Duration(iterations))
		fmt.Printf("  SBE Avg:  %v\n", ordersSBETotal/time.Duration(iterations))
		if ordersSBETotal > 0 {
			fmt.Printf("  Speedup:  %.2fx\n", float64(ordersJSONTotal)/float64(ordersSBETotal))
		}

		fmt.Printf("\nListTrades Service:\n")
		fmt.Printf("  JSON Avg: %v\n", tradesJSONTotal/time.Duration(iterations))
		fmt.Printf("  SBE Avg:  %v\n", tradesSBETotal/time.Duration(iterations))
		if tradesSBETotal > 0 {
			fmt.Printf("  Speedup:  %.2fx\n", float64(tradesJSONTotal)/float64(tradesSBETotal))
		}

		fmt.Printf("\n======================================================\n\n")
	})
}
