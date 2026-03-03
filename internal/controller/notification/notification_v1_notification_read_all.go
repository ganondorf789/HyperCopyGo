package notification

import (
	"context"

	"demo/api/notification/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) NotificationReadAll(ctx context.Context, req *v1.NotificationReadAllReq) (res *v1.NotificationReadAllRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return nil, service.Notification().ReadAll(ctx, userId, req.Category)
}
