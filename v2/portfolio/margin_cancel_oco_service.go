package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginCancelOCOService service to cancel margin account OCO orders
type MarginCancelOCOService struct {
	c                 *Client
	symbol            string
	orderListID       *int64
	listClientOrderID *string
	newClientOrderID  *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *MarginCancelOCOService) Symbol(symbol string) *MarginCancelOCOService {
	s.symbol = symbol
	return s
}

// OrderListID set orderListId
func (s *MarginCancelOCOService) OrderListID(orderListID int64) *MarginCancelOCOService {
	s.orderListID = &orderListID
	return s
}

// ListClientOrderID set listClientOrderId
func (s *MarginCancelOCOService) ListClientOrderID(listClientOrderID string) *MarginCancelOCOService {
	s.listClientOrderID = &listClientOrderID
	return s
}

// NewClientOrderID set newClientOrderId
func (s *MarginCancelOCOService) NewClientOrderID(newClientOrderID string) *MarginCancelOCOService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *MarginCancelOCOService) RecvWindow(recvWindow int64) *MarginCancelOCOService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *MarginCancelOCOService) Do(ctx context.Context) (*MarginCancelOCOResponse, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/margin/orderList",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderListID != nil {
		r.setParam("orderListId", *s.orderListID)
	}
	if s.listClientOrderID != nil {
		r.setParam("listClientOrderId", *s.listClientOrderID)
	}
	if s.newClientOrderID != nil {
		r.setParam("newClientOrderId", *s.newClientOrderID)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(MarginCancelOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginCancelOCOResponse define cancel OCO response
type MarginCancelOCOResponse struct {
	OrderListID       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderID string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderID       int64  `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
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
	} `json:"orderReports"`
}
