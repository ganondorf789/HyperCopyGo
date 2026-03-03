// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/completed_trades/v1"
)

type (
	ICompletedTrades interface {
		List(ctx context.Context, in v1.CompletedTradesListReq) (res *v1.CompletedTradesListRes, err error)
	}
)

var (
	localCompletedTrades ICompletedTrades
)

func CompletedTrades() ICompletedTrades {
	if localCompletedTrades == nil {
		panic("implement not found for interface ICompletedTrades, forgot register?")
	}
	return localCompletedTrades
}

func RegisterCompletedTrades(i ICompletedTrades) {
	localCompletedTrades = i
}
