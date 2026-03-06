package copy_trading

import (
	"context"

	"demo/api/copy_trading/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CopyTradingStop(ctx context.Context, req *v1.CopyTradingStopReq) (res *v1.CopyTradingStopRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	if err = service.CopyTrading().Stop(ctx, userId, req.Id); err != nil {
		return nil, err
	}
	return &v1.CopyTradingStopRes{}, nil
}
