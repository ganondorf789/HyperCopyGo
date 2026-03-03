// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package trader_performance

import (
	"context"

	"demo/api/trader_performance/v1"
)

type ITraderPerformanceV1 interface {
	TraderPerformance(ctx context.Context, req *v1.TraderPerformanceReq) (res *v1.TraderPerformanceRes, err error)
}
