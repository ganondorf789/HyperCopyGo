// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package trader_positions

import (
	"context"

	"demo/api/trader_positions/v1"
)

type ITraderPositionsV1 interface {
	TraderPositionsList(ctx context.Context, req *v1.TraderPositionsListReq) (res *v1.TraderPositionsListRes, err error)
}
