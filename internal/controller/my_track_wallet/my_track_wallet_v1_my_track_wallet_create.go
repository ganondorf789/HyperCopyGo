package my_track_wallet

import (
	"context"

	"demo/api/my_track_wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) MyTrackWalletCreate(ctx context.Context, req *v1.MyTrackWalletCreateReq) (res *v1.MyTrackWalletCreateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.MyTrackWallet().Create(ctx, userId, *req)
}
