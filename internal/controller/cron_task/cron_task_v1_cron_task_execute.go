package cron_task

import (
	"context"

	"demo/api/cron_task/v1"
	"demo/internal/service"
)

func (c *ControllerV1) CronTaskExecute(ctx context.Context, req *v1.CronTaskExecuteReq) (res *v1.CronTaskExecuteRes, err error) {
	return service.CronTask().Execute(ctx, req.Id)
}
