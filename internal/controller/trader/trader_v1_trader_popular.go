package trader

import (
	"context"

	"demo/api/trader/v1"
	"demo/internal/service"
)

func (c *ControllerV1) TraderPopular(ctx context.Context, req *v1.TraderPopularReq) (res *v1.TraderPopularRes, err error) {
	return service.Trader().Popular(ctx)
}
