package futures

import (
	"context"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// DepthService show depth info
type DepthService struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *DepthService) Symbol(symbol string) *DepthService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *DepthService) Limit(limit int) *DepthService {
	s.limit = &limit
	return s
}

// Do send request
func (s *DepthService) Do(ctx context.Context, opts ...RequestOption) (res *DepthResponse, rateLimits map[string]string, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/depth",
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, rateLimits, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, nil, err
	}
	j, err := newJSON(data)
	if err != nil {
		return nil, nil, err
	}
	res = new(DepthResponse)
	res.Time = j.Get("E").MustInt64()
	res.TradeTime = j.Get("T").MustInt64()
	res.LastUpdateID = j.Get("lastUpdateId").MustInt64()
	bidsLen := len(j.Get("bids").MustArray())
	res.Bids = make([]Bid, bidsLen)
	for i := 0; i < bidsLen; i++ {
		item := j.Get("bids").GetIndex(i)
		res.Bids[i] = Bid{
			Price:    item.GetIndex(0).MustString(),
			Quantity: item.GetIndex(1).MustString(),
		}
	}
	asksLen := len(j.Get("asks").MustArray())
	res.Asks = make([]Ask, asksLen)
	for i := 0; i < asksLen; i++ {
		item := j.Get("asks").GetIndex(i)
		res.Asks[i] = Ask{
			Price:    item.GetIndex(0).MustString(),
			Quantity: item.GetIndex(1).MustString(),
		}
	}
	return res, rateLimits, nil
}

// DepthResponse define depth info with bids and asks
type DepthResponse struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Time         int64 `json:"E"`
	TradeTime    int64 `json:"T"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

// Ask is a type alias for PriceLevel.
type Ask = common.PriceLevel

// Bid is a type alias for PriceLevel.
type Bid = common.PriceLevel
