package trader_positions

import (
	"context"

	v1 "demo/api/trader_positions/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterTraderPositions(&sTraderPositions{})
}

type sTraderPositions struct{}

func (s *sTraderPositions) List(ctx context.Context, in v1.TraderPositionsListReq) (res *v1.TraderPositionsListRes, err error) {
	m := dao.TraderPositions.Ctx(ctx)

	if in.Address != "" {
		m = m.Where(entity.TraderPositions{Address: in.Address})
	}
	if in.Coin != "" {
		m = m.Where(entity.TraderPositions{Coin: in.Coin})
	}
	if in.Direction == "long" {
		m = m.Where("szi > 0")
	} else if in.Direction == "short" {
		m = m.Where("szi < 0")
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.TraderPositions
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.TraderPositions.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.TraderPositionItem, 0, len(items))
	for _, e := range items {
		list = append(list, entityToItem(e))
	}

	return &v1.TraderPositionsListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func entityToItem(e entity.TraderPositions) model.TraderPositionItem {
	return model.TraderPositionItem{
		Id:                    e.Id,
		Address:               e.Address,
		Coin:                  e.Coin,
		Szi:                   e.Szi,
		LeverageType:          e.LeverageType,
		Leverage:              e.Leverage,
		EntryPx:               e.EntryPx,
		PositionValue:         e.PositionValue,
		UnrealizedPnl:         e.UnrealizedPnl,
		ReturnOnEquity:        e.ReturnOnEquity,
		LiquidationPx:         e.LiquidationPx,
		MarginUsed:            e.MarginUsed,
		MaxLeverage:           e.MaxLeverage,
		CumFundingAllTime:     e.CumFundingAllTime,
		CumFundingSinceOpen:   e.CumFundingSinceOpen,
		CumFundingSinceChange: e.CumFundingSinceChange,
		CreatedAt:             e.CreatedAt,
		UpdatedAt:             e.UpdatedAt,
	}
}
