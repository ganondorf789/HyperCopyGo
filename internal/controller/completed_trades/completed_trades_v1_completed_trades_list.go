package completed_trades

import (
	"context"

	"demo/api/completed_trades/v1"
	"demo/internal/service"
)

func (c *ControllerV1) CompletedTradesList(ctx context.Context, req *v1.CompletedTradesListReq) (res *v1.CompletedTradesListRes, err error) {
	return service.CompletedTrades().List(ctx, *req)
}
