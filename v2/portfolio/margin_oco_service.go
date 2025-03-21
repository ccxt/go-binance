package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginOCOService service to create OCO orders for margin account
type MarginOCOService struct {
	c                    *Client
	symbol               string
	listClientOrderID    *string
	side                 SideType
	quantity             string
	limitClientOrderID   *string
	price                string
	limitIcebergQty      *string
	stopClientOrderID    *string
	stopPrice            string
	stopLimitPrice       *string
	stopIcebergQty       *string
	stopLimitTimeInForce *TimeInForceType
	newOrderRespType     *NewOrderRespType
	sideEffectType       *SideEffectType
	recvWindow           *int64
}

// Symbol set symbol
func (s *MarginOCOService) Symbol(symbol string) *MarginOCOService {
	s.symbol = symbol
	return s
}

// ListClientOrderID set listClientOrderId
func (s *MarginOCOService) ListClientOrderID(listClientOrderID string) *MarginOCOService {
	s.listClientOrderID = &listClientOrderID
	return s
}

// Side set side
func (s *MarginOCOService) Side(side SideType) *MarginOCOService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *MarginOCOService) Quantity(quantity string) *MarginOCOService {
	s.quantity = quantity
	return s
}

// LimitClientOrderID set limitClientOrderId
func (s *MarginOCOService) LimitClientOrderID(limitClientOrderID string) *MarginOCOService {
	s.limitClientOrderID = &limitClientOrderID
	return s
}

// Price set price
func (s *MarginOCOService) Price(price string) *MarginOCOService {
	s.price = price
	return s
}

// LimitIcebergQty set limitIcebergQty
func (s *MarginOCOService) LimitIcebergQty(limitIcebergQty string) *MarginOCOService {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

// StopClientOrderID set stopClientOrderId
func (s *MarginOCOService) StopClientOrderID(stopClientOrderID string) *MarginOCOService {
	s.stopClientOrderID = &stopClientOrderID
	return s
}

// StopPrice set stop price
func (s *MarginOCOService) StopPrice(stopPrice string) *MarginOCOService {
	s.stopPrice = stopPrice
	return s
}

// StopLimitPrice set stop limit price
func (s *MarginOCOService) StopLimitPrice(stopLimitPrice string) *MarginOCOService {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

// StopIcebergQty set stop iceberg quantity
func (s *MarginOCOService) StopIcebergQty(stopIcebergQty string) *MarginOCOService {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

// StopLimitTimeInForce set stopLimitTimeInForce
func (s *MarginOCOService) StopLimitTimeInForce(timeInForce TimeInForceType) *MarginOCOService {
	s.stopLimitTimeInForce = &timeInForce
	return s
}

// NewOrderRespType set newOrderRespType
func (s *MarginOCOService) NewOrderRespType(newOrderRespType NewOrderRespType) *MarginOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SideEffectType set sideEffectType
func (s *MarginOCOService) SideEffectType(sideEffectType SideEffectType) *MarginOCOService {
	s.sideEffectType = &sideEffectType
	return s
}

// RecvWindow set recvWindow
func (s *MarginOCOService) RecvWindow(recvWindow int64) *MarginOCOService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginOCOService) Do(ctx context.Context, opts ...RequestOption) (res *MarginOCOResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/margin/order/oco",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	r.setParam("quantity", s.quantity)
	r.setParam("price", s.price)
	r.setParam("stopPrice", s.stopPrice)

	if s.listClientOrderID != nil {
		r.setParam("listClientOrderId", *s.listClientOrderID)
	}
	if s.limitClientOrderID != nil {
		r.setParam("limitClientOrderId", *s.limitClientOrderID)
	}
	if s.limitIcebergQty != nil {
		r.setParam("limitIcebergQty", *s.limitIcebergQty)
	}
	if s.stopClientOrderID != nil {
		r.setParam("stopClientOrderId", *s.stopClientOrderID)
	}
	if s.stopLimitPrice != nil {
		r.setParam("stopLimitPrice", *s.stopLimitPrice)
	}
	if s.stopIcebergQty != nil {
		r.setParam("stopIcebergQty", *s.stopIcebergQty)
	}
	if s.stopLimitTimeInForce != nil {
		r.setParam("stopLimitTimeInForce", *s.stopLimitTimeInForce)
	}
	if s.newOrderRespType != nil {
		r.setParam("newOrderRespType", *s.newOrderRespType)
	}
	if s.sideEffectType != nil {
		r.setParam("sideEffectType", *s.sideEffectType)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginOCOResponse defines margin OCO response
type MarginOCOResponse struct {
	OrderListID           int64                  `json:"orderListId"`
	ContingencyType       string                 `json:"contingencyType"`
	ListStatusType        string                 `json:"listStatusType"`
	ListOrderStatus       string                 `json:"listOrderStatus"`
	ListClientOrderID     string                 `json:"listClientOrderId"`
	TransactionTime       int64                  `json:"transactionTime"`
	Symbol                string                 `json:"symbol"`
	MarginBuyBorrowAmount string                 `json:"marginBuyBorrowAmount,omitempty"`
	MarginBuyBorrowAsset  string                 `json:"marginBuyBorrowAsset,omitempty"`
	Orders                []MarginOCOOrder       `json:"orders"`
	OrderReports          []MarginOCOOrderReport `json:"orderReports"`
}

// MarginOCOOrder defines margin OCO order
type MarginOCOOrder struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}

// MarginOCOOrderReport defines margin OCO order report
type MarginOCOOrderReport struct {
	Symbol              string          `json:"symbol"`
	OrderID             int64           `json:"orderId"`
	OrderListID         int64           `json:"orderListId"`
	ClientOrderID       string          `json:"clientOrderId"`
	TransactTime        int64           `json:"transactTime"`
	Price               string          `json:"price"`
	OrigQty             string          `json:"origQty"`
	ExecutedQty         string          `json:"executedQty"`
	CummulativeQuoteQty string          `json:"cummulativeQuoteQty"`
	Status              string          `json:"status"`
	TimeInForce         TimeInForceType `json:"timeInForce"`
	Type                OrderType       `json:"type"`
	Side                SideType        `json:"side"`
	StopPrice           string          `json:"stopPrice,omitempty"`
}
