package wallet

import (
	"context"

	"demo/api/wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) WalletDelete(ctx context.Context, req *v1.WalletDeleteReq) (res *v1.WalletDeleteRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.Wallet().Delete(ctx, userId, req.Id)
	return
}
