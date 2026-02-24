package service

import (
	"context"

	v1 "demo/api/position/v1"
)

type IPosition interface {
	List(ctx context.Context, in v1.PositionListReq) (res *v1.PositionListRes, err error)
	Stats(ctx context.Context, in v1.PositionStatsReq) (res *v1.PositionStatsRes, err error)
	LongShortRatio(ctx context.Context, in v1.PositionLongShortRatioReq) (res *v1.PositionLongShortRatioRes, err error)
}

var localPosition IPosition

func Position() IPosition {
	return localPosition
}

func RegisterPosition(s IPosition) {
	localPosition = s
}
