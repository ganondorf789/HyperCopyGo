package trader

import (
	"context"

	v1 "demo/api/trader/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterTrader(&sTrader{})
}

type sTrader struct{}

type popularTraderRow struct {
	Address        string   `json:"address"         orm:"address"`
	ProfilePicture string   `json:"profilePicture"  orm:"profile_picture"`
	TwitterName    string   `json:"twitterName"     orm:"twitter_name"`
	Labels         []string `json:"labels"          orm:"labels"`
	WinRate        float64  `json:"winRate"          orm:"win_rate"`
	LongRealizedPnl  float64 `json:"longRealizedPnl"  orm:"long_realized_pnl"`
	ShortRealizedPnl float64 `json:"shortRealizedPnl" orm:"short_realized_pnl"`
	TotalValue     float64  `json:"totalValue"       orm:"total_value"`
	PositionCount  float64  `json:"positionCount"    orm:"position_count"`
}

// Popular 获取热门地址列表：IsHotAddress 的 trader LEFT JOIN month 窗口统计
func (s *sTrader) Popular(ctx context.Context) (res *v1.TraderPopularRes, err error) {
	var (
		tTable  = dao.Traders.Table()
		tsTable = dao.TraderStatistics.Table()
		tsCls   = dao.TraderStatistics.Columns()
	)

	var rows []popularTraderRow
	err = g.Model(tTable, "t").
		LeftJoin(tsTable, "ts", "t.address = ts.address AND ts."+tsCls.Window+" = 'month'").
		FieldsPrefix("t", "address", "profile_picture", "twitter_name", "labels").
		FieldsPrefix("ts", "win_rate", "long_realized_pnl", "short_realized_pnl", "total_value", "position_count").
		Where("t.is_hot_address", true).
		Ctx(ctx).
		Scan(&rows)
	if err != nil {
		return nil, err
	}

	list := make([]model.PopularTraderItem, 0, len(rows))
	for _, r := range rows {
		labels := r.Labels
		if labels == nil {
			labels = make([]string, 0)
		}
		list = append(list, model.PopularTraderItem{
			Address:           r.Address,
			UserPhoto:         r.ProfilePicture,
			WinRate:           r.WinRate,
			RealizedPnl:       r.LongRealizedPnl + r.ShortRealizedPnl,
			AccountTotalValue: r.TotalValue,
			CurrentPosition:   r.PositionCount,
			Labels:            labels,
			Remark:            r.TwitterName,
		})
	}

	return &v1.TraderPopularRes{List: list}, nil
}
