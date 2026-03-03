// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CronTask is the golang structure of table cron_task for DAO operations like Where/Data.
type CronTask struct {
	g.Meta      `orm:"table:cron_task, do:true"`
	Id          any         // 主键ID
	Name        any         // 任务名称
	CronExpr    any         // Cron表达式
	TaskType    any         // 任务类型 sync_leaderboard/sync_positions/sync_fills/copy_trade/track_wallet/market_alert
	Params      any         // 任务参数(JSON)
	LastRunAt   *gtime.Time // 上次执行时间
	LastRunCost any         // 上次执行耗时(毫秒)
	LastError   any         // 上次执行错误信息
	RunCount    any         // 累计执行次数
	Status      any         // 状态 1:启用 0:停用
	Remark      any         // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
