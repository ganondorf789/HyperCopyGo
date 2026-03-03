package cron_task

import (
	"context"

	"demo/api/cron_task/v1"
	"demo/internal/service"
)

func (c *ControllerV1) CronTaskList(ctx context.Context, req *v1.CronTaskListReq) (res *v1.CronTaskListRes, err error) {
	return service.CronTask().List(ctx, *req)
}
