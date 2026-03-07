package trade

import (
	"context"

	"demo/api/trade/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) TradeOpenOrders(ctx context.Context, req *v1.TradeOpenOrdersReq) (res *v1.TradeOpenOrdersRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Trade().OpenOrders(ctx, userId, *req)
}
