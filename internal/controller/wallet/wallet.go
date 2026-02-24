package wallet

import (
	"context"

	v1 "demo/api/wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) WalletCreate(ctx context.Context, req *v1.WalletCreateReq) (res *v1.WalletCreateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Wallet().Create(ctx, userId, *req)
}

func (c *Controller) WalletUpdate(ctx context.Context, req *v1.WalletUpdateReq) (res *v1.WalletUpdateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.Wallet().Update(ctx, userId, *req)
	return
}

func (c *Controller) WalletDelete(ctx context.Context, req *v1.WalletDeleteReq) (res *v1.WalletDeleteRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.Wallet().Delete(ctx, userId, req.Id)
	return
}

func (c *Controller) WalletDetail(ctx context.Context, req *v1.WalletDetailReq) (res *v1.WalletDetailRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Wallet().Detail(ctx, userId, req.Id)
}

func (c *Controller) WalletList(ctx context.Context, req *v1.WalletListReq) (res *v1.WalletListRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Wallet().List(ctx, userId, *req)
}
