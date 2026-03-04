package trader

import (
	"context"

	"demo/api/trader/v1"
	"demo/internal/service"
)

func (c *ControllerV1) TraderDiscover(ctx context.Context, req *v1.TraderDiscoverReq) (res *v1.TraderDiscoverRes, err error) {
	return service.Trader().Discover(ctx, *req)
}
