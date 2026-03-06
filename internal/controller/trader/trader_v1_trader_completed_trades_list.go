package trader

import (
	"context"

	v1 "demo/api/trader/v1"
	"demo/internal/service"
)

func (c *ControllerV1) TraderCompletedTradesList(ctx context.Context, req *v1.TraderCompletedTradesListReq) (res *v1.TraderCompletedTradesListRes, err error) {
	return service.Trader().CompletedTradesList(ctx, *req)
}
