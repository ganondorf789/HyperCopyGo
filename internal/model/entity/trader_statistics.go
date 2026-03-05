// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderStatistics is the golang structure for table trader_statistics.
type TraderStatistics struct {
	Id                 int64       `json:"id"                 orm:"id"                   description:"主键ID"`                         // 主键ID
	Address            string      `json:"address"            orm:"address"              description:"钱包地址"`                         // 钱包地址
	Window             string      `json:"window"             orm:"window"               description:"统计窗口（day/week/month/allTime）"` // 统计窗口（day/week/month/allTime）
	Sharpe             float64     `json:"sharpe"             orm:"sharpe"               description:"夏普比率"`                         // 夏普比率
	Drawdown           float64     `json:"drawdown"           orm:"drawdown"             description:"最大回撤"`                         // 最大回撤
	PositionCount      float64     `json:"positionCount"      orm:"position_count"       description:"持仓数"`                          // 持仓数
	TotalValue         float64     `json:"totalValue"         orm:"total_value"          description:"账户总价值"`                        // 账户总价值
	PerpValue          float64     `json:"perpValue"          orm:"perp_value"           description:"永续合约总价值"`                      // 永续合约总价值
	PositionValue      float64     `json:"positionValue"      orm:"position_value"       description:"持仓价值"`                         // 持仓价值
	LongPositionValue  float64     `json:"longPositionValue"  orm:"long_position_value"  description:"多仓仓位价值"`                       // 多仓仓位价值
	ShortPositionValue float64     `json:"shortPositionValue" orm:"short_position_value" description:"空仓仓位价值"`                       // 空仓仓位价值
	MarginUsage        float64     `json:"marginUsage"        orm:"margin_usage"         description:"保证金使用率"`                       // 保证金使用率
	UsedMargin         float64     `json:"usedMargin"         orm:"used_margin"          description:"已用保证金"`                        // 已用保证金
	ProfitCount        float64     `json:"profitCount"        orm:"profit_count"         description:"盈利次数"`                         // 盈利次数
	WinRate            float64     `json:"winRate"            orm:"win_rate"             description:"胜率"`                           // 胜率
	TotalPnl           float64     `json:"totalPnl"           orm:"total_pnl"            description:"总盈亏"`                          // 总盈亏
	LongCount          float64     `json:"longCount"          orm:"long_count"           description:"多仓数"`                          // 多仓数
	LongRealizedPnl    float64     `json:"longRealizedPnl"    orm:"long_realized_pnl"    description:"多仓已实现盈亏"`                      // 多仓已实现盈亏
	LongWinRate        float64     `json:"longWinRate"        orm:"long_win_rate"        description:"多仓胜率"`                         // 多仓胜率
	ShortCount         float64     `json:"shortCount"         orm:"short_count"          description:"空仓数"`                          // 空仓数
	ShortRealizedPnl   float64     `json:"shortRealizedPnl"   orm:"short_realized_pnl"   description:"空仓已实现盈亏"`                      // 空仓已实现盈亏
	ShortWinRate       float64     `json:"shortWinRate"       orm:"short_win_rate"       description:"空仓胜率"`                         // 空仓胜率
	UnrealizedPnl      float64     `json:"unrealizedPnl"      orm:"unrealized_pnl"       description:"未实现盈亏"`                        // 未实现盈亏
	AvgLeverage        float64     `json:"avgLeverage"        orm:"avg_leverage"         description:"平均杠杆"`                         // 平均杠杆
	Coins              []string    `json:"coins"              orm:"coins"                description:"交易过的币种"`                       // 交易过的币种
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:"创建时间"`                         // 创建时间
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:"更新时间"`                         // 更新时间
}
