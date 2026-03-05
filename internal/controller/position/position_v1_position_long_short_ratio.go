package position

import (
	"context"

	"demo/api/position/v1"
	"demo/internal/service"
)

func (c *ControllerV1) PositionLongShortRatio(ctx context.Context, req *v1.PositionLongShortRatioReq) (res *v1.PositionLongShortRatioRes, err error) {
	return service.Position().LongShortRatio(ctx, *req)
}
