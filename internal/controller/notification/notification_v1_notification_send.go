package notification

import (
	"context"

	"demo/api/notification/v1"
	"demo/internal/service"
)

func (c *ControllerV1) NotificationSend(ctx context.Context, req *v1.NotificationSendReq) (res *v1.NotificationSendRes, err error) {
	return service.Notification().Send(ctx, *req)
}
