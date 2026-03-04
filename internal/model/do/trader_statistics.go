// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderStatistics is the golang structure of table trader_statistics for DAO operations like Where/Data.
type TraderStatistics struct {
	g.Meta             `orm:"table:trader_statistics, do:true"`
	Id                 any         // 主键ID
	Address            any         // 钱包地址
	Window             any         // 统计窗口（day/week/month/allTime）
	Sharpe             any         // 夏普比率
	Drawdown           any         // 最大回撤
	PositionCount      any         // 持仓数
	TotalValue         any         // 账户总价值
	PerpValue          any         // 永续合约总价值
	PositionValue      any         // 持仓价值
	LongPositionValue  any         // 多仓仓位价值
	ShortPositionValue any         // 空仓仓位价值
	MarginUsage        any         // 保证金使用率
	UsedMargin         any         // 已用保证金
	ProfitCount        any         // 盈利次数
	WinRate            any         // 胜率
	TotalPnl           any         // 总盈亏
	LongCount          any         // 多仓数
	LongRealizedPnl    any         // 多仓已实现盈亏
	LongWinRate        any         // 多仓胜率
	ShortCount         any         // 空仓数
	ShortRealizedPnl   any         // 空仓已实现盈亏
	ShortWinRate       any         // 空仓胜率
	UnrealizedPnl      any         // 未实现盈亏
	AvgLeverage        any         // 平均杠杆
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
	Coins              []string    // 交易过的币种
}
