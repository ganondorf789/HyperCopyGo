// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package copy_trading

import (
	"context"

	"demo/api/copy_trading/v1"
)

type ICopyTradingV1 interface {
	CopyTradingList(ctx context.Context, req *v1.CopyTradingListReq) (res *v1.CopyTradingListRes, err error)
	CopyTradingStop(ctx context.Context, req *v1.CopyTradingStopReq) (res *v1.CopyTradingStopRes, err error)
}
