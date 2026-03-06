// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package leaderboard

import (
	"context"

	"demo/api/leaderboard/v1"
)

type ILeaderboardV1 interface {
	LeaderboardTop(ctx context.Context, req *v1.LeaderboardTopReq) (res *v1.LeaderboardTopRes, err error)
	LeaderboardDecline(ctx context.Context, req *v1.LeaderboardDeclineReq) (res *v1.LeaderboardDeclineRes, err error)
	LeaderboardVolume(ctx context.Context, req *v1.LeaderboardVolumeReq) (res *v1.LeaderboardVolumeRes, err error)
	LeaderboardProfit(ctx context.Context, req *v1.LeaderboardProfitReq) (res *v1.LeaderboardProfitRes, err error)
}
