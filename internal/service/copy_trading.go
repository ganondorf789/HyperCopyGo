package service

import (
	"context"

	v1 "demo/api/copy_trading/v1"
)

type ICopyTrading interface {
	Create(ctx context.Context, userId int64, in v1.CopyTradingCreateReq) (res *v1.CopyTradingCreateRes, err error)
	Update(ctx context.Context, userId int64, in v1.CopyTradingUpdateReq) error
	Delete(ctx context.Context, userId int64, id int64) error
	Detail(ctx context.Context, userId int64, id int64) (res *v1.CopyTradingDetailRes, err error)
	List(ctx context.Context, userId int64, in v1.CopyTradingListReq) (res *v1.CopyTradingListRes, err error)
}

var localCopyTrading ICopyTrading

func CopyTrading() ICopyTrading {
	return localCopyTrading
}

func RegisterCopyTrading(s ICopyTrading) {
	localCopyTrading = s
}
