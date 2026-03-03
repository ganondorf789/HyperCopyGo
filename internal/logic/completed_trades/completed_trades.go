package completed_trades

import (
	"context"
	"time"

	v1 "demo/api/completed_trades/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterCompletedTrades(&sCompletedTrades{})
}

type sCompletedTrades struct{}

func (s *sCompletedTrades) List(ctx context.Context, in v1.CompletedTradesListReq) (res *v1.CompletedTradesListRes, err error) {
	m := dao.CompletedTrades.Ctx(ctx)

	if in.Address != "" {
		m = m.Where(do.CompletedTrades{Address: in.Address})
	}
	if in.Coin != "" {
		m = m.Where(do.CompletedTrades{Coin: in.Coin})
	}
	if in.Direction != "" {
		m = m.Where(do.CompletedTrades{Direction: in.Direction})
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
		list = append(list, entityToItem(e))
	}

	return &v1.CompletedTradesListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func entityToItem(e entity.CompletedTrades) model.CompletedTradeItem {
	return model.CompletedTradeItem{
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
	}
}
