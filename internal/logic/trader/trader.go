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
	Address          string   `orm:"address"`
	ProfilePicture   string   `orm:"profile_picture"`
	TwitterName      string   `orm:"twitter_name"`
	Labels           []string `orm:"labels"`
	WinRate          float64  `orm:"win_rate"`
	LongRealizedPnl  float64  `orm:"long_realized_pnl"`
	ShortRealizedPnl float64  `orm:"short_realized_pnl"`
	TotalValue       float64  `orm:"total_value"`
	PositionCount    float64  `orm:"position_count"`
}

type kolTraderRow struct {
	TwitterName    string   `orm:"twitter_name"`
	Username       string   `orm:"username"`
	Address        string   `orm:"address"`
	ProfilePicture string   `orm:"profile_picture"`
	Labels         []string `orm:"labels"`
	TotalValue     float64  `orm:"total_value"`
	WinRate        float64  `orm:"win_rate"`
	PositionCount  float64  `orm:"position_count"`
	TotalPnl       float64  `orm:"total_pnl"`
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

// KolList 获取 X KOL 列表，支持分页和窗口筛选
func (s *sTrader) KolList(ctx context.Context, in v1.TraderKolListReq) (res *v1.TraderKolListRes, err error) {
	var (
		tTable  = dao.Traders.Table()
		tsTable = dao.TraderStatistics.Table()
		tsCls   = dao.TraderStatistics.Columns()
	)

	orm := g.Model(tTable, "t").
		LeftJoin(tsTable, "ts", "t.address = ts.address AND ts."+tsCls.Window+" = '"+in.Window+"'").
		FieldsPrefix("t", "twitter_name", "username", "address", "profile_picture", "labels").
		FieldsPrefix("ts", "total_value", "win_rate", "position_count", "total_pnl").
		Where("t.is_twitter_kol", true).
		Ctx(ctx)

	total, err := orm.Count()
	if err != nil {
		return nil, err
	}

	var rows []kolTraderRow
	err = orm.Page(in.Page, in.PageSize).Scan(&rows)
	if err != nil {
		return nil, err
	}

	list := make([]model.KolTraderItem, 0, len(rows))
	for _, r := range rows {
		labels := r.Labels
		if labels == nil {
			labels = make([]string, 0)
		}
		list = append(list, model.KolTraderItem{
			TwitterName:       r.TwitterName,
			Username:          r.Username,
			Address:           r.Address,
			AccountTotalValue: r.TotalValue,
			WinRate:           r.WinRate,
			PositionCount:     r.PositionCount,
			TotalPnl:          r.TotalPnl,
			ProfilePicture:    r.ProfilePicture,
			Labels:            labels,
		})
	}

	return &v1.TraderKolListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}
