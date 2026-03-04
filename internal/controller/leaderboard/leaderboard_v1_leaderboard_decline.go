package leaderboard

import (
	"context"

	"demo/api/leaderboard/v1"
	"demo/internal/service"
)

func (c *ControllerV1) LeaderboardDecline(ctx context.Context, req *v1.LeaderboardDeclineReq) (res *v1.LeaderboardDeclineRes, err error) {
	return service.Leaderboard().Decline(ctx)
}
