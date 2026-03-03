package trader_positions

import (
	"context"

	"demo/api/trader_positions/v1"
	"demo/internal/service"
)

func (c *ControllerV1) TraderPositionsList(ctx context.Context, req *v1.TraderPositionsListReq) (res *v1.TraderPositionsListRes, err error) {
	return service.TraderPositions().List(ctx, *req)
}
