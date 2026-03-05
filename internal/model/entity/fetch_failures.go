// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FetchFailures is the golang structure for table fetch_failures.
type FetchFailures struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键ID"`                             // 主键ID
	Type        string      `json:"type"        orm:"type"         description:"数据类型（fills/orders/funding）"`       // 数据类型（fills/orders/funding）
	Reason      string      `json:"reason"      orm:"reason"       description:"失败原因（exceeds_limit/rate_limited）"` // 失败原因（exceeds_limit/rate_limited）
	Address     string      `json:"address"     orm:"address"      description:"钱包地址"`                             // 钱包地址
	StartMs     int64       `json:"startMs"     orm:"start_ms"     description:"失败时间窗口开始（毫秒时间戳）"`                  // 失败时间窗口开始（毫秒时间戳）
	EndMs       int64       `json:"endMs"       orm:"end_ms"       description:"失败时间窗口结束（毫秒时间戳）"`                  // 失败时间窗口结束（毫秒时间戳）
	RecordCount int64       `json:"recordCount" orm:"record_count" description:"该窗口返回的记录数"`                        // 该窗口返回的记录数
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`                             // 创建时间
}
