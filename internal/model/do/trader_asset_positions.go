// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderAssetPositions is the golang structure of table trader_asset_positions for DAO operations like Where/Data.
type TraderAssetPositions struct {
	g.Meta                `orm:"table:trader_asset_positions, do:true"`
	Id                    any         // 主键ID
	Address               any         // 钱包地址
	Coin                  any         // 币种
	Szi                   any         // 仓位大小（正值为多头，负值为空头）
	LeverageType          any         // 杠杆类型（cross/isolated）
	Leverage              any         // 杠杆倍数
	EntryPx               any         // 入场价
	PositionValue         any         // 持仓价值
	UnrealizedPnl         any         // 未实现盈亏
	ReturnOnEquity        any         // 权益回报率
	LiquidationPx         any         // 清算价
	MarginUsed            any         // 已用保证金
	MaxLeverage           any         // 最大允许杠杆
	CumFundingAllTime     any         // 累计资金费（全部时间）
	CumFundingSinceOpen   any         // 累计资金费（开仓以来）
	CumFundingSinceChange any         // 累计资金费（最近变更以来）
	CreatedAt             *gtime.Time // 创建时间
	UpdatedAt             *gtime.Time // 更新时间
}
