package trader

import (
	"context"

	"demo/api/trader/v1"
	"demo/internal/service"
)

func (c *ControllerV1) TraderKolList(ctx context.Context, req *v1.TraderKolListReq) (res *v1.TraderKolListRes, err error) {
	return service.Trader().KolList(ctx, *req)
}
