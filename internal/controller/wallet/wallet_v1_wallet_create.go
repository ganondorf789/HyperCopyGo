package wallet

import (
	"context"

	"demo/api/wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) WalletCreate(ctx context.Context, req *v1.WalletCreateReq) (res *v1.WalletCreateRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Wallet().Create(ctx, userId, *req)
}
