package leaderboard

import (
	"context"

	v1 "demo/api/leaderboard/v1"
	"demo/internal/dao"
	"demo/internal/model"
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

func (s *sLeaderboard) Profit(ctx context.Context, in v1.LeaderboardProfitReq) (res *v1.LeaderboardProfitRes, err error) {
	m := dao.TraderPerformances.DB().Model("trader_performances AS tp").
		Ctx(ctx).
		Fields(`
			tp.address                        AS eth_address,
			NULLIF(TRIM(t.twitter_name), '')  AS display_name,
			COALESCE(ts.total_value, 0)       AS account_value,
			tp.window,
			tp.pnl,
			tp.roi,
			tp.vlm
		`).
		LeftJoin("traders AS t", "t.address = tp.address").
		LeftJoin("trader_statistics AS ts", "ts.address = tp.address AND ts.window = tp.window").
		Where("tp.window = ?", in.Window).
		OrderDesc("tp.pnl")

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	type profitRow struct {
		EthAddress   string   `json:"ethAddress"`
		DisplayName  *string  `json:"displayName"`
		AccountValue float64  `json:"accountValue"`
		Window       string   `json:"window"`
		Pnl          float64  `json:"pnl"`
		Roi          float64  `json:"roi"`
		Vlm          float64  `json:"vlm"`
	}

	var rows []profitRow
	err = m.Page(in.Page, in.PageSize).Scan(&rows)
	if err != nil {
		return nil, err
	}

	list := make([]model.LeaderboardProfitItem, 0, len(rows))
	for _, r := range rows {
		list = append(list, model.LeaderboardProfitItem{
			EthAddress:   r.EthAddress,
			DisplayName:  r.DisplayName,
			AccountValue: r.AccountValue,
			Window:       r.Window,
			Pnl:          r.Pnl,
			Roi:          r.Roi,
			Vlm:          r.Vlm,
		})
	}

	return &v1.LeaderboardProfitRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}
