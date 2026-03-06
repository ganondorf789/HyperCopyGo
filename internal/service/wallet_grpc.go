package service

import (
	"context"

	v1 "demo/api/wallet_grpc/v1"
)

type IWalletGrpc interface {
	GetWalletList(ctx context.Context, appId, appSecret string) (list []*v1.WalletItem, err error)
}

var localWalletGrpc IWalletGrpc

func WalletGrpc() IWalletGrpc {
	if localWalletGrpc == nil {
		panic("implement not found for interface IWalletGrpc, forgot register?")
	}
	return localWalletGrpc
}

func RegisterWalletGrpc(s IWalletGrpc) {
	localWalletGrpc = s
}
