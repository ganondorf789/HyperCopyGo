// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CompletedTrades is the golang structure for table completed_trades.
type CompletedTrades struct {
	Id         int64       `json:"id"         orm:"id"          description:""`              //
	User       string      `json:"user"       orm:"user"        description:"交易员钱包地址"`       // 交易员钱包地址
	Coin       string      `json:"coin"       orm:"coin"        description:"币种"`            // 币种
	Side       string      `json:"side"       orm:"side"        description:"方向 long/short"` // 方向 long/short
	EntryPx    float64     `json:"entryPx"    orm:"entry_px"    description:"开仓均价"`          // 开仓均价
	ClosePx    float64     `json:"closePx"    orm:"close_px"    description:"平仓均价"`          // 平仓均价
	Sz         float64     `json:"sz"         orm:"sz"          description:"交易数量"`          // 交易数量
	ClosedPnl  float64     `json:"closedPnl"  orm:"closed_pnl"  description:"已实现盈亏"`         // 已实现盈亏
	TotalFee   float64     `json:"totalFee"   orm:"total_fee"   description:"总手续费"`          // 总手续费
	OpenTime   int64       `json:"openTime"   orm:"open_time"   description:"开仓时间戳(ms)"`     // 开仓时间戳(ms)
	CloseTime  int64       `json:"closeTime"  orm:"close_time"  description:"平仓时间戳(ms)"`     // 平仓时间戳(ms)
	DurationMs int64       `json:"durationMs" orm:"duration_ms" description:"持仓时长(ms)"`      // 持仓时长(ms)
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`              //
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`              //
}
