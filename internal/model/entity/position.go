// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Position is the golang structure for table position.
type Position struct {
	Id               int64       `json:"id"               orm:"id"                 description:"主键ID"`                 // 主键ID
	User             string      `json:"user"             orm:"user"               description:"用户钱包地址"`               // 用户钱包地址
	Symbol           string      `json:"symbol"           orm:"symbol"             description:"交易对符号"`                // 交易对符号
	PositionSize     float64     `json:"positionSize"     orm:"position_size"      description:"持仓数量（负数为空头）"`          // 持仓数量（负数为空头）
	EntryPrice       float64     `json:"entryPrice"       orm:"entry_price"        description:"开仓均价"`                 // 开仓均价
	MarkPrice        float64     `json:"markPrice"        orm:"mark_price"         description:"标记价格"`                 // 标记价格
	LiqPrice         float64     `json:"liqPrice"         orm:"liq_price"          description:"强平价格"`                 // 强平价格
	Leverage         int64       `json:"leverage"         orm:"leverage"           description:"杠杆倍数"`                 // 杠杆倍数
	MarginBalance    float64     `json:"marginBalance"    orm:"margin_balance"     description:"保证金余额"`                // 保证金余额
	PositionValueUsd float64     `json:"positionValueUsd" orm:"position_value_usd" description:"持仓价值(USD)"`            // 持仓价值(USD)
	UnrealizedPnl    float64     `json:"unrealizedPnl"    orm:"unrealized_pnl"     description:"未实现盈亏"`                // 未实现盈亏
	FundingFee       float64     `json:"fundingFee"       orm:"funding_fee"        description:"资金费用"`                 // 资金费用
	MarginMode       string      `json:"marginMode"       orm:"margin_mode"        description:"保证金模式 cross/isolated"` // 保证金模式 cross/isolated
	Labels           string      `json:"labels"           orm:"labels"             description:"标签,逗号分隔"`              // 标签,逗号分隔
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`                 // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:"更新时间"`                 // 更新时间
}
