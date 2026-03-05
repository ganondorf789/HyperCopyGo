package trader_performance

import (
	"context"
	"math"
	"sort"
	"time"

	v1 "demo/api/trader_performance/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterTraderPerformance(&sTraderPerformance{})
}

type sTraderPerformance struct{}

func (s *sTraderPerformance) Performance(ctx context.Context, in v1.TraderPerformanceReq) (res *v1.TraderPerformanceRes, err error) {
	m := dao.CompletedTrades.Ctx(ctx).Where(entity.CompletedTrades{Address: in.Address})

	switch in.Window {
	case "day":
		m = m.Where("end_time >= ?", time.Now().Add(-24*time.Hour).UnixMilli())
	case "week":
		m = m.Where("end_time >= ?", time.Now().Add(-7*24*time.Hour).UnixMilli())
	case "month":
		m = m.Where("end_time >= ?", time.Now().Add(-30*24*time.Hour).UnixMilli())
	}

	var items []entity.CompletedTrades
	if err = m.OrderDesc("pnl").Scan(&items); err != nil {
		return nil, err
	}

	res = &v1.TraderPerformanceRes{
		BestTrades:        make([]model.TradePerformanceBestTrade, 0),
		PerformanceAssets: make([]model.TradePerformanceAsset, 0),
	}
	if len(items) == 0 {
		return
	}

	var totalPnl, fees, longPnl, shortPnl float64
	var totalDurationMs int64
	minDurationMs := int64(math.MaxInt64)
	var maxDurationMs int64
	var winning int

	coinMap := make(map[string]*model.TradePerformanceAsset)

	for _, e := range items {
		totalPnl += e.Pnl
		fees += e.TotalFee

		if e.Direction == "long" {
			longPnl += e.Pnl
		} else {
			shortPnl += e.Pnl
		}

		if e.Pnl > 0 {
			winning++
		}

		duration := e.EndTime - e.StartTime
		totalDurationMs += duration
		if duration < minDurationMs {
			minDurationMs = duration
		}
		if duration > maxDurationMs {
			maxDurationMs = duration
		}

		res.BestTrades = append(res.BestTrades, model.TradePerformanceBestTrade{
			Coin:      e.Coin,
			Direction: e.Direction,
			Duration:  duration,
			CreateAt:  e.StartTime,
			Pnl:       e.Pnl,
		})

		asset, ok := coinMap[e.Coin]
		if !ok {
			asset = &model.TradePerformanceAsset{Coin: e.Coin}
			coinMap[e.Coin] = asset
		}
		asset.Trades++
		asset.Pnl += e.Pnl
		asset.Fees += e.TotalFee
	}

	total := len(items)

	winRate := 0.0
	if total > 0 {
		winRate = float64(winning) / float64(total)
	}

	performanceAssets := make([]model.TradePerformanceAsset, 0, len(coinMap))
	for _, v := range coinMap {
		performanceAssets = append(performanceAssets, *v)
	}
	sort.Slice(performanceAssets, func(i, j int) bool {
		return performanceAssets[i].Pnl > performanceAssets[j].Pnl
	})

	res.TotalPnl = totalPnl
	res.Gross = totalPnl + fees
	res.Fees = fees
	res.LongPnl = longPnl
	res.ShortPnl = shortPnl
	res.WinRate = winRate
	res.Winning = winning
	res.Total = total
	res.TradeDuration = totalDurationMs / 1000
	res.MinDuration = minDurationMs / 1000
	res.MaxDuration = maxDurationMs / 1000
	res.PerformanceAssets = performanceAssets

	return
}

func (s *sTraderPerformance) Summary(ctx context.Context, in v1.TraderPerformanceSummaryReq) (res *v1.TraderPerformanceSummaryRes, err error) {
	res = &v1.TraderPerformanceSummaryRes{}

	windowStart := time.Now().Add(-7 * 24 * time.Hour).UnixMilli()
	tradesQuery := dao.CompletedTrades.Ctx(ctx).
		Where(entity.CompletedTrades{Address: in.Address}).
		Where("end_time >= ?", windowStart)

	var trades []entity.CompletedTrades
	if err = tradesQuery.OrderAsc("end_time").Scan(&trades); err != nil {
		return nil, err
	}

	closePosCount := len(trades)
	res.ClosePosCount = closePosCount

	if closePosCount > 0 {
		var totalPnl float64
		var totalDurationMs int64
		var winning int

		var peakCumPnl float64
		var maxDrawdown float64
		var cumPnl float64

		for _, e := range trades {
			totalPnl += e.Pnl
			totalDurationMs += e.EndTime - e.StartTime

			if e.Pnl > 0 {
				winning++
			}

			cumPnl += e.Pnl
			if cumPnl > peakCumPnl {
				peakCumPnl = cumPnl
			}
			if peakCumPnl > 0 {
				dd := (peakCumPnl - cumPnl) / peakCumPnl
				if dd > maxDrawdown {
					maxDrawdown = dd
				}
			}
		}

		res.WinRate = float64(winning) / float64(closePosCount)
		res.TotalPnl = totalPnl
		res.AvgPosDuration = totalDurationMs / int64(closePosCount) / 1000
		res.MaxDrawdown = maxDrawdown
	}

	ordersQuery := dao.TraderOrders.Ctx(ctx).
		Where(entity.TraderOrders{Address: in.Address}).
		Where("timestamp >= ?", windowStart)
	orderCount, err := ordersQuery.Count()
	if err != nil {
		return nil, err
	}
	res.OrderCount = orderCount

	return
}

