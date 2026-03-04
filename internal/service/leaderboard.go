// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/leaderboard/v1"
)

type (
	ILeaderboard interface {
		Top(ctx context.Context) (res *v1.LeaderboardTopRes, err error)
		Decline(ctx context.Context) (res *v1.LeaderboardDeclineRes, err error)
		Volume(ctx context.Context) (res *v1.LeaderboardVolumeRes, err error)
	}
)

var (
	localLeaderboard ILeaderboard
)

func Leaderboard() ILeaderboard {
	if localLeaderboard == nil {
		panic("implement not found for interface ILeaderboard, forgot register?")
	}
	return localLeaderboard
}

func RegisterLeaderboard(i ILeaderboard) {
	localLeaderboard = i
}
