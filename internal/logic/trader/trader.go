package trader

import (
	"context"

	v1 "demo/api/trader/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterTrader(&sTrader{})
}

type sTrader struct{}

// Popular 获取热门地址列表：IsHotAddress 的 trader + month 窗口统计
func (s *sTrader) Popular(ctx context.Context) (res *v1.TraderPopularRes, err error) {
	var traders []entity.Traders
	err = dao.Traders.Ctx(ctx).
		Where(do.Traders{IsHotAddress: true}).
		Scan(&traders)
	if err != nil {
		return nil, err
	}

	if len(traders) == 0 {
		return &v1.TraderPopularRes{List: make([]model.PopularTraderItem, 0)}, nil
	}

	addresses := make([]string, 0, len(traders))
	for _, t := range traders {
		addresses = append(addresses, t.Address)
	}

	var stats []entity.TraderStatistics
	err = dao.TraderStatistics.Ctx(ctx).
		Where(do.TraderStatistics{Window: "month"}).
		WhereIn(dao.TraderStatistics.Columns().Address, addresses).
		Scan(&stats)
	if err != nil {
		return nil, err
	}

	statsMap := make(map[string]*entity.TraderStatistics, len(stats))
	for i := range stats {
		statsMap[stats[i].Address] = &stats[i]
	}

	list := make([]model.PopularTraderItem, 0, len(traders))
	for _, t := range traders {
		item := model.PopularTraderItem{
			Address:   t.Address,
			UserPhoto: t.ProfilePicture,
			Labels:    t.Labels,
			Remark:    t.TwitterName,
		}
		if item.Labels == nil {
			item.Labels = make([]string, 0)
		}
		if st, ok := statsMap[t.Address]; ok {
			item.WinRate = st.WinRate
			item.RealizedPnl = st.LongRealizedPnl + st.ShortRealizedPnl
			item.AccountTotalValue = st.TotalValue
			item.CurrentPosition = st.PositionCount
		}
		list = append(list, item)
	}

	return &v1.TraderPopularRes{List: list}, nil
}
