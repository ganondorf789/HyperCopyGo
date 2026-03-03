package model

import "github.com/gogf/gf/v2/os/gtime"

// BaseCronTask 定时任务公共字段，供 API 请求/响应复用
type BaseCronTask struct {
	Name     string `json:"name"`
	CronExpr string `json:"cronExpr"`
	TaskType string `json:"taskType"`
	Params   string `json:"params"`
	Remark   string `json:"remark"`
}

// CronTaskItem 定时任务列表项
type CronTaskItem struct {
	Id int64 `json:"id"`
	BaseCronTask
	LastRunAt   *gtime.Time `json:"lastRunAt"`
	LastRunCost int64       `json:"lastRunCost"`
	LastError   string      `json:"lastError"`
	RunCount    int64       `json:"runCount"`
	Status      int         `json:"status"`
	CreatedAt   *gtime.Time `json:"createdAt"`
	UpdatedAt   *gtime.Time `json:"updatedAt"`
}
