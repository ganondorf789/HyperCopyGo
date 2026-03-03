package notification

import (
	"context"

	"demo/api/notification/v1"
	"demo/internal/service"
)

func (c *ControllerV1) NotificationDelete(ctx context.Context, req *v1.NotificationDeleteReq) (res *v1.NotificationDeleteRes, err error) {
	return nil, service.Notification().Delete(ctx, req.Id)
}
