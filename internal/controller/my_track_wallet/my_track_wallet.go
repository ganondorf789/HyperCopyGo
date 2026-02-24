package my_track_wallet

import (
	"context"

	v1 "demo/api/my_track_wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) MyTrackWalletCreate(ctx context.Context, req *v1.MyTrackWalletCreateReq) (res *v1.MyTrackWalletCreateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.MyTrackWallet().Create(ctx, userId, *req)
}

func (c *Controller) MyTrackWalletUpdate(ctx context.Context, req *v1.MyTrackWalletUpdateReq) (res *v1.MyTrackWalletUpdateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.MyTrackWallet().Update(ctx, userId, *req)
	return
}

func (c *Controller) MyTrackWalletDelete(ctx context.Context, req *v1.MyTrackWalletDeleteReq) (res *v1.MyTrackWalletDeleteRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.MyTrackWallet().Delete(ctx, userId, req.Id)
	return
}

func (c *Controller) MyTrackWalletDetail(ctx context.Context, req *v1.MyTrackWalletDetailReq) (res *v1.MyTrackWalletDetailRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.MyTrackWallet().Detail(ctx, userId, req.Id)
}

func (c *Controller) MyTrackWalletList(ctx context.Context, req *v1.MyTrackWalletListReq) (res *v1.MyTrackWalletListRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.MyTrackWallet().List(ctx, userId, *req)
}

func (c *Controller) MyTrackWalletExport(ctx context.Context, req *v1.MyTrackWalletExportReq) (res *v1.MyTrackWalletExportRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.MyTrackWallet().Export(ctx, userId)
}

func (c *Controller) MyTrackWalletImport(ctx context.Context, req *v1.MyTrackWalletImportReq) (res *v1.MyTrackWalletImportRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.MyTrackWallet().Import(ctx, userId, *req)
}
