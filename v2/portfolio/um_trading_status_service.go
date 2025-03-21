package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMTradingStatusService get UM trading quantitative rules indicators
type GetUMTradingStatusService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *GetUMTradingStatusService) Symbol(symbol string) *GetUMTradingStatusService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *GetUMTradingStatusService) Do(ctx context.Context, opts ...RequestOption) (*TradingStatus, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/apiTradingStatus",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(TradingStatus)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TradingStatus define trading status
type TradingStatus struct {
	Indicators map[string][]Indicator `json:"indicators"`
	UpdateTime int64                  `json:"updateTime"`
}

// Indicator define trading indicator
type Indicator struct {
	IsLocked           bool    `json:"isLocked"`
	PlannedRecoverTime int64   `json:"plannedRecoverTime"`
	Indicator          string  `json:"indicator"`    // UFR/IFER/GCR/DR/TMV
	Value              float64 `json:"value"`        // Current value
	TriggerValue       float64 `json:"triggerValue"` // Trigger value
}
