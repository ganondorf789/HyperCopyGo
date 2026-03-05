package wallet

import (
	"context"

	"demo/api/wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) WalletDetail(ctx context.Context, req *v1.WalletDetailReq) (res *v1.WalletDetailRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Wallet().Detail(ctx, userId, req.Id)
}
