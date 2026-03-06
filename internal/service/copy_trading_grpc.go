package service

import (
	"context"

	"demo/internal/model"
)

type ICopyTradingGrpc interface {
	GetAutoCopyTradingList(ctx context.Context, appId, appSecret string) (list []model.CopyTradingItem, err error)
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
