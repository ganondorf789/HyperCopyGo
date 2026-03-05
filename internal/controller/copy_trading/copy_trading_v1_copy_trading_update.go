package copy_trading

import (
	"context"

	"demo/api/copy_trading/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CopyTradingUpdate(ctx context.Context, req *v1.CopyTradingUpdateReq) (res *v1.CopyTradingUpdateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.CopyTrading().Update(ctx, userId, *req)
	return
}
