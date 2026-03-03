// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CronTask is the golang structure for table cron_task.
type CronTask struct {
	Id          int64       `json:"id"          orm:"id"            description:"主键ID"`                                                                                 // 主键ID
	Name        string      `json:"name"        orm:"name"          description:"任务名称"`                                                                                 // 任务名称
	CronExpr    string      `json:"cronExpr"    orm:"cron_expr"     description:"Cron表达式"`                                                                              // Cron表达式
	TaskType    string      `json:"taskType"    orm:"task_type"     description:"任务类型 sync_leaderboard/sync_positions/sync_fills/copy_trade/track_wallet/market_alert"` // 任务类型 sync_leaderboard/sync_positions/sync_fills/copy_trade/track_wallet/market_alert
	Params      string      `json:"params"      orm:"params"        description:"任务参数(JSON)"`                                                                           // 任务参数(JSON)
	LastRunAt   *gtime.Time `json:"lastRunAt"   orm:"last_run_at"   description:"上次执行时间"`                                                                               // 上次执行时间
	LastRunCost int64       `json:"lastRunCost" orm:"last_run_cost" description:"上次执行耗时(毫秒)"`                                                                           // 上次执行耗时(毫秒)
	LastError   string      `json:"lastError"   orm:"last_error"    description:"上次执行错误信息"`                                                                             // 上次执行错误信息
	RunCount    int64       `json:"runCount"    orm:"run_count"     description:"累计执行次数"`                                                                               // 累计执行次数
	Status      int         `json:"status"      orm:"status"        description:"状态 1:启用 0:停用"`                                                                         // 状态 1:启用 0:停用
	Remark      string      `json:"remark"      orm:"remark"        description:"备注"`                                                                                   // 备注
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:"创建时间"`                                                                                 // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:"更新时间"`                                                                                 // 更新时间
}
