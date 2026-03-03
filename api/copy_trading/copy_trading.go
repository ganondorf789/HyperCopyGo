// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package copy_trading

import (
	"context"

	"demo/api/copy_trading/v1"
)

type ICopyTradingV1 interface {
	CopyTradingCreate(ctx context.Context, req *v1.CopyTradingCreateReq) (res *v1.CopyTradingCreateRes, err error)
	CopyTradingUpdate(ctx context.Context, req *v1.CopyTradingUpdateReq) (res *v1.CopyTradingUpdateRes, err error)
	CopyTradingDelete(ctx context.Context, req *v1.CopyTradingDeleteReq) (res *v1.CopyTradingDeleteRes, err error)
	CopyTradingDetail(ctx context.Context, req *v1.CopyTradingDetailReq) (res *v1.CopyTradingDetailRes, err error)
	CopyTradingList(ctx context.Context, req *v1.CopyTradingListReq) (res *v1.CopyTradingListRes, err error)
}
