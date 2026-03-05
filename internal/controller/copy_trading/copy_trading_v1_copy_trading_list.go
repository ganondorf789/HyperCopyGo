package copy_trading

import (
	"context"

	"demo/api/copy_trading/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CopyTradingList(ctx context.Context, req *v1.CopyTradingListReq) (res *v1.CopyTradingListRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.CopyTrading().List(ctx, userId, *req)
}
