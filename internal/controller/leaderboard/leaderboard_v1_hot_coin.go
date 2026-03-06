package leaderboard

import (
	"context"

	"demo/api/leaderboard/v1"
	"demo/internal/service"
)

func (c *ControllerV1) HotCoin(ctx context.Context, req *v1.HotCoinReq) (res *v1.HotCoinRes, err error) {
	return service.Leaderboard().HotCoin(ctx)
}
