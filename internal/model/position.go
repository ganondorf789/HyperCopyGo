package model

// PositionItem 持仓列表项
type PositionItem struct {
	Id               int64    `json:"id"`
	User             string   `json:"user"`             // 用户钱包地址
	Symbol           string   `json:"symbol"`           // 交易对符号
	PositionSize     float64  `json:"positionSize"`     // 持仓数量（负数为空头）
	EntryPrice       float64  `json:"entryPrice"`       // 开仓均价
	MarkPrice        float64  `json:"markPrice"`        // 标记价格
	LiqPrice         float64  `json:"liqPrice"`         // 强平价格
	Leverage         int      `json:"leverage"`         // 杠杆倍数
	MarginBalance    float64  `json:"marginBalance"`    // 保证金余额
	PositionValueUsd float64  `json:"positionValueUsd"` // 持仓价值(USD)
	UnrealizedPnL    float64  `json:"unrealizedPnL"`    // 未实现盈亏
	FundingFee       float64  `json:"fundingFee"`       // 资金费用
	MarginMode       string   `json:"marginMode"`       // 保证金模式 cross/isolated
	CreateTime       int64    `json:"createTime"`       // 创建时间戳(ms)
	UpdateTime       int64    `json:"updateTime"`       // 更新时间戳(ms)
	Labels           []string `json:"labels"`           // 标签列表
}

// PositionStatsPoint 持仓统计数据点
type PositionStatsPoint struct {
	Timestamp     int64   `json:"timestamp"`     // 时间桶时间戳(ms)
	LongAccounts  int     `json:"longAccounts"`  // 多头账户数
	ShortAccounts int     `json:"shortAccounts"` // 空头账户数
	LongLiqValue  float64 `json:"longLiqValue"`  // 多头清算价值
	ShortLiqValue float64 `json:"shortLiqValue"` // 空头清算价值
}
