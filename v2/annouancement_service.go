package binance

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2/common"
)

// CreateAnnouncementParam creates a new WsAnnouncementParam for use with WsAnnouncementServe.
//
// Currently supports only WithRecvWindow option, which defaults to 6000 milliseconds
// if not specified.
func (c *Client) CreateAnnouncementParam(opts ...RequestOption) (WsAnnouncementParam, error) {
	if c.APIKey == "" || c.SecretKey == "" {
		return WsAnnouncementParam{}, errors.New("miss apikey or secret key")
	}
	kt := c.KeyType
	if kt == "" {
		kt = common.KeyTypeHmac
	}
	req := new(request)
	for _, opt := range opts {
		opt(req)
	}
	if req.recvWindow == 0 {
		req.recvWindow = 6000
	}

	sf, err := common.SignFunc(kt)
	if err != nil {
		return WsAnnouncementParam{}, err
	}
	r := make([]byte, 16)
	rand.Read(r)
	random := hex.EncodeToString(r)
	timestamp := time.Now().UnixMilli()
	recvWindow := req.recvWindow

	param := WsAnnouncementParam{
		Random:     random,
		Topic:      "com_announcement_en",
		RecvWindow: recvWindow,
		Timestamp:  timestamp,
		ApiKey:     c.APIKey,
	}
	signature, err := sf(c.SecretKey, fmt.Sprintf("random=%s&topic=%s&recvWindow=%d&timestamp=%d", param.Random, param.Topic, param.RecvWindow, param.Timestamp))
	if err != nil {
		return WsAnnouncementParam{}, err
	}
	param.Signature = *signature
	return param, nil
}
