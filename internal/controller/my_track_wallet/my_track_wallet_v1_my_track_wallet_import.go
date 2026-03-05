package my_track_wallet

import (
	"context"

	"demo/api/my_track_wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) MyTrackWalletImport(ctx context.Context, req *v1.MyTrackWalletImportReq) (res *v1.MyTrackWalletImportRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.MyTrackWallet().Import(ctx, userId, *req)
}
