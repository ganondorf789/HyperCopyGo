package my_track_wallet

import (
	"context"

	"demo/api/my_track_wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) MyTrackWalletDetail(ctx context.Context, req *v1.MyTrackWalletDetailReq) (res *v1.MyTrackWalletDetailRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.MyTrackWallet().Detail(ctx, userId, req.Id)
}
