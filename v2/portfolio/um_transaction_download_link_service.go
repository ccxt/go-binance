package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMTransactionDownloadLinkService get UM futures transaction download link by Id
type GetUMTransactionDownloadLinkService struct {
	c          *Client
	downloadID string
	recvWindow *int64
}

// DownloadID set downloadId
func (s *GetUMTransactionDownloadLinkService) DownloadID(downloadID string) *GetUMTransactionDownloadLinkService {
	s.downloadID = downloadID
	return s
}

// RecvWindow set recvWindow
func (s *GetUMTransactionDownloadLinkService) RecvWindow(recvWindow int64) *GetUMTransactionDownloadLinkService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetUMTransactionDownloadLinkService) Do(ctx context.Context, opts ...RequestOption) (*UMTransactionDownloadLink, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/income/asyn/id",
		secType:  secTypeSigned,
	}
	r.setParam("downloadId", s.downloadID)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(UMTransactionDownloadLink)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMTransactionDownloadLink define download link response
type UMTransactionDownloadLink struct {
	DownloadID          string `json:"downloadId"`
	Status              string `json:"status"` // Enum：completed，processing
	URL                 string `json:"url"`    // The link is mapped to download id
	S3Link              string `json:"s3Link"`
	Notified            bool   `json:"notified"`            // ignore
	ExpirationTimestamp int64  `json:"expirationTimestamp"` // The link would expire after this timestamp
	IsExpired           *bool  `json:"isExpired"`
}
