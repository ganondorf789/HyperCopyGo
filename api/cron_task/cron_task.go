// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package cron_task

import (
	"context"

	"demo/api/cron_task/v1"
)

type ICronTaskV1 interface {
	CronTaskCreate(ctx context.Context, req *v1.CronTaskCreateReq) (res *v1.CronTaskCreateRes, err error)
	CronTaskUpdate(ctx context.Context, req *v1.CronTaskUpdateReq) (res *v1.CronTaskUpdateRes, err error)
	CronTaskDelete(ctx context.Context, req *v1.CronTaskDeleteReq) (res *v1.CronTaskDeleteRes, err error)
	CronTaskDetail(ctx context.Context, req *v1.CronTaskDetailReq) (res *v1.CronTaskDetailRes, err error)
	CronTaskList(ctx context.Context, req *v1.CronTaskListReq) (res *v1.CronTaskListRes, err error)
}
