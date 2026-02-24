package copy_trading

import (
	"context"

	v1 "demo/api/copy_trading/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) CopyTradingCreate(ctx context.Context, req *v1.CopyTradingCreateReq) (res *v1.CopyTradingCreateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.CopyTrading().Create(ctx, userId, *req)
}

func (c *Controller) CopyTradingUpdate(ctx context.Context, req *v1.CopyTradingUpdateReq) (res *v1.CopyTradingUpdateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.CopyTrading().Update(ctx, userId, *req)
	return
}

func (c *Controller) CopyTradingDelete(ctx context.Context, req *v1.CopyTradingDeleteReq) (res *v1.CopyTradingDeleteRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.CopyTrading().Delete(ctx, userId, req.Id)
	return
}

func (c *Controller) CopyTradingDetail(ctx context.Context, req *v1.CopyTradingDetailReq) (res *v1.CopyTradingDetailRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.CopyTrading().Detail(ctx, userId, req.Id)
}

func (c *Controller) CopyTradingList(ctx context.Context, req *v1.CopyTradingListReq) (res *v1.CopyTradingListRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.CopyTrading().List(ctx, userId, *req)
}
