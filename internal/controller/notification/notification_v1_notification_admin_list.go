package notification

import (
	"context"

	"demo/api/notification/v1"
	"demo/internal/service"
)

func (c *ControllerV1) NotificationAdminList(ctx context.Context, req *v1.NotificationAdminListReq) (res *v1.NotificationAdminListRes, err error) {
	return service.Notification().AdminList(ctx, *req)
}
