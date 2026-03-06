package copy_trade_config

import (
	"context"

	"demo/api/copy_trade_config/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CopyTradeConfigDetail(ctx context.Context, req *v1.CopyTradeConfigDetailReq) (res *v1.CopyTradeConfigDetailRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.CopyTradeConfig().Detail(ctx, userId, req.Id)
}

