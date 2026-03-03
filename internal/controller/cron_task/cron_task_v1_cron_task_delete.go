package cron_task

import (
	"context"

	"demo/api/cron_task/v1"
	"demo/internal/service"
)

func (c *ControllerV1) CronTaskDelete(ctx context.Context, req *v1.CronTaskDeleteReq) (res *v1.CronTaskDeleteRes, err error) {
	return nil, service.CronTask().Delete(ctx, req.Id)
}
