package position

import (
	"context"

	v1 "demo/api/position/v1"
	"demo/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) PositionList(ctx context.Context, req *v1.PositionListReq) (res *v1.PositionListRes, err error) {
	return service.Position().List(ctx, *req)
}

func (c *Controller) PositionStats(ctx context.Context, req *v1.PositionStatsReq) (res *v1.PositionStatsRes, err error) {
	return service.Position().Stats(ctx, *req)
}

func (c *Controller) PositionLongShortRatio(ctx context.Context, req *v1.PositionLongShortRatioReq) (res *v1.PositionLongShortRatioRes, err error) {
	return service.Position().LongShortRatio(ctx, *req)
}
