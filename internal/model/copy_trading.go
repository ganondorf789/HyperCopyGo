package model

import "github.com/gogf/gf/v2/os/gtime"

// BaseCopyTrading 跟单配置公共字段，供 API 请求/响应复用
type BaseCopyTrading struct {
	TargetWallet                   string   `json:"targetWallet"`
	TargetWalletPlatform           string   `json:"targetWalletPlatform"`
	Remark                         string   `json:"remark"`
	FollowType                     int64    `json:"followType"`
	FollowOnce                     int64    `json:"followOnce"`
	PositionConditions             string   `json:"positionConditions"`
	TraderConditions               string   `json:"traderConditions"`
	TagAccountValue                string   `json:"tagAccountValue"`
	TagProfitScale                 string   `json:"tagProfitScale"`
	TagDirection                   string   `json:"tagDirection"`
	TagTradingRhythm               string   `json:"tagTradingRhythm"`
	TagProfitStatus                string   `json:"tagProfitStatus"`
	TagTradingStyles               []string `json:"tagTradingStyles"`
	TraderMetricPeriod             string   `json:"traderMetricPeriod"`
	FollowMarginMode               int64    `json:"followMarginMode"`
	FollowSymbol                   string   `json:"followSymbol"`
	Leverage                       int64    `json:"leverage"`
	MarginMode                     int64    `json:"marginMode"`
	FollowModel                    int64    `json:"followModel"`
	FollowModelValue               float64  `json:"followModelValue"`
	MinValue                       float64  `json:"minValue"`
	MaxValue                       float64  `json:"maxValue"`
	MaxMarginUsage                 float64  `json:"maxMarginUsage"`
	TpValue                        float64  `json:"tpValue"`
	SlValue                        float64  `json:"slValue"`
	OptReverseFollowOrder          int64    `json:"optReverseFollowOrder"`
	OptFollowupDecrease            int64    `json:"optFollowupDecrease"`
	OptFollowupIncrease            int64    `json:"optFollowupIncrease"`
	OptForcedLiquidationProtection int64    `json:"optForcedLiquidationProtection"`
	OptPositionIncreaseOpening     int64    `json:"optPositionIncreaseOpening"`
	OptSlippageProtection          int64    `json:"optSlippageProtection"`
	SymbolListType                 string   `json:"symbolListType"`
	SymbolList                     string   `json:"symbolList"`
	MainWallet                     string   `json:"mainWallet"`
	MainWalletPlatform             string   `json:"mainWalletPlatform"`
}

// CopyTradingItem 跟单配置列表项
type CopyTradingItem struct {
	Id int64 `json:"id"`
	BaseCopyTrading
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// CopyTradeRecordItem 跟单记录列表项
type CopyTradeRecordItem struct {
	Id            int64       `json:"id"`
	Address       string      `json:"address"`
	Coin          string      `json:"coin"`
	Direction     string      `json:"direction"`
	Size          float64     `json:"size"`
	Price         float64     `json:"price"`
	ClosedPnl     float64     `json:"closedPnl"`
	ExecuteStatus int64       `json:"executeStatus"`
	OrderStatus   string      `json:"orderStatus"`
	ErrorMsg      string      `json:"errorMsg"`
	TradeTime     *gtime.Time `json:"tradeTime"`
	CreatedAt     *gtime.Time `json:"createdAt"`
}
