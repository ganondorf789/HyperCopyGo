// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package completed_trades

import (
	"context"

	"demo/api/completed_trades/v1"
)

type ICompletedTradesV1 interface {
	CompletedTradesList(ctx context.Context, req *v1.CompletedTradesListReq) (res *v1.CompletedTradesListRes, err error)
}
