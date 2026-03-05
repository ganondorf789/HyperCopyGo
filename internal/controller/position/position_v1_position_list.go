package position

import (
	"context"

	"demo/api/position/v1"
	"demo/internal/service"
)

func (c *ControllerV1) PositionList(ctx context.Context, req *v1.PositionListReq) (res *v1.PositionListRes, err error) {
	return service.Position().List(ctx, *req)
}
