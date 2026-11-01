package portfolio

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// CMAccountTradesService service to get CM account trade list
type CMAccountTradesService struct {
	c          *Client
	symbol     *string
	pair       *string
	startTime  *int64
	endTime    *int64
	fromID     *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *CMAccountTradesService) Symbol(symbol string) *CMAccountTradesService {
	s.symbol = &symbol
	return s
}

// Pair set pair
func (s *CMAccountTradesService) Pair(pair string) *CMAccountTradesService {
	s.pair = &pair
	return s
}

// StartTime set startTime
func (s *CMAccountTradesService) StartTime(startTime int64) *CMAccountTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *CMAccountTradesService) EndTime(endTime int64) *CMAccountTradesService {
	s.endTime = &endTime
	return s
}

// FromID set fromID
func (s *CMAccountTradesService) FromID(fromID int64) *CMAccountTradesService {
	s.fromID = &fromID
	return s
}

// Limit set limit
func (s *CMAccountTradesService) Limit(limit int) *CMAccountTradesService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *CMAccountTradesService) RecvWindow(recvWindow int64) *CMAccountTradesService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMAccountTradesService) Do(ctx context.Context) ([]*CMAccountTrade, error) {
	if s.symbol == nil && s.pair == nil {
		return nil, errors.New("either symbol or pair must be sent")
	}
	if s.symbol != nil && s.pair != nil {
		return nil, errors.New("symbol and pair cannot be sent together")
	}
	if s.pair != nil && s.fromID != nil {
		return nil, errors.New("pair and fromId cannot be sent together")
	}
	if s.fromID != nil && (s.startTime != nil || s.endTime != nil) {
		return nil, errors.New("fromId cannot be sent with startTime or endTime")
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/userTrades",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.pair != nil {
		r.setParam("pair", *s.pair)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*CMAccountTrade
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMAccountTrade define CM account trade
type CMAccountTrade struct {
	Symbol          string `json:"symbol"`
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	Pair            string `json:"pair"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	RealizedPnl     string `json:"realizedPnl"`
	MarginAsset     string `json:"marginAsset"`
	BaseQty         string `json:"baseQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	PositionSide    string `json:"positionSide"`
	Buyer           bool   `json:"buyer"`
	Maker           bool   `json:"maker"`
}
