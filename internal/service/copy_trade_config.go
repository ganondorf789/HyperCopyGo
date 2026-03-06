package service

import (
	"context"

	v1 "demo/api/copy_trade_config/v1"
)

type ICopyTradeConfig interface {
	Create(ctx context.Context, userId int64, in v1.CopyTradeConfigCreateReq) (res *v1.CopyTradeConfigCreateRes, err error)
	Update(ctx context.Context, userId int64, in v1.CopyTradeConfigUpdateReq) error
	Delete(ctx context.Context, userId int64, id int64) error
	Detail(ctx context.Context, userId int64, id int64) (res *v1.CopyTradeConfigDetailRes, err error)
	List(ctx context.Context, userId int64, in v1.CopyTradeConfigListReq) (res *v1.CopyTradeConfigListRes, err error)
	RecordList(ctx context.Context, userId int64, in v1.CopyTradeConfigRecordListReq) (res *v1.CopyTradeConfigRecordListRes, err error)
}

var localCopyTradeConfig ICopyTradeConfig

func CopyTradeConfig() ICopyTradeConfig {
	return localCopyTradeConfig
}

func RegisterCopyTradeConfig(s ICopyTradeConfig) {
	localCopyTradeConfig = s
}

