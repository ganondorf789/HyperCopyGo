package trader_performance

import (
	"context"

	"demo/api/trader_performance/v1"
	"demo/internal/service"
)

func (c *ControllerV1) TraderPerformance(ctx context.Context, req *v1.TraderPerformanceReq) (res *v1.TraderPerformanceRes, err error) {
	return service.TraderPerformance().Performance(ctx, *req)
}
