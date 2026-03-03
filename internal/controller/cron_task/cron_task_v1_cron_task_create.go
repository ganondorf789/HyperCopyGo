package cron_task

import (
	"context"

	"demo/api/cron_task/v1"
	"demo/internal/service"
)

func (c *ControllerV1) CronTaskCreate(ctx context.Context, req *v1.CronTaskCreateReq) (res *v1.CronTaskCreateRes, err error) {
	return service.CronTask().Create(ctx, *req)
}
