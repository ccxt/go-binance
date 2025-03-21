package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMTransactionHistoryDownloadIDService get download id for UM futures transaction history
type GetUMTransactionHistoryDownloadIDService struct {
	c          *Client
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

// StartTime set startTime
func (s *GetUMTransactionHistoryDownloadIDService) StartTime(startTime int64) *GetUMTransactionHistoryDownloadIDService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetUMTransactionHistoryDownloadIDService) EndTime(endTime int64) *GetUMTransactionHistoryDownloadIDService {
	s.endTime = &endTime
	return s
}

// RecvWindow set recvWindow
func (s *GetUMTransactionHistoryDownloadIDService) RecvWindow(recvWindow int64) *GetUMTransactionHistoryDownloadIDService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetUMTransactionHistoryDownloadIDService) Do(ctx context.Context, opts ...RequestOption) (*UMTransactionHistoryDownloadID, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/income/asyn",
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
	res := new(UMTransactionHistoryDownloadID)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMTransactionHistoryDownloadID define download id response
type UMTransactionHistoryDownloadID struct {
	AvgCostTimestampOfLast30d int64  `json:"avgCostTimestampOfLast30d"`
	DownloadID                string `json:"downloadId"`
}
