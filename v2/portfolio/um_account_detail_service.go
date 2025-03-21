package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMAccountDetailService get current UM account asset and position information
type GetUMAccountDetailService struct {
	c *Client
}

// Do send request
func (s *GetUMAccountDetailService) Do(ctx context.Context, opts ...RequestOption) (*UMAccountDetail, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/account",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(UMAccountDetail)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMAccountDetail define UM account detail
type UMAccountDetail struct {
	Assets    []UMAsset    `json:"assets"`
	Positions []UMPosition `json:"positions"`
}

// UMAsset define UM asset info
type UMAsset struct {
	Asset                  string `json:"asset"`                  // asset name
	CrossWalletBalance     string `json:"crossWalletBalance"`     // wallet balance
	CrossUnPnl             string `json:"crossUnPnl"`             // unrealized profit
	MaintMargin            string `json:"maintMargin"`            // maintenance margin required
	InitialMargin          string `json:"initialMargin"`          // total initial margin required with current mark price
	PositionInitialMargin  string `json:"positionInitialMargin"`  // initial margin required for positions with current mark price
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // initial margin required for open orders with current mark price
	UpdateTime             int64  `json:"updateTime"`             // last update time
}
