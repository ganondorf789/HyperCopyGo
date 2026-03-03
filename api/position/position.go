// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package position

import (
	"context"

	"demo/api/position/v1"
)

type IPositionV1 interface {
	PositionList(ctx context.Context, req *v1.PositionListReq) (res *v1.PositionListRes, err error)
	PositionStats(ctx context.Context, req *v1.PositionStatsReq) (res *v1.PositionStatsRes, err error)
	PositionLongShortRatio(ctx context.Context, req *v1.PositionLongShortRatioReq) (res *v1.PositionLongShortRatioRes, err error)
}
