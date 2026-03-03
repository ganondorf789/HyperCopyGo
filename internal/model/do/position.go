// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Position is the golang structure of table position for DAO operations like Where/Data.
type Position struct {
	g.Meta           `orm:"table:position, do:true"`
	Id               any         // 主键ID
	User             any         // 用户钱包地址
	Symbol           any         // 交易对符号
	PositionSize     any         // 持仓数量（负数为空头）
	EntryPrice       any         // 开仓均价
	MarkPrice        any         // 标记价格
	LiqPrice         any         // 强平价格
	Leverage         any         // 杠杆倍数
	MarginBalance    any         // 保证金余额
	PositionValueUsd any         // 持仓价值(USD)
	UnrealizedPnl    any         // 未实现盈亏
	FundingFee       any         // 资金费用
	MarginMode       any         // 保证金模式 cross/isolated
	Labels           any         // 标签,逗号分隔
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
