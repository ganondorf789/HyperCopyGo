package service

import (
	"context"

	v1 "demo/api/wallet/v1"
)

type IWallet interface {
	Create(ctx context.Context, userId int64, in v1.WalletCreateReq) (res *v1.WalletCreateRes, err error)
	Update(ctx context.Context, userId int64, in v1.WalletUpdateReq) error
	Delete(ctx context.Context, userId int64, id int64) error
	Detail(ctx context.Context, userId int64, id int64) (res *v1.WalletDetailRes, err error)
	List(ctx context.Context, userId int64, in v1.WalletListReq) (res *v1.WalletListRes, err error)
}

var localWallet IWallet

func Wallet() IWallet {
	return localWallet
}

func RegisterWallet(s IWallet) {
	localWallet = s
}
