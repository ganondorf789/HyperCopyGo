// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FetchFailures is the golang structure of table fetch_failures for DAO operations like Where/Data.
type FetchFailures struct {
	g.Meta      `orm:"table:fetch_failures, do:true"`
	Id          any         // 主键ID
	Type        any         // 数据类型（fills/orders/funding）
	Reason      any         // 失败原因（exceeds_limit/rate_limited）
	Address     any         // 钱包地址
	StartMs     any         // 失败时间窗口开始（毫秒时间戳）
	EndMs       any         // 失败时间窗口结束（毫秒时间戳）
	RecordCount any         // 该窗口返回的记录数
	CreatedAt   *gtime.Time // 创建时间
}
