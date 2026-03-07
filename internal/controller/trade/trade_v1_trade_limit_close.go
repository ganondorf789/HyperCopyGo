package trade

import (
	"context"

	"demo/api/trade/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) TradeLimitClose(ctx context.Context, req *v1.TradeLimitCloseReq) (res *v1.TradeLimitCloseRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Trade().LimitClose(ctx, userId, *req)
}
