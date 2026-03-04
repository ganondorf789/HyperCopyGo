package leaderboard

import (
	"context"

	"demo/api/leaderboard/v1"
	"demo/internal/service"
)

func (c *ControllerV1) LeaderboardTop(ctx context.Context, req *v1.LeaderboardTopReq) (res *v1.LeaderboardTopRes, err error) {
	return service.Leaderboard().Top(ctx)
}
