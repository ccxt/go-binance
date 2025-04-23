package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiTradingStatusService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (service *ApiTradingStatusService) Symbol(symbol string) *ApiTradingStatusService {
	service.symbol = symbol
	return service
}

// Do send request
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Trading-Quantitative-Rules-Indicators
func (s *ApiTradingStatusService) Do(ctx context.Context, opts ...RequestOption) (res *TradingStatusIndicators, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/apiTradingStatus",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(TradingStatusIndicators)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Trading Status Indicator
type TradingStatusIndicators struct {
	Indicators map[string][]*IndicatorInfo `json:"indicators"`
	UpdateTime int64                       `json:"updateTime"`
}

type IndicatorInfo struct {
	IsLocked           bool    `json:"isLocked"`
	PlannedRecoverTime int64   `json:"plannedRecoverTime"`
	Indicator          string  `json:"indicator"`
	Value              float64 `json:"value"`
	TriggerValue       float64 `json:"triggerValue"`
}
