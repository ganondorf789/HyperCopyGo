package model

// BaseCronTask 定时任务公共字段，供 API 请求/响应复用
type BaseCronTask struct {
	Name     string `json:"name"`
	CronExpr string `json:"cronExpr"`
	TaskType string `json:"taskType"`
	Params   string `json:"params"`
	Remark   string `json:"remark"`
}
