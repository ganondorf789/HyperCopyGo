// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderPerformances is the golang structure of table trader_performances for DAO operations like Where/Data.
type TraderPerformances struct {
	g.Meta    `orm:"table:trader_performances, do:true"`
	Id        any         // 主键ID
	Address   any         // 钱包地址
	Window    any         // 统计窗口（day/week/month/allTime）
	Pnl       any         // 盈亏
	Roi       any         // 收益率
	Vlm       any         // 交易量
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
