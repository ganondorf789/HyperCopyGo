// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CompletedTrades is the golang structure of table completed_trades for DAO operations like Where/Data.
type CompletedTrades struct {
	g.Meta     `orm:"table:completed_trades, do:true"`
	Id         any         // 主键ID
	Address    any         // 钱包地址
	Coin       any         // 币种
	MarginMode any         // 保证金模式（isolated/cross）
	Direction  any         // 方向（long/short）
	Size       any         // 最大持仓量
	EntryPrice any         // 加权平均入场价
	ClosePrice any         // 加权平均平仓价
	StartTime  any         // 开仓时间（毫秒时间戳）
	EndTime    any         // 平仓时间（毫秒时间戳）
	TotalFee   any         // 总手续费
	Pnl        any         // 已实现盈亏（closedPnl 之和）
	FillCount  any         // 成交笔数
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
