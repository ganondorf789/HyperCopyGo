package service

import (
	"context"

	v1 "demo/api/my_track_wallet/v1"
)

type IMyTrackWallet interface {
	Create(ctx context.Context, userId int64, in v1.MyTrackWalletCreateReq) (res *v1.MyTrackWalletCreateRes, err error)
	Update(ctx context.Context, userId int64, in v1.MyTrackWalletUpdateReq) error
	Delete(ctx context.Context, userId int64, id int64) error
	Detail(ctx context.Context, userId int64, id int64) (res *v1.MyTrackWalletDetailRes, err error)
	List(ctx context.Context, userId int64, in v1.MyTrackWalletListReq) (res *v1.MyTrackWalletListRes, err error)
}

var localMyTrackWallet IMyTrackWallet

func MyTrackWallet() IMyTrackWallet {
	return localMyTrackWallet
}

func RegisterMyTrackWallet(s IMyTrackWallet) {
	localMyTrackWallet = s
}
