package service

import (
	"context"

	v1 "demo/api/copy_trading_grpc/v1"
)

type ICopyTradingGrpc interface {
	CreateCopyTrading(ctx context.Context, appId, appSecret string, in *v1.CreateCopyTradingReq) (int64, error)
	GetCopyTradingDetail(ctx context.Context, appId, appSecret string, id int64) (*v1.CopyTradingItem, error)
	GetCopyTradingList(ctx context.Context, appId, appSecret string, copyTradingId int64) (list []*v1.CopyTradingItem, err error)
}

var localCopyTradingGrpc ICopyTradingGrpc

func CopyTradingGrpc() ICopyTradingGrpc {
	if localCopyTradingGrpc == nil {
		panic("implement not found for interface ICopyTradingGrpc, forgot register?")
	}
	return localCopyTradingGrpc
}

func RegisterCopyTradingGrpc(s ICopyTradingGrpc) {
	localCopyTradingGrpc = s
}
