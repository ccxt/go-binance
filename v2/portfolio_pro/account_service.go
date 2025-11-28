package portfolio_pro

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetAccountService struct {
	c *Client
}

type Account struct {
	UniMMR                string `json:"uniMMR"`                // 经典统一账户模式维持保证金率
	AccountEquity         string `json:"accountEquity"`         // 经典统一账户总权益，单位为USD
	ActualEquity          string `json:"actualEquity"`          // 不考虑质押率经典统一账户总权益，单位为USD
	AccountMaintMargin    string `json:"accountMaintMargin"`    // 经典统一账户维持保证金，即账户开仓及借贷总共需要的维持保证金，单位为USD
	AccountInitialMargin  string `json:"accountInitialMargin"`  // PM PRO and PM PRO SPAN请忽略
	TotalAvailableBalance string `json:"totalAvailableBalance"` // PM PRO and PM PRO SPAN请忽略
	AccountStatus         string `json:"accountStatus"`         // 经典统一账户当前账户状态:"NORMAL"正常状态, "MARGIN_CALL"补充保证金, "SUPPLY_MARGIN"再一次补充保证金, "REDUCE_ONLY"触发交易限制, "ACTIVE_LIQUIDATION"手动强制平仓, "FORCE_LIQUIDATION"强制平仓, "BANKRUPTED"破产
	AccountType           string `json:"accountType"`           // PM_1统一账户专业版, PM_2统一账户， PM_2统一账户专业版SPAN
}

func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (*Account, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/portfolio/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAccountBalanceService get account balance
type GetAccountBalanceService struct {
	c *Client
}

type AccountBalance struct {
	Asset               string `json:"asset"`               // 资产
	TotalWalletBalance  string `json:"totalWalletBalance"`  // 钱包余额 =  全仓杠杆未锁定 + 全仓杠杆锁定 + u本位合约钱包余额 + 币本位合约钱包余额
	CrossMarginAsset    string `json:"crossMarginAsset"`    // 全仓资产 = 全仓杠杆未锁定 + 全仓杠杆锁定
	CrossMarginBorrowed string `json:"crossMarginBorrowed"` // 全仓杠杆借贷
	CrossMarginFree     string `json:"crossMarginFree"`     // 全仓杠杆未锁定
	CrossMarginInterest string `json:"crossMarginInterest"` // 全仓杠杆利息
	CrossMarginLocked   string `json:"crossMarginLocked"`   // 全仓杠杆锁定
	UmWalletBalance     string `json:"umWalletBalance"`     // u本位合约钱包余额
	UmUnrealizedPNL     string `json:"umUnrealizedPNL"`     // u本位未实现盈亏
	CmWalletBalance     string `json:"cmWalletBalance"`     // 币本位合约钱包余额
	CmUnrealizedPNL     string `json:"cmUnrealizedPNL"`     // 币本位未实现盈亏
	UpdateTime          int64  `json:"updateTime"`          // 更新时间
	NegativeBalance     string `json:"negativeBalance"`     // 负余额
	OptionWalletBalance string `json:"optionWalletBalance"` // 仅适用于PM PRO SPAN
	OptionEquity        string `json:"optionEquity"`        // 仅适用于PM PRO SPAN
}

func (s *GetAccountBalanceService) Do(ctx context.Context, opts ...RequestOption) ([]*AccountBalance, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/portfolio/balance",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := make([]*AccountBalance, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
