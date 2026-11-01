package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginCancelAllOrdersService service to cancel all margin account open orders on a symbol
type MarginCancelAllOrdersService struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

// Symbol set symbol
func (s *MarginCancelAllOrdersService) Symbol(symbol string) *MarginCancelAllOrdersService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *MarginCancelAllOrdersService) RecvWindow(recvWindow int64) *MarginCancelAllOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginCancelAllOrdersService) Do(ctx context.Context) ([]*MarginCancelAllOrdersResponse, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/margin/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*MarginCancelAllOrdersResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginCancelAllOrdersResponse define cancel all orders response
type MarginCancelAllOrdersResponse struct {
	Symbol              string   `json:"symbol"`
	OrigClientOrderID   string   `json:"origClientOrderId,omitempty"`
	OrderID             int64    `json:"orderId,omitempty"`
	OrderListID         int64    `json:"orderListId"`
	ClientOrderID       string   `json:"clientOrderId,omitempty"`
	Price               string   `json:"price,omitempty"`
	OrigQty             string   `json:"origQty,omitempty"`
	ExecutedQty         string   `json:"executedQty,omitempty"`
	CummulativeQuoteQty string   `json:"cummulativeQuoteQty,omitempty"`
	Status              string   `json:"status,omitempty"`
	TimeInForce         string   `json:"timeInForce,omitempty"`
	Type                string   `json:"type,omitempty"`
	Side                string   `json:"side,omitempty"`
	ContingencyType     string   `json:"contingencyType,omitempty"`
	ListStatusType      string   `json:"listStatusType,omitempty"`
	ListOrderStatus     string   `json:"listOrderStatus,omitempty"`
	ListClientOrderID   string   `json:"listClientOrderId,omitempty"`
	TransactionTime     int64    `json:"transactionTime,omitempty"`
	Orders              []Order  `json:"orders,omitempty"`
	OrderReports        []Report `json:"orderReports,omitempty"`
}

// Order define order
type Order struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}

// Report define order report
type Report struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderID   string `json:"origClientOrderId"`
	OrderID             int64  `json:"orderId"`
	OrderListID         int64  `json:"orderListId"`
	ClientOrderID       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice,omitempty"`
	IcebergQty          string `json:"icebergQty,omitempty"`
}
