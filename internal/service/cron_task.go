// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/cron_task/v1"
)

type (
	ICronTask interface {
		Create(ctx context.Context, in v1.CronTaskCreateReq) (res *v1.CronTaskCreateRes, err error)
		Update(ctx context.Context, in v1.CronTaskUpdateReq) error
		Delete(ctx context.Context, id int64) error
		Detail(ctx context.Context, id int64) (res *v1.CronTaskDetailRes, err error)
		List(ctx context.Context, in v1.CronTaskListReq) (res *v1.CronTaskListRes, err error)
		// Execute 手动执行指定定时任务
		Execute(ctx context.Context, id int64) (res *v1.CronTaskExecuteRes, err error)
		// TaskTypes 返回所有已注册的可用任务类型
		TaskTypes(ctx context.Context) (res *v1.CronTaskTypesRes, err error)
		// StartAll 从数据库加载所有启用的定时任务，根据 cron_expr 注册到 gcron 调度器
		StartAll(ctx context.Context)
	}
)

var (
	localCronTask ICronTask
)

func CronTask() ICronTask {
	if localCronTask == nil {
		panic("implement not found for interface ICronTask, forgot register?")
	}
	return localCronTask
}

func RegisterCronTask(i ICronTask) {
	localCronTask = i
}
