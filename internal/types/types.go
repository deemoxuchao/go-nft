// Code generated by goctl. DO NOT EDIT.
package types

type BalanceResp struct {
	Balance int64 `json:"balance"`
}

type CheckArrivedReq struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

type CheckArrivedResp struct {
	Arrived bool   `json:"arrived"`
	Address string `json:"address"`
	Balance int64  `json:"balance"`
}

type CheckERC20ArrivedReq struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type CheckERC20ArrivedResp struct {
	Arrived bool   `json:"arrived"`
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type CheckEthArrivedReq struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

type CheckEthArrivedResp struct {
	Arrived bool   `json:"arrived"`
	Address string `json:"address"`
	Balance int64  `json:"balance"`
}

type NewAccountResp struct {
	Address string `json:"address"`
}

type RechargePlatformReq struct {
	ToAddress string `json:"to"`
}

type RechargePlatformResp struct {
	Msg  string `json:"msg"`
	Code int64  `json:"code"`
}
