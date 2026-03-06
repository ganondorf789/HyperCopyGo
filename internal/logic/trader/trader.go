package trader

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	v1 "demo/api/trader/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterTrader(&sTrader{})
}

type sTrader struct{}

// sort/filter 字段白名单 → 带表前缀的 SQL 列名
var discoverFieldMap = map[string]string{
	"sharpe":             "ts.sharpe",
	"drawdown":           "ts.drawdown",
	"positionCount":      "ts.position_count",
	"totalValue":         "ts.total_value",
	"perpValue":          "ts.perp_value",
	"positionValue":      "ts.position_value",
	"longPositionValue":  "ts.long_position_value",
	"shortPositionValue": "ts.short_position_value",
	"marginUsage":        "ts.margin_usage",
	"usedMargin":         "ts.used_margin",
	"profitCount":        "ts.profit_count",
	"winRate":            "ts.win_rate",
	"totalPnl":           "ts.total_pnl",
	"longCount":          "ts.long_count",
	"longRealizedPnl":    "ts.long_realized_pnl",
	"longWinRate":        "ts.long_win_rate",
	"shortCount":         "ts.short_count",
	"shortRealizedPnl":   "ts.short_realized_pnl",
	"shortWinRate":       "ts.short_win_rate",
	"unrealizedPnl":      "ts.unrealized_pnl",
	"avgLeverage":        "ts.avg_leverage",
	"snapEffLeverage":        "t.snap_eff_leverage",
	"snapLongPositionCount":  "t.snap_long_position_count",
	"snapLongPositionValue":  "t.snap_long_position_value",
	"snapMarginUsageRate":    "t.snap_margin_usage_rate",
	"snapPerpValue":          "t.snap_perp_value",
	"snapPositionCount":      "t.snap_position_count",
	"snapPositionValue":      "t.snap_position_value",
	"snapShortPositionCount": "t.snap_short_position_count",
	"snapShortPositionValue": "t.snap_short_position_value",
	"snapSpotValue":          "t.snap_spot_value",
	"snapTotalMarginUsed":    "t.snap_total_margin_used",
	"snapTotalValue":         "t.snap_total_value",
	"snapUnrealizedPnl":      "t.snap_unrealized_pnl",
}

var allowedFilterOps = map[string]bool{
	"<": true, "=": true, ">": true, ">=": true, "<=": true, "!=": true, "<>": true,
}

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

type discoverRow struct {
	Address                string   `orm:"address"`
	Labels                 []string `orm:"labels"`
	SnapEffLeverage        float64  `orm:"snap_eff_leverage"`
	SnapLongPositionCount  int64    `orm:"snap_long_position_count"`
	SnapLongPositionValue  float64  `orm:"snap_long_position_value"`
	SnapMarginUsageRate    float64  `orm:"snap_margin_usage_rate"`
	SnapPerpValue          float64  `orm:"snap_perp_value"`
	SnapPositionCount      int64    `orm:"snap_position_count"`
	SnapPositionValue      float64  `orm:"snap_position_value"`
	SnapShortPositionCount int64    `orm:"snap_short_position_count"`
	SnapShortPositionValue float64  `orm:"snap_short_position_value"`
	SnapSpotValue          float64  `orm:"snap_spot_value"`
	SnapTotalMarginUsed    float64  `orm:"snap_total_margin_used"`
	SnapTotalValue         float64  `orm:"snap_total_value"`
	SnapUnrealizedPnl      float64  `orm:"snap_unrealized_pnl"`
	LongPnl                float64  `orm:"long_pnl"`
	ShortPnl               float64  `orm:"short_pnl"`
	TotalPnl               float64  `orm:"total_pnl"`
	Sharpe                 float64  `orm:"sharpe"`
	Drawdown               float64  `orm:"drawdown"`
	WinRate                float64  `orm:"win_rate"`
	AvgLeverage            float64  `orm:"avg_leverage"`
	TsLongWinRate          float64  `orm:"ts_long_win_rate"`
	TsShortWinRate         float64  `orm:"ts_short_win_rate"`
}

// Discover 发现交易员，支持排序/筛选/标签/币种/分页
func (s *sTrader) Discover(ctx context.Context, in v1.TraderDiscoverReq) (res *v1.TraderDiscoverRes, err error) {
	var (
		tTable  = dao.Traders.Table()
		tsTable = dao.TraderStatistics.Table()
		tsCls   = dao.TraderStatistics.Columns()
	)

	selectFields := `
		t.address, t.labels,
		t.snap_eff_leverage, t.snap_long_position_count, t.snap_long_position_value,
		t.snap_margin_usage_rate, t.snap_perp_value, t.snap_position_count,
		t.snap_position_value, t.snap_short_position_count, t.snap_short_position_value,
		t.snap_spot_value, t.snap_total_margin_used, t.snap_total_value,
		t.snap_unrealized_pnl, t.long_pnl, t.short_pnl, t.total_pnl,
		ts.sharpe, ts.drawdown, ts.win_rate, ts.avg_leverage,
		ts.long_win_rate AS ts_long_win_rate,
		ts.short_win_rate AS ts_short_win_rate`

	orm := g.Model(tTable, "t").
		LeftJoin(tsTable, "ts", "t.address = ts.address AND ts."+tsCls.Window+" = '"+in.Window+"'").
		Fields(selectFields).
		Ctx(ctx)

	// tags 筛选（labels jsonb 包含）
	if len(in.Tags) > 0 {
		tagsJSON, _ := json.Marshal(in.Tags)
		orm = orm.Where("t.labels @> ?::jsonb", string(tagsJSON))
	}

	// coins 筛选（coins jsonb 有交集）
	if len(in.Coins) > 0 {
		coinsJSON, _ := json.Marshal(in.Coins)
		orm = orm.Where("ts.coins @> ?::jsonb", string(coinsJSON))
	}

	// 动态 filters
	for _, f := range in.Filters {
		col, ok := discoverFieldMap[f.Field]
		if !ok {
			return nil, fmt.Errorf("不支持的筛选字段: %s", f.Field)
		}
		if f.Op == "exist" {
			orm = orm.Where(fmt.Sprintf("%s IS NOT NULL", col))
			continue
		}
		if !allowedFilterOps[f.Op] {
			return nil, fmt.Errorf("不支持的操作符: %s", f.Op)
		}
		orm = orm.Where(fmt.Sprintf("%s %s ?", col, f.Op), f.Val)
	}

	// 排序
	if in.Sort != nil && in.Sort.Field != "" {
		col, ok := discoverFieldMap[in.Sort.Field]
		if !ok {
			return nil, fmt.Errorf("不支持的排序字段: %s", in.Sort.Field)
		}
		dir := "DESC"
		if in.Sort.Dir == "ASC" {
			dir = "ASC"
		}
		orm = orm.Order(col + " " + dir)
	}

	total, err := orm.Count()
	if err != nil {
		return nil, err
	}

	var rows []discoverRow
	err = orm.Page(in.Page, in.PageSize).Scan(&rows)
	if err != nil {
		return nil, err
	}

	// 批量获取 pnl 历史
	addresses := make([]string, 0, len(rows))
	for _, r := range rows {
		addresses = append(addresses, r.Address)
	}
	pnlMap := make(map[string][]model.PnlPoint)
	if len(addresses) > 0 {
		var pnlRecords []entity.TraderPnlHistories
		err = dao.TraderPnlHistories.Ctx(ctx).
			WhereIn("address", addresses).
			Where("window", in.Window).
			Scan(&pnlRecords)
		if err != nil {
			return nil, err
		}
		for _, rec := range pnlRecords {
			pnlMap[rec.Address] = parsePnlHistory(rec.History)
		}
	}

	list := make([]model.DiscoverTraderItem, 0, len(rows))
	for _, r := range rows {
		tags := r.Labels
		if tags == nil {
			tags = make([]string, 0)
		}
		pnlList := pnlMap[r.Address]
		if pnlList == nil {
			pnlList = make([]model.PnlPoint, 0)
		}
		list = append(list, model.DiscoverTraderItem{
			Address:                r.Address,
			AvgLeverage:            r.AvgLeverage,
			DdDrawdown:             r.Drawdown,
			LongPnl:                r.LongPnl,
			LongWinRate:            r.TsLongWinRate,
			PnlList:                pnlList,
			Sharpe:                 r.Sharpe,
			ShortPnl:               r.ShortPnl,
			ShortWinRate:           r.TsShortWinRate,
			SnapEffLeverage:        r.SnapEffLeverage,
			SnapLongPositionCount:  r.SnapLongPositionCount,
			SnapLongPositionValue:  r.SnapLongPositionValue,
			SnapMarginUsageRate:    r.SnapMarginUsageRate,
			SnapPerpValue:          r.SnapPerpValue,
			SnapPositionCount:      r.SnapPositionCount,
			SnapPositionValue:      r.SnapPositionValue,
			SnapShortPositionCount: r.SnapShortPositionCount,
			SnapShortPositionValue: r.SnapShortPositionValue,
			SnapSpotValue:          r.SnapSpotValue,
			SnapTotalMarginUsed:    r.SnapTotalMarginUsed,
			SnapTotalValue:         r.SnapTotalValue,
			SnapUnrealizedPnl:      r.SnapUnrealizedPnl,
			Tags:                   tags,
			TotalPnl:               r.TotalPnl,
			WinRate:                r.WinRate,
		})
	}

	return &v1.TraderDiscoverRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

// parsePnlHistory 解析 "[[ts,val],[ts,val],...]" 为 PnlPoint 切片
func parsePnlHistory(raw string) []model.PnlPoint {
	if raw == "" {
		return nil
	}
	var pairs [][]json.RawMessage
	if err := json.Unmarshal([]byte(raw), &pairs); err != nil {
		return nil
	}
	points := make([]model.PnlPoint, 0, len(pairs))
	for _, p := range pairs {
		if len(p) < 2 {
			continue
		}
	var ts int64
	var v string
	_ = json.Unmarshal(p[0], &ts)
	_ = json.Unmarshal(p[1], &v)
	points = append(points, model.PnlPoint{Ts: ts, V: v})
	}
	return points
}

func (s *sTrader) CompletedTradesList(ctx context.Context, in v1.TraderCompletedTradesListReq) (res *v1.TraderCompletedTradesListRes, err error) {
	m := dao.CompletedTrades.Ctx(ctx)

	if in.Address != "" {
		m = m.Where("address = ?", in.Address)
	}
	if in.Coin != "" {
		m = m.Where("coin = ?", in.Coin)
	}
	if in.Direction != "" {
		m = m.Where("direction = ?", in.Direction)
	}

	switch in.Window {
	case "day":
		m = m.Where("end_time >= ?", time.Now().Add(-24*time.Hour).UnixMilli())
	case "week":
		m = m.Where("end_time >= ?", time.Now().Add(-7*24*time.Hour).UnixMilli())
	case "month":
		m = m.Where("end_time >= ?", time.Now().Add(-30*24*time.Hour).UnixMilli())
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.CompletedTrades
	err = m.Page(in.Page, in.PageSize).
		OrderDesc("end_time").
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.CompletedTradeItem, 0, len(items))
	for _, e := range items {
		list = append(list, model.CompletedTradeItem{
			Id:         e.Id,
			Address:    e.Address,
			Coin:       e.Coin,
			MarginMode: e.MarginMode,
			Direction:  e.Direction,
			Size:       e.Size,
			EntryPrice: e.EntryPrice,
			ClosePrice: e.ClosePrice,
			StartTime:  e.StartTime,
			EndTime:    e.EndTime,
			TotalFee:   e.TotalFee,
			Pnl:        e.Pnl,
			FillCount:  e.FillCount,
			CreatedAt:  e.CreatedAt,
			UpdatedAt:  e.UpdatedAt,
		})
	}

	return &v1.TraderCompletedTradesListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}
