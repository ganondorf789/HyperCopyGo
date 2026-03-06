// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/copy_trading/v1"
)

type (
	ICopyTrading interface {
		List(ctx context.Context, userId int64, in v1.CopyTradingListReq) (res *v1.CopyTradingListRes, err error)
		Stop(ctx context.Context, userId int64, id int64) error
	}
)

var (
	localCopyTrading ICopyTrading
)

func CopyTrading() ICopyTrading {
	if localCopyTrading == nil {
		panic("implement not found for interface ICopyTrading, forgot register?")
	}
	return localCopyTrading
}

func RegisterCopyTrading(i ICopyTrading) {
	localCopyTrading = i
}
