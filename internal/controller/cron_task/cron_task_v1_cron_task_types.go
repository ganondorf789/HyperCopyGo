package cron_task

import (
	"context"

	"demo/api/cron_task/v1"
	"demo/internal/service"
)

func (c *ControllerV1) CronTaskTypes(ctx context.Context, req *v1.CronTaskTypesReq) (res *v1.CronTaskTypesRes, err error) {
	return service.CronTask().TaskTypes(ctx)
}
