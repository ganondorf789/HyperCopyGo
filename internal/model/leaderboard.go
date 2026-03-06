package model

// LeaderboardProfitItem 盈利榜列表项
type LeaderboardProfitItem struct {
	EthAddress   string   `json:"ethAddress"`
	DisplayName  *string  `json:"displayName"`
	AccountValue float64  `json:"accountValue"`
	Window       string   `json:"window"`
	Pnl          float64  `json:"pnl"`
	Roi          float64  `json:"roi"`
	Vlm          float64  `json:"vlm"`
}
