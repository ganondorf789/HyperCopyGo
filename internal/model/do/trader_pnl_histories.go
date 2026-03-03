// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderPnlHistories is the golang structure of table trader_pnl_histories for DAO operations like Where/Data.
type TraderPnlHistories struct {
	g.Meta    `orm:"table:trader_pnl_histories, do:true"`
	Id        any         // 主键ID
	Address   any         // 钱包地址
	Window    any         // 统计窗口（day/week/month/allTime）
	History   any         // 盈亏历史数据（[[timestamp, value], ...]）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
