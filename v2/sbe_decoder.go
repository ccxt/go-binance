package binance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"

	sbe "github.com/adshao/go-binance/v2/sbe/spot_3_1"
)

// SBEDecoder handles decoding of SBE responses into Go types
type SBEDecoder struct {
	marshaller *sbe.SbeGoMarshaller
}

// NewSBEDecoder creates a new SBE decoder
func NewSBEDecoder() *SBEDecoder {
	return &SBEDecoder{
		marshaller: sbe.NewSbeGoMarshaller(),
	}
}

// sbeDecoder is a package-level decoder instance for SBE responses
var sbeDecoder = NewSBEDecoder()

// DecodeResponse decodes SBE binary data into the target type
// This is the main entry point for all SBE decoding
func (d *SBEDecoder) DecodeResponse(data []byte, target interface{}) error {
	reader := bytes.NewReader(data)

	// Decode message header
	var header sbe.MessageHeader
	if err := header.Decode(d.marshaller, reader, 0); err != nil {
		return fmt.Errorf("failed to decode SBE header: %w", err)
	}

	// Route to appropriate decoder based on template ID and target type
	return d.decodeByTemplateID(header.TemplateId, header.Version, header.BlockLength, reader, target)
}

// decodeByTemplateID routes to the appropriate SBE message decoder
func (d *SBEDecoder) decodeByTemplateID(templateID uint16, version uint16, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	switch templateID {
	// Market Data - Order Book & Trades
	case 200: // DepthResponse
		return d.decodeDepth(version, blockLength, reader, target)
	case 201: // TradesResponse
		return d.decodeTrades(version, blockLength, reader, target)
	case 202: // AggTradesResponse
		return d.decodeAggTrades(version, blockLength, reader, target)
	case 203: // KlinesResponse
		return d.decodeKlines(version, blockLength, reader, target)
	case 204: // AveragePriceResponse
		return d.decodeAveragePrice(version, blockLength, reader, target)

	// PriceTicker - Symbol and List
	case 205: // Ticker24hSymbolFullResponse
		return d.decodeTicker24hSymbolFull(version, blockLength, reader, target)
	case 206: // Ticker24hFullResponse
		return d.decodeTicker24hFullList(version, blockLength, reader, target)
	case 209: // PriceTickerSymbolResponse
		return d.decodePriceTickerSymbol(version, blockLength, reader, target)
	case 210: // PriceTickerResponse
		return d.decodePriceTickerList(version, blockLength, reader, target)

	// BookTicker - Symbol and List
	case 211: // BookTickerSymbolResponse
		return d.decodeBookTickerSymbol(version, blockLength, reader, target)
	case 212: // BookTickerResponse
		return d.decodeBookTickerList(version, blockLength, reader, target)

	// SymbolTicker - Full and Mini variants
	case 213: // TickerSymbolFullResponse
		return d.decodeTickerSymbolFull(version, blockLength, reader, target)
	case 214: // TickerFullResponse
		return d.decodeTickerFullList(version, blockLength, reader, target)
	case 215: // TickerSymbolMiniResponse
		return d.decodeTickerSymbolMini(version, blockLength, reader, target)
	case 216: // TickerMiniResponse
		return d.decodeTickerMiniList(version, blockLength, reader, target)

	// Trading - Order Placement (300-302)
	case 300: // NewOrderAckResponse
		return d.decodeNewOrderAck(version, blockLength, reader, target)
	case 301: // NewOrderResultResponse
		return d.decodeNewOrderResult(version, blockLength, reader, target)
	case 302: // NewOrderFullResponse
		return d.decodeNewOrderFull(version, blockLength, reader, target)

	// Trading - Order Management (304-308)
	case 304: // OrderResponse
		return d.decodeOrder(version, blockLength, reader, target)
	case 305: // CancelOrderResponse
		return d.decodeCancelOrder(version, blockLength, reader, target)
	case 306: // CancelOpenOrdersResponse
		return d.decodeCancelOpenOrders(version, blockLength, reader, target)
	case 308: // OrdersResponse
		return d.decodeOrders(version, blockLength, reader, target)

	// Account (400-401)
	case 400: // AccountResponse
		return d.decodeAccount(version, blockLength, reader, target)
	case 401: // AccountTradesResponse
		return d.decodeAccountTrades(version, blockLength, reader, target)

	default:
		return fmt.Errorf("unsupported template ID: %d", templateID)
	}
}

// Decoder functions for each template type

// decodeDepth decodes template 200 - DepthResponse (order book)
func (d *SBEDecoder) decodeDepth(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.DepthResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	// Convert to Go Depth response
	result := &DepthResponse{
		LastUpdateID: sbeResp.LastUpdateId,
		Bids:         make([]Bid, len(sbeResp.Bids)),
		Asks:         make([]Ask, len(sbeResp.Asks)),
	}

	for i, bid := range sbeResp.Bids {
		result.Bids[i] = Bid{
			Price:    convertSBEPrice(bid.Price, sbeResp.PriceExponent),
			Quantity: convertSBEPrice(bid.Qty, sbeResp.QtyExponent),
		}
	}

	for i, ask := range sbeResp.Asks {
		result.Asks[i] = Ask{
			Price:    convertSBEPrice(ask.Price, sbeResp.PriceExponent),
			Quantity: convertSBEPrice(ask.Qty, sbeResp.QtyExponent),
		}
	}

	return assignTarget(target, result)
}

// decodeTrades decodes template 201 - TradesResponse (recent trades)
func (d *SBEDecoder) decodeTrades(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.TradesResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*Trade, len(sbeResp.Trades))
	for i, trade := range sbeResp.Trades {
		result[i] = &Trade{
			ID:            trade.Id,
			Price:         convertSBEPrice(trade.Price, sbeResp.PriceExponent),
			Quantity:      convertSBEPrice(trade.Qty, sbeResp.QtyExponent),
			QuoteQuantity: convertSBEPrice(trade.QuoteQty, sbeResp.QtyExponent),
			Time:          trade.Time / 1000, // Convert microseconds to milliseconds
			IsBuyerMaker:  trade.IsBuyerMaker == 1,
			IsBestMatch:   trade.IsBestMatch == 1,
		}
	}

	return assignTarget(target, result)
}

// decodeAggTrades decodes template 202 - AggTradesResponse (aggregated trades)
func (d *SBEDecoder) decodeAggTrades(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.AggTradesResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*AggTrade, len(sbeResp.AggTrades))
	for i, trade := range sbeResp.AggTrades {
		result[i] = &AggTrade{
			AggTradeID:       trade.AggTradeId,
			Price:            convertSBEPrice(trade.Price, sbeResp.PriceExponent),
			Quantity:         convertSBEPrice(trade.Qty, sbeResp.QtyExponent),
			FirstTradeID:     trade.FirstTradeId,
			LastTradeID:      trade.LastTradeId,
			Timestamp:        trade.Time / 1000, // Convert microseconds to milliseconds
			IsBuyerMaker:     trade.IsBuyerMaker == 1,
			IsBestPriceMatch: trade.IsBestMatch == 1,
		}
	}

	return assignTarget(target, result)
}

// decodeKlines decodes template 203 - KlinesResponse (candlestick data)
func (d *SBEDecoder) decodeKlines(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.KlinesResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*Kline, len(sbeResp.Klines))
	for i, kline := range sbeResp.Klines {
		result[i] = &Kline{
			OpenTime:                 kline.OpenTime / 1000, // Convert microseconds to milliseconds
			Open:                     convertSBEPrice(kline.OpenPrice, sbeResp.PriceExponent),
			High:                     convertSBEPrice(kline.HighPrice, sbeResp.PriceExponent),
			Low:                      convertSBEPrice(kline.LowPrice, sbeResp.PriceExponent),
			Close:                    convertSBEPrice(kline.ClosePrice, sbeResp.PriceExponent),
			Volume:                   convertSBEString(kline.Volume[:]),
			CloseTime:                kline.CloseTime / 1000, // Convert microseconds to milliseconds
			QuoteAssetVolume:         convertSBEString(kline.QuoteVolume[:]),
			TradeNum:                 kline.NumTrades,
			TakerBuyBaseAssetVolume:  convertSBEString(kline.TakerBuyBaseVolume[:]),
			TakerBuyQuoteAssetVolume: convertSBEString(kline.TakerBuyQuoteVolume[:]),
		}
	}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodeAveragePrice(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.AveragePriceResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := &AvgPrice{
		Mins:  sbeResp.Mins,
		Price: convertSBEPrice(sbeResp.Price, sbeResp.PriceExponent),
	}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodeBookTickerSymbol(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.BookTickerSymbolResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := []*BookTicker{{
		Symbol:      convertSBEString(sbeResp.Symbol[:]),
		BidPrice:    convertSBEPrice(sbeResp.BidPrice, sbeResp.PriceExponent),
		BidQuantity: convertSBEPrice(sbeResp.BidQty, sbeResp.QtyExponent),
		AskPrice:    convertSBEPrice(sbeResp.AskPrice, sbeResp.PriceExponent),
		AskQuantity: convertSBEPrice(sbeResp.AskQty, sbeResp.QtyExponent),
	}}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodeBookTickerList(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.BookTickerResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*BookTicker, len(sbeResp.Tickers))
	for i, ticker := range sbeResp.Tickers {
		result[i] = &BookTicker{
			Symbol:      convertSBEString(ticker.Symbol[:]),
			BidPrice:    convertSBEPrice(ticker.BidPrice, ticker.PriceExponent),
			BidQuantity: convertSBEPrice(ticker.BidQty, ticker.QtyExponent),
			AskPrice:    convertSBEPrice(ticker.AskPrice, ticker.PriceExponent),
			AskQuantity: convertSBEPrice(ticker.AskQty, ticker.QtyExponent),
		}
	}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodePriceTickerSymbol(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.PriceTickerSymbolResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := []*SymbolPrice{{
		Symbol: convertSBEString(sbeResp.Symbol[:]),
		Price:  convertSBEPrice(sbeResp.Price, sbeResp.PriceExponent),
	}}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodePriceTickerList(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.PriceTickerResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*SymbolPrice, len(sbeResp.Tickers))
	for i, ticker := range sbeResp.Tickers {
		result[i] = &SymbolPrice{
			Symbol: convertSBEString(ticker.Symbol[:]),
			Price:  convertSBEPrice(ticker.Price, ticker.PriceExponent),
		}
	}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodeTicker24hSymbolFull(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.Ticker24hSymbolFullResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := []*PriceChangeStats{{
		Symbol:             convertSBEString(sbeResp.Symbol[:]),
		PriceChange:        convertSBEPrice(sbeResp.PriceChange, sbeResp.PriceExponent),
		PriceChangePercent: fmt.Sprintf("%.3f", sbeResp.PriceChangePercent),
		WeightedAvgPrice:   convertSBEPrice(sbeResp.WeightedAvgPrice, sbeResp.PriceExponent),
		PrevClosePrice:     convertSBEPrice(sbeResp.PrevClosePrice, sbeResp.PriceExponent),
		LastPrice:          convertSBEPrice(sbeResp.LastPrice, sbeResp.PriceExponent),
		LastQty:            convertSBEString(sbeResp.LastQty[:]),
		BidPrice:           convertSBEPrice(sbeResp.BidPrice, sbeResp.PriceExponent),
		BidQty:             convertSBEPrice(sbeResp.BidQty, sbeResp.QtyExponent),
		AskPrice:           convertSBEPrice(sbeResp.AskPrice, sbeResp.PriceExponent),
		AskQty:             convertSBEPrice(sbeResp.AskQty, sbeResp.QtyExponent),
		OpenPrice:          convertSBEPrice(sbeResp.OpenPrice, sbeResp.PriceExponent),
		HighPrice:          convertSBEPrice(sbeResp.HighPrice, sbeResp.PriceExponent),
		LowPrice:           convertSBEPrice(sbeResp.LowPrice, sbeResp.PriceExponent),
		Volume:             convertSBEString(sbeResp.Volume[:]),
		QuoteVolume:        convertSBEString(sbeResp.QuoteVolume[:]),
		OpenTime:           sbeResp.OpenTime / 1000,
		CloseTime:          sbeResp.CloseTime / 1000,
		FirstID:            sbeResp.FirstId,
		LastID:             sbeResp.LastId,
		Count:              sbeResp.NumTrades,
	}}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodeTicker24hFullList(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.Ticker24hFullResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*PriceChangeStats, len(sbeResp.Tickers))
	for i, ticker := range sbeResp.Tickers {
		result[i] = &PriceChangeStats{
			Symbol:             convertSBEString(ticker.Symbol[:]),
			PriceChange:        convertSBEPrice(ticker.PriceChange, ticker.PriceExponent),
			PriceChangePercent: fmt.Sprintf("%.3f", ticker.PriceChangePercent),
			WeightedAvgPrice:   convertSBEPrice(ticker.WeightedAvgPrice, ticker.PriceExponent),
			PrevClosePrice:     convertSBEPrice(ticker.PrevClosePrice, ticker.PriceExponent),
			LastPrice:          convertSBEPrice(ticker.LastPrice, ticker.PriceExponent),
			LastQty:            convertSBEString(ticker.LastQty[:]),
			BidPrice:           convertSBEPrice(ticker.BidPrice, ticker.PriceExponent),
			BidQty:             convertSBEPrice(ticker.BidQty, ticker.QtyExponent),
			AskPrice:           convertSBEPrice(ticker.AskPrice, ticker.PriceExponent),
			AskQty:             convertSBEPrice(ticker.AskQty, ticker.QtyExponent),
			OpenPrice:          convertSBEPrice(ticker.OpenPrice, ticker.PriceExponent),
			HighPrice:          convertSBEPrice(ticker.HighPrice, ticker.PriceExponent),
			LowPrice:           convertSBEPrice(ticker.LowPrice, ticker.PriceExponent),
			Volume:             convertSBEString(ticker.Volume[:]),
			QuoteVolume:        convertSBEString(ticker.QuoteVolume[:]),
			OpenTime:           ticker.OpenTime / 1000,
			CloseTime:          ticker.CloseTime / 1000,
			FirstID:            ticker.FirstId,
			LastID:             ticker.LastId,
			Count:              ticker.NumTrades,
		}
	}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodeTickerSymbolFull(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.TickerSymbolFullResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := []*SymbolTicker{{
		Symbol:             convertSBEString(sbeResp.Symbol[:]),
		PriceChange:        convertSBEPrice(sbeResp.PriceChange, sbeResp.PriceExponent),
		PriceChangePercent: fmt.Sprintf("%.3f", sbeResp.PriceChangePercent),
		WeightedAvgPrice:   convertSBEPrice(sbeResp.WeightedAvgPrice, sbeResp.PriceExponent),
		OpenPrice:          convertSBEPrice(sbeResp.OpenPrice, sbeResp.PriceExponent),
		HighPrice:          convertSBEPrice(sbeResp.HighPrice, sbeResp.PriceExponent),
		LowPrice:           convertSBEPrice(sbeResp.LowPrice, sbeResp.PriceExponent),
		LastPrice:          convertSBEPrice(sbeResp.LastPrice, sbeResp.PriceExponent),
		Volume:             convertSBEString(sbeResp.Volume[:]),
		QuoteVolume:        convertSBEString(sbeResp.QuoteVolume[:]),
		OpenTime:           sbeResp.OpenTime / 1000,
		CloseTime:          sbeResp.CloseTime / 1000,
		FirstId:            sbeResp.FirstId,
		LastId:             sbeResp.LastId,
		Count:              sbeResp.NumTrades,
	}}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodeTickerFullList(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.TickerFullResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*SymbolTicker, len(sbeResp.Tickers))
	for i, ticker := range sbeResp.Tickers {
		result[i] = &SymbolTicker{
			Symbol:             convertSBEString(ticker.Symbol[:]),
			PriceChange:        convertSBEPrice(ticker.PriceChange, ticker.PriceExponent),
			PriceChangePercent: fmt.Sprintf("%.3f", ticker.PriceChangePercent),
			WeightedAvgPrice:   convertSBEPrice(ticker.WeightedAvgPrice, ticker.PriceExponent),
			OpenPrice:          convertSBEPrice(ticker.OpenPrice, ticker.PriceExponent),
			HighPrice:          convertSBEPrice(ticker.HighPrice, ticker.PriceExponent),
			LowPrice:           convertSBEPrice(ticker.LowPrice, ticker.PriceExponent),
			LastPrice:          convertSBEPrice(ticker.LastPrice, ticker.PriceExponent),
			Volume:             convertSBEString(ticker.Volume[:]),
			QuoteVolume:        convertSBEString(ticker.QuoteVolume[:]),
			OpenTime:           ticker.OpenTime / 1000,
			CloseTime:          ticker.CloseTime / 1000,
			FirstId:            ticker.FirstId,
			LastId:             ticker.LastId,
			Count:              ticker.NumTrades,
		}
	}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodeTickerSymbolMini(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.TickerSymbolMiniResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := []*SymbolTicker{{
		Symbol:      convertSBEString(sbeResp.Symbol[:]),
		OpenPrice:   convertSBEPrice(sbeResp.OpenPrice, sbeResp.PriceExponent),
		HighPrice:   convertSBEPrice(sbeResp.HighPrice, sbeResp.PriceExponent),
		LowPrice:    convertSBEPrice(sbeResp.LowPrice, sbeResp.PriceExponent),
		LastPrice:   convertSBEPrice(sbeResp.LastPrice, sbeResp.PriceExponent),
		Volume:      convertSBEString(sbeResp.Volume[:]),
		QuoteVolume: convertSBEString(sbeResp.QuoteVolume[:]),
		OpenTime:    sbeResp.OpenTime / 1000,
		CloseTime:   sbeResp.CloseTime / 1000,
		FirstId:     sbeResp.FirstId,
		LastId:      sbeResp.LastId,
		Count:       sbeResp.NumTrades,
	}}

	return assignTarget(target, result)
}

func (d *SBEDecoder) decodeTickerMiniList(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.TickerMiniResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*SymbolTicker, len(sbeResp.Tickers))
	for i, ticker := range sbeResp.Tickers {
		result[i] = &SymbolTicker{
			Symbol:      convertSBEString(ticker.Symbol[:]),
			OpenPrice:   convertSBEPrice(ticker.OpenPrice, ticker.PriceExponent),
			HighPrice:   convertSBEPrice(ticker.HighPrice, ticker.PriceExponent),
			LowPrice:    convertSBEPrice(ticker.LowPrice, ticker.PriceExponent),
			LastPrice:   convertSBEPrice(ticker.LastPrice, ticker.PriceExponent),
			Volume:      convertSBEString(ticker.Volume[:]),
			QuoteVolume: convertSBEString(ticker.QuoteVolume[:]),
			OpenTime:    ticker.OpenTime / 1000,
			CloseTime:   ticker.CloseTime / 1000,
			FirstId:     ticker.FirstId,
			LastId:      ticker.LastId,
			Count:       ticker.NumTrades,
		}
	}

	return assignTarget(target, result)
}

// Trading Decoders - Order Placement

// decodeNewOrderAck decodes template 300 - NewOrderAckResponse
func (d *SBEDecoder) decodeNewOrderAck(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.NewOrderAckResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := &CreateOrderResponse{
		Symbol:        convertSBEString(sbeResp.Symbol),
		OrderID:       sbeResp.OrderId,
		ClientOrderID: convertSBEString(sbeResp.ClientOrderId),
		TransactTime:  sbeResp.TransactTime / 1000, // Convert microseconds to milliseconds
	}

	return assignTarget(target, result)
}

// decodeNewOrderResult decodes template 301 - NewOrderResultResponse
func (d *SBEDecoder) decodeNewOrderResult(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.NewOrderResultResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := &CreateOrderResponse{
		Symbol:                   convertSBEString(sbeResp.Symbol),
		OrderID:                  sbeResp.OrderId,
		ClientOrderID:            convertSBEString(sbeResp.ClientOrderId),
		TransactTime:             sbeResp.TransactTime / 1000,
		Price:                    convertSBEPrice(sbeResp.Price, sbeResp.PriceExponent),
		OrigQuantity:             convertSBEPrice(sbeResp.OrigQty, sbeResp.QtyExponent),
		ExecutedQuantity:         convertSBEPrice(sbeResp.ExecutedQty, sbeResp.QtyExponent),
		CummulativeQuoteQuantity: convertSBEPrice(sbeResp.CummulativeQuoteQty, sbeResp.QtyExponent),
		Status:                   OrderStatusType(convertSBEOrderStatus(sbeResp.Status)),
		TimeInForce:              TimeInForceType(convertSBETimeInForce(sbeResp.TimeInForce)),
		Type:                     OrderType(convertSBEOrderType(sbeResp.OrderType)),
		Side:                     SideType(convertSBEOrderSide(sbeResp.Side)),
		SelfTradePreventionMode:  SelfTradePreventionMode(convertSBESelfTradePreventionMode(sbeResp.SelfTradePreventionMode)),
	}

	return assignTarget(target, result)
}

// decodeNewOrderFull decodes template 302 - NewOrderFullResponse
func (d *SBEDecoder) decodeNewOrderFull(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.NewOrderFullResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	// Convert fills
	fills := make([]*Fill, len(sbeResp.Fills))
	for i, fill := range sbeResp.Fills {
		fills[i] = &Fill{
			TradeID:         fill.TradeId,
			Price:           convertSBEPrice(fill.Price, sbeResp.PriceExponent),
			Quantity:        convertSBEPrice(fill.Qty, sbeResp.QtyExponent),
			Commission:      convertSBEPrice(fill.Commission, fill.CommissionExponent),
			CommissionAsset: convertSBEString(fill.CommissionAsset),
		}
	}

	result := &CreateOrderResponse{
		Symbol:                   convertSBEString(sbeResp.Symbol),
		OrderID:                  sbeResp.OrderId,
		ClientOrderID:            convertSBEString(sbeResp.ClientOrderId),
		TransactTime:             sbeResp.TransactTime / 1000,
		Price:                    convertSBEPrice(sbeResp.Price, sbeResp.PriceExponent),
		OrigQuantity:             convertSBEPrice(sbeResp.OrigQty, sbeResp.QtyExponent),
		ExecutedQuantity:         convertSBEPrice(sbeResp.ExecutedQty, sbeResp.QtyExponent),
		CummulativeQuoteQuantity: convertSBEPrice(sbeResp.CummulativeQuoteQty, sbeResp.QtyExponent),
		Status:                   OrderStatusType(convertSBEOrderStatus(sbeResp.Status)),
		TimeInForce:              TimeInForceType(convertSBETimeInForce(sbeResp.TimeInForce)),
		Type:                     OrderType(convertSBEOrderType(sbeResp.OrderType)),
		Side:                     SideType(convertSBEOrderSide(sbeResp.Side)),
		Fills:                    fills,
		SelfTradePreventionMode:  SelfTradePreventionMode(convertSBESelfTradePreventionMode(sbeResp.SelfTradePreventionMode)),
	}

	return assignTarget(target, result)
}

// Trading Decoders - Order Management

// decodeOrder decodes template 304 - OrderResponse (query single order)
func (d *SBEDecoder) decodeOrder(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.OrderResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := &Order{
		Symbol:                   convertSBEString(sbeResp.Symbol),
		OrderID:                  sbeResp.OrderId,
		ClientOrderID:            convertSBEString(sbeResp.ClientOrderId),
		Price:                    convertSBEPrice(sbeResp.Price, sbeResp.PriceExponent),
		OrigQuantity:             convertSBEPrice(sbeResp.OrigQty, sbeResp.QtyExponent),
		ExecutedQuantity:         convertSBEPrice(sbeResp.ExecutedQty, sbeResp.QtyExponent),
		CummulativeQuoteQuantity: convertSBEPrice(sbeResp.CummulativeQuoteQty, sbeResp.QtyExponent),
		Status:                   OrderStatusType(convertSBEOrderStatus(sbeResp.Status)),
		TimeInForce:              TimeInForceType(convertSBETimeInForce(sbeResp.TimeInForce)),
		Type:                     OrderType(convertSBEOrderType(sbeResp.OrderType)),
		Side:                     SideType(convertSBEOrderSide(sbeResp.Side)),
		Time:                     sbeResp.Time / 1000,
		UpdateTime:               sbeResp.UpdateTime / 1000,
	}

	return assignTarget(target, result)
}

// decodeCancelOrder decodes template 305 - CancelOrderResponse
func (d *SBEDecoder) decodeCancelOrder(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.CancelOrderResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := &CancelOrderResponse{
		Symbol:                   convertSBEString(sbeResp.Symbol),
		OrderID:                  sbeResp.OrderId,
		ClientOrderID:            convertSBEString(sbeResp.ClientOrderId),
		OrigClientOrderID:        convertSBEString(sbeResp.OrigClientOrderId),
		Price:                    convertSBEPrice(sbeResp.Price, sbeResp.PriceExponent),
		OrigQuantity:             convertSBEPrice(sbeResp.OrigQty, sbeResp.QtyExponent),
		ExecutedQuantity:         convertSBEPrice(sbeResp.ExecutedQty, sbeResp.QtyExponent),
		CummulativeQuoteQuantity: convertSBEPrice(sbeResp.CummulativeQuoteQty, sbeResp.QtyExponent),
		Status:                   OrderStatusType(convertSBEOrderStatus(sbeResp.Status)),
		TimeInForce:              TimeInForceType(convertSBETimeInForce(sbeResp.TimeInForce)),
		Type:                     OrderType(convertSBEOrderType(sbeResp.OrderType)),
		Side:                     SideType(convertSBEOrderSide(sbeResp.Side)),
		SelfTradePreventionMode:  SelfTradePreventionMode(convertSBESelfTradePreventionMode(sbeResp.SelfTradePreventionMode)),
	}

	return assignTarget(target, result)
}

// decodeCancelOpenOrders decodes template 306 - CancelOpenOrdersResponse
func (d *SBEDecoder) decodeCancelOpenOrders(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.CancelOpenOrdersResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	// This returns an array of cancelled orders (each response is JSON encoded)
	result := make([]*CancelOrderResponse, len(sbeResp.Responses))
	for i, resp := range sbeResp.Responses {
		var cancelResp CancelOrderResponse
		if err := json.Unmarshal(resp.Response, &cancelResp); err != nil {
			return err
		}
		result[i] = &cancelResp
	}

	return assignTarget(target, result)
}

// decodeOrders decodes template 308 - OrdersResponse (list orders)
func (d *SBEDecoder) decodeOrders(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.OrdersResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*Order, len(sbeResp.Orders))
	for i, order := range sbeResp.Orders {
		result[i] = &Order{
			Symbol:                   convertSBEString(order.Symbol),
			OrderID:                  order.OrderId,
			ClientOrderID:            convertSBEString(order.ClientOrderId),
			Price:                    convertSBEPrice(order.Price, order.PriceExponent),
			OrigQuantity:             convertSBEPrice(order.OrigQty, order.QtyExponent),
			ExecutedQuantity:         convertSBEPrice(order.ExecutedQty, order.QtyExponent),
			CummulativeQuoteQuantity: convertSBEPrice(order.CummulativeQuoteQty, order.QtyExponent),
			Status:                   OrderStatusType(convertSBEOrderStatus(order.Status)),
			TimeInForce:              TimeInForceType(convertSBETimeInForce(order.TimeInForce)),
			Type:                     OrderType(convertSBEOrderType(order.OrderType)),
			Side:                     SideType(convertSBEOrderSide(order.Side)),
			Time:                     order.Time / 1000,
			UpdateTime:               order.UpdateTime / 1000,
		}
	}

	return assignTarget(target, result)
}

// Account Decoders

// decodeAccount decodes template 400 - AccountResponse
func (d *SBEDecoder) decodeAccount(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.AccountResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	// Convert balances
	balances := make([]Balance, len(sbeResp.Balances))
	for i, bal := range sbeResp.Balances {
		balances[i] = Balance{
			Asset:  convertSBEString(bal.Asset),
			Free:   convertSBEPrice(bal.Free, bal.Exponent),
			Locked: convertSBEPrice(bal.Locked, bal.Exponent),
		}
	}

	// Convert permissions
	permissions := make([]string, len(sbeResp.Permissions))
	for i, perm := range sbeResp.Permissions {
		permissions[i] = convertSBEString(perm.Permission)
	}

	result := &Account{
		MakerCommission:  sbeResp.CommissionRateMaker,
		TakerCommission:  sbeResp.CommissionRateTaker,
		BuyerCommission:  sbeResp.CommissionRateBuyer,
		SellerCommission: sbeResp.CommissionRateSeller,
		CommissionRates: CommissionRates{
			Maker:  convertSBEPrice(sbeResp.CommissionRateMaker, sbeResp.CommissionExponent),
			Taker:  convertSBEPrice(sbeResp.CommissionRateTaker, sbeResp.CommissionExponent),
			Buyer:  convertSBEPrice(sbeResp.CommissionRateBuyer, sbeResp.CommissionExponent),
			Seller: convertSBEPrice(sbeResp.CommissionRateSeller, sbeResp.CommissionExponent),
		},
		CanTrade:    sbeResp.CanTrade == 1,
		CanWithdraw: sbeResp.CanWithdraw == 1,
		CanDeposit:  sbeResp.CanDeposit == 1,
		UpdateTime:  uint64(sbeResp.UpdateTime / 1000),
		AccountType: convertSBEAccountType(sbeResp.AccountType),
		Balances:    balances,
		Permissions: permissions,
		UID:         sbeResp.Uid,
	}

	return assignTarget(target, result)
}

// decodeAccountTrades decodes template 401 - AccountTradesResponse
func (d *SBEDecoder) decodeAccountTrades(version, blockLength uint16, reader *bytes.Reader, target interface{}) error {
	var sbeResp sbe.AccountTradesResponse
	if err := sbeResp.Decode(d.marshaller, reader, version, blockLength, false); err != nil {
		return err
	}

	result := make([]*TradeV3, len(sbeResp.Trades))
	for i, trade := range sbeResp.Trades {
		result[i] = &TradeV3{
			ID:              trade.Id,
			Symbol:          convertSBEString(trade.Symbol),
			OrderID:         trade.OrderId,
			OrderListId:     trade.OrderListId,
			Price:           convertSBEPrice(trade.Price, trade.PriceExponent),
			Quantity:        convertSBEPrice(trade.Qty, trade.QtyExponent),
			QuoteQuantity:   convertSBEPrice(trade.QuoteQty, trade.QtyExponent),
			Commission:      convertSBEPrice(trade.Commission, trade.CommissionExponent),
			CommissionAsset: convertSBEString(trade.CommissionAsset),
			Time:            trade.Time / 1000,
			IsBuyer:         trade.IsBuyer == 1,
			IsMaker:         trade.IsMaker == 1,
			IsBestMatch:     trade.IsBestMatch == 1,
		}
	}

	return assignTarget(target, result)
}

// Helper functions

// assignTarget assigns the decoded result to the target interface
func assignTarget(target interface{}, result interface{}) error {
	// Use JSON marshaling as a generic way to copy data
	// This avoids reflection complexity and type assertions
	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed to marshal result: %w", err)
	}
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal to target: %w", err)
	}
	return nil
}

// convertSBEPrice converts an SBE price (mantissa + exponent) to a string
// SBE encodes decimals as mantissa * 10^exponent
func convertSBEPrice(mantissa int64, exponent int8) string {
	if mantissa == math.MinInt64 {
		return "" // Null value
	}

	value := float64(mantissa) * math.Pow10(int(exponent))
	precision := -int(exponent)
	if precision < 0 {
		precision = 0
	}

	return fmt.Sprintf("%.*f", precision, value)
}

// convertSBEString converts a byte array to a string, trimming null bytes
// SBE stores strings as fixed-length byte arrays padded with null bytes
func convertSBEString(data []byte) string {
	for i, b := range data {
		if b == 0 {
			return string(data[:i])
		}
	}
	return string(data)
}

// convertSBEOrderStatus converts SBE OrderStatusEnum to string
func convertSBEOrderStatus(status sbe.OrderStatusEnum) string {
	switch status {
	case sbe.OrderStatus.New:
		return "NEW"
	case sbe.OrderStatus.PartiallyFilled:
		return "PARTIALLY_FILLED"
	case sbe.OrderStatus.Filled:
		return "FILLED"
	case sbe.OrderStatus.Canceled:
		return "CANCELED"
	case sbe.OrderStatus.PendingCancel:
		return "PENDING_CANCEL"
	case sbe.OrderStatus.Rejected:
		return "REJECTED"
	case sbe.OrderStatus.Expired:
		return "EXPIRED"
	case sbe.OrderStatus.ExpiredInMatch:
		return "EXPIRED_IN_MATCH"
	default:
		return "UNKNOWN"
	}
}

// convertSBETimeInForce converts SBE TimeInForceEnum to string
func convertSBETimeInForce(tif sbe.TimeInForceEnum) string {
	switch tif {
	case sbe.TimeInForce.Gtc:
		return "GTC"
	case sbe.TimeInForce.Ioc:
		return "IOC"
	case sbe.TimeInForce.Fok:
		return "FOK"
	default:
		return "GTC"
	}
}

// convertSBEOrderType converts SBE OrderTypeEnum to string
func convertSBEOrderType(orderType sbe.OrderTypeEnum) string {
	switch orderType {
	case sbe.OrderType.Market:
		return "MARKET"
	case sbe.OrderType.Limit:
		return "LIMIT"
	case sbe.OrderType.StopLoss:
		return "STOP_LOSS"
	case sbe.OrderType.StopLossLimit:
		return "STOP_LOSS_LIMIT"
	case sbe.OrderType.TakeProfit:
		return "TAKE_PROFIT"
	case sbe.OrderType.TakeProfitLimit:
		return "TAKE_PROFIT_LIMIT"
	case sbe.OrderType.LimitMaker:
		return "LIMIT_MAKER"
	default:
		return "LIMIT"
	}
}

// convertSBEOrderSide converts SBE OrderSideEnum to string
func convertSBEOrderSide(side sbe.OrderSideEnum) string {
	switch side {
	case sbe.OrderSide.Buy:
		return "BUY"
	case sbe.OrderSide.Sell:
		return "SELL"
	default:
		return "BUY"
	}
}

// convertSBESelfTradePreventionMode converts SBE SelfTradePreventionModeEnum to string
func convertSBESelfTradePreventionMode(mode sbe.SelfTradePreventionModeEnum) string {
	switch mode {
	case sbe.SelfTradePreventionMode.None:
		return "NONE"
	case sbe.SelfTradePreventionMode.ExpireTaker:
		return "EXPIRE_TAKER"
	case sbe.SelfTradePreventionMode.ExpireMaker:
		return "EXPIRE_MAKER"
	case sbe.SelfTradePreventionMode.ExpireBoth:
		return "EXPIRE_BOTH"
	default:
		return "NONE"
	}
}

// convertSBEAccountType converts SBE AccountTypeEnum to string
func convertSBEAccountType(accountType sbe.AccountTypeEnum) string {
	switch accountType {
	case sbe.AccountType.Spot:
		return "SPOT"
	default:
		return "SPOT"
	}
}
