package notification

import (
	"context"

	"demo/api/notification/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) NotificationRead(ctx context.Context, req *v1.NotificationReadReq) (res *v1.NotificationReadRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return nil, service.Notification().Read(ctx, userId, req.Ids)
}
