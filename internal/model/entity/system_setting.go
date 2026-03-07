// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemSetting is the golang structure for table system_setting.
type SystemSetting struct {
	Id                     int64       `json:"id"                     orm:"id"                        description:"主键ID"`            // 主键ID
	MarketMinutes          int64       `json:"marketMinutes"          orm:"market_minutes"            description:"行情监控时间窗口（分钟）"`    // 行情监控时间窗口（分钟）
	MarketNewPositionCount int64       `json:"marketNewPositionCount" orm:"market_new_position_count" description:"时间窗口内新仓位数量阈值"`    // 时间窗口内新仓位数量阈值
	CreatedAt              *gtime.Time `json:"createdAt"              orm:"created_at"                description:"创建时间"`            // 创建时间
	UpdatedAt              *gtime.Time `json:"updatedAt"              orm:"updated_at"                description:"更新时间"`            // 更新时间
	LimitTradingWallet     int64       `json:"limitTradingWallet"     orm:"limit_trading_wallet"      description:"交易钱包数量限制（0=不限制）"` // 交易钱包数量限制（0=不限制）
	LimitCopyTrading       int64       `json:"limitCopyTrading"       orm:"limit_copy_trading"        description:"跟单交易数量限制（0=不限制）"` // 跟单交易数量限制（0=不限制）
	LimitWatchedAddress    int64       `json:"limitWatchedAddress"    orm:"limit_watched_address"     description:"监控地址数量限制（0=不限制）"` // 监控地址数量限制（0=不限制）
}
