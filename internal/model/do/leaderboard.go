// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Leaderboard is the golang structure of table leaderboard for DAO operations like Where/Data.
type Leaderboard struct {
	g.Meta       `orm:"table:leaderboard, do:true"`
	Id           any         // 主键ID
	EthAddress   any         // 钱包地址
	AccountValue any         // 账户价值
	Pnl          any         // 盈亏
	Roi          any         // 投资回报率
	Vlm          any         // 交易量
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	Window       any         // 统计窗口 day/week/month/allTime
}
