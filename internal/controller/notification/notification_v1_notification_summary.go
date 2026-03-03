package notification

import (
	"context"

	"demo/api/notification/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) NotificationSummary(ctx context.Context, req *v1.NotificationSummaryReq) (res *v1.NotificationSummaryRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Notification().Summary(ctx, userId)
}
