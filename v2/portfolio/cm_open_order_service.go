package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMOpenOrderService service to get current CM open order
type CMOpenOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *CMOpenOrderService) Symbol(symbol string) *CMOpenOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CMOpenOrderService) OrderID(orderID int64) *CMOpenOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderId
func (s *CMOpenOrderService) OrigClientOrderID(origClientOrderID string) *CMOpenOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *CMOpenOrderService) RecvWindow(recvWindow int64) *CMOpenOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMOpenOrderService) Do(ctx context.Context) (*CMOpenOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/openOrder",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(CMOpenOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMOpenOrderResponse define open order response
type CMOpenOrderResponse struct {
	AvgPrice      string `json:"avgPrice"`
	ClientOrderID string `json:"clientOrderId"`
	CumBase       string `json:"cumBase"`
	ExecutedQty   string `json:"executedQty"`
	OrderID       int64  `json:"orderId"`
	OrigQty       string `json:"origQty"`
	OrigType      string `json:"origType"`
	Price         string `json:"price"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	Status        string `json:"status"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	Time          int64  `json:"time"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	UpdateTime    int64  `json:"updateTime"`
}
