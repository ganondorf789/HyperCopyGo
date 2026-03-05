package my_track_wallet

import (
	"context"

	"demo/api/my_track_wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) MyTrackWalletUpdate(ctx context.Context, req *v1.MyTrackWalletUpdateReq) (res *v1.MyTrackWalletUpdateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.MyTrackWallet().Update(ctx, userId, *req)
	return
}
