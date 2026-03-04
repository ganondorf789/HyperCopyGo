package leaderboard

import (
	"context"

	"demo/api/leaderboard/v1"
	"demo/internal/service"
)

func (c *ControllerV1) LeaderboardVolume(ctx context.Context, req *v1.LeaderboardVolumeReq) (res *v1.LeaderboardVolumeRes, err error) {
	return service.Leaderboard().Volume(ctx)
}
