package position

import (
	"context"
	"fmt"
	"strings"
	"time"

	v1 "demo/api/position/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterPosition(&sPosition{})
}

type sPosition struct{}

func (s *sPosition) List(ctx context.Context, in v1.PositionListReq) (res *v1.PositionListRes, err error) {
	m := dao.Position.Ctx(ctx)

	// Coin 筛选
	if in.User != "" {
		m = m.Where("user = ?", in.User)
	}
	if in.Symbol != "" {
		m = m.Where("symbol = ?", in.Symbol)
	}

	// Direction 筛选: long (position_size > 0) / short (position_size < 0)
	if in.Direction == "long" {
		m = m.Where("position_size > 0")
	} else if in.Direction == "short" {
		m = m.Where("position_size < 0")
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 排序：uPnL / Funding Fee / 默认按 id 倒序
	if in.UpnlSort == "asc" {
		m = m.OrderAsc("unrealized_pnl")
	} else if in.UpnlSort == "desc" {
		m = m.OrderDesc("unrealized_pnl")
	} else if in.FundingSort == "asc" {
		m = m.OrderAsc("funding_fee")
	} else if in.FundingSort == "desc" {
		m = m.OrderDesc("funding_fee")
	} else {
		m = m.OrderDesc(dao.Position.Columns().Id)
	}

	var items []entity.Position
	err = m.Page(in.Page, in.PageSize).Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.PositionItem, 0, len(items))
	for _, item := range items {
		list = append(list, s.entityToItem(item))
	}

	return &v1.PositionListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func (s *sPosition) entityToItem(e entity.Position) model.PositionItem {
	// labels: 逗号分隔字符串 -> []string
	var labels []string
	if e.Labels != "" {
		labels = strings.Split(e.Labels, ",")
		for i := range labels {
			labels[i] = strings.TrimSpace(labels[i])
		}
	} else {
		labels = make([]string, 0)
	}

	var createTime, updateTime int64
	if e.CreatedAt != nil {
		createTime = e.CreatedAt.UnixMilli()
	}
	if e.UpdatedAt != nil {
		updateTime = e.UpdatedAt.UnixMilli()
	}

	return model.PositionItem{
		Id:               e.Id,
		User:             e.User,
		Symbol:           e.Symbol,
		PositionSize:     e.PositionSize,
		EntryPrice:       e.EntryPrice,
		MarkPrice:        e.MarkPrice,
		LiqPrice:         e.LiqPrice,
		Leverage:         int(e.Leverage),
		MarginBalance:    e.MarginBalance,
		PositionValueUsd: e.PositionValueUsd,
		UnrealizedPnL:    e.UnrealizedPnl,
		FundingFee:       e.FundingFee,
		MarginMode:       e.MarginMode,
		CreateTime:       createTime,
		UpdateTime:       updateTime,
		Labels:           labels,
	}
}

// intervalConfig 时间框架配置
type intervalConfig struct {
	truncSQL string        // PostgreSQL date_trunc / time_bucket 表达式
	lookback time.Duration // 查询回溯时间范围
}

// getIntervalConfig 根据 interval 返回 SQL 截断表达式和回溯范围
func getIntervalConfig(interval string) (intervalConfig, error) {
	switch interval {
	case "5m":
		return intervalConfig{
			truncSQL: "to_timestamp(floor(extract(epoch from created_at) / 300) * 300)",
			lookback: 24 * time.Hour,
		}, nil
	case "30m":
		return intervalConfig{
			truncSQL: "to_timestamp(floor(extract(epoch from created_at) / 1800) * 1800)",
			lookback: 3 * 24 * time.Hour,
		}, nil
	case "1h":
		return intervalConfig{
			truncSQL: "date_trunc('hour', created_at)",
			lookback: 7 * 24 * time.Hour,
		}, nil
	case "4h":
		return intervalConfig{
			truncSQL: "to_timestamp(floor(extract(epoch from created_at) / 14400) * 14400)",
			lookback: 14 * 24 * time.Hour,
		}, nil
	case "12h":
		return intervalConfig{
			truncSQL: "to_timestamp(floor(extract(epoch from created_at) / 43200) * 43200)",
			lookback: 30 * 24 * time.Hour,
		}, nil
	case "1D":
		return intervalConfig{
			truncSQL: "date_trunc('day', created_at)",
			lookback: 90 * 24 * time.Hour,
		}, nil
	default:
		return intervalConfig{}, fmt.Errorf("不支持的时间框架: %s", interval)
	}
}

// statsRow 原生 SQL 查询结果行
type statsRow struct {
	Bucket        time.Time `json:"bucket"`
	LongAccounts  int       `json:"long_accounts"`
	ShortAccounts int       `json:"short_accounts"`
	LongLiqValue  float64   `json:"long_liq_value"`
	ShortLiqValue float64   `json:"short_liq_value"`
}

func (s *sPosition) Stats(ctx context.Context, in v1.PositionStatsReq) (res *v1.PositionStatsRes, err error) {
	cfg, err := getIntervalConfig(in.Interval)
	if err != nil {
		return nil, err
	}

	since := time.Now().Add(-cfg.lookback)

	symbolFilter := ""
	if in.Symbol != "" {
		symbolFilter = fmt.Sprintf("AND symbol = '%s'", strings.ReplaceAll(in.Symbol, "'", "''"))
	}

	sql := fmt.Sprintf(`
		SELECT
			%s AS bucket,
			COUNT(DISTINCT CASE WHEN position_size > 0 THEN "user" END) AS long_accounts,
			COUNT(DISTINCT CASE WHEN position_size < 0 THEN "user" END) AS short_accounts,
			COALESCE(SUM(CASE WHEN position_size > 0 THEN ABS(position_value_usd) ELSE 0 END), 0) AS long_liq_value,
			COALESCE(SUM(CASE WHEN position_size < 0 THEN ABS(position_value_usd) ELSE 0 END), 0) AS short_liq_value
		FROM position
		WHERE created_at >= $1 %s
		GROUP BY bucket
		ORDER BY bucket ASC
	`, cfg.truncSQL, symbolFilter)

	var rows []statsRow
	err = dao.Position.DB().GetScan(ctx, &rows, sql, since)
	if err != nil {
		return nil, err
	}

	list := make([]model.PositionStatsPoint, 0, len(rows))
	for _, r := range rows {
		list = append(list, model.PositionStatsPoint{
			Timestamp:     r.Bucket.UnixMilli(),
			LongAccounts:  r.LongAccounts,
			ShortAccounts: r.ShortAccounts,
			LongLiqValue:  r.LongLiqValue,
			ShortLiqValue: r.ShortLiqValue,
		})
	}

	return &v1.PositionStatsRes{List: list}, nil
}

// ratioRow 多空比率原生 SQL 查询结果行
type ratioRow struct {
	Bucket     time.Time `json:"bucket"`
	LongValue  float64   `json:"long_value"`
	ShortValue float64   `json:"short_value"`
}

func (s *sPosition) LongShortRatio(ctx context.Context, in v1.PositionLongShortRatioReq) (res *v1.PositionLongShortRatioRes, err error) {
	// 仅支持 1h, 4h, 1D
	var truncSQL string
	var lookback time.Duration
	switch in.Interval {
	case "1h":
		truncSQL = "date_trunc('hour', created_at)"
		lookback = 7 * 24 * time.Hour
	case "4h":
		truncSQL = "to_timestamp(floor(extract(epoch from created_at) / 14400) * 14400)"
		lookback = 14 * 24 * time.Hour
	case "1D":
		truncSQL = "date_trunc('day', created_at)"
		lookback = 90 * 24 * time.Hour
	default:
		return nil, fmt.Errorf("不支持的时间框架: %s", in.Interval)
	}

	since := time.Now().Add(-lookback)

	symbolFilter := ""
	if in.Symbol != "" {
		symbolFilter = fmt.Sprintf("AND symbol = '%s'", strings.ReplaceAll(in.Symbol, "'", "''"))
	}

	sql := fmt.Sprintf(`
		SELECT
			%s AS bucket,
			COALESCE(SUM(CASE WHEN position_size > 0 THEN ABS(position_value_usd) ELSE 0 END), 0) AS long_value,
			COALESCE(SUM(CASE WHEN position_size < 0 THEN ABS(position_value_usd) ELSE 0 END), 0) AS short_value
		FROM position
		WHERE created_at >= $1 %s
		GROUP BY bucket
		ORDER BY bucket ASC
	`, truncSQL, symbolFilter)

	var rows []ratioRow
	err = dao.Position.DB().GetScan(ctx, &rows, sql, since)
	if err != nil {
		return nil, err
	}

	list := make([]model.LongShortRatioPoint, 0, len(rows))
	for _, r := range rows {
		ratio := 0.0
		if r.ShortValue > 0 {
			ratio = r.LongValue / r.ShortValue
		}
		list = append(list, model.LongShortRatioPoint{
			Timestamp:               r.Bucket.UnixMilli(),
			LongShortRatio:          ratio,
			PositionValueDifference: r.LongValue - r.ShortValue,
		})
	}

	return &v1.PositionLongShortRatioRes{List: list}, nil
}
