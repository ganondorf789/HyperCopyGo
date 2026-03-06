package copy_trade_config

import (
	"context"

	v1 "demo/api/copy_trade_config/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CopyTradeConfigRecordList(ctx context.Context, req *v1.CopyTradeConfigRecordListReq) (res *v1.CopyTradeConfigRecordListRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.CopyTradeConfig().RecordList(ctx, userId, *req)
}

