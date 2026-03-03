// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/trader_performance/v1"
)

type (
	ITraderPerformance interface {
		Performance(ctx context.Context, in v1.TraderPerformanceReq) (res *v1.TraderPerformanceRes, err error)
		Summary(ctx context.Context, in v1.TraderPerformanceSummaryReq) (res *v1.TraderPerformanceSummaryRes, err error)
	}
)

var (
	localTraderPerformance ITraderPerformance
)

func TraderPerformance() ITraderPerformance {
	if localTraderPerformance == nil {
		panic("implement not found for interface ITraderPerformance, forgot register?")
	}
	return localTraderPerformance
}

func RegisterTraderPerformance(i ITraderPerformance) {
	localTraderPerformance = i
}
