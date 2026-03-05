package my_track_wallet

import (
	"context"

	"demo/api/my_track_wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) MyTrackWalletDelete(ctx context.Context, req *v1.MyTrackWalletDeleteReq) (res *v1.MyTrackWalletDeleteRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.MyTrackWallet().Delete(ctx, userId, req.Id)
	return
}
