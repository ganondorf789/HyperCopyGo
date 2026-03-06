package copy_trade_config

import (
	"context"

	"demo/api/copy_trade_config/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CopyTradeConfigUpdate(ctx context.Context, req *v1.CopyTradeConfigUpdateReq) (res *v1.CopyTradeConfigUpdateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.CopyTradeConfig().Update(ctx, userId, *req)
	return
}

