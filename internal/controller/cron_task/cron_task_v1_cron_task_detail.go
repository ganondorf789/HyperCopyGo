package cron_task

import (
	"context"

	"demo/api/cron_task/v1"
	"demo/internal/service"
)

func (c *ControllerV1) CronTaskDetail(ctx context.Context, req *v1.CronTaskDetailReq) (res *v1.CronTaskDetailRes, err error) {
	return service.CronTask().Detail(ctx, req.Id)
}
