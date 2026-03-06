package copy_trading

import (
	"context"

	v1 "demo/api/copy_trading/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CopyTradingRecordList(ctx context.Context, req *v1.CopyTradingRecordListReq) (res *v1.CopyTradingRecordListRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.CopyTrading().RecordList(ctx, userId, *req)
}
