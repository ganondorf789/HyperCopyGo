package my_track_wallet

import (
	"context"

	"demo/api/my_track_wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) MyTrackWalletList(ctx context.Context, req *v1.MyTrackWalletListReq) (res *v1.MyTrackWalletListRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.MyTrackWallet().List(ctx, userId, *req)
}
