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
