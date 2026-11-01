package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMTradeHistoryDownloadIDService get download id for UM futures trade history
type GetUMTradeHistoryDownloadIDService struct {
	c          *Client
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

// StartTime set startTime
func (s *GetUMTradeHistoryDownloadIDService) StartTime(startTime int64) *GetUMTradeHistoryDownloadIDService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetUMTradeHistoryDownloadIDService) EndTime(endTime int64) *GetUMTradeHistoryDownloadIDService {
	s.endTime = &endTime
	return s
}

// RecvWindow set recvWindow
func (s *GetUMTradeHistoryDownloadIDService) RecvWindow(recvWindow int64) *GetUMTradeHistoryDownloadIDService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetUMTradeHistoryDownloadIDService) Do(ctx context.Context, opts ...RequestOption) (*UMTradeHistoryDownloadID, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/trade/asyn",
		secType:  secTypeSigned,
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(UMTradeHistoryDownloadID)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMTradeHistoryDownloadID define download id response
type UMTradeHistoryDownloadID struct {
	AvgCostTimestampOfLast30d int64  `json:"avgCostTimestampOfLast30d"`
	DownloadID                string `json:"downloadId"`
}
