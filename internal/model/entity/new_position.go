// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NewPosition is the golang structure for table new_position.
type NewPosition struct {
	Id                    int64       `json:"id"                    orm:"id"                       description:"主键ID"`                 // 主键ID
	Address               string      `json:"address"               orm:"address"                  description:"钱包地址"`                 // 钱包地址
	Coin                  string      `json:"coin"                  orm:"coin"                     description:"币种"`                   // 币种
	Szi                   float64     `json:"szi"                   orm:"szi"                      description:"仓位大小（正值为多头，负值为空头）"`    // 仓位大小（正值为多头，负值为空头）
	LeverageType          string      `json:"leverageType"          orm:"leverage_type"            description:"杠杆类型（cross/isolated）"` // 杠杆类型（cross/isolated）
	Leverage              int64       `json:"leverage"              orm:"leverage"                 description:"杠杆倍数"`                 // 杠杆倍数
	EntryPx               float64     `json:"entryPx"               orm:"entry_px"                 description:"入场价"`                  // 入场价
	PositionValue         float64     `json:"positionValue"         orm:"position_value"           description:"持仓价值"`                 // 持仓价值
	UnrealizedPnl         float64     `json:"unrealizedPnl"         orm:"unrealized_pnl"           description:"未实现盈亏"`                // 未实现盈亏
	ReturnOnEquity        float64     `json:"returnOnEquity"        orm:"return_on_equity"         description:"权益回报率"`                // 权益回报率
	LiquidationPx         float64     `json:"liquidationPx"         orm:"liquidation_px"           description:"清算价"`                  // 清算价
	MarginUsed            float64     `json:"marginUsed"            orm:"margin_used"              description:"已用保证金"`                // 已用保证金
	MaxLeverage           int64       `json:"maxLeverage"           orm:"max_leverage"             description:"最大允许杠杆"`               // 最大允许杠杆
	CumFundingAllTime     float64     `json:"cumFundingAllTime"     orm:"cum_funding_all_time"     description:"累计资金费（全部时间）"`          // 累计资金费（全部时间）
	CumFundingSinceOpen   float64     `json:"cumFundingSinceOpen"   orm:"cum_funding_since_open"   description:"累计资金费（开仓以来）"`          // 累计资金费（开仓以来）
	CumFundingSinceChange float64     `json:"cumFundingSinceChange" orm:"cum_funding_since_change" description:"累计资金费（最近变更以来）"`        // 累计资金费（最近变更以来）
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"               description:"创建时间"`                 // 创建时间
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"               description:"更新时间"`                 // 更新时间
}
