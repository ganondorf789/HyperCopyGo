package copy_trading

import (
	"context"

	"demo/api/copy_trading/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CopyTradingCreate(ctx context.Context, req *v1.CopyTradingCreateReq) (res *v1.CopyTradingCreateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.CopyTrading().Create(ctx, userId, *req)
}
