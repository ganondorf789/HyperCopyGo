// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemSetting is the golang structure of table system_setting for DAO operations like Where/Data.
type SystemSetting struct {
	g.Meta                 `orm:"table:system_setting, do:true"`
	Id                     any         // 主键ID
	MarketMinutes          any         // 行情监控时间窗口（分钟）
	MarketNewPositionCount any         // 时间窗口内新仓位数量阈值
	CreatedAt              *gtime.Time // 创建时间
	UpdatedAt              *gtime.Time // 更新时间
	LimitTradingWallet     any         // 交易钱包数量限制（0=不限制）
	LimitCopyTrading       any         // 跟单交易数量限制（0=不限制）
	LimitWatchedAddress    any         // 监控地址数量限制（0=不限制）
}
