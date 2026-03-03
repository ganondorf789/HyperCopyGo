// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/trader_positions/v1"
)

type (
	ITraderPositions interface {
		List(ctx context.Context, in v1.TraderPositionsListReq) (res *v1.TraderPositionsListRes, err error)
	}
)

var (
	localTraderPositions ITraderPositions
)

func TraderPositions() ITraderPositions {
	if localTraderPositions == nil {
		panic("implement not found for interface ITraderPositions, forgot register?")
	}
	return localTraderPositions
}

func RegisterTraderPositions(i ITraderPositions) {
	localTraderPositions = i
}
