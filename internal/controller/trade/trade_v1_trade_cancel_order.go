package trade

import (
	"context"

	"demo/api/trade/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) TradeCancelOrder(ctx context.Context, req *v1.TradeCancelOrderReq) (res *v1.TradeCancelOrderRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	if err = service.Trade().CancelOrder(ctx, userId, *req); err != nil {
		return nil, err
	}
	return &v1.TradeCancelOrderRes{}, nil
}
