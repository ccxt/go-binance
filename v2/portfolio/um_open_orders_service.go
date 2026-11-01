package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMOpenOrdersService service to get all current UM open orders
type UMOpenOrdersService struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}

// Symbol set symbol
func (s *UMOpenOrdersService) Symbol(symbol string) *UMOpenOrdersService {
	s.symbol = &symbol
	return s
}

// RecvWindow set recvWindow
func (s *UMOpenOrdersService) RecvWindow(recvWindow int64) *UMOpenOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMOpenOrdersService) Do(ctx context.Context) ([]*UMOpenOrdersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*UMOpenOrdersResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMOpenOrdersResponse define open orders response
type UMOpenOrdersResponse struct {
	AvgPrice                string `json:"avgPrice"`
	ClientOrderID           string `json:"clientOrderId"`
	CumQuote                string `json:"cumQuote"`
	ExecutedQty             string `json:"executedQty"`
	OrderID                 int64  `json:"orderId"`
	OrigQty                 string `json:"origQty"`
	OrigType                string `json:"origType"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	Status                  string `json:"status"`
	Symbol                  string `json:"symbol"`
	Time                    int64  `json:"time"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	UpdateTime              int64  `json:"updateTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
	PriceMatch              string `json:"priceMatch"`
}
