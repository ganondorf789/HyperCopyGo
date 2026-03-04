package leaderboard

import (
	"context"

	v1 "demo/api/leaderboard/v1"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterLeaderboard(&sLeaderboard{})
}

type sLeaderboard struct{}

func (s *sLeaderboard) Top(ctx context.Context) (res *v1.LeaderboardTopRes, err error) {
	var list []entity.Leaderboard
	err = dao.Leaderboard.Ctx(ctx).
		OrderDesc(dao.Leaderboard.Columns().Roi).
		Limit(10).
		Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = make([]entity.Leaderboard, 0)
	}
	return &v1.LeaderboardTopRes{List: list}, nil
}

func (s *sLeaderboard) Decline(ctx context.Context) (res *v1.LeaderboardDeclineRes, err error) {
	var list []entity.Leaderboard
	err = dao.Leaderboard.Ctx(ctx).
		OrderAsc(dao.Leaderboard.Columns().Roi).
		Limit(10).
		Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = make([]entity.Leaderboard, 0)
	}
	return &v1.LeaderboardDeclineRes{List: list}, nil
}

func (s *sLeaderboard) Volume(ctx context.Context) (res *v1.LeaderboardVolumeRes, err error) {
	var list []entity.Leaderboard
	err = dao.Leaderboard.Ctx(ctx).
		OrderDesc(dao.Leaderboard.Columns().Vlm).
		Limit(10).
		Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = make([]entity.Leaderboard, 0)
	}
	return &v1.LeaderboardVolumeRes{List: list}, nil
}
