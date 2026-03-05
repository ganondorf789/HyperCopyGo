package position

import (
	"context"

	"demo/api/position/v1"
	"demo/internal/service"
)

func (c *ControllerV1) PositionStats(ctx context.Context, req *v1.PositionStatsReq) (res *v1.PositionStatsRes, err error) {
	return service.Position().Stats(ctx, *req)
}
