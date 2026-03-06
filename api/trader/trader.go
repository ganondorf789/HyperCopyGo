// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package trader

import (
	"context"

	"demo/api/trader/v1"
)

type ITraderV1 interface {
	TraderPopular(ctx context.Context, req *v1.TraderPopularReq) (res *v1.TraderPopularRes, err error)
	TraderDiscover(ctx context.Context, req *v1.TraderDiscoverReq) (res *v1.TraderDiscoverRes, err error)
	TraderKolList(ctx context.Context, req *v1.TraderKolListReq) (res *v1.TraderKolListRes, err error)
	TraderCompletedTradesList(ctx context.Context, req *v1.TraderCompletedTradesListReq) (res *v1.TraderCompletedTradesListRes, err error)
}
