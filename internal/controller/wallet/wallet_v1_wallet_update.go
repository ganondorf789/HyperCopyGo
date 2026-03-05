package wallet

import (
	"context"

	"demo/api/wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) WalletUpdate(ctx context.Context, req *v1.WalletUpdateReq) (res *v1.WalletUpdateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	err = service.Wallet().Update(ctx, userId, *req)
	return
}
