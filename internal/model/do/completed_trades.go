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
	Id         any         //
	User       any         // 交易员钱包地址
	Coin       any         // 币种
	Side       any         // 方向 long/short
	EntryPx    any         // 开仓均价
	ClosePx    any         // 平仓均价
	Sz         any         // 交易数量
	ClosedPnl  any         // 已实现盈亏
	TotalFee   any         // 总手续费
	OpenTime   any         // 开仓时间戳(ms)
	CloseTime  any         // 平仓时间戳(ms)
	DurationMs any         // 持仓时长(ms)
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
