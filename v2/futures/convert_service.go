package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListConvertExchangeInfoService struct {
	c         *Client
	fromAsset string
	toAsset   string
}

type ConvertExchangeInfo struct {
	FromAsset          string `json:"fromAsset"`
	ToAsset            string `json:"toAsset"`
	FromAssetMinAmount string `json:"fromAssetMinAmount"`
	FromAssetMaxAmount string `json:"fromAssetMaxAmount"`
	ToAssetMinAmount   string `json:"toAssetMinAmount"`
	ToAssetMaxAmount   string `json:"toAssetMaxAmount"`
}

func (l *ListConvertExchangeInfoService) FromAsset(fromAsset string) *ListConvertExchangeInfoService {
	l.fromAsset = fromAsset
	return l
}

func (l *ListConvertExchangeInfoService) ToAsset(toAsset string) *ListConvertExchangeInfoService {
	l.toAsset = toAsset
	return l
}

func (l *ListConvertExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (res []ConvertExchangeInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/convert/exchangeInfo",
		secType:  secTypeNone,
	}
	if l.fromAsset != "" {
		r.setParam("fromAsset", l.fromAsset)
	}
	if l.toAsset != "" {
		r.setParam("toAsset", l.toAsset)
	}

	data, _, err := l.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]ConvertExchangeInfo, 0, 50)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type ConvertValidTime string

const (
	ConvertValidTime10S ConvertValidTime = "10s"
	ConvertValidTime30S ConvertValidTime = "30s"
	ConvertValidTime1M  ConvertValidTime = "1m"
	ConvertValidTime2M  ConvertValidTime = "2m"
)

type CreateConvertQuoteService struct {
	c          *Client
	fromAsset  string
	toAsset    string
	fromAmount string
	toAmount   string
	validTime  ConvertValidTime
}

func (c *CreateConvertQuoteService) FromAsset(fromAsset string) *CreateConvertQuoteService {
	c.fromAsset = fromAsset
	return c
}

func (c *CreateConvertQuoteService) ToAsset(toAsset string) *CreateConvertQuoteService {
	c.toAsset = toAsset
	return c
}

func (c *CreateConvertQuoteService) FromAmount(fromAmount string) *CreateConvertQuoteService {
	c.fromAmount = fromAmount
	return c
}

func (c *CreateConvertQuoteService) ToAmount(toAmount string) *CreateConvertQuoteService {
	c.toAmount = toAmount
	return c
}

func (c *CreateConvertQuoteService) ValidTime(validTime ConvertValidTime) *CreateConvertQuoteService {
	c.validTime = validTime
	return c
}

type ConvertQuote struct {
	QuoteId        string `json:"quoteId"`
	Ratio          string `json:"ratio"`
	InverseRatio   string `json:"inverseRatio"`
	ValidTimestamp int64  `json:"validTimestamp"`
	ToAmount       string `json:"toAmount"`
	FromAmount     string `json:"fromAmount"`
}

func (c *CreateConvertQuoteService) Do(ctx context.Context, opts ...RequestOption) (res *ConvertQuote, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/convert/getQuote",
		secType:  secTypeSigned,
	}
	m := params{
		"fromAsset": c.fromAsset,
		"toAsset":   c.toAsset,
	}
	if c.fromAmount != "" {
		m["fromAmount"] = c.fromAmount
	}
	if c.toAmount != "" {
		m["toAmount"] = c.toAmount
	}
	if c.validTime != "" {
		m["validTime"] = c.validTime
	}
	r.setFormParams(m)
	data, _, err := c.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ConvertQuote)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

type ConvertAcceptService struct {
	c       *Client
	quoteId string
}

func (c *ConvertAcceptService) QuoteId(quoteId string) *ConvertAcceptService {
	c.quoteId = quoteId
	return c
}

type ConvertAcceptStatus string

const (
	ConvertAcceptStatusProcess       ConvertAcceptStatus = "PROCESS"
	ConvertAcceptStatusAcceptSuccess ConvertAcceptStatus = "ACCEPT_SUCCESS"
	ConvertAcceptStatusSuccess       ConvertAcceptStatus = "SUCCESS"
	ConvertAcceptStatusFailed        ConvertAcceptStatus = "FAILED"
)

type ConvertResult struct {
	OrderId     string              `json:"orderId"`
	CreateTime  int64               `json:"createTime"`
	OrderStatus ConvertAcceptStatus `json:"orderStatus"`
}

func (c *ConvertAcceptService) Do(ctx context.Context, opts ...RequestOption) (res *ConvertResult, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/convert/acceptQuote",
		secType:  secTypeSigned,
	}
	m := params{
		"quoteId": c.quoteId,
	}
	r.setFormParams(m)
	data, _, err := c.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ConvertResult)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

type ConvertStatusService struct {
	c       *Client
	quoteId string
	orderId string
}

func (c *ConvertStatusService) QuoteId(quoteId string) *ConvertStatusService {
	c.quoteId = quoteId
	return c
}

func (c *ConvertStatusService) OrderId(orderId string) *ConvertStatusService {
	c.orderId = orderId
	return c
}

type ConvertStatusResult struct {
	OrderId      string              `json:"orderId"`
	OrderStatus  ConvertAcceptStatus `json:"orderStatus"`
	FromAsset    string              `json:"fromAsset"`
	FromAmount   string              `json:"fromAmount"`
	ToAsset      string              `json:"toAsset"`
	ToAmount     string              `json:"toAmount"`
	Ratio        string              `json:"ratio"`
	InverseRatio string              `json:"inverseRatio"`
	CreateTime   int64               `json:"createTime"`
}

func (c *ConvertStatusService) Do(ctx context.Context, opts ...RequestOption) (res *ConvertStatusResult, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/convert/orderStatus",
		secType:  secTypeSigned,
	}
	m := params{
		"quoteId": c.quoteId,
	}
	if c.orderId != "" {
		m["orderId"] = c.orderId
	}
	r.setParam("quoteId", c.quoteId)
	r.setFormParams(m)
	data, _, err := c.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ConvertStatusResult)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}
