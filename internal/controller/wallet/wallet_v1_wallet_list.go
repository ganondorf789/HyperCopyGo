package wallet

import (
	"context"

	"demo/api/wallet/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) WalletList(ctx context.Context, req *v1.WalletListReq) (res *v1.WalletListRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Wallet().List(ctx, userId, *req)
}
