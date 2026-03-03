package trader_performance

import (
	"context"

	"demo/api/trader_performance/v1"
	"demo/internal/service"
)

func (c *ControllerV1) TraderPerformanceSummary(ctx context.Context, req *v1.TraderPerformanceSummaryReq) (res *v1.TraderPerformanceSummaryRes, err error) {
	return service.TraderPerformance().Summary(ctx, *req)
}
