package service

import (
	"context"

	v1 "demo/api/copy_trading_grpc/v1"
)

type ICopyTradingGrpc interface {
	GetAutoCopyTradingList(ctx context.Context, appId, appSecret string) (list []*v1.CopyTradingItem, err error)
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
