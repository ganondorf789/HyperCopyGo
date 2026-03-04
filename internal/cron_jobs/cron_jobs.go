package cron_jobs

import "context"

// TaskHandler 定时任务处理函数签名，params 为数据库中配置的 JSON 参数
type TaskHandler func(ctx context.Context, params string)

// Handlers 任务类型 → 处理函数 的注册表
var Handlers = map[string]TaskHandler{}

// Register 注册一个任务类型对应的处理函数
func Register(taskType string, handler TaskHandler) {
	Handlers[taskType] = handler
}
