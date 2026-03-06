package leaderboard

import (
	"context"

	v1 "demo/api/leaderboard/v1"
	"demo/internal/service"
)

func (c *ControllerV1) LeaderboardProfit(ctx context.Context, req *v1.LeaderboardProfitReq) (res *v1.LeaderboardProfitRes, err error) {
	return service.Leaderboard().Profit(ctx, *req)
}
