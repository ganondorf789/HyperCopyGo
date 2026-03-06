// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package copy_trade_config

import (
	"context"

	"demo/api/copy_trade_config/v1"
)

type ICopyTradeConfigV1 interface {
	CopyTradeConfigCreate(ctx context.Context, req *v1.CopyTradeConfigCreateReq) (res *v1.CopyTradeConfigCreateRes, err error)
	CopyTradeConfigUpdate(ctx context.Context, req *v1.CopyTradeConfigUpdateReq) (res *v1.CopyTradeConfigUpdateRes, err error)
	CopyTradeConfigDelete(ctx context.Context, req *v1.CopyTradeConfigDeleteReq) (res *v1.CopyTradeConfigDeleteRes, err error)
	CopyTradeConfigDetail(ctx context.Context, req *v1.CopyTradeConfigDetailReq) (res *v1.CopyTradeConfigDetailRes, err error)
	CopyTradeConfigList(ctx context.Context, req *v1.CopyTradeConfigListReq) (res *v1.CopyTradeConfigListRes, err error)
	CopyTradeConfigRecordList(ctx context.Context, req *v1.CopyTradeConfigRecordListReq) (res *v1.CopyTradeConfigRecordListRes, err error)
}
