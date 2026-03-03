package notification

import (
	"context"

	"demo/api/notification/v1"
	"demo/internal/service"
)

func (c *ControllerV1) NotificationUpdate(ctx context.Context, req *v1.NotificationUpdateReq) (res *v1.NotificationUpdateRes, err error) {
	return nil, service.Notification().Update(ctx, *req)
}
