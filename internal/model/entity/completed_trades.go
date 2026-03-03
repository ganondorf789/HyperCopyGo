// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CompletedTrades is the golang structure for table completed_trades.
type CompletedTrades struct {
	Id         int64       `json:"id"         orm:"id"          description:"主键ID"`                  // 主键ID
	Address    string      `json:"address"    orm:"address"     description:"钱包地址"`                  // 钱包地址
	Coin       string      `json:"coin"       orm:"coin"        description:"币种"`                    // 币种
	MarginMode string      `json:"marginMode" orm:"margin_mode" description:"保证金模式（isolated/cross）"` // 保证金模式（isolated/cross）
	Direction  string      `json:"direction"  orm:"direction"   description:"方向（long/short）"`        // 方向（long/short）
	Size       float64     `json:"size"       orm:"size"        description:"最大持仓量"`                 // 最大持仓量
	EntryPrice float64     `json:"entryPrice" orm:"entry_price" description:"加权平均入场价"`               // 加权平均入场价
	ClosePrice float64     `json:"closePrice" orm:"close_price" description:"加权平均平仓价"`               // 加权平均平仓价
	StartTime  int64       `json:"startTime"  orm:"start_time"  description:"开仓时间（毫秒时间戳）"`           // 开仓时间（毫秒时间戳）
	EndTime    int64       `json:"endTime"    orm:"end_time"    description:"平仓时间（毫秒时间戳）"`           // 平仓时间（毫秒时间戳）
	TotalFee   float64     `json:"totalFee"   orm:"total_fee"   description:"总手续费"`                  // 总手续费
	Pnl        float64     `json:"pnl"        orm:"pnl"         description:"已实现盈亏（closedPnl 之和）"`   // 已实现盈亏（closedPnl 之和）
	FillCount  int64       `json:"fillCount"  orm:"fill_count"  description:"成交笔数"`                  // 成交笔数
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`                  // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`                  // 更新时间
}
