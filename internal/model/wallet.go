package model

// WalletItem 钱包列表项（含 Hyperliquid 链上数据）
type WalletItem struct {
	Id               int64   `json:"id"`
	Address          string  `json:"wallet"`           // 钱包地址
	ApiWalletAddress string  `json:"apiWalletAddress"` // API Wallet Address
	Remark           string  `json:"remark"`           // 备注
	Balance          string  `json:"balance"`          // 账户余额
	TotalMarginUsed  string  `json:"totalMarginUsed"`  // 已使用的保证金总额
	Withdrawable     string  `json:"withdrawable"`     // 可提现金额
	Upnl             string  `json:"upnl"`             // 未实现盈亏
	DepositWallet    string  `json:"depositWallet"`    // 充值钱包地址（API Wallet Address）
	ArbWithdrawAble  *string `json:"arbWithdrawAble"`  // Arbitrum 网络可提现金额
	BscWithdrawAble  *string `json:"bscWithdrawAble"`  // BSC 网络可提现金额
	CreatedAt        string  `json:"createdAt"`
	UpdatedAt        string  `json:"updatedAt"`
}
