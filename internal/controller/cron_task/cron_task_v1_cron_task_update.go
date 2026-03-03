package cron_task

import (
	"context"

	"demo/api/cron_task/v1"
	"demo/internal/service"
)

func (c *ControllerV1) CronTaskUpdate(ctx context.Context, req *v1.CronTaskUpdateReq) (res *v1.CronTaskUpdateRes, err error) {
	return nil, service.CronTask().Update(ctx, *req)
}
