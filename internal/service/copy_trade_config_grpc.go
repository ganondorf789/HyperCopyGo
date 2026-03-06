package service

import (
	"context"

	v1 "demo/api/copy_trade_config_grpc/v1"
)

type ICopyTradeConfigGrpc interface {
	GetAutoCopyTradeConfigList(ctx context.Context, appId, appSecret string) (list []*v1.CopyTradeConfigItem, err error)
}

var localCopyTradeConfigGrpc ICopyTradeConfigGrpc

func CopyTradeConfigGrpc() ICopyTradeConfigGrpc {
	if localCopyTradeConfigGrpc == nil {
		panic("implement not found for interface ICopyTradeConfigGrpc, forgot register?")
	}
	return localCopyTradeConfigGrpc
}

func RegisterCopyTradeConfigGrpc(s ICopyTradeConfigGrpc) {
	localCopyTradeConfigGrpc = s
}

