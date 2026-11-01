package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetCommissionRatesService queries user's commission rates for a given symbol
type GetCommissionRatesService struct {
	c      *Client
	symbol string
}

// Symbol sets the trading symbol (e.g., BTCUSDT)
func (s *GetCommissionRatesService) Symbol(symbol string) *GetCommissionRatesService {
	s.symbol = symbol
	return s
}

// Do sends the request
func (s *GetCommissionRatesService) Do(ctx context.Context, opts ...RequestOption) (res *CommissionRatesResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/account/commission",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CommissionRatesResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CommissionRatesResponse holds the full response for commission rates
type CommissionRatesResponse struct {
	Symbol             string          `json:"symbol"`
	StandardCommission CommissionGroup `json:"standardCommission"`
	TaxCommission      CommissionGroup `json:"taxCommission"`
	Discount           DiscountInfo    `json:"discount"`
}

// CommissionGroup defines a group of commission rates
type CommissionGroup struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

// DiscountInfo defines BNB discount info
type DiscountInfo struct {
	EnabledForAccount bool   `json:"enabledForAccount"`
	EnabledForSymbol  bool   `json:"enabledForSymbol"`
	DiscountAsset     string `json:"discountAsset"`
	Discount          string `json:"discount"`
}
